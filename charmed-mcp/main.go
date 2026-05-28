// charmed-mcp is the MCP server for the charmed plugin.
// It exposes Charm ecosystem TUI intelligence as MCP tools and resources.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/B-A-M-N/charmed/charmed-mcp/internal/scanner"
	"github.com/B-A-M-N/charmed/charmed-mcp/internal/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.SetOutput(os.Stderr)
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Try to create ATree client from environment or well-known paths
	var atreeClient scanner.AtreeClient
	if dbPath := findAtreeDB(); dbPath != "" {
		client, err := scanner.NewAtreeClient(dbPath)
		if err != nil {
			log.Printf("ATree client unavailable: %v", err)
		} else {
			atreeClient = client
			log.Printf("ATree connected: %s", dbPath)
		}
	}

	// Create scanner
	s, err := scanner.NewScanner(atreeClient)
	if err != nil {
		log.Fatalf("Failed to create scanner: %v", err)
	}

	// Create MCP server
	cs := server.New(s)

	// Serve on stdio
	log.Println("charmed-mcp starting on stdio...")
	if err := cs.ServeContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

// findAtreeDB locates the ATree SQLite database from the environment.
func findAtreeDB() string {
	if path := os.Getenv("ATREE_DB_PATH"); path != "" {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	return ""
}
