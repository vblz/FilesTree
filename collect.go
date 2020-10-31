package main

import (
	"github.com/github.com/vblz/FilesTree/store"
	"github.com/github.com/vblz/FilesTree/tree"
	"github.com/go-pkgz/lgr"
	"os"
)

type collectCommand struct {
	Args struct {
		Path string
	} `positional-args:"yes"`

	commonOptions
}

func (c *collectCommand) Execute(args []string) error {
	if len(args) > 0 {
		flagParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	if c.Args.Path == "" {
		c.Args.Path = "."
	}

	setupLogs(c.Verbose)

	t := tree.NewTraverser(func(path string, size int64) {
		lgr.Printf("[DEBUG] %s %d", path, size)
	})

	result, err := t.Traverse(c.Args.Path)
	if err != nil {
		lgr.Fatalf("traversing error: %s", err)
		os.Exit(2)
	}

	err = store.Write(c.DatabasePath, result)

	lgr.Printf("[DEBUG] finished")

	return nil
}
