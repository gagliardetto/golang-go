// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mips

import (
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc"
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa"
	"github.com/gagliardetto/golang-go/cmd/internal/obj/mips"
	"github.com/gagliardetto/golang-go/cmd/internal/objabi"
)

func Init(arch *gc.Arch) {
	arch.LinkArch = &mips.Linkmips
	if objabi.GOARCH == "mipsle" {
		arch.LinkArch = &mips.Linkmipsle
	}
	arch.REGSP = mips.REGSP
	arch.MAXWIDTH = (1 << 31) - 1
	arch.SoftFloat = (objabi.GOMIPS == "softfloat")
	arch.ZeroRange = zerorange
	arch.Ginsnop = ginsnop
	arch.Ginsnopdefer = ginsnop
	arch.SSAMarkMoves = func(s *gc.SSAGenState, b *ssa.Block) {}
	arch.SSAGenValue = ssaGenValue
	arch.SSAGenBlock = ssaGenBlock
}
