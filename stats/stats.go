package stats

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

type Gathere interface {
	Gather(path string) (Stats, error)
}
