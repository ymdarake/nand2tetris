package assembler

import (
	"fmt"
	"strconv"
)

type Bits struct {
	data []int
}

func (b Bits) String() string {
	out := ""
	for _, v := range b.data {
		out += strconv.Itoa(v)
	}
	return out
}

// d1 d2 d3 の各ビットが出力先を示す。
// A  D  M
func GenDest(dest string) Bits {
	switch dest {
	case "null":
		return Bits{[]int{0, 0, 0}}
	case "M":
		return Bits{[]int{0, 0, 1}}
	case "D":
		return Bits{[]int{0, 1, 0}}
	case "MD":
		return Bits{[]int{0, 1, 1}}
	case "A":
		return Bits{[]int{1, 0, 0}}
	case "AM":
		return Bits{[]int{1, 0, 1}}
	case "AD":
		return Bits{[]int{1, 1, 0}}
	case "AMD":
		return Bits{[]int{1, 1, 1}}
	}

	fmt.Printf("ERROR: GenDest: unknown dest %s\n", dest)
	return Bits{[]int{0, 0, 0}}
}

// a c1c2c3c4c5c6
func GenComp(comp string) Bits {
	switch comp {
	case "0":
		return Bits{[]int{0, 1, 0, 1, 0, 1, 0}}
	case "1":
		return Bits{[]int{0, 1, 1, 1, 1, 1, 1}}
	case "-1":
		return Bits{[]int{0, 1, 1, 1, 0, 1, 0}}
	case "D":
		return Bits{[]int{0, 0, 0, 1, 1, 0, 0}}
	case "A":
		return Bits{[]int{0, 1, 1, 0, 0, 0, 0}}
	case "M":
		return Bits{[]int{1, 1, 1, 0, 0, 0, 0}}
	case "!D":
		return Bits{[]int{0, 0, 0, 1, 1, 0, 1}}
	case "!A":
		return Bits{[]int{0, 1, 1, 0, 0, 0, 1}}
	case "!M":
		return Bits{[]int{1, 1, 1, 0, 0, 0, 1}}
	case "-D":
		return Bits{[]int{0, 0, 0, 1, 1, 1, 1}}
	case "-A":
		return Bits{[]int{0, 1, 1, 0, 0, 1, 1}}
	case "-M":
		return Bits{[]int{1, 1, 1, 0, 0, 1, 1}}
	case "D+1":
		return Bits{[]int{0, 0, 1, 1, 1, 1, 1}}
	case "A+1":
		return Bits{[]int{0, 1, 1, 0, 1, 1, 1}}
	case "M+1":
		return Bits{[]int{1, 1, 1, 0, 1, 1, 1}}
	case "D-1":
		return Bits{[]int{0, 0, 0, 1, 1, 1, 0}}
	case "A-1":
		return Bits{[]int{0, 1, 1, 0, 0, 1, 0}}
	case "M-1":
		return Bits{[]int{1, 1, 1, 0, 0, 1, 0}}
	case "D+A":
		return Bits{[]int{0, 0, 0, 0, 0, 1, 0}}
	case "D+M":
		return Bits{[]int{1, 0, 0, 0, 0, 1, 0}}
	case "D-A":
		return Bits{[]int{0, 0, 1, 0, 0, 1, 1}}
	case "D-M":
		return Bits{[]int{1, 0, 1, 0, 0, 1, 1}}
	case "A-D":
		return Bits{[]int{0, 0, 0, 0, 1, 1, 1}}
	case "M-D":
		return Bits{[]int{1, 0, 0, 0, 1, 1, 1}}
	case "D&A":
		return Bits{[]int{0, 0, 0, 0, 0, 0, 0}}
	case "D&M":
		return Bits{[]int{1, 0, 0, 0, 0, 0, 0}}
	case "D|A":
		return Bits{[]int{0, 0, 1, 0, 1, 0, 1}}
	case "D|M":
		return Bits{[]int{1, 0, 1, 0, 1, 0, 1}}
	}

	fmt.Printf("ERROR: GenComp: unknown code given: %s\n", comp)
	return Bits{[]int{0, 0, 0, 0, 0, 0, 0}}
}
func GenJump(jump string) Bits {
	switch jump {
	case "null":
		return Bits{[]int{0, 0, 0}}
	case "JGT":
		return Bits{[]int{0, 0, 1}}
	case "JEQ":
		return Bits{[]int{0, 1, 0}}
	case "JGE":
		return Bits{[]int{0, 1, 1}}
	case "JLT":
		return Bits{[]int{1, 0, 0}}
	case "JNE":
		return Bits{[]int{1, 0, 1}}
	case "JLE":
		return Bits{[]int{1, 1, 0}}
	case "JMP":
		return Bits{[]int{1, 1, 1}}
	}

	fmt.Printf("ERROR: GenJump: unknown jump %s\n", jump)
	return Bits{[]int{0, 0, 0}}
}
