package day03

import (
	"regexp"
	"slices"
)

const (
	eMu = iota
	eDo
	eDn
)

type Instruction struct {
	inst  string
	index int
	typ   int
}

func createInsts(re *regexp.Regexp, s string, typ int) []Instruction {
	idxs := re.FindAllStringIndex(s, -1)
	strs := re.FindAllString(s, -1)
	insts := []Instruction{}
	for idx, str := range strs {
		insts = append(insts, Instruction{str, idxs[idx][0], typ})
	}
	return insts
}

func d03p02(inputString string) int {
	muRe := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	doRe := regexp.MustCompile(`do\(\)`)
	dnRe := regexp.MustCompile(`don't\(\)`)

	muInsts := createInsts(muRe, inputString, eMu)
	doInsts := createInsts(doRe, inputString, eDo)
	dnInsts := createInsts(dnRe, inputString, eDn)

	allInsts := append([]Instruction{}, muInsts...)
	allInsts = append(allInsts, doInsts...)
	allInsts = append(allInsts, dnInsts...)

	// solution isn't the quickest since we have to sort here when we can just read instructions in order
	slices.SortFunc(allInsts, func(a Instruction, b Instruction) int {
		return a.index - b.index
	})

	total := 0
	enabled := true
	for _, inst := range allInsts {
		switch inst.typ {
		case eMu:
			if enabled {
				total += doMul(inst.inst)
			}
		case eDo:
			enabled = true
		case eDn:
			enabled = false
		}
	}
	return total
}
