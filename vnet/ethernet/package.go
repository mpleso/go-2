// Copyright 2016 Platina Systems, Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ethernet

import (
	"github.com/platinasystems/go/vnet"
)

var packageIndex uint

type Main struct {
	vnet.Package
	ipNeighborMain
	nodeMain
	pgMain
}

func Init(v *vnet.Vnet) {
	m := &Main{}
	packageIndex = v.AddPackage("ethernet", m)
	m.DependsOn("pg")
}

func GetMain(v *vnet.Vnet) *Main { return v.GetPackage(packageIndex).(*Main) }

func (m *Main) Init() (err error) {
	v := m.Vnet
	m.ipNeighborMain.init(v)
	m.nodeInit(v)
	m.pgMain.pgInit(v)
	return
}
