package gochip8

// Mmu is the Memory Mapped Unit for the Chip8 system.
type Mmu struct {
	memory  [4096]byte
	fontset [80]byte
}

// Word represents the size of a MMU word.
type Word uint16

// NewMmu returns an initialized Chip8 MMU.
func NewMmu() Mmu {
	var mmu = Mmu{}

	mmu.Reset()

	return mmu
}

// Reset re-initializes the MMU.
func (mmu *Mmu) Reset() {
	for i := range mmu.memory {
		mmu.memory[i] = 0
	}

	fontset := GetFontset()

	for i, value := range fontset {
		mmu.memory[i] = value
	}
}

func (mmu *Mmu) WriteByte(address int, value byte) {
	mmu.memory[address] = value
}

func (mmu *Mmu) ReadByte(address int) byte {
	return mmu.memory[address]
}

func (mmu *Mmu) ReadWord(address int) Word {
	return Word(mmu.memory[address])<<8 | Word(mmu.memory[address+1])
}

func (mmu *Mmu) LoadRom(rom []byte) {
	for i, value := range rom {
		mmu.memory[i+512] = value
	}
}
