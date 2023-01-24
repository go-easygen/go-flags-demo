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

// *** Sub-command: build ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The BuildCommand type defines all the configurable options from cli.
type BuildCommand struct {
	Dir string `long:"dir" description:"source code root dir" default:"./"`
}

var buildCommand BuildCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("build",
		"Build the network application",
		"Usage:\n  redo build [Options] Arch(i386|amd64)",
		&buildCommand)
}

func (x *BuildCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Build the network application\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022-2023, Myself <me@mine.org>\n\n")
	clis.Setup("redo::build", opts.Verbose)
	clis.Verbose(1, "Doing Build, with %+v, %+v", opts, args)
	// fmt.Println(x.Dir)
	return x.Exec(args)
}

// // Exec implements the business logic of command `build`
// func (x *BuildCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("build::Exec", err)
// 	// or,
// 	// clis.AbortOn("build::Exec", err)
// 	return nil
// }
