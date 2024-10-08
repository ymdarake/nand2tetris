package assembler

type SymbolTable interface {
	AddEntry(key string, address int)
	Contains(key string) bool
	GetAddress(key string) int
}

type symbolTableImpl struct {
	entries map[string]int
}

func NewSymbolTable() SymbolTable {
	return &symbolTableImpl{
		entries: map[string]int{
			"SP":     0,
			"LCL":    1,
			"ARG":    2,
			"THIS":   3,
			"THAT":   4,
			"R0":     0,
			"R1":     1,
			"R2":     2,
			"R3":     3,
			"R4":     4,
			"R5":     5,
			"R6":     6,
			"R7":     7,
			"R8":     8,
			"R9":     9,
			"R10":    10,
			"R11":    11,
			"R12":    12,
			"R13":    13,
			"R14":    14,
			"R15":    15,
			"SCREEN": 0x4000,
			"KBD":    0x6000,
		},
	}
}

// AddEntry implements SymbolTable.
func (s *symbolTableImpl) AddEntry(key string, address int) {
	s.entries[key] = address
}

// Contains implements SymbolTable.
func (s *symbolTableImpl) Contains(key string) bool {
	_, ok := s.entries[key]
	return ok
}

// GetAddress implements SymbolTable.
func (s *symbolTableImpl) GetAddress(key string) int {
	v, ok := s.entries[key]
	if !ok {
		return -1
	}
	return v
}
