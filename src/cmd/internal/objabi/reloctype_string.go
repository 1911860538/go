// Code generated by "stringer -type=RelocType"; DO NOT EDIT.

package objabi

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[R_ADDR-1]
	_ = x[R_ADDRPOWER-2]
	_ = x[R_ADDRARM64-3]
	_ = x[R_ADDRMIPS-4]
	_ = x[R_ADDROFF-5]
	_ = x[R_SIZE-6]
	_ = x[R_CALL-7]
	_ = x[R_CALLARM-8]
	_ = x[R_CALLARM64-9]
	_ = x[R_CALLIND-10]
	_ = x[R_CALLPOWER-11]
	_ = x[R_CALLMIPS-12]
	_ = x[R_CONST-13]
	_ = x[R_PCREL-14]
	_ = x[R_TLS_LE-15]
	_ = x[R_TLS_IE-16]
	_ = x[R_GOTOFF-17]
	_ = x[R_PLT0-18]
	_ = x[R_PLT1-19]
	_ = x[R_PLT2-20]
	_ = x[R_USEFIELD-21]
	_ = x[R_USETYPE-22]
	_ = x[R_USEIFACE-23]
	_ = x[R_USEIFACEMETHOD-24]
	_ = x[R_USENAMEDMETHOD-25]
	_ = x[R_METHODOFF-26]
	_ = x[R_KEEP-27]
	_ = x[R_POWER_TOC-28]
	_ = x[R_GOTPCREL-29]
	_ = x[R_JMPMIPS-30]
	_ = x[R_DWARFSECREF-31]
	_ = x[R_ARM64_TLS_LE-32]
	_ = x[R_ARM64_TLS_IE-33]
	_ = x[R_ARM64_GOTPCREL-34]
	_ = x[R_ARM64_GOT-35]
	_ = x[R_ARM64_PCREL-36]
	_ = x[R_ARM64_PCREL_LDST8-37]
	_ = x[R_ARM64_PCREL_LDST16-38]
	_ = x[R_ARM64_PCREL_LDST32-39]
	_ = x[R_ARM64_PCREL_LDST64-40]
	_ = x[R_ARM64_LDST8-41]
	_ = x[R_ARM64_LDST16-42]
	_ = x[R_ARM64_LDST32-43]
	_ = x[R_ARM64_LDST64-44]
	_ = x[R_ARM64_LDST128-45]
	_ = x[R_POWER_TLS_LE-46]
	_ = x[R_POWER_TLS_IE-47]
	_ = x[R_POWER_TLS-48]
	_ = x[R_POWER_TLS_IE_PCREL34-49]
	_ = x[R_POWER_TLS_LE_TPREL34-50]
	_ = x[R_ADDRPOWER_DS-51]
	_ = x[R_ADDRPOWER_GOT-52]
	_ = x[R_ADDRPOWER_GOT_PCREL34-53]
	_ = x[R_ADDRPOWER_PCREL-54]
	_ = x[R_ADDRPOWER_TOCREL-55]
	_ = x[R_ADDRPOWER_TOCREL_DS-56]
	_ = x[R_ADDRPOWER_D34-57]
	_ = x[R_ADDRPOWER_PCREL34-58]
	_ = x[R_RISCV_JAL-59]
	_ = x[R_RISCV_JAL_TRAMP-60]
	_ = x[R_RISCV_CALL-61]
	_ = x[R_RISCV_PCREL_ITYPE-62]
	_ = x[R_RISCV_PCREL_STYPE-63]
	_ = x[R_RISCV_TLS_IE-64]
	_ = x[R_RISCV_TLS_LE-65]
	_ = x[R_RISCV_GOT_HI20-66]
	_ = x[R_RISCV_GOT_PCREL_ITYPE-67]
	_ = x[R_RISCV_PCREL_HI20-68]
	_ = x[R_RISCV_PCREL_LO12_I-69]
	_ = x[R_RISCV_PCREL_LO12_S-70]
	_ = x[R_RISCV_BRANCH-71]
	_ = x[R_RISCV_RVC_BRANCH-72]
	_ = x[R_RISCV_RVC_JUMP-73]
	_ = x[R_PCRELDBL-74]
	_ = x[R_LOONG64_ADDR_HI-75]
	_ = x[R_LOONG64_ADDR_LO-76]
	_ = x[R_LOONG64_TLS_LE_HI-77]
	_ = x[R_LOONG64_TLS_LE_LO-78]
	_ = x[R_CALLLOONG64-79]
	_ = x[R_LOONG64_TLS_IE_HI-80]
	_ = x[R_LOONG64_TLS_IE_LO-81]
	_ = x[R_LOONG64_GOT_HI-82]
	_ = x[R_LOONG64_GOT_LO-83]
	_ = x[R_LOONG64_ADD64-84]
	_ = x[R_LOONG64_SUB64-85]
	_ = x[R_JMP16LOONG64-86]
	_ = x[R_JMP21LOONG64-87]
	_ = x[R_JMPLOONG64-88]
	_ = x[R_ADDRMIPSU-89]
	_ = x[R_ADDRMIPSTLS-90]
	_ = x[R_ADDRCUOFF-91]
	_ = x[R_WASMIMPORT-92]
	_ = x[R_XCOFFREF-93]
	_ = x[R_PEIMAGEOFF-94]
	_ = x[R_INITORDER-95]
	_ = x[R_DWTXTADDR_U1-96]
	_ = x[R_DWTXTADDR_U2-97]
	_ = x[R_DWTXTADDR_U3-98]
	_ = x[R_DWTXTADDR_U4-99]
}

const _RelocType_name = "R_ADDRR_ADDRPOWERR_ADDRARM64R_ADDRMIPSR_ADDROFFR_SIZER_CALLR_CALLARMR_CALLARM64R_CALLINDR_CALLPOWERR_CALLMIPSR_CONSTR_PCRELR_TLS_LER_TLS_IER_GOTOFFR_PLT0R_PLT1R_PLT2R_USEFIELDR_USETYPER_USEIFACER_USEIFACEMETHODR_USENAMEDMETHODR_METHODOFFR_KEEPR_POWER_TOCR_GOTPCRELR_JMPMIPSR_DWARFSECREFR_ARM64_TLS_LER_ARM64_TLS_IER_ARM64_GOTPCRELR_ARM64_GOTR_ARM64_PCRELR_ARM64_PCREL_LDST8R_ARM64_PCREL_LDST16R_ARM64_PCREL_LDST32R_ARM64_PCREL_LDST64R_ARM64_LDST8R_ARM64_LDST16R_ARM64_LDST32R_ARM64_LDST64R_ARM64_LDST128R_POWER_TLS_LER_POWER_TLS_IER_POWER_TLSR_POWER_TLS_IE_PCREL34R_POWER_TLS_LE_TPREL34R_ADDRPOWER_DSR_ADDRPOWER_GOTR_ADDRPOWER_GOT_PCREL34R_ADDRPOWER_PCRELR_ADDRPOWER_TOCRELR_ADDRPOWER_TOCREL_DSR_ADDRPOWER_D34R_ADDRPOWER_PCREL34R_RISCV_JALR_RISCV_JAL_TRAMPR_RISCV_CALLR_RISCV_PCREL_ITYPER_RISCV_PCREL_STYPER_RISCV_TLS_IER_RISCV_TLS_LER_RISCV_GOT_HI20R_RISCV_GOT_PCREL_ITYPER_RISCV_PCREL_HI20R_RISCV_PCREL_LO12_IR_RISCV_PCREL_LO12_SR_RISCV_BRANCHR_RISCV_RVC_BRANCHR_RISCV_RVC_JUMPR_PCRELDBLR_LOONG64_ADDR_HIR_LOONG64_ADDR_LOR_LOONG64_TLS_LE_HIR_LOONG64_TLS_LE_LOR_CALLLOONG64R_LOONG64_TLS_IE_HIR_LOONG64_TLS_IE_LOR_LOONG64_GOT_HIR_LOONG64_GOT_LOR_LOONG64_ADD64R_LOONG64_SUB64R_JMP16LOONG64R_JMP21LOONG64R_JMPLOONG64R_ADDRMIPSUR_ADDRMIPSTLSR_ADDRCUOFFR_WASMIMPORTR_XCOFFREFR_PEIMAGEOFFR_INITORDERR_DWTXTADDR_U1R_DWTXTADDR_U2R_DWTXTADDR_U3R_DWTXTADDR_U4"

var _RelocType_index = [...]uint16{0, 6, 17, 28, 38, 47, 53, 59, 68, 79, 88, 99, 109, 116, 123, 131, 139, 147, 153, 159, 165, 175, 184, 194, 210, 226, 237, 243, 254, 264, 273, 286, 300, 314, 330, 341, 354, 373, 393, 413, 433, 446, 460, 474, 488, 503, 517, 531, 542, 564, 586, 600, 615, 638, 655, 673, 694, 709, 728, 739, 756, 768, 787, 806, 820, 834, 850, 873, 891, 911, 931, 945, 963, 979, 989, 1006, 1023, 1042, 1061, 1074, 1093, 1112, 1128, 1144, 1159, 1174, 1188, 1202, 1214, 1225, 1238, 1249, 1261, 1271, 1283, 1294, 1308, 1322, 1336, 1350}

func (i RelocType) String() string {
	i -= 1
	if i < 0 || i >= RelocType(len(_RelocType_index)-1) {
		return "RelocType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RelocType_name[_RelocType_index[i]:_RelocType_index[i+1]]
}
