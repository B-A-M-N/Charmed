// Package scanner — AST-based semantic graph builder and detector functions.
// This replaces the ATree dependency for graph extraction using Go's standard
// go/parser and go/ast packages.
package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"

	"github.com/B-A-M-N/charmed/charmed-mcp/internal/ir"
)

// buildGraphFromAST parses Go files and builds an AtreeSemanticGraph using
// only the standard library — no ATree dependency required.
func buildGraphFromAST(goFiles []string) (*AtreeSemanticGraph, error) {
	graph := &AtreeSemanticGraph{}
	fset := token.NewFileSet()

	for _, path := range goFiles {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			continue // skip unparseable files
		}

		fileSem := AtreeFileSem{
			Path:     path,
			Language: "go",
		}

		// Collect imports
		for _, imp := range f.Imports {
			fullPath := strings.Trim(imp.Path.Value, `"`)
			localName := filepath.Base(fullPath)
			isAlias := false
			if imp.Name != nil {
				localName = imp.Name.Name
				isAlias = true
			}
			graph.Imports = append(graph.Imports, AtreeImport{
				File:      path,
				LocalName: localName,
				FullPath:  fullPath,
				IsAlias:   isAlias,
			})
		}

		// Current enclosing function/method scope for nested inspections
		var currentScope string

		// Walk declarations
		for _, decl := range f.Decls {
			switch d := decl.(type) {
			case *ast.GenDecl:
				// Handle struct type declarations
				for _, spec := range d.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					if _, isStruct := ts.Type.(*ast.StructType); isStruct {
						pos := fset.Position(ts.Pos())
						sym := AtreeSymbol{
							Name:      ts.Name.Name,
							Qualified: ts.Name.Name,
							Kind:      "DefinitionStruct",
							File:      path,
							Line:      pos.Line,
						}
						fileSem.Defs = append(fileSem.Defs, sym)
					}
				}

			case *ast.FuncDecl:
				if d.Name == nil {
					continue
				}

				pos := fset.Position(d.Pos())
				bodyStart := pos.Line
				bodyEnd := pos.Line
				if d.Body != nil {
					bodyEnd = fset.Position(d.Body.End()).Line
				}

				sym := AtreeSymbol{
					Name:      d.Name.Name,
					File:      path,
					Line:      pos.Line,
					BodyStart: bodyStart,
					BodyEnd:   bodyEnd,
				}

				if d.Recv != nil && len(d.Recv.List) > 0 {
					// Method
					recv := d.Recv.List[0]
					recvName := receiverTypeName(recv)
					sym.Receiver = recvName
					sym.IsPtrRecv = isPointerReceiver(recv)
					sym.Qualified = recvName + "." + d.Name.Name
					sym.Kind = "DefinitionMethod"
					currentScope = recvName + "." + d.Name.Name
				} else {
					// Function
					sym.Qualified = d.Name.Name
					sym.Kind = "DefinitionFunction"
					currentScope = d.Name.Name
				}

				// Collect parameters
				if d.Type.Params != nil {
					for _, field := range d.Type.Params.List {
						typStr := exprToString(field.Type)
						if len(field.Names) == 0 {
							sym.Params = append(sym.Params, AtreeParam{Name: "_", Type: typStr})
						} else {
							for _, name := range field.Names {
								sym.Params = append(sym.Params, AtreeParam{Name: name.Name, Type: typStr})
							}
						}
					}
				}

				// Collect return types
				if d.Type.Results != nil {
					for _, field := range d.Type.Results.List {
						sym.Returns = append(sym.Returns, exprToString(field.Type))
					}
				}

				fileSem.Defs = append(fileSem.Defs, sym)

				// Collect switch cases and function calls within the body
				if d.Body != nil {
					scope := currentScope
					collectBodyInfo(fset, d.Body, path, scope, &fileSem, graph)
				}
			}
		}

		graph.Files = append(graph.Files, fileSem)
	}

	return graph, nil
}

// collectBodyInfo walks a function body collecting switch cases, calls, and assignments.
func collectBodyInfo(fset *token.FileSet, body *ast.BlockStmt, file, scope string, fileSem *AtreeFileSem, graph *AtreeSemanticGraph) {
	ast.Inspect(body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.TypeSwitchStmt:
			sw := AtreeSwitch{
				Scope: scope,
				File:  file,
				Line:  fset.Position(node.Pos()).Line,
			}
			// Check if assign is "msg.(type)" or similar
			if assign, ok := node.Assign.(*ast.AssignStmt); ok {
				for _, rhs := range assign.Rhs {
					if ta, ok := rhs.(*ast.TypeAssertExpr); ok {
						sw.VarType = exprToString(ta.X)
					}
				}
			} else if expr, ok := node.Assign.(*ast.ExprStmt); ok {
				if ta, ok := expr.X.(*ast.TypeAssertExpr); ok {
					sw.VarType = exprToString(ta.X)
				}
			}
			for _, stmt := range node.Body.List {
				cs, ok := stmt.(*ast.CaseClause)
				if !ok {
					continue
				}
				for _, caseType := range cs.List {
					casePos := fset.Position(cs.Pos())
					bodyEndPos := fset.Position(cs.End()).Line
					sw.Cases = append(sw.Cases, AtreeCase{
						Type:      exprToString(caseType),
						Line:      casePos.Line,
						BodyStart: casePos.Line,
						BodyEnd:   bodyEndPos,
						File:      file,
					})
				}
			}
			fileSem.SwitchCases = append(fileSem.SwitchCases, sw)

		case *ast.CallExpr:
			name, recv := callExprInfo(node)
			if name == "" {
				return true
			}
			callPos := fset.Position(node.Pos())
			fileSem.FunctionCalls = append(fileSem.FunctionCalls, AtreeFuncCall{
				Name:  name,
				Recv:  recv,
				Scope: scope,
				File:  file,
				Line:  callPos.Line,
			})
			// Add to global Calls as well
			graph.Calls = append(graph.Calls, AtreeCall{
				CallerScope: scope,
				CalleeName:  name,
				Receiver:    recv,
				Resolved:    true,
				Confidence:  1.0,
				File:        file,
				Line:        callPos.Line,
			})

		case *ast.AssignStmt:
			for _, lhs := range node.Lhs {
				if sel, ok := lhs.(*ast.SelectorExpr); ok {
					field := exprToString(sel.X) + "." + sel.Sel.Name
					assignPos := fset.Position(node.Pos())
					graph.Assignments = append(graph.Assignments, AtreeAssign{
						Field:      field,
						InFunction: scope,
						File:       file,
						Line:       assignPos.Line,
					})
				}
			}
		}
		return true
	})
}

// receiverTypeName extracts the type name from a receiver field, unwrapping pointers.
func receiverTypeName(field *ast.Field) string {
	return extractTypeName(field.Type)
}

func extractTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return extractTypeName(t.X)
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t)
	case *ast.IndexExpr:
		return extractTypeName(t.X)
	}
	return ""
}

// isPointerReceiver returns true if the receiver is a pointer type.
func isPointerReceiver(field *ast.Field) bool {
	_, ok := field.Type.(*ast.StarExpr)
	return ok
}

// exprToString converts an AST expression to a string representation.
func exprToString(expr ast.Expr) string {
	if expr == nil {
		return ""
	}
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.StarExpr:
		return "*" + exprToString(e.X)
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToString(e.Elt)
	case *ast.MapType:
		return "map[" + exprToString(e.Key) + "]" + exprToString(e.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.FuncType:
		return "func"
	case *ast.ChanType:
		return "chan " + exprToString(e.Value)
	case *ast.Ellipsis:
		return "..." + exprToString(e.Elt)
	case *ast.IndexExpr:
		return exprToString(e.X) + "[" + exprToString(e.Index) + "]"
	case *ast.ParenExpr:
		return exprToString(e.X)
	case *ast.TypeAssertExpr:
		if e.Type == nil {
			return exprToString(e.X) + ".(type)"
		}
		return exprToString(e.X) + ".(" + exprToString(e.Type) + ")"
	}
	return ""
}

// callExprInfo extracts the function name and receiver from a call expression.
func callExprInfo(call *ast.CallExpr) (name, receiver string) {
	switch fn := call.Fun.(type) {
	case *ast.Ident:
		name = fn.Name
	case *ast.SelectorExpr:
		name = fn.Sel.Name
		receiver = exprToString(fn.X)
	}
	return
}

// ── Detector functions ──

// detectModelsFromGraph finds Bubble Tea model structs.
// A model has both Update (1 param, 2 returns) and View (0 extra params, 1 return).
func detectModelsFromGraph(graph *AtreeSemanticGraph) ([]ir.ModelDef, error) {
	// Group method defs by receiver
	type methodSig struct {
		params    int
		returns   int
		bodyStart int
		bodyEnd   int
		file      string
		line      int
	}

	receiverMethods := make(map[string]map[string]methodSig)
	structFiles := make(map[string]string) // structName → file
	structLines := make(map[string]int)    // structName → line

	for _, fileSem := range graph.Files {
		for _, def := range fileSem.Defs {
			switch def.Kind {
			case "DefinitionStruct":
				structFiles[def.Name] = def.File
				structLines[def.Name] = def.Line
			case "DefinitionMethod":
				if def.Receiver == "" {
					continue
				}
				if receiverMethods[def.Receiver] == nil {
					receiverMethods[def.Receiver] = make(map[string]methodSig)
				}
				receiverMethods[def.Receiver][def.Name] = methodSig{
					params:    len(def.Params),
					returns:   len(def.Returns),
					bodyStart: def.BodyStart,
					bodyEnd:   def.BodyEnd,
					file:      def.File,
					line:      def.Line,
				}
			}
		}
	}

	// Find structs that implement the tea.Model interface:
	//   Update(tea.Msg) (tea.Model, tea.Cmd)  → params=1, returns=2
	//   View() string                          → params=0, returns=1
	var models []ir.ModelDef
	modelNames := make(map[string]bool)

	for structName := range structFiles {
		methods, ok := receiverMethods[structName]
		if !ok {
			continue
		}
		updateM, hasUpdate := methods["Update"]
		viewM, hasView := methods["View"]

		if !hasUpdate || !hasView {
			continue
		}
		// Update should have 1 param (tea.Msg) and 2 returns (Model, Cmd)
		if updateM.params < 1 || updateM.returns < 2 {
			continue
		}
		// View should have 0 params and 1 return (string)
		if viewM.returns < 1 {
			continue
		}

		impls := []string{"Update", "View"}
		if _, hasInit := methods["Init"]; hasInit {
			impls = append([]string{"Init"}, impls...)
		}

		model := ir.ModelDef{
			Name:             structName,
			File:             structFiles[structName],
			Line:             structLines[structName],
			Implements:       impls,
			UpdateComplexity: updateM.bodyEnd - updateM.bodyStart,
			ViewComplexity:   viewM.bodyEnd - viewM.bodyStart,
			Role:             "standalone",
		}
		models = append(models, model)
		modelNames[structName] = true
	}

	// Determine parent/child roles by looking for model-type fields in other models
	// We do a simple string search: if a struct's fields contain another model's name
	// Build a map of struct field types
	structFieldTypes := make(map[string][]string) // structName → list of field type names
	for _, fileSem := range graph.Files {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fileSem.Path, nil, 0)
		if err != nil {
			continue
		}
		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range genDecl.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				var fieldTypes []string
				for _, field := range st.Fields.List {
					fieldTypes = append(fieldTypes, exprToString(field.Type))
				}
				structFieldTypes[ts.Name.Name] = fieldTypes
			}
		}
	}

	// Mark models that appear as fields in other models as children
	childModels := make(map[string]bool)
	for parentName, fieldTypes := range structFieldTypes {
		if !modelNames[parentName] {
			continue
		}
		for _, ft := range fieldTypes {
			// Strip pointer prefix
			cleanFT := strings.TrimPrefix(ft, "*")
			if modelNames[cleanFT] && cleanFT != parentName {
				childModels[cleanFT] = true
			}
		}
	}

	// Apply roles
	for i := range models {
		if childModels[models[i].Name] {
			models[i].Role = "child"
		} else if len(models) > 1 {
			// If there are children, the non-child with parent field is root
			hasChildField := false
			for _, ft := range structFieldTypes[models[i].Name] {
				cleanFT := strings.TrimPrefix(ft, "*")
				if childModels[cleanFT] {
					hasChildField = true
					break
				}
			}
			if hasChildField {
				models[i].Role = "root"
			}
		}
	}

	return models, nil
}

// detectMessageFlowFromGraph finds custom message types from type switch cases in Update().
func detectMessageFlowFromGraph(graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.MessageDef, []ir.MessageFlowEdge, error) {
	modelNames := make(map[string]bool)
	for _, m := range models {
		modelNames[m.Name] = true
	}

	msgMap := make(map[string]*ir.MessageDef) // type name → def
	var edges []ir.MessageFlowEdge

	for _, fileSem := range graph.Files {
		for _, sw := range fileSem.SwitchCases {
			// Only interested in type switches in Update() methods of known models
			parts := strings.SplitN(sw.Scope, ".", 2)
			if len(parts) != 2 || parts[1] != "Update" {
				continue
			}
			receiverName := parts[0]

			for _, c := range sw.Cases {
				typeName := c.Type
				// Skip built-ins and tea.* types
				if typeName == "" || strings.HasPrefix(typeName, "tea.") || strings.HasPrefix(typeName, "bubbletea.") {
					continue
				}
				// Skip default/nil cases
				if typeName == "nil" {
					continue
				}

				// Strip package prefix for local types
				localName := typeName
				if idx := strings.LastIndex(typeName, "."); idx >= 0 {
					localName = typeName[idx+1:]
				}

				if msgMap[localName] == nil {
					msgMap[localName] = &ir.MessageDef{
						Name: localName,
						File: c.File,
						Line: c.Line,
					}
				}
				msgMap[localName].Handlers = append(msgMap[localName].Handlers, ir.Handler{
					Model:   receiverName,
					Pattern: "type_switch",
				})

				// Add flow edge: message → model
				edges = append(edges, ir.MessageFlowEdge{
					From: localName,
					To:   receiverName,
					Via:  "type_switch",
				})
			}

			// Also check for child model Update() forwarding
			if modelNames[receiverName] {
				// look for calls to childModel.Update in this scope
				for _, call := range fileSem.FunctionCalls {
					if call.Scope == sw.Scope && call.Name == "Update" && call.Recv != "" {
						edges = append(edges, ir.MessageFlowEdge{
							From: receiverName,
							To:   call.Recv,
							Via:  "delegate",
						})
					}
				}
			}
		}
	}

	var messages []ir.MessageDef
	for _, m := range msgMap {
		messages = append(messages, *m)
	}

	return messages, edges, nil
}

// detectComponentsFromGraph finds Charm component usages in model structs.
func detectComponentsFromGraph(graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.ComponentBinding, error) {
	modelNames := make(map[string]bool)
	for _, m := range models {
		modelNames[m.Name] = true
	}

	// Known component short names (local import names → package path)
	knownComponents := map[string]string{
		"viewport":   "github.com/charmbracelet/bubbles/viewport",
		"list":       "github.com/charmbracelet/bubbles/list",
		"textarea":   "github.com/charmbracelet/bubbles/textarea",
		"textinput":  "github.com/charmbracelet/bubbles/textinput",
		"spinner":    "github.com/charmbracelet/bubbles/spinner",
		"progress":   "github.com/charmbracelet/bubbles/progress",
		"help":       "github.com/charmbracelet/bubbles/help",
		"table":      "github.com/charmbracelet/bubbles/table",
		"filepicker": "github.com/charmbracelet/bubbles/filepicker",
	}

	var bindings []ir.ComponentBinding
	seen := make(map[string]bool) // dedup model+component pairs

	for _, fileSem := range graph.Files {
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fileSem.Path, nil, 0)
		if err != nil {
			continue
		}
		for _, decl := range f.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range genDecl.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok || !modelNames[ts.Name.Name] {
					continue
				}
				st, ok := ts.Type.(*ast.StructType)
				if !ok {
					continue
				}
				for _, field := range st.Fields.List {
					typeStr := exprToString(field.Type)
					cleanType := strings.TrimPrefix(typeStr, "*")
					// Check selector form: "viewport.Model", "list.Model" etc.
					parts := strings.SplitN(cleanType, ".", 2)
					if len(parts) != 2 || parts[1] != "Model" {
						continue
					}
					pkgName := parts[0]
					pkgPath, known := knownComponents[pkgName]
					if !known {
						continue
					}
					key := ts.Name.Name + "+" + pkgName
					if seen[key] {
						continue
					}
					seen[key] = true
					bindings = append(bindings, ir.ComponentBinding{
						Model:     ts.Name.Name,
						Component: pkgName + ".Model",
						Package:   pkgPath,
					})
				}
			}
		}
	}

	return bindings, nil
}

// detectAsyncBoundariesFromGraph finds async patterns (tea.Cmd, goroutines, tickers).
func detectAsyncBoundariesFromGraph(graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.AsyncBoundary, error) {
	modelNames := make(map[string]bool)
	for _, m := range models {
		modelNames[m.Name] = true
	}

	var boundaries []ir.AsyncBoundary

	for _, fileSem := range graph.Files {
		for _, call := range fileSem.FunctionCalls {
			// Determine the model from scope
			parts := strings.SplitN(call.Scope, ".", 2)
			if len(parts) < 2 {
				continue
			}
			modelName := parts[0]
			if !modelNames[modelName] {
				continue
			}

			var boundary *ir.AsyncBoundary

			switch {
			case call.Recv == "tea" && call.Name == "Batch":
				boundary = &ir.AsyncBoundary{
					Model: modelName,
					Line:  call.Line,
					Kind:  "sequence",
				}
			case call.Recv == "tea" && call.Name == "Sequence":
				boundary = &ir.AsyncBoundary{
					Model: modelName,
					Line:  call.Line,
					Kind:  "sequence",
				}
			case call.Recv == "tea" && call.Name == "Tick":
				boundary = &ir.AsyncBoundary{
					Model: modelName,
					Line:  call.Line,
					Kind:  "ticker",
				}
			case call.Recv == "tea" && call.Name == "Every":
				boundary = &ir.AsyncBoundary{
					Model: modelName,
					Line:  call.Line,
					Kind:  "ticker",
				}
			case call.Name == "Cmd" && call.Recv == "tea":
				boundary = &ir.AsyncBoundary{
					Model: modelName,
					Line:  call.Line,
					Kind:  "command",
				}
			}

			if boundary != nil {
				boundaries = append(boundaries, *boundary)
			}
		}

		// Detect goroutine spawns by looking for go statements in function bodies
		// We re-parse here to find GoStmt nodes
		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, fileSem.Path, nil, 0)
		if err != nil {
			continue
		}
		for _, decl := range f.Decls {
			funcDecl, ok := decl.(*ast.FuncDecl)
			if !ok || funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
				continue
			}
			recvName := receiverTypeName(funcDecl.Recv.List[0])
			if !modelNames[recvName] {
				continue
			}
			if funcDecl.Body == nil {
				continue
			}
			ast.Inspect(funcDecl.Body, func(n ast.Node) bool {
				goStmt, ok := n.(*ast.GoStmt)
				if !ok {
					return true
				}
				goPos := fset.Position(goStmt.Pos())
				boundaries = append(boundaries, ir.AsyncBoundary{
					Model:            recvName,
					Line:             goPos.Line,
					Kind:             "goroutine",
					CreatesGoroutine: true,
				})
				return true
			})
		}
	}

	return boundaries, nil
}

// analyzeStateMutationFromGraph tracks field mutations in Update() methods.
func analyzeStateMutationFromGraph(graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.StateOwnership, []ir.MutationPath, error) {
	modelNames := make(map[string]bool)
	for _, m := range models {
		modelNames[m.Name] = true
	}

	ownershipMap := make(map[string]*ir.StateOwnership) // field → ownership
	var mutations []ir.MutationPath

	for _, assign := range graph.Assignments {
		// Parse scope to find owning model
		parts := strings.SplitN(assign.InFunction, ".", 2)
		if len(parts) < 2 {
			continue
		}
		modelName := parts[0]
		if !modelNames[modelName] {
			continue
		}

		// field is "m.FieldName" or "someVar.FieldName"
		fieldParts := strings.SplitN(assign.Field, ".", 2)
		if len(fieldParts) < 2 {
			continue
		}
		fieldName := fieldParts[1]

		key := modelName + "." + fieldName
		if ownershipMap[key] == nil {
			ownershipMap[key] = &ir.StateOwnership{
				Field: fieldName,
				Owner: modelName,
			}
		}

		mutations = append(mutations, ir.MutationPath{
			Field:   fieldName,
			Source:  modelName,
			Trigger: parts[1], // method name
		})
	}

	var ownership []ir.StateOwnership
	for _, o := range ownershipMap {
		ownership = append(ownership, *o)
	}

	return ownership, mutations, nil
}

// analyzeViewsFromGraph analyzes View() methods for rendering patterns.
func analyzeViewsFromGraph(graph *AtreeSemanticGraph, models []ir.ModelDef) ([]ir.RenderNode, []ir.LayoutNode, error) {
	modelNames := make(map[string]bool)
	for _, m := range models {
		modelNames[m.Name] = true
	}

	var renderNodes []ir.RenderNode

	for _, fileSem := range graph.Files {
		modelViewCalls := make(map[string][]AtreeFuncCall) // modelName → calls in View()

		for _, call := range fileSem.FunctionCalls {
			parts := strings.SplitN(call.Scope, ".", 2)
			if len(parts) != 2 || parts[1] != "View" {
				continue
			}
			modelName := parts[0]
			if !modelNames[modelName] {
				continue
			}
			modelViewCalls[modelName] = append(modelViewCalls[modelName], call)
		}

		for modelName, calls := range modelViewCalls {
			node := ir.RenderNode{
				Model:      modelName,
				OutputType: "string",
			}
			for _, call := range calls {
				if call.Recv == "lipgloss" {
					node.UsesLipgloss = true
				}
				if call.Name == "View" && call.Recv != "" {
					// Calls a sub-component's View
					node.OutputType = "composite"
				}
				if strings.Contains(call.Recv, "viewport") || call.Name == "View" && strings.Contains(call.Scope, "viewport") {
					node.UsesViewport = true
				}
			}
			renderNodes = append(renderNodes, node)
		}
	}

	// Build simple layout tree from render nodes
	var layoutNodes []ir.LayoutNode
	if len(renderNodes) > 0 {
		root := ir.LayoutNode{
			Type: "vertical",
		}
		for range renderNodes {
			root.Children = append(root.Children, ir.LayoutNode{Type: "stack"})
		}
		layoutNodes = append(layoutNodes, root)
	}

	return renderNodes, layoutNodes, nil
}
