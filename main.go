package main

import (
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func main() {

    dirPtr := flag.String("dir", ".", "Directory to trace (default: current)")
    includeFiles := flag.Bool("files", true, "Include files in the output (default: true)")

    flag.Parse()

    if _, err := os.Stat(*dirPtr); os.IsNotExist(err) {
        fmt.Printf("The directory '%s' does not exist.\n", *dirPtr)
        return
    }

    tree, err := generateTree(*dirPtr, *includeFiles, "")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(tree)
}

func generateTree(path string, includeFiles bool, prefix string) (string, error) {
    var result strings.Builder

    entries, err := os.ReadDir(path)
    if err != nil {
        return "", err
    }

    for i, entry := range entries {

        connector := "├── "
        if i == len(entries)-1 {
            connector = "└── "
        }

        if entry.IsDir() {

            result.WriteString(prefix + connector + entry.Name() + "\n")

            var subPrefix string
            if connector == "├── " {
                subPrefix = prefix + "│   "
            } else {
                subPrefix = prefix + "    "
            }

            subtree, err := generateTree(filepath.Join(path, entry.Name()), includeFiles, subPrefix)
            if err != nil {
                return "", err
            }
            result.WriteString(subtree)
        } else if includeFiles {

            result.WriteString(prefix + connector + entry.Name() + "\n")
        }
    }

    return result.String(), nil
}