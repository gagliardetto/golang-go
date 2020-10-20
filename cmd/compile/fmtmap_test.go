// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the knownFormats map which records the valid
// formats for a given type. The valid formats must correspond to
// supported compiler formats implemented in fmt.go, or whatever
// other format verbs are implemented for the given type. The map may
// also be used to change the use of a format verb across all compiler
// sources automatically (for instance, if the implementation of fmt.go
// changes), by using the -r option together with the new formats in the
// map. To generate this file automatically from the existing source,
// run: go test -run Formats -u.
//
// See the package comment in fmt_test.go for additional information.

package main_test

// knownFormats entries are of the form "typename format" -> "newformat".
// An absent entry means that the format is not recognized as valid.
// An empty new format means that the format should remain unchanged.
var knownFormats = map[string]string{
	"*bytes.Buffer %s":                                "",
	"*cmd/compile/internal/gc.EscLocation %v":         "",
	"*cmd/compile/internal/gc.Mpflt %v":               "",
	"*cmd/compile/internal/gc.Mpint %v":               "",
	"*cmd/compile/internal/gc.Node %#v":               "",
	"*cmd/compile/internal/gc.Node %+S":               "",
	"*cmd/compile/internal/gc.Node %+v":               "",
	"*cmd/compile/internal/gc.Node %L":                "",
	"*cmd/compile/internal/gc.Node %S":                "",
	"*cmd/compile/internal/gc.Node %j":                "",
	"*cmd/compile/internal/gc.Node %p":                "",
	"*cmd/compile/internal/gc.Node %v":                "",
	"*cmd/compile/internal/ssa.Block %s":              "",
	"*cmd/compile/internal/ssa.Block %v":              "",
	"*cmd/compile/internal/ssa.Func %s":               "",
	"*cmd/compile/internal/ssa.Func %v":               "",
	"*cmd/compile/internal/ssa.Register %s":           "",
	"*cmd/compile/internal/ssa.Register %v":           "",
	"*cmd/compile/internal/ssa.SparseTreeNode %v":     "",
	"*cmd/compile/internal/ssa.Value %s":              "",
	"*cmd/compile/internal/ssa.Value %v":              "",
	"*cmd/compile/internal/ssa.sparseTreeMapEntry %v": "",
	"*cmd/compile/internal/types.Field %p":            "",
	"*cmd/compile/internal/types.Field %v":            "",
	"*cmd/compile/internal/types.Sym %0S":             "",
	"*cmd/compile/internal/types.Sym %S":              "",
	"*cmd/compile/internal/types.Sym %p":              "",
	"*cmd/compile/internal/types.Sym %v":              "",
	"*cmd/compile/internal/types.Type %#L":            "",
	"*cmd/compile/internal/types.Type %#v":            "",
	"*cmd/compile/internal/types.Type %+v":            "",
	"*cmd/compile/internal/types.Type %-S":            "",
	"*cmd/compile/internal/types.Type %0S":            "",
	"*cmd/compile/internal/types.Type %L":             "",
	"*cmd/compile/internal/types.Type %S":             "",
	"*cmd/compile/internal/types.Type %p":             "",
	"*cmd/compile/internal/types.Type %s":             "",
	"*cmd/compile/internal/types.Type %v":             "",
	"*cmd/internal/obj.Addr %v":                       "",
	"*cmd/internal/obj.LSym %v":                       "",
	"*math/big.Float %f":                              "",
	"*math/big.Int %#x":                               "",
	"*math/big.Int %s":                                "",
	"*math/big.Int %v":                                "",
	"[16]byte %x":                                     "",
	"[]*cmd/compile/internal/ssa.Block %v":            "",
	"[]*cmd/compile/internal/ssa.Value %v":            "",
	"[][]string %q":                                   "",
	"[]byte %s":                                       "",
	"[]byte %x":                                       "",
	"[]cmd/compile/internal/ssa.Edge %v":              "",
	"[]cmd/compile/internal/ssa.ID %v":                "",
	"[]cmd/compile/internal/ssa.posetNode %v":         "",
	"[]cmd/compile/internal/ssa.posetUndo %v":         "",
	"[]cmd/compile/internal/syntax.token %s":          "",
	"[]string %v":                                     "",
	"[]uint32 %v":                                     "",
	"bool %v":                                         "",
	"byte %08b":                                       "",
	"byte %c":                                         "",
	"byte %q":                                         "",
	"byte %v":                                         "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/arm.shift %d":               "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Class %d":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Class %s":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Class %v":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Ctype %d":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Ctype %v":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Nodes %#v":               "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Nodes %+v":               "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Nodes %.v":               "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Nodes %v":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Op %#v":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Op %v":                   "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Val %#v":                 "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Val %T":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.Val %v":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.fmtMode %d":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.initKind %d":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/gc.itag %v":                 "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.BranchPrediction %d":    "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.Edge %v":                "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.GCNode %v":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.ID %d":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.ID %v":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.LocPair %s":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.LocalSlot %s":           "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.LocalSlot %v":           "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.Location %T":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.Location %s":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.Op %s":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.Op %v":                  "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.ValAndOff %s":           "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.domain %v":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.posetNode %v":           "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.posetTestOp %v":         "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.rbrank %d":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.regMask %d":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.register %d":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/ssa.relation %s":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Error %q":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Expr %#v":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.LitKind %d":          "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Node %T":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Operator %s":         "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Pos %s":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.Pos %v":              "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.position %s":         "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.token %q":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/syntax.token %s":            "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/types.EType %d":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/types.EType %s":             "",
	"github.com/gagliardetto/golang-go/cmd/compile/internal/types.EType %v":             "",
	"github.com/gagliardetto/golang-go/cmd/internal/obj.ABI %v":                         "",
	"error %v":                                        "",
	"float64 %.2f":                                    "",
	"float64 %.3f":                                    "",
	"float64 %.6g":                                    "",
	"float64 %g":                                      "",
	"int %-12d":                                       "",
	"int %-6d":                                        "",
	"int %-8o":                                        "",
	"int %02d":                                        "",
	"int %6d":                                         "",
	"int %c":                                          "",
	"int %d":                                          "",
	"int %v":                                          "",
	"int %x":                                          "",
	"int16 %d":                                        "",
	"int16 %x":                                        "",
	"int32 %d":                                        "",
	"int32 %v":                                        "",
	"int32 %x":                                        "",
	"int64 %+d":                                       "",
	"int64 %-10d":                                     "",
	"int64 %.5d":                                      "",
	"int64 %X":                                        "",
	"int64 %d":                                        "",
	"int64 %v":                                        "",
	"int64 %x":                                        "",
	"int8 %d":                                         "",
	"int8 %v":                                         "",
	"int8 %x":                                         "",
	"interface{} %#v":                                 "",
	"interface{} %T":                                  "",
	"interface{} %p":                                  "",
	"interface{} %q":                                  "",
	"interface{} %s":                                  "",
	"interface{} %v":                                  "",
	"map[*cmd/compile/internal/gc.Node]*cmd/compile/internal/ssa.Value %v": "",
	"map[*cmd/compile/internal/gc.Node][]*cmd/compile/internal/gc.Node %v": "",
	"map[cmd/compile/internal/ssa.ID]uint32 %v":                            "",
	"map[int64]uint32 %v":  "",
	"math/big.Accuracy %s": "",
	"reflect.Type %s":      "",
	"rune %#U":             "",
	"rune %c":              "",
	"rune %q":              "",
	"string %-*s":          "",
	"string %-16s":         "",
	"string %-6s":          "",
	"string %q":            "",
	"string %s":            "",
	"string %v":            "",
	"time.Duration %d":     "",
	"time.Duration %v":     "",
	"uint %04x":            "",
	"uint %5d":             "",
	"uint %d":              "",
	"uint %x":              "",
	"uint16 %d":            "",
	"uint16 %x":            "",
	"uint32 %#U":           "",
	"uint32 %#x":           "",
	"uint32 %d":            "",
	"uint32 %v":            "",
	"uint32 %x":            "",
	"uint64 %08x":          "",
	"uint64 %b":            "",
	"uint64 %d":            "",
	"uint64 %x":            "",
	"uint8 %d":             "",
	"uint8 %v":             "",
	"uint8 %x":             "",
	"uintptr %d":           "",
}
