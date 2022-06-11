package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/mod/modfile"
)

var keyStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00ff00"))
var valStyle = lipgloss.NewStyle().Italic(true).Foreground(lipgloss.Color("#f0f"))

type overview struct {
	ModuleName        string
	CountDependencies int
}

func (o overview) Init() tea.Cmd {
	return nil
}

func (o overview) View() string {
	s := fmt.Sprintf("%s %s\n", keyStyle.Render("Module:"), valStyle.Render(o.ModuleName))
	s += fmt.Sprintf("%s %s\n", keyStyle.Render("Dependencies:"), valStyle.Render(fmt.Sprintf("%d", o.CountDependencies)))

	return s
}

func (o overview) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return o, tea.Quit
}

func initModel(path string) overview {
	mi, err := GetModuleInfo(path)
	if err != nil {
		log.Fatalf("failed to get module info: %v", err)
	}
	return overview{
		ModuleName:        mi.ModuleName,
		CountDependencies: mi.CountDependencies,
	}
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)
	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <path>", os.Args[0])
	}

	p := tea.NewProgram(initModel(os.Args[1]))
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type ModeuleInfo struct {
	ModuleName        string
	CountDependencies int
}

// GetModuleInfo receives path to a directory with a go module, parses go.mod file in it
// and returns the module name and count of required direct modules.
func GetModuleInfo(path string) (*ModeuleInfo, error) {
	f, err := os.Open(path + "/go.mod")
	if err != nil {
		return nil, fmt.Errorf("failed to open go.mod file: %w", err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read go.mod file: %w", err)
	}

	mf, err := modfile.Parse(path, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to parse go.mod file: %w", err)
	}

	countOfDirectModules := 0
	for _, req := range mf.Require {
		if req.Indirect {
			continue
		}
		countOfDirectModules++
	}

	mi := &ModeuleInfo{
		ModuleName:        mf.Module.Mod.String(),
		CountDependencies: countOfDirectModules,
	}

	return mi, nil
}
