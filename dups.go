package main

import (
	"fmt"
	"github.com/github.com/vblz/FilesTree/store"
	"os"
	"strings"
)

type dupsCommand struct {
	commonOptions
}

func (d *dupsCommand) Execute(args []string) error {
	if len(args) > 0 {
		flagParser.WriteHelp(os.Stdout)
		os.Exit(1)
	}

	files, err := store.Read(d.DatabasePath)
	if err != nil {
		return fmt.Errorf("database open error: %w", err)
	}

	res := findDuplicates(files)
	if len(res) == 0 {
		fmt.Printf("no duplicates find")
		return nil
	}

	for _, dups := range res {
		for _, filePath := range dups {
			fmt.Println(filePath)
		}
		fmt.Println(strings.Repeat("-", 32))
	}

	return nil
}

func findDuplicates(files map[string]os.FileInfo) [][]string {
	sizes := make(map[int64][]string)

	for k, v := range files {
		size := v.Size()
		if !v.IsDir() && size != 0 {
			l, ok := sizes[size]
			if !ok {
				l = make([]string, 0, 1)
			}
			l = append(l, k)
			sizes[size] = l
		}
	}

	result := make([][]string, 0)

	for _, v := range sizes {
		if len(v) > 1 {
			result = append(result, v)
		}
	}

	return result
}
