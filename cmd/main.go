package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <path>", os.Args[0])
	}

	PrintModuleInfo(os.Args[1])
	PrintFileInfo(os.Args[1])
}

// PrintModuleInfo receives path to a directory with a go module, parses go.mod file in it
// and prints the module name and count of required direct modules.
func PrintModuleInfo(path string) {
	f, err := os.Open(path + "/go.mod")
	if err != nil {
		log.Fatalf("failed to open go.mod file: %v", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("failed to read go.mod file: %v", err)
	}

	mf, err := modfile.Parse(path, data, nil)
	if err != nil {
		log.Fatalf("failed to parse go.mod file: %v", err)
	}

	countOfDirectModules := 0
	for _, req := range mf.Require {
		if req.Indirect {
			continue
		}
		countOfDirectModules++
	}

	log.Printf("Module: %s\nCount of dependencies: %d", mf.Module.Mod.Path, countOfDirectModules)
}

// PrintFileInfo receives path to a directory with a go module and prints count of go files int it and in its subdirectories recursively.
func PrintFileInfo(path string) {
	countOfGoFiles := 0
	err := filepath.WalkDir(path, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Fatalf("failed to walk path %s: %v", path, err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".go" {
			countOfGoFiles++
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to walk path %s: %v", path, err)
	}

	log.Printf("Count of go files: %d", countOfGoFiles)
}
