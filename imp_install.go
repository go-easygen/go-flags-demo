////////////////////////////////////////////////////////////////////////////
// Program: redo
// Purpose: global option redo
// Authors: Myself <me@mine.org> (c) 2022-2023, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	//"github.com/mkideal/cli/clis"
	"github.com/go-easygen/go-flags/clis"
)

// *** Sub-command: install ***
// Exec implements the business logic of command `install`
func (x *InstallCommand) Exec(args []string) error {
	clis.WarnOn("install::Exec", fmt.Errorf("sample warning message"))
	// or,
	clis.AbortOn("install::Exec", fmt.Errorf("critical error message"))
	return nil
}
