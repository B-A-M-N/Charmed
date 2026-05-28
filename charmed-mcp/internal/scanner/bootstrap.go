// Package scanner implements the scan_project pipeline.
package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/B-A-M-N/charmed/charmed-mcp/internal/graph"
)

// BootstrapResult holds Stage 1 output.
type BootstrapResult struct {
	ProjectRoot   string
	GoModFound    bool
	CharmDeps     []string
	GoFiles       []string
	AtreeAvailable bool
	AtreeDBPath   string
	Warnings      []string
}

// Bootstrap validates the project and locates dependencies.
func Bootstrap(projectPath string, includeVendor bool, analyzeTests bool) (*BootstrapResult, error) {
	result := &BootstrapResult{}

	// Resolve to absolute path
	abs, err := filepath.Abs(projectPath)
	if err != nil {
		return nil, fmt.Errorf("resolve path: %w", err)
	}
	result.ProjectRoot = abs

	// Check go.mod
	goModPath := filepath.Join(abs, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		result.GoModFound = true
		// Parse charm deps from go.mod
		result.CharmDeps = parseCharmDeps(goModPath)
	} else {
		result.Warnings = append(result.Warnings, "no_go_mod: go.mod not found at "+abs)
	}

	// List Go files
	result.GoFiles = listGoFiles(abs, includeVendor, analyzeTests)

	// Locate ATree DB
	result.AtreeDBPath, result.AtreeAvailable = locateAtreeDB(abs)

	if !result.AtreeAvailable {
		result.Warnings = append(result.Warnings, "atree_unavailable: no ATree DB found")
	}

	if len(result.CharmDeps) == 0 && result.GoModFound {
		result.Warnings = append(result.Warnings, "no_charm_deps: no charmbracelet imports in go.mod")
	}

	return result, nil
}

// parseCharmDeps extracts charmbracelet/* imports from go.mod.
func parseCharmDeps(goModPath string) []string {
	data, err := os.ReadFile(goModPath)
	if err != nil {
		return nil
	}

	var deps []string
	seen := map[string]bool{}

	for _, line := range strings.Split(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		for _, prefix := range graph.CharmImportPrefixes {
			if strings.Contains(trimmed, prefix) && !seen[prefix] {
				seen[prefix] = true
				deps = append(deps, prefix)
			}
		}
	}
	return deps
}

// listGoFiles recursively finds .go files, excluding vendor/ and *_test.go.
func listGoFiles(root string, includeVendor bool, analyzeTests bool) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip unreadable paths
		}
		if info.IsDir() {
			if !includeVendor && info.Name() == "vendor" {
				return filepath.SkipDir
			}
			if info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		if !analyzeTests && strings.HasSuffix(path, "_test.go") {
			return nil
		}
		files = append(files, path)
		return nil
	})
	return files
}

// locateAtreeDB finds the ATree SQLite database.
func locateAtreeDB(projectRoot string) (string, bool) {
	// Check project-local path
	candidate := filepath.Join(projectRoot, ".atree", "index.sqlite")
	if _, err := os.Stat(candidate); err == nil {
		return candidate, true
	}
	// Check env var
	if envPath := os.Getenv("ATREE_DB_PATH"); envPath != "" {
		if _, err := os.Stat(envPath); err == nil {
			return envPath, true
		}
	}
	return "", false
}
