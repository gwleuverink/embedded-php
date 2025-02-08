package main

import (
    "embed"
    "fmt"
    "os"
    "os/exec"
)

// Embed both PHP binary and app.php
//go:embed php_binary/php
var phpBinaryFS embed.FS

//go:embed app.php
var appScriptFS embed.FS

func main() {
    // Extract PHP binary
    phpBinary, err := phpBinaryFS.ReadFile("php_binary/php")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading embedded PHP binary: %v\n", err)
        os.Exit(1)
    }

    phpTemp, err := os.CreateTemp("", "php-binary")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating PHP binary temp file: %v\n", err)
        os.Exit(1)
    }
    defer os.Remove(phpTemp.Name())

    if err := os.WriteFile(phpTemp.Name(), phpBinary, 0644); err != nil {
        fmt.Fprintf(os.Stderr, "Error writing PHP binary: %v\n", err)
        os.Exit(1)
    }

    if err := os.Chmod(phpTemp.Name(), 0755); err != nil {
        fmt.Fprintf(os.Stderr, "Error setting executable permission: %v\n", err)
        os.Exit(1)
    }

    // Extract app.php script
    appScript, err := appScriptFS.ReadFile("app.php")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading embedded PHP script: %v\n", err)
        os.Exit(1)
    }

    appTemp, err := os.CreateTemp("", "app-script")
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error creating PHP script temp file: %v\n", err)
        os.Exit(1)
    }
    defer os.Remove(appTemp.Name())

    if err := os.WriteFile(appTemp.Name(), appScript, 0644); err != nil {
        fmt.Fprintf(os.Stderr, "Error writing PHP script: %v\n", err)
        os.Exit(1)
    }

    // Execute the PHP binary with the embedded app.php script
    cmd := exec.Command(phpTemp.Name(), appTemp.Name())

    // Forward any arguments passed to the Go binary to PHP
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
