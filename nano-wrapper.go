package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nano <file>")
		os.Exit(1)
	}

	// Get the first argument, which is the file path
	filePath := os.Args[1]

	// Check if the file path is a Windows-style path (e.g., C:\text.txt)
	isWindowsPath := regexp.MustCompile(`^[A-Za-z]:\\`).MatchString(filePath)

	if isWindowsPath {
		// Convert Windows path to WSL path by removing ':' and converting '\' to '/'
		wslPath := strings.ToLower(strings.Replace(filePath, `\`, `/`, -1))
		wslPath = "/mnt/" + strings.ToLower(wslPath[:1]) + wslPath[2:]

		// Run nano with WSL for a Windows path, ensuring it runs interactively
		cmd := exec.Command("wsl", "bash", "-i", "-c", "nano "+wslPath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running nano:", err)
			os.Exit(1)
		}
	} else {
		// Run nano directly with the given path (Unix-style or relative path), ensuring it's interactive
		cmd := exec.Command("bash", "-i", "-c", "nano "+filePath)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running nano:", err)
			os.Exit(1)
		}
	}
}
