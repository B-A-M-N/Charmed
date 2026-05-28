// Package graph provides the normalized project graph — the internal semantic
// substrate that all analysis stages consume. No ATree formats leak past this layer.
package graph

// ProjectGraph is the normalized representation of a Go project's code graph.
// Built during Stage 3 (Graph Normalization) from ATree MCP responses.
type ProjectGraph struct {
	Files       map[string]*FileInfo   // keyed by relative path
	Symbols     map[string]*Symbol     // keyed by qualified name
	Structs     []*StructInfo
	Methods     map[string][]*MethodInfo // receiver type name → methods
	Interfaces  []*InterfaceInfo
	Calls       []*CallEdge
	Imports     map[string]*ImportMap  // file path → import map
	Heritage    []*HeritageEdge
	Assignments []*Assignment
	Scopes      map[int]*ScopeInfo     // scope ID → scope
	Unresolved  []UnresolvedRef
}

// FileInfo describes a single source file.
type FileInfo struct {
	Path        string
	Language    string
	HasCharmDep bool
}

// Symbol is a resolved code entity.
type Symbol struct {
	Name       string
	Qualified  string // e.g. "module/path.Type.Method" or "Type.Method"
	Tag        string // "struct" | "method" | "function" | "interface" | "const" | "var" | "property"
	File       string
	Line       int
	Col        int
	IsExported bool
}

// StructInfo is a struct definition with its fields, methods, and embeds.
type StructInfo struct {
	Symbol
	Fields  []*FieldInfo
	Methods []*MethodInfo
	Embeds  []string // embedded type qualified names
	File    string
	Line    int
}

// FieldInfo is a struct field.
type FieldInfo struct {
	Name      string
	Type      string // fully qualified, import aliases expanded
	StructTag string
	Line      int
	IsEmbed   bool
}

// MethodInfo is a method definition with full signature.
type MethodInfo struct {
	Symbol
	Receiver      string       // receiver type name
	ReceiverIsPtr bool
	Params        []ParamInfo  // name + type
	Returns       []string     // return types
	File          string
	Line          int
	BodyStart     int          // scope line_start
	BodyEnd       int          // scope line_end
}

// ParamInfo is a method/function parameter.
type ParamInfo struct {
	Name string
	Type string
}

// InterfaceInfo is an interface definition.
type InterfaceInfo struct {
	Symbol
	Methods []*MethodInfo
	File    string
	Line    int
}

// CallEdge is a resolved call relationship.
type CallEdge struct {
	Caller     string
	Callee     string
	Receiver   string  // method receiver, if method call
	File       string
	Line       int
	Col        int
	Confidence float64
	Resolved   bool
}

// ImportMap resolves local import names to full paths per file.
type ImportMap struct {
	Imports map[string]*ImportEntry // local name → entry
}

// ImportEntry is a single import.
type ImportEntry struct {
	LocalName string
	FullPath  string
	IsAlias   bool
}

// HeritageEdge is a type relationship (extends, implements, embeds).
type HeritageEdge struct {
	Child  string
	Parent string
	Kind   string // "extends" | "implements" | "embeds"
	File   string
	Line   int
}

// Assignment tracks field mutations.
type Assignment struct {
	Field      string
	Source     string // what's assigned
	InFunction string // which method/function
	File       string
	Line       int
}

// ScopeInfo is a lexical scope.
type ScopeInfo struct {
	ID         int
	ParentID   int
	OwnerSym   string // qualified symbol name
	Kind       string // "method" | "function" | "struct" | "block"
	File       string
	LineStart  int
	LineEnd    int
}

// UnresolvedRef tracks symbols that couldn't be resolved.
type UnresolvedRef struct {
	Name   string
	File   string
	Line   int
	Reason string
}

// CharmComponent is a known Charm ecosystem component type.
var CharmComponentTypes = map[string]struct{}{
	"github.com/charmbracelet/bubbles/viewport.Model":  {},
	"github.com/charmbracelet/bubbles/list.Model":       {},
	"github.com/charmbracelet/bubbles/textarea.Model":   {},
	"github.com/charmbracelet/bubbles/textinput.Model":  {},
	"github.com/charmbracelet/bubbles/spinner.Model":    {},
	"github.com/charmbracelet/bubbles/progress.Model":   {},
	"github.com/charmbracelet/bubbles/help.Model":       {},
	"github.com/charmbracelet/bubbles/table.Model":      {},
	"github.com/charmbracelet/bubbles/filepicker.Model": {},
	"github.com/charmbracelet/bubbles/key.Binding":      {},
	"github.com/charmbracelet/lipgloss.Style":            {},
}

// CharmImportPrefixes are import paths that indicate Charm ecosystem usage.
var CharmImportPrefixes = []string{
	"github.com/charmbracelet/bubbletea",
	"github.com/charmbracelet/bubbles",
	"github.com/charmbracelet/lipgloss",
	"github.com/charmbracelet/huh",
	"github.com/charmbracelet/glamour",
	"github.com/charmbracelet/glow",
	"github.com/charmbracelet/wish",
}
