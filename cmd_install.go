////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: install ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The InstallCommand type defines all the configurable options from cli.
type InstallCommand struct {
	Dir    string `short:"d" description:"source code root dir" default:"./"`
	Suffix string `long:"suffix" description:"source file suffix" default:".go,.c,.s"`
}

var installCommand InstallCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("install",
		"Install the network application",
		"The add command adds a file to the repository. Use -a to add all files",
		&installCommand)
}

func (x *InstallCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Install the network application\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022-2023, Myself <me@mine.org>\n\n")
	clis.Setup("redo::install", opts.Verbose)
	clis.Verbose(1, "Doing Install, with %+v, %+v", opts, args)
	// fmt.Println(x.Dir, x.Suffix)
	return x.Exec(args)
}

// // Exec implements the business logic of command `install`
// func (x *InstallCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("install::Exec", err)
// 	// or,
// 	// clis.AbortOn("install::Exec", err)
// 	return nil
// }
