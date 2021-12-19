package main

type InstructionType uint8

type Instruction struct {
	instruction uint16
}

const (
	// Registers
	R_R0 = iota
	R_R1
	R_R2
	R_R3
	R_R4
	R_R5
	R_R6
	R_R7
	R_PC /* program counter */
	R_COND
	R_COUNT
)

// operation code
const (
	OP_BR   = iota /* branch */
	OP_ADD         /* add  */
	OP_LD          /* load */
	OP_ST          /* store */
	OP_JSR         /* jump register */
	OP_AND         /* bitwise and */
	OP_LDR         /* load register */
	OP_STR         /* store register */
	OP_RTI         /* unused */
	OP_NOT         /* bitwise not */
	OP_LDI         /* load indirect */
	OP_STI         /* store indirect */
	OP_JMP         /* jump */
	OP_RES         /* reserved (unused) */
	OP_LEA         /* load effective address */
	OP_TRAP        /* execute trap */
)

// condition flags
const (
	FL_POS = 1 << 0 /* P */
	FL_ZRO = 1 << 1 /* Z */
	FL_NEG = 1 << 2 /* N */
)

const (
	StackSize = 65536 // 2^16
)

var (
	VMStack         = make([]uint16, StackSize)
	Register        = make([]uint16, R_COUNT)
	PC       uint16 = 0x0010
	SP       uint16 = 0x0000
)

func (ins *Instruction) Type() InstructionType {
	return InstructionType(ins.instruction >> 12)
}
