// Package server — DeepWiki resource handler.
// Exposes the embedded DeepWiki reference docs as queryable MCP resources.
package server

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// DeepWikiIndex maps resource URIs to file paths.
type DeepWikiIndex struct {
	Entries []DeepWikiEntry `json:"entries"`
}

// DeepWikiEntry is a single DeepWiki document.
type DeepWikiEntry struct {
	URI      string `json:"uri"`       // e.g. "charmed://deepwiki/bubbletea/overview"
	Library  string `json:"library"`   // e.g. "bubbletea", "bubbles", "lipgloss"
	Topic    string `json:"topic"`     // e.g. "overview", "lifecycle", "viewport"
	Title    string `json:"title"`     // human-readable title
	FilePath string `json:"file_path"` // absolute path to the .md file
}

var (
	deepwiki     *DeepWikiIndex
	deepwikiOnce sync.Once
)

// loadDeepWikiIndex loads or builds the DeepWiki index.
func loadDeepWikiIndex() *DeepWikiIndex {
	deepwikiOnce.Do(func() {
		// Try loading pre-built index first
		candidates := []string{
			"knowledge/deepwiki-index.json",
			"../../charmed/knowledge/deepwiki-index.json",
		}
		for _, path := range candidates {
			if data, err := os.ReadFile(path); err == nil {
				idx := &DeepWikiIndex{}
				if err := json.Unmarshal(data, idx); err == nil && len(idx.Entries) > 0 {
					deepwiki = idx
					return
				}
			}
		}
		// Fallback: scan directory
		deepwiki = buildDeepWikiIndex()
	})
	return deepwiki
}

// DeepWikiBaseDirs returns candidate base directories for DeepWiki references.
func DeepWikiBaseDirs() []string {
	dirs := []string{
		"../../charmed/DeepWiki References",
		"../charmed/DeepWiki References",
		"charmed/DeepWiki References",
	}
	if envPath := os.Getenv("DEEPWIKI_PATH"); envPath != "" {
		dirs = append([]string{envPath}, dirs...)
	}
	return dirs
}

// buildDeepWikiIndex scans the DeepWiki directory and builds an index.
func buildDeepWikiIndex() *DeepWikiIndex {
	idx := &DeepWikiIndex{}

	for _, base := range DeepWikiBaseDirs() {
		if _, err := os.Stat(base); err != nil {
			continue
		}
		entries, _ := os.ReadDir(base)
		for _, libEntry := range entries {
			if !libEntry.IsDir() {
				continue
			}
			libName := libEntry.Name()
			libPath := filepath.Join(base, libName)

			filepath.Walk(libPath, func(path string, info os.FileInfo, err error) error {
				if err != nil || info.IsDir() || !strings.HasSuffix(path, ".md") {
					return nil
				}
				rel, _ := filepath.Rel(libPath, path)
				topic := strings.TrimSuffix(rel, ".md")
				topic = strings.ReplaceAll(topic, string(filepath.Separator), "/")

				entry := DeepWikiEntry{
					URI:      fmt.Sprintf("charmed://deepwiki/%s/%s", libName, topic),
					Library:  libName,
					Topic:    topic,
					Title:    deepwikiTitle(topic),
					FilePath: path,
				}
				idx.Entries = append(idx.Entries, entry)
				return nil
			})
		}
		if len(idx.Entries) > 0 {
			break
		}
	}

	return idx
}

// deepwikiTitle generates a human-readable title from a file path.
func deepwikiTitle(path string) string {
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, ".md")
	// Remove numeric prefix like "2.1-"
	parts := strings.SplitN(base, "-", 2)
	if len(parts) > 1 && strings.ContainsAny(parts[0], "0123456789.") {
		base = parts[1]
	}
	base = strings.ReplaceAll(base, "_", " ")
	base = strings.ReplaceAll(base, "-", " ")
	return strings.TrimSpace(base)
}

// resolveDeepWiki resolves a URI to file contents.
func resolveDeepWiki(uri string) ([]byte, error) {
	idx := loadDeepWikiIndex()
	for _, entry := range idx.Entries {
		if entry.URI == uri {
			return os.ReadFile(entry.FilePath)
		}
	}
	return nil, fmt.Errorf("deepwiki resource not found: %s", uri)
}

// listDeepWiki returns all available DeepWiki resources, optionally filtered by library.
func listDeepWiki(library string) []DeepWikiEntry {
	idx := loadDeepWikiIndex()
	if library == "" {
		return idx.Entries
	}
	var filtered []DeepWikiEntry
	for _, e := range idx.Entries {
		if strings.EqualFold(e.Library, library) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
