// Code generated by mkbuiltin.go. DO NOT EDIT.

package gc

import "github.com/gagliardetto/golang-go/cmd/compile/internal/types"

var runtimeDecls = [...]struct {
	name string
	tag  int
	typ  int
}{
	{"newobject", funcTag, 4},
	{"panicdivide", funcTag, 5},
	{"panicshift", funcTag, 5},
	{"panicmakeslicelen", funcTag, 5},
	{"throwinit", funcTag, 5},
	{"panicwrap", funcTag, 5},
	{"gopanic", funcTag, 7},
	{"gorecover", funcTag, 10},
	{"goschedguarded", funcTag, 5},
	{"goPanicIndex", funcTag, 12},
	{"goPanicIndexU", funcTag, 14},
	{"goPanicSliceAlen", funcTag, 12},
	{"goPanicSliceAlenU", funcTag, 14},
	{"goPanicSliceAcap", funcTag, 12},
	{"goPanicSliceAcapU", funcTag, 14},
	{"goPanicSliceB", funcTag, 12},
	{"goPanicSliceBU", funcTag, 14},
	{"goPanicSlice3Alen", funcTag, 12},
	{"goPanicSlice3AlenU", funcTag, 14},
	{"goPanicSlice3Acap", funcTag, 12},
	{"goPanicSlice3AcapU", funcTag, 14},
	{"goPanicSlice3B", funcTag, 12},
	{"goPanicSlice3BU", funcTag, 14},
	{"goPanicSlice3C", funcTag, 12},
	{"goPanicSlice3CU", funcTag, 14},
	{"printbool", funcTag, 16},
	{"printfloat", funcTag, 18},
	{"printint", funcTag, 20},
	{"printhex", funcTag, 22},
	{"printuint", funcTag, 22},
	{"printcomplex", funcTag, 24},
	{"printstring", funcTag, 26},
	{"printpointer", funcTag, 27},
	{"printiface", funcTag, 27},
	{"printeface", funcTag, 27},
	{"printslice", funcTag, 27},
	{"printnl", funcTag, 5},
	{"printsp", funcTag, 5},
	{"printlock", funcTag, 5},
	{"printunlock", funcTag, 5},
	{"concatstring2", funcTag, 30},
	{"concatstring3", funcTag, 31},
	{"concatstring4", funcTag, 32},
	{"concatstring5", funcTag, 33},
	{"concatstrings", funcTag, 35},
	{"cmpstring", funcTag, 36},
	{"intstring", funcTag, 39},
	{"slicebytetostring", funcTag, 41},
	{"slicebytetostringtmp", funcTag, 42},
	{"slicerunetostring", funcTag, 45},
	{"stringtoslicebyte", funcTag, 46},
	{"stringtoslicerune", funcTag, 49},
	{"slicecopy", funcTag, 51},
	{"slicestringcopy", funcTag, 52},
	{"decoderune", funcTag, 53},
	{"countrunes", funcTag, 54},
	{"convI2I", funcTag, 55},
	{"convT16", funcTag, 57},
	{"convT32", funcTag, 57},
	{"convT64", funcTag, 57},
	{"convTstring", funcTag, 57},
	{"convTslice", funcTag, 57},
	{"convT2E", funcTag, 58},
	{"convT2Enoptr", funcTag, 58},
	{"convT2I", funcTag, 58},
	{"convT2Inoptr", funcTag, 58},
	{"assertE2I", funcTag, 55},
	{"assertE2I2", funcTag, 59},
	{"assertI2I", funcTag, 55},
	{"assertI2I2", funcTag, 59},
	{"panicdottypeE", funcTag, 60},
	{"panicdottypeI", funcTag, 60},
	{"panicnildottype", funcTag, 61},
	{"ifaceeq", funcTag, 63},
	{"efaceeq", funcTag, 63},
	{"fastrand", funcTag, 65},
	{"makemap64", funcTag, 67},
	{"makemap", funcTag, 68},
	{"makemap_small", funcTag, 69},
	{"mapaccess1", funcTag, 70},
	{"mapaccess1_fast32", funcTag, 71},
	{"mapaccess1_fast64", funcTag, 71},
	{"mapaccess1_faststr", funcTag, 71},
	{"mapaccess1_fat", funcTag, 72},
	{"mapaccess2", funcTag, 73},
	{"mapaccess2_fast32", funcTag, 74},
	{"mapaccess2_fast64", funcTag, 74},
	{"mapaccess2_faststr", funcTag, 74},
	{"mapaccess2_fat", funcTag, 75},
	{"mapassign", funcTag, 70},
	{"mapassign_fast32", funcTag, 71},
	{"mapassign_fast32ptr", funcTag, 71},
	{"mapassign_fast64", funcTag, 71},
	{"mapassign_fast64ptr", funcTag, 71},
	{"mapassign_faststr", funcTag, 71},
	{"mapiterinit", funcTag, 76},
	{"mapdelete", funcTag, 76},
	{"mapdelete_fast32", funcTag, 77},
	{"mapdelete_fast64", funcTag, 77},
	{"mapdelete_faststr", funcTag, 77},
	{"mapiternext", funcTag, 78},
	{"mapclear", funcTag, 79},
	{"makechan64", funcTag, 81},
	{"makechan", funcTag, 82},
	{"chanrecv1", funcTag, 84},
	{"chanrecv2", funcTag, 85},
	{"chansend1", funcTag, 87},
	{"closechan", funcTag, 27},
	{"writeBarrier", varTag, 89},
	{"typedmemmove", funcTag, 90},
	{"typedmemclr", funcTag, 91},
	{"typedslicecopy", funcTag, 92},
	{"selectnbsend", funcTag, 93},
	{"selectnbrecv", funcTag, 94},
	{"selectnbrecv2", funcTag, 96},
	{"selectsetpc", funcTag, 61},
	{"selectgo", funcTag, 97},
	{"block", funcTag, 5},
	{"makeslice", funcTag, 98},
	{"makeslice64", funcTag, 99},
	{"growslice", funcTag, 101},
	{"memmove", funcTag, 102},
	{"memclrNoHeapPointers", funcTag, 103},
	{"memclrHasPointers", funcTag, 103},
	{"memequal", funcTag, 104},
	{"memequal0", funcTag, 105},
	{"memequal8", funcTag, 105},
	{"memequal16", funcTag, 105},
	{"memequal32", funcTag, 105},
	{"memequal64", funcTag, 105},
	{"memequal128", funcTag, 105},
	{"f32equal", funcTag, 106},
	{"f64equal", funcTag, 106},
	{"c64equal", funcTag, 106},
	{"c128equal", funcTag, 106},
	{"strequal", funcTag, 106},
	{"interequal", funcTag, 106},
	{"nilinterequal", funcTag, 106},
	{"memhash", funcTag, 107},
	{"memhash0", funcTag, 108},
	{"memhash8", funcTag, 108},
	{"memhash16", funcTag, 108},
	{"memhash32", funcTag, 108},
	{"memhash64", funcTag, 108},
	{"memhash128", funcTag, 108},
	{"f32hash", funcTag, 108},
	{"f64hash", funcTag, 108},
	{"c64hash", funcTag, 108},
	{"c128hash", funcTag, 108},
	{"strhash", funcTag, 108},
	{"interhash", funcTag, 108},
	{"nilinterhash", funcTag, 108},
	{"int64div", funcTag, 109},
	{"uint64div", funcTag, 110},
	{"int64mod", funcTag, 109},
	{"uint64mod", funcTag, 110},
	{"float64toint64", funcTag, 111},
	{"float64touint64", funcTag, 112},
	{"float64touint32", funcTag, 113},
	{"int64tofloat64", funcTag, 114},
	{"uint64tofloat64", funcTag, 115},
	{"uint32tofloat64", funcTag, 116},
	{"complex128div", funcTag, 117},
	{"racefuncenter", funcTag, 118},
	{"racefuncenterfp", funcTag, 5},
	{"racefuncexit", funcTag, 5},
	{"raceread", funcTag, 118},
	{"racewrite", funcTag, 118},
	{"racereadrange", funcTag, 119},
	{"racewriterange", funcTag, 119},
	{"msanread", funcTag, 119},
	{"msanwrite", funcTag, 119},
	{"checkptrAlignment", funcTag, 120},
	{"checkptrArithmetic", funcTag, 122},
	{"libfuzzerTraceCmp1", funcTag, 124},
	{"libfuzzerTraceCmp2", funcTag, 126},
	{"libfuzzerTraceCmp4", funcTag, 127},
	{"libfuzzerTraceCmp8", funcTag, 128},
	{"libfuzzerTraceConstCmp1", funcTag, 124},
	{"libfuzzerTraceConstCmp2", funcTag, 126},
	{"libfuzzerTraceConstCmp4", funcTag, 127},
	{"libfuzzerTraceConstCmp8", funcTag, 128},
	{"x86HasPOPCNT", varTag, 15},
	{"x86HasSSE41", varTag, 15},
	{"x86HasFMA", varTag, 15},
	{"armHasVFPv4", varTag, 15},
	{"arm64HasATOMICS", varTag, 15},
}

func runtimeTypes() []*types.Type {
	var typs [129]*types.Type
	typs[0] = types.Bytetype
	typs[1] = types.NewPtr(typs[0])
	typs[2] = types.Types[TANY]
	typs[3] = types.NewPtr(typs[2])
	typs[4] = functype(nil, []*Node{anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[5] = functype(nil, nil, nil)
	typs[6] = types.Types[TINTER]
	typs[7] = functype(nil, []*Node{anonfield(typs[6])}, nil)
	typs[8] = types.Types[TINT32]
	typs[9] = types.NewPtr(typs[8])
	typs[10] = functype(nil, []*Node{anonfield(typs[9])}, []*Node{anonfield(typs[6])})
	typs[11] = types.Types[TINT]
	typs[12] = functype(nil, []*Node{anonfield(typs[11]), anonfield(typs[11])}, nil)
	typs[13] = types.Types[TUINT]
	typs[14] = functype(nil, []*Node{anonfield(typs[13]), anonfield(typs[11])}, nil)
	typs[15] = types.Types[TBOOL]
	typs[16] = functype(nil, []*Node{anonfield(typs[15])}, nil)
	typs[17] = types.Types[TFLOAT64]
	typs[18] = functype(nil, []*Node{anonfield(typs[17])}, nil)
	typs[19] = types.Types[TINT64]
	typs[20] = functype(nil, []*Node{anonfield(typs[19])}, nil)
	typs[21] = types.Types[TUINT64]
	typs[22] = functype(nil, []*Node{anonfield(typs[21])}, nil)
	typs[23] = types.Types[TCOMPLEX128]
	typs[24] = functype(nil, []*Node{anonfield(typs[23])}, nil)
	typs[25] = types.Types[TSTRING]
	typs[26] = functype(nil, []*Node{anonfield(typs[25])}, nil)
	typs[27] = functype(nil, []*Node{anonfield(typs[2])}, nil)
	typs[28] = types.NewArray(typs[0], 32)
	typs[29] = types.NewPtr(typs[28])
	typs[30] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[25]), anonfield(typs[25])}, []*Node{anonfield(typs[25])})
	typs[31] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25])}, []*Node{anonfield(typs[25])})
	typs[32] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25])}, []*Node{anonfield(typs[25])})
	typs[33] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25]), anonfield(typs[25])}, []*Node{anonfield(typs[25])})
	typs[34] = types.NewSlice(typs[25])
	typs[35] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[34])}, []*Node{anonfield(typs[25])})
	typs[36] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[25])}, []*Node{anonfield(typs[11])})
	typs[37] = types.NewArray(typs[0], 4)
	typs[38] = types.NewPtr(typs[37])
	typs[39] = functype(nil, []*Node{anonfield(typs[38]), anonfield(typs[19])}, []*Node{anonfield(typs[25])})
	typs[40] = types.NewSlice(typs[0])
	typs[41] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[40])}, []*Node{anonfield(typs[25])})
	typs[42] = functype(nil, []*Node{anonfield(typs[40])}, []*Node{anonfield(typs[25])})
	typs[43] = types.Runetype
	typs[44] = types.NewSlice(typs[43])
	typs[45] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[44])}, []*Node{anonfield(typs[25])})
	typs[46] = functype(nil, []*Node{anonfield(typs[29]), anonfield(typs[25])}, []*Node{anonfield(typs[40])})
	typs[47] = types.NewArray(typs[43], 32)
	typs[48] = types.NewPtr(typs[47])
	typs[49] = functype(nil, []*Node{anonfield(typs[48]), anonfield(typs[25])}, []*Node{anonfield(typs[44])})
	typs[50] = types.Types[TUINTPTR]
	typs[51] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2]), anonfield(typs[50])}, []*Node{anonfield(typs[11])})
	typs[52] = functype(nil, []*Node{anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[11])})
	typs[53] = functype(nil, []*Node{anonfield(typs[25]), anonfield(typs[11])}, []*Node{anonfield(typs[43]), anonfield(typs[11])})
	typs[54] = functype(nil, []*Node{anonfield(typs[25])}, []*Node{anonfield(typs[11])})
	typs[55] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2])})
	typs[56] = types.Types[TUNSAFEPTR]
	typs[57] = functype(nil, []*Node{anonfield(typs[2])}, []*Node{anonfield(typs[56])})
	typs[58] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, []*Node{anonfield(typs[2])})
	typs[59] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2])}, []*Node{anonfield(typs[2]), anonfield(typs[15])})
	typs[60] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[1])}, nil)
	typs[61] = functype(nil, []*Node{anonfield(typs[1])}, nil)
	typs[62] = types.NewPtr(typs[50])
	typs[63] = functype(nil, []*Node{anonfield(typs[62]), anonfield(typs[56]), anonfield(typs[56])}, []*Node{anonfield(typs[15])})
	typs[64] = types.Types[TUINT32]
	typs[65] = functype(nil, nil, []*Node{anonfield(typs[64])})
	typs[66] = types.NewMap(typs[2], typs[2])
	typs[67] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[19]), anonfield(typs[3])}, []*Node{anonfield(typs[66])})
	typs[68] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[11]), anonfield(typs[3])}, []*Node{anonfield(typs[66])})
	typs[69] = functype(nil, nil, []*Node{anonfield(typs[66])})
	typs[70] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[3])}, []*Node{anonfield(typs[3])})
	typs[71] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[2])}, []*Node{anonfield(typs[3])})
	typs[72] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3])})
	typs[73] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[3])}, []*Node{anonfield(typs[3]), anonfield(typs[15])})
	typs[74] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[2])}, []*Node{anonfield(typs[3]), anonfield(typs[15])})
	typs[75] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[3]), anonfield(typs[1])}, []*Node{anonfield(typs[3]), anonfield(typs[15])})
	typs[76] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[3])}, nil)
	typs[77] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66]), anonfield(typs[2])}, nil)
	typs[78] = functype(nil, []*Node{anonfield(typs[3])}, nil)
	typs[79] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[66])}, nil)
	typs[80] = types.NewChan(typs[2], types.Cboth)
	typs[81] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[19])}, []*Node{anonfield(typs[80])})
	typs[82] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[11])}, []*Node{anonfield(typs[80])})
	typs[83] = types.NewChan(typs[2], types.Crecv)
	typs[84] = functype(nil, []*Node{anonfield(typs[83]), anonfield(typs[3])}, nil)
	typs[85] = functype(nil, []*Node{anonfield(typs[83]), anonfield(typs[3])}, []*Node{anonfield(typs[15])})
	typs[86] = types.NewChan(typs[2], types.Csend)
	typs[87] = functype(nil, []*Node{anonfield(typs[86]), anonfield(typs[3])}, nil)
	typs[88] = types.NewArray(typs[0], 3)
	typs[89] = tostruct([]*Node{namedfield("enabled", typs[15]), namedfield("pad", typs[88]), namedfield("needed", typs[15]), namedfield("cgo", typs[15]), namedfield("alignme", typs[21])})
	typs[90] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3]), anonfield(typs[3])}, nil)
	typs[91] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[3])}, nil)
	typs[92] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[2]), anonfield(typs[2])}, []*Node{anonfield(typs[11])})
	typs[93] = functype(nil, []*Node{anonfield(typs[86]), anonfield(typs[3])}, []*Node{anonfield(typs[15])})
	typs[94] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[83])}, []*Node{anonfield(typs[15])})
	typs[95] = types.NewPtr(typs[15])
	typs[96] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[95]), anonfield(typs[83])}, []*Node{anonfield(typs[15])})
	typs[97] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[1]), anonfield(typs[11])}, []*Node{anonfield(typs[11]), anonfield(typs[15])})
	typs[98] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[11]), anonfield(typs[11])}, []*Node{anonfield(typs[56])})
	typs[99] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[19]), anonfield(typs[19])}, []*Node{anonfield(typs[56])})
	typs[100] = types.NewSlice(typs[2])
	typs[101] = functype(nil, []*Node{anonfield(typs[1]), anonfield(typs[100]), anonfield(typs[11])}, []*Node{anonfield(typs[100])})
	typs[102] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[50])}, nil)
	typs[103] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[50])}, nil)
	typs[104] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3]), anonfield(typs[50])}, []*Node{anonfield(typs[15])})
	typs[105] = functype(nil, []*Node{anonfield(typs[3]), anonfield(typs[3])}, []*Node{anonfield(typs[15])})
	typs[106] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[56])}, []*Node{anonfield(typs[15])})
	typs[107] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[50]), anonfield(typs[50])}, []*Node{anonfield(typs[50])})
	typs[108] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[50])}, []*Node{anonfield(typs[50])})
	typs[109] = functype(nil, []*Node{anonfield(typs[19]), anonfield(typs[19])}, []*Node{anonfield(typs[19])})
	typs[110] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[21])}, []*Node{anonfield(typs[21])})
	typs[111] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[19])})
	typs[112] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[21])})
	typs[113] = functype(nil, []*Node{anonfield(typs[17])}, []*Node{anonfield(typs[64])})
	typs[114] = functype(nil, []*Node{anonfield(typs[19])}, []*Node{anonfield(typs[17])})
	typs[115] = functype(nil, []*Node{anonfield(typs[21])}, []*Node{anonfield(typs[17])})
	typs[116] = functype(nil, []*Node{anonfield(typs[64])}, []*Node{anonfield(typs[17])})
	typs[117] = functype(nil, []*Node{anonfield(typs[23]), anonfield(typs[23])}, []*Node{anonfield(typs[23])})
	typs[118] = functype(nil, []*Node{anonfield(typs[50])}, nil)
	typs[119] = functype(nil, []*Node{anonfield(typs[50]), anonfield(typs[50])}, nil)
	typs[120] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[1]), anonfield(typs[50])}, nil)
	typs[121] = types.NewSlice(typs[56])
	typs[122] = functype(nil, []*Node{anonfield(typs[56]), anonfield(typs[121])}, nil)
	typs[123] = types.Types[TUINT8]
	typs[124] = functype(nil, []*Node{anonfield(typs[123]), anonfield(typs[123])}, nil)
	typs[125] = types.Types[TUINT16]
	typs[126] = functype(nil, []*Node{anonfield(typs[125]), anonfield(typs[125])}, nil)
	typs[127] = functype(nil, []*Node{anonfield(typs[64]), anonfield(typs[64])}, nil)
	typs[128] = functype(nil, []*Node{anonfield(typs[21]), anonfield(typs[21])}, nil)
	return typs[:]
}
