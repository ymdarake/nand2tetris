package chap06

import "fmt"

// d1 d2 d3 の各ビットが出力先を示す。
// A  D  M
func GenDest(dest string) [3]int {
	switch dest {
	case "null":
		return [3]int{0, 0, 0}
	case "M":
		return [3]int{0, 0, 1}
	case "D":
		return [3]int{0, 1, 0}
	case "MD":
		return [3]int{0, 1, 1}
	case "A":
		return [3]int{1, 0, 0}
	case "AM":
		return [3]int{1, 0, 1}
	case "AD":
		return [3]int{1, 1, 0}
	case "AMD":
		return [3]int{1, 1, 1}
	}

	fmt.Printf("ERROR: GenDest: unknown dest %s\n", dest)
	return [3]int{0, 0, 0}
}
func GenComp(comp string) [7]int {
	switch comp {
	case "0":
		return [7]int{0, 1, 0, 1, 0, 1, 0}
	case "1":
		return [7]int{0, 1, 1, 1, 1, 1, 1}
	case "-1":
		return [7]int{0, 1, 1, 1, 0, 1, 0}
	case "D":
		return [7]int{0, 0, 0, 1, 1, 0, 0}
	case "A":
		return [7]int{0, 1, 1, 0, 0, 0, 0}
	case "M":
		return [7]int{1, 1, 1, 0, 0, 0, 0}
	case "!D":
		return [7]int{0, 0, 0, 1, 1, 0, 1}
	case "!A":
		return [7]int{0, 1, 1, 0, 0, 0, 1}
	case "!M":
		return [7]int{1, 1, 1, 0, 0, 0, 1}
	case "-D":
		return [7]int{0, 0, 0, 1, 1, 1, 1}
	case "-A":
		return [7]int{0, 1, 1, 0, 0, 1, 1}
	case "-M":
		return [7]int{1, 1, 1, 0, 0, 1, 1}
	case "D+1":
		return [7]int{0, 0, 1, 1, 1, 1, 1}
	case "A+1":
		return [7]int{0, 1, 1, 0, 1, 1, 1}
	case "M+1":
		return [7]int{1, 1, 1, 0, 1, 1, 1}
	case "D-1":
		return [7]int{0, 0, 0, 1, 1, 1, 0}
	case "A-1":
		return [7]int{0, 1, 1, 0, 0, 1, 0}
	case "M-1":
		return [7]int{1, 1, 1, 0, 0, 1, 0}
	case "D+A":
		return [7]int{0, 0, 0, 0, 0, 1, 0}
	case "D+M":
		return [7]int{1, 0, 0, 0, 0, 1, 0}
	case "D-A":
		return [7]int{0, 0, 1, 0, 0, 1, 1}
	case "D-M":
		return [7]int{1, 0, 1, 0, 0, 1, 1}
	case "A-D":
		return [7]int{0, 0, 0, 0, 1, 1, 1}
	case "M-D":
		return [7]int{1, 0, 0, 0, 1, 1, 1}
	case "D&A":
		return [7]int{0, 0, 0, 0, 0, 0, 0}
	case "D&M":
		return [7]int{1, 0, 0, 0, 0, 0, 0}
	case "D|A":
		return [7]int{0, 0, 1, 0, 1, 0, 1}
	case "D|M":
		return [7]int{1, 0, 1, 0, 1, 0, 1}
	}

	fmt.Printf("ERROR: GenComp: unknown code given: %s\n", comp)
	return [7]int{0, 0, 0, 0, 0, 0, 0}
}
func GenJump(jump string) [3]int {
	switch jump {
	case "null":
		return [3]int{0, 0, 0}
	case "JGT":
		return [3]int{0, 0, 1}
	case "JEQ":
		return [3]int{0, 1, 0}
	case "JGE":
		return [3]int{0, 1, 1}
	case "JLT":
		return [3]int{1, 0, 0}
	case "JNE":
		return [3]int{1, 0, 1}
	case "JLE":
		return [3]int{1, 1, 0}
	case "JMP":
		return [3]int{1, 1, 1}
	}

	fmt.Printf("ERROR: GenJump: unknown jump %s\n", jump)
	return [3]int{0, 0, 0}
}
