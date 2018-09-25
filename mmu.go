package main

// MMU is the Memory Mapped Unit for the Chip8 system.
type MMU struct {
	memory  Memory
	fontset Fontset
}

// NewMMU returns an initialized Chip8 MMU.
func NewMMU() MMU {
	var mmu = MMU{}

	mmu.Reset()

	return mmu
}

// Reset re-initializes the MMU.
func (mmu *MMU) Reset() {
	mmu.memory = NewMemory()

	for i, value := range NewFontset() {
		mmu.memory[i] = value
	}
}

func (mmu *MMU) WriteByte(address int, value Byte) {
	mmu.memory[address] = value
}

func (mmu *MMU) ReadByte(address int) Byte {
	return mmu.memory[address]
}

func (mmu *MMU) ReadWord(address int) Word {
	return Word(mmu.memory[address])<<8 | Word(mmu.memory[address+1])
}

func (mmu *MMU) LoadRom(rom []Byte) {
	for i, value := range rom {
		mmu.memory[i+512] = value
	}
}
