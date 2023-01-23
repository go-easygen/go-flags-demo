// redo - global option redo

// Redo global option via automatic code-gen

package main

////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

//go:generate sh go-flags-demo_cliGen.sh
//go:generate emd gen -in README.beg.e.md -in README.e.md -in README.end.e.md -out README.md

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags"
)

//////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "redo"
	version  = "0.1.0"
	date     = "2023-01-22"

	// opts store all the configurable options
	opts optsT
)

var parser = flags.NewParser(&opts, flags.Default)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	opts.Version = showVersion
	opts.Verbflg = func() {
		opts.Verbose++
	}

	if _, err := parser.Parse(); err != nil {
		fmt.Println()
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	fmt.Println()
	//DoRedo()
}

func showVersion() {
	fmt.Fprintf(os.Stderr, "redo - global option redo, version %s\n", version)
	fmt.Fprintf(os.Stderr, "Built on %s\n", date)
	fmt.Fprintf(os.Stderr, "Copyright (C) 2022-2023, Myself <me@mine.org>\n\n")
	fmt.Fprintf(os.Stderr, "Redo global option via automatic code-gen\n")
	os.Exit(0)
}
