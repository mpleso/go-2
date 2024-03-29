// Copyright 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style license described in the
// LICENSE file.

package source

import (
	"fmt"
	"os"

	"github.com/platinasystems/go/command"
	"github.com/platinasystems/go/flags"
	"github.com/platinasystems/go/notliner"
	"github.com/platinasystems/go/url"
)

const Name = "source"

type cmd struct{}

func New() cmd { return cmd{} }

func (cmd) String() string { return Name }
func (cmd) Usage() string  { return "source [-x] FILE" }

func (cmd) Main(args ...string) error {
	flag, args := flags.New(args, "-x")
	if len(args) == 0 {
		return fmt.Errorf("FILE: missing")
	}
	if len(args) > 1 {
		return fmt.Errorf("%v: unexpected", args[1:])
	}
	f, err := url.Open(args[0])
	if err != nil {
		return err
	}
	defer f.Close()

	t := os.Getenv("TRACE")

	if flag["-x"] {
		os.Setenv("TRACE", "true")
	}

	err = command.Shell(notliner.New(f, nil))

	if flag["-x"] && t != "true" {
		os.Unsetenv("TRACE")
	}
	return err
}

func (cmd) Apropos() map[string]string {
	return map[string]string{
		"en_US.UTF-8": "import command script",
	}
}

func (cmd) Man() map[string]string {
	return map[string]string{
		"en_US.UTF-8": `NAME
	source - import command script

SYNOPSIS
	source [-x] URL

DESCRIPTION
	Import a command script from the given URL.

OPTIONS
	-x	trace each line executed`,
	}
}
