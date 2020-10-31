package main

import (
	"fmt"
	"github.com/github.com/vblz/FilesTree/store"
	"os"
	"sort"
)

type printCommand struct {
	commonOptions
}

func (p *printCommand) Execute(args []string) error {
	if len(args) > 0 {
		flagParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	fileInfos, err := store.Read(p.DatabasePath)
	if err != nil {
		return fmt.Errorf("database reading error: %w", err)
	}

	sortedKeys := make([]string, 0, len(fileInfos))
	for k := range fileInfos {
		sortedKeys = append(sortedKeys, k)
	}

	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		t := "[file]"
		if fileInfos[k].IsDir() {
			t = "[directory]"
		}
		fmt.Printf("%s %d %s\n", k, fileInfos[k].Size(), t)
	}

	return nil
}
