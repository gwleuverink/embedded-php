package main

import (
    "embed"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
)

//go:embed php_binary/php
var phpFiles embed.FS

const (
    entryScript = "./app.php"
)

func main() {
    // Get the absolute path to the entry script
    scriptPath, err := filepath.Abs(entryScript)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error resolving script path: %v\n", err)
        os.Exit(1)
    }

    // Get the embedded PHP binary
    phpBinary, err := phpFiles.ReadFile("php_binary/php")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading embedded PHP binary: %v\n", err)
        os.Exit(1)
    }

    // Create a temporary file for the PHP binary
    tmpFile, err := os.CreateTemp("", "php-binary")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating temp file: %v\n", err)
        os.Exit(1)
    }
    defer os.Remove(tmpFile.Name())

    // Write the binary and make it executable
    if err := os.WriteFile(tmpFile.Name(), phpBinary, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing PHP binary: %v\n", err)
		os.Exit(1)
	}

	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error setting executable permission: %v\n", err)
		os.Exit(1)
	}

    // Prepare the command
    cmd := exec.Command(tmpFile.Name(), scriptPath)

    // Forward any arguments passed to our binary to the PHP script
    if len(os.Args) > 1 {
        cmd.Args = append(cmd.Args, os.Args[1:]...)
    }

    // Connect standard IO
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin

    // Run the command
    err = cmd.Run()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error executing PHP: %v\n", err)
        if exitError, ok := err.(*exec.ExitError); ok {
            os.Exit(exitError.ExitCode())
        }
        os.Exit(1)
    }
}
