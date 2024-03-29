// Copyright 2015-2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style license described in the
// LICENSE file.

package slice_string

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	var args []string
	args = New(`echo hello\ beautiful\ world`)
	if !reflect.DeepEqual(args, []string{
		"echo",
		"hello beautiful world",
	}) {
		t.Error("unexpected:", args)
	}
	args = New(`echo "hello 'beautiful world'"`)
	if !reflect.DeepEqual(args, []string{
		"echo",
		"hello 'beautiful world'",
	}) {
		t.Error("unexpected:", args)
	}
	args = New(`echo 'hello \"beautiful world\"'`)
	if !reflect.DeepEqual(args, []string{
		"echo",
		`hello \"beautiful world\"`,
	}) {
		t.Error("unexpected:", args)
	}
}
