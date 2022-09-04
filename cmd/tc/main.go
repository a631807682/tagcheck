package main

import (
	"log"
	"os"

	"github.com/a631807682/tagcheck/tagcheck"
	"honnef.co/go/tools/lintcmd"
	"honnef.co/go/tools/lintcmd/version"
	"honnef.co/go/tools/quickfix"
	"honnef.co/go/tools/unused"
)

func main() {
	cmd := lintcmd.NewCommand("tc")
	cmd.SetVersion(version.Version, version.MachineVersion)

	fs := cmd.FlagSet()
	debug := fs.String("debug.unused-graph", "", "Write unused's object graph to `file`")
	qf := fs.Bool("debug.run-quickfix-analyzers", false, "Run quickfix analyzers")

	cmd.ParseFlags(os.Args[1:])

	cmd.AddAnalyzers(tagcheck.Analyzers...)

	if *qf {
		cmd.AddAnalyzers(quickfix.Analyzers...)
	}

	if *debug != "" {
		f, err := os.OpenFile(*debug, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			log.Fatal(err)
		}
		unused.Debug = f
	}

	cmd.Run()
}
