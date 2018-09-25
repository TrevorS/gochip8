package main

type CPU struct {
	mmu        MMU
	registers  Registers
	stack      Stack
	video      VideoBuffer
	input      InputBuffer
	delayTimer Byte
	soundTimer Byte
	sp         Word
	pc         Word
	i          Word
}

func NewCPU(mmu MMU) CPU {
	return CPU{
		mmu:        mmu,
		registers:  NewRegisters(),
		stack:      NewStack(),
		video:      NewVideoBuffer(),
		input:      NewInputBuffer(),
		delayTimer: 0,
		soundTimer: 0,
		sp:         0,
		pc:         0x200,
		i:          0,
	}
}
