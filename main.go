package main

import (
	"github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	"os"
)

type commonOptions struct {
	Verbose      bool   `short:"v" long:"verbose" description:"Show current file"`
	DatabasePath string `short:"d" long:"database" required:"true" description:"path to result file"`
}

var options struct {
	Collect           collectCommand `command:"collect" description:"gather data to database"`
	PrintCmd          printCommand   `command:"print" description:"print content of database"`
	FindDuplicatesCmd dupsCommand    `command:"dups" description:"find duplicates in database basing on file size"`
}

var flagParser = flags.NewParser(&options, flags.Default)

func main() {
	_, err := flagParser.Parse()
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func setupLogs(verbose bool) {
	if verbose {
		lgr.Setup(lgr.Debug)
	}
}
