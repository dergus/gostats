package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/samber/lo"
	"golang.org/x/mod/modfile"
)

type ModuleStats struct {
	Name                 string
	DirecteDependencies  int
	IndirectDependencies int
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	if len(os.Args) != 2 {
		log.Fatalf("usage: gostats <path>\n")
	}

	path := os.Args[1]

	modFilePath := filepath.Join(path, "go.mod")
	modFileData, err := os.ReadFile(modFilePath)

	if err != nil {
		if err == os.ErrNotExist {
			fatalErr("go.mod file not found")
		}

		fatalErr("can't open go.mod faile: %v", err)
	}

	modFile, err := modfile.ParseLax(modFilePath, modFileData, nil)

	if err != nil {
		fatalErr("cant' parse go.mod file: %s", err)
	}

	dirDeps := lo.Reduce(modFile.Require, func(agg int, item *modfile.Require, _ int) int {
		if !item.Indirect {
			return agg + 1
		}

		return agg
	}, 0)

	moduleStats := ModuleStats{
		Name:                 modFile.Module.Mod.String(),
		DirecteDependencies:  dirDeps,
		IndirectDependencies: len(modFile.Require) - dirDeps,
	}

	fmt.Printf("%+v\n", moduleStats)
}

func fatalErr(msg string, args ...any) {
	fmt.Printf(msg, args...)
	os.Exit(1)
}

type Stats struct {
	Module struct {
		Name                      string
		CountOfDirectDependencies int
		CountOfAllDependencies    int
		CountOfPackages           int
		CountOfFiles              int
		CountOfStructs            int
		CountOfGlobalVariables    int
		CountOfConsts             int
		CountOfLines              int
		CountOfCodeLines          int
		CountOfFunctions          int
		CountOfMethods            int
	}

	Package struct {
		NameChars       CommonStats
		Files           CommonStats
		Structs         CommonStats
		GlobalVariables CommonStats
		Consts          CommonStats
		Functions       CommonStats
		Lines           CommonStats
		CodeLines       CommonStats
		Methods         CommonStats
		Imports         CommonStats
		PublicTypes     CommonStats
		PublicFunctions CommonStats
		PublicMethods   CommonStats
	}

	File struct {
		NameChars       CommonStats
		Files           CommonStats
		Structs         CommonStats
		GlobalVariables CommonStats
		Consts          CommonStats
		Functions       CommonStats
		Lines           CommonStats
		CodeLines       CommonStats
		Methods         CommonStats
		Imports         CommonStats
		PublicTypes     CommonStats
		PublicFunctions CommonStats
		PublicMethods   CommonStats
	}
}

type CommonStats struct {
	CountMax    int
	CountMin    int
	CountAvg    int
	CountMedian int
}
