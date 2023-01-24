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

// *** Sub-command: publish ***

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The PublishCommand type defines all the configurable options from cli.
type PublishCommand struct {
	Dir    string   `short:"d" long:"dir" description:"publish dir" required:"true"`
	Suffix []string `short:"s" long:"suffix" description:"source file suffix for publish" choice:".go" choice:".c" choice:".h"`
	Out    string   `short:"o" long:"out" description:"output filename"`

	// Example of positional arguments
	Args struct {
		ID   string
		Num  int
		Rest []string
	} `positional-args:"yes" required:"yes"`
}

var publishCommand PublishCommand

////////////////////////////////////////////////////////////////////////////
// Function definitions

func init() {
	gfParser.AddCommand("publish",
		"Publish the network application",
		"Publish the built network application to central repo",
		&publishCommand)
}

func (x *PublishCommand) Execute(args []string) error {
	fmt.Fprintf(os.Stderr, "Publish the network application\n")
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2022-2023, Myself <me@mine.org>\n\n")
	clis.Setup("redo::publish", opts.Verbose)
	clis.Verbose(1, "Doing Publish, with %+v, %+v", opts, args)
	// fmt.Println(x.Dir, x.Suffix, x.Out, x.Args)
	return x.Exec(args)
}

// // Exec implements the business logic of command `publish`
// func (x *PublishCommand) Exec(args []string) error {
// 	// err := ...
// 	// clis.WarnOn("publish::Exec", err)
// 	// or,
// 	// clis.AbortOn("publish::Exec", err)
// 	return nil
// }
