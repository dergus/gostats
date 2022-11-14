package gatherer

import (
	"io/fs"
	"path/filepath"
)

type Gatherer struct {
}

type CodeStats struct {
}

type Node struct {
	Kind     string
	Children []Node
	Stats    []Stat
}

func (g Gatherer) Gather(path string) (CodeStats, error) {
	filepath.WalkDir(func(path string, d fs.DirEntry, err error) error {

	})
}
