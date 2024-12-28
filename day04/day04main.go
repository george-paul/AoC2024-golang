package day04

import (
	"AoC_2024/util"
	"strings"
)

// corresponds to each direction (cul = const up left)
const (
	cul = iota
	cuu
	cur
	cll
	crr
	cdl
	cdd
	cdr
)

var dirs = []uint8{0, 1, 2, 3, 4, 5, 6, 7}

const (
	un = iota // unknown
	ye
	no
)

const (
	xRune = 88
	mRune = 77
	aRune = 65
	sRune = 83
)

type PresList [8]uint8
type Cell struct {
	pres   PresList // presence list; holds list of un/ye/no; in the same order as direction constants
	letter uint8    // X=1;M=2...
}
type Cells [dimH][dimW]Cell

var letterMap = map[rune]uint8{
	xRune: 1,
	mRune: 2,
	aRune: 3,
	sRune: 4,
}
var initialPres = PresList{un, un, un, un, un, un, un, un}
var initialPresX = PresList{ye, ye, ye, ye, ye, ye, ye, ye}

func Day04Main() {

	cells := Cells{}

	for rIdx, line := range strings.Split(d04input, "\n") {
		for cIdx, char := range line {
			if char == xRune {
				cells[rIdx][cIdx] = Cell{initialPresX, letterMap[char]}
			} else {
				cells[rIdx][cIdx] = Cell{initialPres, letterMap[char]}
			}
		}
	}

	util.ResultString(4, d04p01(cells), d04p02(cells))
}

const dimH = 140 // input is 140 chars tall
const dimW = 140 // input is 140 chars wide
// const dimH = 10 // input is 140 chars tall
// const dimW = 10 // input is 140 chars wide

const d04inputTest = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

const d04input = `XXXMMSAMXXMASXASAMMSMMSAAMXMAXSAXMSXMXSAXMMXXXXMXAXXXMASMMSSSSSMMMSMSMMSMSMSMMSXMXMMSMMMMAMXSMMMMSSXMASXSMAMSMAXMAMXSXAMSSMSSSMXSMSXXXXAXMMM
AMSMMASAMXXXXAXMAXAAMSAXXMXAMAMXAMMAMAMMSMSMXXASXSMMSMMSAMAAAMXMAMAAMMMSASXXMASAMXMMMAMXSAMMMAXMMAMAMASAMXASAMMMSAMXMMMAAAAAAAMASMXMSMSMMSAM
MAAAMAXXASMMASASMMSSSXAXMXSAMASMSMMSMAXAAAAAMSMMAMXAAAAXAMMMMMASASMSMMAMAMSXMASAMMXASAMXSASMSSMSMSXSMAMXMSXSASAXMAMAAAXMSSMMSMMASAMXSAAAASAS
XSMSMSSMXAAMAXMASMXMMMMSAMXXSMSAMXAMSSSSMSMSMAAMAMMXSXMMSMASASXSXSXMAMXXSAMXMASXXASMSASMSAMAXMAAXMAMMSMSXMXSXMASXXSXMXSAMXMAMAMXSASAMXMMMSXM
XXSXXAAXSSSMSSXMMMAXMASMMMMXMAMMMMXXAMAXAAAAMMMMASAAXASAXMASASXSASASXMAMAASMSAMAMXSMSAMXMAMMMMXMSMMMAMXMAMAMAMSSXMAMMSAMXAMMMAMASAMMSASAMXAA
XMASMMSMMAAAXXXMASXMMMMAAXXAMAMAXXMMAMMMSMSMSXXMSMMMSAMASMAMAMAMAMAMSMSSSMSXSAMXMASXXMAMXAMXAMAXXAXMAMMSAMAMAXAXAAAAMAMMSSXSXXMAMAMXMAXASMMM
SMAMXMAXXAMMMMMXAXAXSSSSSSSXSSSXSASXMMAAAAAAXMMXAXXAMXMSMMSSXMXMAMAMXMAAAAXASAMXMAMXMAAMXXXMAMMSSMMMMSASXSXSAMMXXMXSXAXMAMASAMSSSMMMXMMMMAXA
XMMSSSMSSMSAXAAMSSSMMAXAAAAAXAAASAMXASMSSSMMMAMSMMSAMXXAMAAXAMXSXSAAXSMXMMXMXAMXMMSMAMAMAASXXSAAXXMAMXMXXMAAXMMASMMMMXXMAMAMAMAASAAMASAMXSMM
MAXAXAAAXXSXXMXSAAAMMMMMMMMMMMMMMAMSMMMAAAAASXMAAASAMMSMMMSSMMAMMMMMXMMXSAMSSSMXMMAAASAMSMSAASMSSXSASMSSMMMAMSXXXAAMSSMSMMXSMMMMMSMSASXSAAAM
AAMXSMMMXMXMASXMMXMMSXMMXMAMAAXASAMAXAMMSMMMMASMSMSXMAMXAXMMMAXXAAASMSAXMASAAAMMXSASXXAXMAMMMMMAMXSASAAMXAMMASMSSMMSAAXAASMMMASMMAXMAMAMMSMM
SMSAXAXAMSAMSAAMAXMASASXMMXMSMSXSXSMSXSXMXXXMXMXAMMMMXMSASXASMSSSXSAAMSSXMMMSMMAXSAXXSMMMMMMXXXXXAMAMMASXXSMAMAAAMSMMSMMXSAASASAMSSMAMAMAXAM
MAMASAMXMSASXSXMASMASMMAMMAMAAXMSASXSMMMXMAMSMSXMAMAMSMSSMMXSAAAXXMMXMASAAMAMXMSMMASMMSASAAMSSSMAMMSMMAXAAXMAMMMXMXAXAAXASMMMMXAMMAMAMAMMSXM
MMMMMXMSMMXMAXXMAXAAMASXMSASAXSAMAMMXAMSSXSAAAXAXMSMMAMSAMAMMMMXMXXXXMASMMMAXMXMSMAMAAXXMSXSAAAXXMAAMMMMAMSMXXASASMSSSSMXSXMASXSMSAMXSASMSXS
MXAAAXMAMXSMMMMMMSMXSXMAXXAXASMAMAMAASAAXAXSMSSMMXAXSSXSAMXXAXSSSMMSAMAXXAMAXSAMAXSSMMSSMMMMMMMMXMXXSAMSAMXMASXSASAAXMAXAXASXSAAAMMSMMASAMXM
XSSMSSSMSMMAMXMAXAMMAASXMMSMMMMAMXASAMMMMAMAXXAAASAMXMASMMSMMMSAAAAMXMMMSXSXXSASMXMAXMAAAASMSAXAMMMASAMSAMAMMMMMMMMMMSMMSSMMAMXMMMMAASXMMMAS
AMAMAXXXAAXAMXXXXMSAMXMAAAXAXASMSXAMXXXXMAXXMMSMMSMSSMAMXAAAAMMXMMMSMXAAXMAXMXXMXMSASMSSSMSASAXSASAAMAMSAMAMAAAAMSMMAMAAAAXMXMSXAAXMMMAAMSAS
MSAMMMMSMSSMMSSSSMAXXASXMASXSMSAAXXXMAXSMSSSMMASXXAAMAMXMMSSMSASXSXAMSMMXAMXMMSXMAXASAMAXAMXMXAMMXMSSSMSASMSSSSSMSAMMXMMSSMMAMAMSMSMMSMMMSXM
XSXMAAAAAMXXSAAAASASMXMXMXMAAAMXMMXMAMXSAMAAXXAMMMMMSMSMMMMMMMMMAASAMAMXXSASAASAXMMMMMMAMXMSSMSMSMXXAMAMXMAAAXAXASAMMASXMXAXXSMXAMAXMAXMMMAM
XMAMSSXMSMMMAMMMMMMMXAMXAAMSMMMAXAXSSMAMAMXMMMSMSASAMAAAASASXMMMXMSXMSXSXXASMASMMXXXAAMXSSMMAAAAAAMMMMAMSMMMSMMMAXXMSAMASMSMASXSASAMSXMSASAM
MSSMMMMMAAAXMXSAMMSMSMAMSXXXAMSMSXMAMXASXMAXXXMASASAMSMSMSAMMXSAMXMMMMMMAMMMMAMMSAMMSMMAXAAMMXMSMXMMAMAAMAMXMXXSSMMMXMMAMAXMAMAXAMAXMMASASXS
MXMXSAAXSSMSMAXASAAMAMAMXMMXSMMXSXXSXSXSMSMSXMMMMXMXMXAAAMAMSASAMAMAAAAMAMAXMMSAXMAXAMMMSSMMAMMAXXASXSSMSAMAXAXMAASAAAXSMMXMMMMMXMXAAAXMAMAS
XASAMXMMMXMAMMXMMXXSMXAMAAMAXAAASAMXMMMSAASAASAMXAMASMSMSMMXMASMMASMSSSSMSMSMXMASMXSASMXAMXMMMMAMMXMXXAASAMASMSSXMMMXMSAAXXXXAAXXMMSMSAMXMSM
SMMMSAXAMMSAMSASXSMAMSMSMSMASAMXSMAAAXASMMXSXMAASXSASAAMXMXXMXMMSASAMXMMMMXSXAMAMXAMASXAMMSASASXSXMSMMMMMAMXAXMAMSMSMSMXMXMAMSMMASAAAXXMXMAM
AXAXXMXASAXXMSXXAAMMMAAAAMMMMAMXSXSSSMMSASXAMMMXSXMASXSXMMSMXSAMXAMMMAXMASAMMXASMMXMAMMMMASXMASAAMXAAAXXXSMSXMMAMAASXSAMXAMMXAMMAMSXMXAMXMAM
MSMSAAMXMAXMMSXMSMMXSXSSXSAMXAMMSAMXMAMSAMXMAASXXAMAMAMXSXMAAMMMMXMSSXSMAMAMAXAMASXMASMXMXXAMXMMMASXMMSSXMAAXMSMSMSMAMAMSSSSSMSASMXXMSAMXSAS
XAAMXMASMMAMXMAAAAMAXAAAXSASXSMXMXMAMMMMXMXXSMSMSSMASMMAMAMMMXASXSMXMXMMASMMSAASAMXSAMXAXMMMMAXXMMASAXAXAMSMMMAMAXXMXMAMAAXAXASAXAXXAAAXMMAS
SMXMAXMAAASMMSSMMXXAMSMMMSAMAXMXMSSMXSASXMMMXXSAMASAXXMXXAMAMXAXAXMAMAMMASMAAMMMMMAMASMSSMASMMMXXMASXMASMMMAXMAXMSSMXXMMMXMAMMMMMMMAMSMSSMAM
XSSXSSSMMMMAAAMMSAMSXXAXXMSMAMSMMAAMASMSAAAAMAMAXAMMSMSMSMSASMSMSMSASXXMXSAMSXMASMMSMMXAXMASASXSXMAMAAXAMASAMXXSXAAXSMMMSMMXAXAAAAAMMAAXAMSS
MXSAAAXXSXMXMASAXMAXAMXMSXMMMMMXMXMMMXMSXMMSSXXAMSSMAXAAAXMASAAAAASAMAASAMAXMMSAMXMMASMMSMMSAAAXXMAMMASMSASASXXAMSSMAAXAMAMMXSSSSSMSAMXSMMAX
XAMMMMMSAAASXMMXSXSMXMAAXASXMASMSSMXXMXSMMMAMAMSSMAMXSMSMSMMMMMSMXMMMSMMASXMAXSSSXXMASAXXAAMMMMMMMSXSASXMAMMMMXSAMXMSMMMSAMSMMMXMXASXSXMAMXM
MMMXXXXMMSMXASMXMAXAMXSMSAMASXSAAAXMXMXMASMAXAAXAMAMMMXAMXAAAXXMMXMXMAMMAAXSMMMAMMXMASMMSMMXSAMASXXAMXSAMSMSAMMAAAXMAXMAXXSXASMSSMXMASASXMAS
AASMMMSXMXXSMMMAMAMAMMAXXAMAMSMMSMMSMSMSAMMSSMMSAMXMAMSSSMMMSXMXSAXSAMXMMSMAAAMAMSASAMXMXAAAXASASMMXMAXXMAMSAMAMMMMSMSMMSMMMMMAAAXMXMSAMMSAA
XXAAAAXXMXMMMASMMMSSMAMSXMMMMXXAMXASMSAMXSAXXXMAMXXSXMAXSAMXMXMASMSMXXMASMAXMMMXMMAMXAAMMSMMSAMMMAXSMSSMMAMSXMMSAAAAMAMMAAMMMMMMMMMAAMAMAMSS
MSSXMSXXMAAASXSMAAMAMMXSMMAXXSMSSMMSAMSMAMASMMSMASXAMMSMSSMAMAMXXXXMAMMXMASMMSMSMMAMMSMSAMAXXAMXSSMSAMAMSMXXXAMMMMSSSSSSMSAAASXMAMXAMMAMXXAX
XAAMSAMSMMXASMXXMXSAAXXMAXMSXXAXSAAMXMAMXMXXAAXMASMMXAAAXAMSSXMSMSMXSMMXXAAMAAAAMXASAAAMASASMMMXMXAMSMAMMMSSMSMMSMXAAXAXAAMXMXASXMSASXAAMMSS
SMSMMASAASMMXXAXMAMXSMSMSAXAAMMMSMMSXMAXMMMSMSSXXXAXMSMSMXMMAMXAAAMMXAMXSSMMSMSMSSXSMMMMXMXXAASAMMAMXMSSXAAXMASAAAMMSMMMXMMSSSMMXAAAMXXMAAAM
XXAAMSMMAMAMAMMSSSXAAXAAMMAXXMMASAMXASXXSAAAMMMMAMSXXMAMXAXSMMSMMMSASAMAMAXAXXMAMAAMAXXSAMASMMMASXMXMAAMMMXSXMMSSXSAAAASMXMAAAXSMSMMMXMSMMXX
XSSMMMAXMSSMAXXAAXMAMXMXMXXMSXMAMMMSMMAASMXMMAAMAMXMSMMMSXXAXMXSAMMMSAMASXMMMSMSMSXSAMMAMMXAMXSXXAXAMMMSXSMXAXXXXMMMSSMSMAMMMMMMAAXAXXAAXMSS
XMASMSMMMAMXMSMMSMAAAASXXSAAMAMSMSASMMMMMSMASXSSMSAXXAXASMXMMAASXXAXXMSMSXAAAAAMXMXMASAAXSAMXMAMXMMXSAMXAMXSMMMAMSAAMAMXMSSMXSXMSMSSSMSSSXAA
MMAMXASXMAXAMAMAAXMMXXSAAXMASAMMAMASMSMMAXMAMMAAASMXSAMXXAMXXMXSASMXMXAAMXSMSMMXAMXAXXMASXSXXAAXMXSSMMSMAXAMAAXASMMMSASMMMXXMXAMXAAMAXAAMXMM
SMAMSMMXSASXSSMSXSXMSSMMMMXAXASMMMMMMAAMXSMASMMMMMMXMXASXSMSASMMAMAASXSSXAXXXXAMXMSXMSSMSAMSSSMAMAMAAAXMSMMSSMSXSAAXXASXSXAMMSMMMMMMSMMAMAXX
SMMXAAXMMASAAXMMASAMXAAXMASMSMMMXAXASXSMAXXAXXXMAMXAMXSMSAAXASAMXMSMMAAMMMMSAMSAMXSAAAAAMAMAMAMAMASMMMSXXAAAAAXXSMMSMSMAMMMSAAXAAAMXXXSASXSX
MAMXSMXAMAMMMMAMMSMMMXMASASMAXMASMSMXMAMMMMSSMMSAMSASMMSSMMMXMXXAAXXAMMMAMAAAXMASXSXMMSMSSMMSASMSAMAAXAMSMMSMMMXXMAMMAMAXMASMSSSSMSXMASAMMMM
MAMXMSXMAMXMMSMSXMMSAXXXMASASMSAMXAMAXAMSAMXAAXMMASAMASXXMASMMSSMSMSXXXXSMXMXMXMXXMASMMAAAAASAMAMXXMMMMAAMMXXMMMAMMXSASXMMXSAXXAMXSAMAMSMSAM
SASAAMMSMSMSAAXSAAAXMAXMXSXXMASAASMSMSMSSSSSMMXAMAMASXMMASASAAAXAAAASAMXMXSXSMSMMXSAMAXMSSMMMMMSXMMSAMXSMSMAMXAMASXASASXASAMMMXMMXSAMASAXSAS
SSSMSMAAAAAMSSXSAMMSXSXAXXXXMAMSMSAMAAXMXAMXAXSXMXSAMMASAMSSMMMMXMXMMXMAMMAAAAAAMAMASAMXAMXSXXAXAMAAAMMMXAMSMSXSASMXMAMMXMASAMSAMASAMXXMASAM
XMXAMMSMSMSMAMASXMAXAMMXMSMSMSMMMSMMXMSXXSMSSMMASAXMASXMASAXAMMSAMAXAAMAMAMSMXXMMASAMXXMASASXMSSXMMSSMASXMMMMMAMASAMMSMAMSSMAMXAMASAMXSXXMAS
MAMXMAMXMMMMXXAMXMXMXMAAMAASXMASAMXMSMMMMMAMXASXMASXMSAMXMMSXMASASMSSSSSSSXAXAXSSMMAXXXMAMXSAXAMAMMMAMASAXAAAMAMXMXMAAXXAXAMSMSSMXSAMSAMXSAM
XSXSMMSMSAAAXMSMMAAXAMMXSMMMAMAMAMMMAAAAAMMMSMMMMMMAASMMASXMAMMXXAAAXXAXAMXSMXMAAASMMAXMAMMMAMXSAXAMAMXSMSSSSMMSXSSMSMSASXXMMAAXXXXXMAXXAMXS
XXAMXMAASXMSSXMAMSMSSSMAXMASXMSSXMXMMMMXXSMMAAAMAAMMMMXMXSAMXMXXSMMMSMSMMMMAASMSMMAAMAMSSMMXSXMSMSMSXSAXXAMMXAAXAXMAAAAXXMASMMMXMSMSMSAXXMMA
XMSMAXMXMMSXMASMMAMMAAMSMXXMAAAMXSMSMXMSMXAAMXSSSSSXSXMXASMMAMXMMAAAXXXAAAAMSMAXAMSAMAXAAASAXAAXAXAAXMAMMMSAXMMMSMMMMMMSMSMSMSMSAAXAAMMSMSXA
SAMMSXSAAXMASAMXMAMMSAMMASXSMMMMMXAAAASAMSMMXAXMAAXXMASMAMMMMXAASXMSSSSSMSXXMMMMXAMAMMMSMMMAXMMMXMXMXMXXSXMXSSXXAAMASXMMAAAMAAAMXMMMMXXAAAMX
SASAXXSSXMMXMXXMSXSAMAXSAMXAASAAMMSMSMSAMMXAMSMMMMMSXAAMSMSASXMXMAAXAXAAXMASXXAMSXSSMXMMMSMMMAXAMXAXASAMXASAAAXSASXASASMSMSMMMSMAXXAASMMSMSA
SAMASMMMXSSSMMMMAXMAMAMMMXSMMMSMSAXMXASXMSMSXMASAMXSMSSXAASAXXXXSMMMXMSMMSMMASMXMMAAAMMMAAAASMMSASMXAXAMSMMMMMMSAMMXMAMMASAMXAAMSMMXMSAAAAMM
MAMAMAXXXAAAXAAMXXSAMAMASMSXSXAXMXSAMMMAMMAMAXMSASAXAMMMMXMAMSMXMASXMMMXAMXXMMMAASMMMMXMSMSMSAMAXAMMMMSAMXXAXSAMASMSMAMMMMSMMSMAAAMXMSMMMXMX
MXMXMAMSMMSMSSSMMASXSMMMSAMAXXXMXXMMMAMMMMAMMMXSXMMMXSAAMASMMSAAXXAXAXMMMMXMXAXSMMAAAXXMXXXXSMXSMMXAAAMXMXSMSMASMMAAXSXXAAXXXXMSSSMAAMMXXSXM
SASMMSMSAAXAAAMASAMXAAAXMMMSMMMMXMSXSSSSXSASMSAMXSMAAAMAXAXAAXSSSMMASMMXSASMSSMXASMMMMXSAAMXMXMXAMSSSMSASASMXXAMAMSMSMMSMSSMSXMAMXMMSXMSMMMA
SASAAMASMMMMMXSXMASASMMSAASAMMAMAXMAAXAMASAAXMASAMMMXMSSMSSMMXAMAAMMXMAAMASAAMAXAMXMAMMSMMMAXAASAMXAXASXMASAMMSSSMXSMAAAMMAAAXMASAMSMXMAAASM
MXMMXMAMASXXAXMAMXXAMAAMMMSASMSSMSMMMMAMAMXMASAMXSMMSMAAAXXASXSSXMMXXMMSSSMMMMSMMXAMAXAXAAXMSMMMXAMXMXMMMXMXMAMMXXMAXMXSMSMMMSXAXAXAAAMSXMXX
MASAXMASAMAMXXXAMXMSMMMMMXSXMAMAXAXMAAXMXSXMMMMSAMAASMSMMMSMMAMMASMMMSAAXAASXAXMAMMMMMMMMMXXAMAXASXSMXMAMASMSMXXXMSSMSAXAMSAAXMMSSSMSMXMMSSS
SASAMXAMMSAMSSMSSSMASAAXXAMAXSSMMMMSASMSAMAXMAXMAMMMSAAAAAXXMXMSAMAAAMAMXSMMMAMSAMXASASXAXMASMXSAMAASXSMSAMAAMXMSAAAAMASASMMMXMMAAAAMXAAMXAA
MAMMXMASMSAXXAXXAMXAXMMMSSMMMAAMMXASMMXMASMMSSMXAMSAMXMMSMMXMXMMMSSSMXSSMMSSXMASAXSAMAXMXXMXMAXMAMSMMXAAMAMSMMAAAMMMSMXMMMAAMSSMMXMMMXMXMMSM
SSMSXSASASMMSSMMAMMMSAAAAAMXXMAXMMMXAXSSMMAAAMMSMSMSSSMXXAMSSSMAMAMAXAMAXAXSAMXSAMMSMSMSSSSSMMMMAMXAXMMMMAMXASMSSSSXMXMMMSSMSAXAXASXSMSAAMAS
SAAXAMAMAXMAXAMSSMAXSXMXXSASMXSMAMASMMXAASMMMSAAXSXXMASAXAMXAASMMASAMXSMMMSSXMAMXMAXSAASAAAXAXASASXMMMXAXMMSAXAAAXMASMMAAAAMMMSMMMAAAAMMMMAS
SMMSXMAMSMMSSMMAASMMMASXMAMAAAXSAMXAXMSSMMXXSMMSXMAAAAMMSSSMSMMMMMMXSAXXAXAMAMXSMMXSXMXMMMMSXSASXSXSASXMASASXMMMSMSAMASMSSMMAXAAAXMMMMMXSMSM
XSAMXMXSXAAAAXMSMMXASAMAAMXMMAMMXMAMXXAAASAMXAMMXSAMMSSXAAMMASASMMAMMMSSSMXXAMXXAXMAMXXMXMAAMMAMASASASAMXMMMXMXAAXMASAMXXAXSAMXSSSXSXXMAMXAX
MMMMXMXMXMXMMMMXXMMMMAXAMXSXSSSSXSAXSMSSMMMASMMMAXASXAAMMSMSAXMSAMXSAAXAMAMSMXMASMMSXMXXAXMSXMAMAMMMXMMSXMXSAMMSMSAASASMMMXMXXMXXMAXAMMMSSSS
XMASAMXXAMSMSSXSASAASMSSMMSAMAAAASXMXXMAXXMMMSAMXXSMMMSXAMXMMSMSXMAMMXXMMAMAASXMXAAMAMASMSXMASAXAMXMASMAAMAXASMAMXMASAMAAAMMMMMMAMAMSMMSAXAM
MXAAAAXMMAAAAMAXAMXMXXAASAMXMMMMMMMSMMAMSSXAASMSAXMAAXXMASAMXAASXMASXSMXSASXMXAXSMMMAMAAXSAXMMMXAXXMASMSSMASMMMASMSAMXMSMMSMAAAXAMXXMASMMMMX
AMASAMXAXMMXMXAMSMSXMMMMMMMXSXMXXSAMASXMAXSMMMAMMAXSMAAXXSASMMXMASAXAAXASASXXXMMSAMSXSSMMMAMSAMSAMXMAXXAMXAXMASXXAXXAAMASXAXSSMSMXSAMXMAAXXM
XAXXXAXSSSSMMMSSMAAASAMXMSAXMASXMMXSAMXMSMMXSMAMXMAAXXMAXMMMASXMMMSMSMMMSXXMMXSAXAMAXMAASMSMSAMAAAAMASMMXMAMSMMASXSAMXMAXXMMMAMSXAAAXAXSMSMA
MSAMMXMMAMXAAAXAMAMMMXSAAXMASMMASAAMMMSAXAMAXMXXAXSXMMXMMASMSMAXSAMXXAAAXMMXMAMAMSMMMMSMMAXMXXMMSSSMMSASXSSMXAXXAASASMMSSXSAMXASMMMSMSXXMSMS
XMASAASMMSSSMMSAMXXAXASMSMXXXASMMMMXASMMSMMMMASXSXMASMAMMMMMAMSMMASASXMSSXMAMAMXMXASMMMMMMMXXSAMMXMAXSAMAAAAMSASMMSAMMAAAMSXMSSXAAXXAXASASAA
XSAMXMSAAXMASAXASMMSAMXAXAMXSMMASXSSMSAXXMAMMAMAXMXXMSASASAAAAXASXMMSASAXASXSAXSASAMMAAXAASMMMASMMXXMMAMMMMSMMAXXAMXMMMMMMXXASASXMSMAMAMAMSM
XMSSMXSMMSXAMSSXMAAXMXMMMMAXMASAMXMAMSXMXMASMSMSMSMMMXXSASAXXSMXMAAXMMMASMMMSMSAMMSMSSSXSMXAXSAMMASXMSXSAMXXAMXMMAXMSASASMXMMSAMXXXMAMAMXMAX
SXAXMASAMXMMSAXMMMMMAAXAAXMMMMMAMXSAMSASXMAMXAAAAAXSXSMMMMXMSXXASMMMXXMAMXSXSSMAXAXMXAAMMASXMMASMAMAAAAXXAMMMMMXSXSAAASAAMAMXMAMXMSSMMAXAMAX
SMMSMMSAMAASMXSMSAAXSSMSSXSAASMSMAXXSMMAMAMSSMMMSMMASMASMAAMMASMSXMMMXMMMAMMMAMSMMXXMMMMXAMMASAMMSSMMMSMSMSASMSMMMAMXAMXSSXSASAMAMAAXSMSSMSS
XAXAXXMXMAMSAMSASMMMAMAMAMXXAXAMMMMMMXXXXXXAXAAMAMXMASAMMXXAMXMXSMSAAXSAMASASMMMASAASXSSMSMXXMASAMMXAXAAAASAXMAAASMMMXXXMAASASASXMXSMAAAAAAX
SSMMMXMAXMMMAMMMMSXXAAMMAMXMSMMXASAMXSXMMAMMSMMSAMAXAMXSASXSXMAXXASAMXMXSASMSXXSAMSSMASMAMXMASXMASXSAXMMMMMAMSMSMSXXMAMXMMMMXMAMMMSXXMASMMMS
XMASMASAMXASXMXSMMXSMMXSXMXAAAXSMMAXAMMXMAXMAAMSASXSMMMSASAXAMXSMAMMSAMXMASAXAMSXMXMMMMMAMAXXMAXAMMMSXASMMSSMAAMXMMMSAMMXMXMAMXMXMAXAMXAMXXX
XSAMSMMAXXAXAXAAAAAXASAMMMXXXXMAMMSMAXAAXMMSXSMSAMAAAXAMXMAMXMAXAMAASAMMMSMAMMMMXXXXASAMMSSMXMMMMXSAAMAMXAAMMSSMASAXSAMAAAASASASXMMSAMMSMMMX
MMMXASMASMSMAMSSMMXSAMMMXSAMXXSASAMSSSSMSAASAXMMSMXSMMMSMMAMAMAMMMMMXMMSSMXSAMAASMMSMMASAAAMXXSAMAMSSSXMMMMXAAAMXMSXSAMSXSMSAMAMAAXMXMAXMASM
SAMXAXMASAMMSMXAAAXMXMAMAMASAASAMMMAMXAMXMAMAXXAAXAMXXXAASXSMSASMAXXSMMMAMAMASMXSAASMSMMMSSMMAMAMXXAMXAXMSSMMSSMXXMMMAMXMMMMMSMSAMXSAMXXXAXA
SASAMXMXMAMAAMXMMMXXAXAMASASMAMXSAMXMMSMXSSMSMMSSSMMASXXXXAMXSASXMAMAAASAMXMAMXSSMMMAAXMMAAMAAXAMXMASXSSMAMXAMXXXSAAXAMXAXXAAMXAAXMMMSMSMMSM
SXMAXXAXMMMSSMASASMSMMSXMMASXXMAXXMSAXAAXAAAXSAMXAAAXSASXMMMAMMMMMSSSSMSXSSMSSMMXSAXXMASMSSMSMSMMMSMMAAAMASXMAAXMXMSSSSMSASMSSXMMMAAMAAAAAAX
XMSSXXSMMSAAAXMSAXAAMAXAMMMMMAMXSAXAASMSMMMMMMMMSSMSAMAMMAXMASAAAAAMXXAXAAAAAXMMASXSXAAMAAMXXXAAAMASXMMMMASAXMMSMSAMAAAAXXMAMXMXSXSSSMXMSSMS
SXAAMXMAAMMSMMMMXMSMMASAMASASXSMMXMMXAMAAAAAAXAAMAMAAMSSSSMAMXMSSMSSSMMMSMMMMMSMAXASAMXMMMSMXMMXMSASMSAXMASXMASAAXAMMSMMMSSSMMAMXAMAAAXXAMXX
AMAXMAMMMSXXASXXMAAXMASXMASASAAXXAAXMMSMSMMXMXMSSMMSXMAAXMAASXMXMAMXAMXAXXXXXAMMSMXMXSAAAAAASMSSXMAMASMSMMMMAMMMSMXSAXASAAMAAXSMMSMSMMMMAXXX
MMAMSASMMMASAMXMASXSMMXAMAMXMMMMMMSSMMAXXMXSMSMAXAAXAMMXMXSMMAXXMXMSMMXSXMMMMSAAXAMMSSMSMSSXXAMAXMAMXMXAMAAMSXAAAXMMASAMMSSSMMMAAXAAXSXSSMSM
XMAMMASAAMMMMMSAXMXMAMSMMSSMASXAAAAAAXMXMSXSAAMASMXSXMMXXAXMSMMMSAMAXSMMMSAAAXMASAAXXXMAMXMAMSMMXSASMMMMSSSSMXMSSXXMXMMMAAAAMASMMMMMSXAAAAAM
XSMSMXSMMXSAXAXSMMMSAMAAAXAMASMSMSSSMMSMAAAMAMMMXMASMAXMMSSMASXAMASMMMAAAXMMMXMASXMMMMSMSAMXMMAXXMXSXMXXAXMAMAMAXMXSSSXMMXMMMXMXXASMXMMMMSMX
MSAXMAXASASXMMSAMAAXASMSMSXMASAMAXAMXAAMSMSAAMAXASASMMMXAAAMAMMMSAMXASXMMSXSMXMAMXXAAAXXMXSSMSMMMMMSAMAMMSSSMMMASXXMAMMMSAMSXSMMSASMSASXMAMM
AMMMMASXMAMXSXSXSMSSXMAAMMAMAMAMMMSMMSSMXAMXSSXSASAMAMSMMSSMXSAMAMASXMASMSAAAXMASMSMSMMMMMMMAAMAXAASXMASXAXAAXMAMXMMAMAAMAXSAMSAAAXAXASXSASX
SAXXMMMMMSMMMAMAXAXMMMMMASAMAMSSXAXXXAXMAXSAMXAMMMMMAMAAMXXMXAMSASXMASAMAMSMMSSXSAMAMXMAAAASXMSSMMXSAMXSMXSSMMMASAXSASMSSSMMAMMSSSMMMSMAMASX
XXMSMAAXAXAAMAMXMAMMSAMXMSSSXSAAMMSMMMSMMMMAMMXMMAMSSSSSMMSXMMMSAMMSAMMMSMXMAXASMAMAMSMSXSMSAXAXSMAXMAMMMMMAMSSMSXXMMMAXAMXMMMAMXMAXSAMSMAMX
MAMASXXMSSSMSMSAMXXMAMXAXXMMXMXMAAAMAAAAAXSXMXXXMXMAAMMMXAMASXXMXMAMXXXAMASAMXMMSXMSASMXXXXMXMAXMMXSXSAMXAXAMXAXSXSASMMMAMSSMMASXSXMMAMAMXSX
MASASMSMMAMXXAXAMXXXMSMMMSAMXXMSMXSMMSSSMMSASAMSASMSMXAMMMSAMAXMMMMMSXMASMMAXAXAMXMMXXAMASMSMSSSSMAMAXMAMSMSXSAMXASAMSXMXMAAXSXSAAMSSSMMXASX
SMXXMXAMXMAXMSMSSMMMMAAAAXASMSMAXXMAXAXAXAMASMASMAAAMXMMMXMASMMSAXAASAXASAMXMMMMXAXSSMSMAMAAASMAAMXSASXSXMAMAXAAMMMAMAXSXSMSMMASMXMMAMSXMMMS
XXXMSSMSMSSSMXXMAXAASMSMSMXMAXXAXMMXMASXMMMAMXXMXMSMSASAAMAAMXAXMSMSSXMXSAMMASAASXSAAAXMAMSXSMMSMMAMASAMXMSMMSMMMXASMMASAMMMAXAXMAXMAMMMXAAS
SXSAAAAXMAMAXXMSMASMSMMAXMMMSMSSXSAASAMASXMMXSAXAXMASASXSAMMSMXXXAMAMMSASXMASXXXAMXMMMXMMMMMXXAXASAMMMXMSMXAXAASMMXXAXAMAMSSMMSXSXXMASXSSMSS
SASMXMMMMMSSMSMAXMMXMAMAMASXXAMAAMSMMASMMAMMAXMSSSMXMAMAXXXSAMMSSXSAAAMASMMMAMMSAMXXAMXMMAMMMMMSMMAXMAXSAMMMMSSMAMASMMMSMMAAXAMMSAAMASMAASAM
MMMSAASAMXAAAXSASXMASXMMSAMXMMMMMMXMSAMXXAMMMXXAMXXAMASMMSMSASAMMASXXAMXMASMAMASXMXMAXAAMASMAMMAMMXMMAMAMMXSAMAMMMXMAMAAXMSMMXSAXSSMASMSXMSS
MXAXSAMXSMSSSXXMXAMASMAAMAMXAXAXAMAMMXSASMSSSSMSXSAMSAMAMSAMAMMSMAMXSAMXSAMMAMXSASASXMSXSASXSMMASXSMMSSMXAAMMXAMXMMMXMSSXMAAASMMMAMXAXAXAMXM
MMSMXMAXXAMMXMAXSAMXSXMMMAXSSSMSMSASMAMXSAAXAAAMASXMMMSAMMAMSMMSMMXAAMSXMASMSSMSAMMSAAAAMAMMMMSASAAMAMXXMMXXSMSSSMAAAAMAMSSMMSAMMAMXAMXMAXAS
SXMAMMSSMAMAXSAMXSAMXAMSMMMMAAMAAXXXMASMMMMMSMMMAMAXAMMXXMXMAAXSASMXSAAXXXMAXAXMAMAMMMMXMAAAAXMAMXMMSSSMSSSMSXAAMXAMASMXMAXAAMMMMXSSMAMSMSAS
AAMSMAAAMXMAMMMSSXMMXMASASAMSMMMSMSMMAXXAAMXAMSMMSAMXSSSSMSSSXMSAMAMAMMXXXMASMMSXMSXSASXSXMMXXMAMXSXMAAXAAAAXMMSMSMSAMXSMXXMMSXMMAMAMSAAMMXM
MSMAMMSSMXMASAXMAMXXAXMSAAMMMMMMMASAMSSSSMSXSXXAAXXAAXAAAAAMMAXMMMAMXMXMAXMASMAMXAASAAAXMMSSSMSMXAAAXXXMMSMMMAXXAAMMXMAAXMASAXAMMSSXMMMXMMAX
MXXMMMXMAXXMSMSMXMXSASMMMXXAMASAMMMAMAAXAASAXASMMSMMMMMSMMXAXXMXXXXMAXAAMXMAXMAMXMMSXMMSXMAAAAXMMSSMMMMAXXAXXSMMMMSSMMSMAMXMASXMAXMMXSXMASXS
SMMSMMMXMMSMMAAAASASXSXAXMMXXAMASMSMMMSMMMMASASAAXAAAAXXXMSXSAMXSXASMXMSMMMSSSMSMAXMXMXXAMMXMMMXAMAAAAASMXMXMAXAMXMXAAAMAAXAAMAMSSMASAASXMAS
SAAAAAAMXMAAMSMSMMAMASXXSASXMSSMMXAAAAAXAAMXMASMMSSMMXXAAXAASASAMSMMSAAAAXAMAAAMAXAMAXSSMMSXSSSMSSSMSMSAMASAMMSSSMMSMMSSSSSMXSXMAXAXXSMMSMAM
SMMSSMXXASXSMAXXXMSMSMAMSXMAXAAXMSSSMSSSMMMAMMMAAAASMMXSMSMXSAMXMAXAXMSSXMXMSMMSSXMSMSMXMXMAXAAAMAXAXMXAMXMAMMAAAMMAXXXMAAXMMMMSMMSXMMXAXMAS
XAAXMMXMMMAAXAMXAAAXAMSMSASMMSAMXAAMMMAAXSSSSSSMMSXMMSAAAXXXMXMXMSMMSXXMASAMXXAXXMMAMXMAXAMSMXMSMMMMMMSXMASMMMMMMMSMSMXMMMMAMAAAXAMAMXMMMXSS
SMMSMAAXXMSMMMMMMMMSSMXXSXMXMMMSMMSMAMMMMMAMAAMMXXXAAMXMMMMXMAMAXAAXSAXSAMAXMMMSMXSAMSSXSASAMXXAAXAXAXMASMSXMAMMMXAAAAMXMASASMSSMMSSMXMXSMXS
XMAXMSXSAXXXMMMAAAXXMASAMSMAMMASXAXXXSSMAMXMMMMSAMMMMMMXMASASAXXXMSMXAXMASXMXAMXAAXAMAAASXSAMXMSSSXSAMMXSAMAMAMSSSMSMSAASMSXSXMAAXAMXXSAMXAX
SXSXAAMMXMXASASMSMSAMXMAMAMAXMAXXAMXXXAXXMASXXAMASXSAMXAMMMAMMMSXSAMMXMSAMMASMXMSMXMMMSMMAMXMAAXMMXMXSSMMASAMASXAAAMXXMAMXMAXAMSMMASXXMMMMXS
MAMMMMSSMXSAMASXAASMMSXMMMXSXMXSAMASMSMMMSMSMSSSMMXSMSSSSMMSMMASMSASMAMMSSMXMSAXAMASAMMAAXMSSSSMMSAMSAAASAMXSXSMMMAMMXXXMASAXSMMAMXXMASAASMA
MAMMMAXAAXMXMAMMMMMSMAASXMAXASAMXSXMAAXAXAAXAAXMAMXMMXXMAXAAAMAXASAAXAXXAXXXAMMSXSASASXMXASXAAXAASASMSSMMMSAMAXAXXMXSAAMMMSXAMXMAXAAAMMMMXAM
MSSXAMSMMMMAMAMAXAAAMSMMAMASXMAMXSAMSMSSSMSMMMMXAMXXXMASXMSSSMMSXMXMSSSMMSSMAMXMAMASMMAMSMAMMMMSMMMMXAMMAAXASMSMMMMAMSSMAXXMAMAMAXSXMXASMSSX
XAMMSAMXAAMASMSMMMSXMAXSSMAMMSMMAMMMXAXAMXMASMSSSXSAMMMAAAAXXMXMMAAXXXAXXAAXSMMMMMAMASXMASAAASAMXSMSMSSSMSSMMAAXAAMAMAXMXMAMMSAMAAMAMSMSAAXA
MSSMMAXMXXSASAAXXAAMXMAMMXAMAAAMAXSAMXSMSASXMAAAAAXMMAXMMMMSXMASXSMSASAMXSMMXMAXAMMSMXMXAMMSMXAMASAAMXMXXXXMMSMSXSSXMASXMXMMMAXMMXSAMAAXMMXA
AXAASAMMAAMAMMMXMASMSSXSASMMSSXSAXMMSXMASASAMMMSMSSXSAMSXSAMXMAXAAAXMMMSMMXXXASMMXMAMMSMSSMAXMAMAMSMSXMMSMSMAXXXAXXXMAMAXASASAMXSASMSMSMXXSS
SSMMMAXMMSSMXXSSMMMMXAAXAAAAMAMXMSMMSXMAMXMAXMXMXMMAXAXSAMMSMXSMSMMSMAXAASMSSMMASMSMXASXMAXAMXAMXXAMAAXAAAXMAXSMSMASMAMSMMSAMAXXMASAAMAXSAXA
MAXMMSMXMMAMSXMAAXXSMMMMSMSXSXSMAAAXSAMMXSMMMMAAAMMSMSMMAMASMAMMMAAXXASMXMAASXSAMAASMXSAXAMSSSSSMMAASMMSSXSASXSAMXXMSSXMAMMMMMMMMAMXMSASMAMX
AXMXMAXXMMAMSASXMMAXAASXMAMMMMSASMSMSAMAAMASASXSMMAAAAAXXMASMXMAXMXMMMMXSMMMSXMAMXAMMMMXMSSXAMXAASXMAXAMAASAMXMSMSMMXMASXMAAMASXMAMAXMXMMAAX
SMSAXXSMSSSMSAMXMXMMXMSAMMMMXAMXMAAASXMMMSASAMMAASXMSMXSAMASXSAMXSXMASXMXXAXMASMSMSMSSSSXMXMMMSSMMMSMMXSMXMAMSMMMXAXAMAMXMMXMASXXXMSMMMMXMXM
XAXXXSAAAAMAMMMMSAMXSAXXMAAMMSMSMXMMMXAAMMMMMMMMMXMAXAAXXMMSMSAMMSAXAMAMMSSMXXMAXAXAAASAMXXMXAXMMSXAXMASMMMAMAAMMMSXSMMXMSMSMASMSAAAAASMMXAA
MAMSXMAMMMMAMAAASASAAAXSSSMSAXMSAMXMXMMSSMXMAMXXAMMMMMMSMSXSXSXMASAMSSXMAXAXSSXAMSMMMSMMXMASASXMASXXMAMXAXSASXSMAAMXMAMXASAAMXMASMSMSMSAASXS
MAMAAAASMXSSSSMXMMMMMMMMAAAMASAXXXXAAXXAMXAASMXSASMXSAMXXMASAMSMXMASAAAMXMMMAAMSXMAXMAMXMAXMAXXMASMMSSXSSMSAMXAMXMMASAMXSMSMSAMMMAXAMAMMMMAA
SXMXXMXMAAMAAMXASMASXAXMMMMMAMMMMMSSSSMAMXXXXAAXMMMAMAXMSMAMAMXMASMMXSXMAMAMAXAAASXMSASAXXXMAMAMMSMAAAMAMXMAMSSMSSXASAMMMAMMSASXSXMXMAMAAMXM
AMSXSAMXMMMMMMMAASAXMXMXAAAMMXXAAMAAAXXXSSSXSAMXSAMMSSMMAMASAMXMAAAAAXAMASXSXSSMMMSASASXMMXMMSMSSXMMXSMAMAMAMXAAAXMAXAMAMAMXMAMAMAMXMSXMMXSX
MAXMASMXAAXASXXAMMMSMAXXXSAMXSSSSSSSMSSXMAAXSAAXSAXXAMASMXMSXXXMSSMMSXMMASAAAXAASAMMMMMAMSAMXAMXAMASMXXAXMMMMMMMMSASXMMSXMXAMAMXSAMXXMMXSAMX
XSSMAMXXAMXMMMMMXAXAMMXSMMXSAMXAXAMXXXMAMXMXSAMMSSMMASXMSAMXMMMXMAMSMASMMMXMXMMMMAMXAAMMAMASMSSMXMAXMAXSMSAAAMAXXXMXAMAAAMSSSMSMSXSXASAAMMSM
XMAAASXSMSAAASXSSSSSXSASAAAMXSMSMMMXMASMMASXMXMXMAMSAMXASMSAXSAMXMSXMAMAXAAXAMXMSMMSSSSXSSMMAAMMAMXXMSMXASMSMXSAMXASMMSSMMAAAXAAXASMMMMMSAAM
MSAMMSASASXSMSAAAAMXAMASMMMXSAMXAMXAMMMAXASAMMMASMMMXSMXMXXMMMXSXMMAMMXSMSMMMXAXAAMMAAXXMAMMMMMSMMMSAXAMXMAMAMXAXAXASXMAXXXSXMMSMXMAXSXMMXSS
AXASAXMMAMAMAMMMMSMMXMXMASAXMXMXAMSXXAXAMXSMMASASXAAMSMMXSMMAMXMASMSMAAMXMXMSSMMSSMMMMXSAMXSXMAMAAXMAMXMMMXMAMXMASAMAXSAMXMMMSXAXMSXMAMAMXXM
SXSMMSAMXMMXXMMXAMXMXMASMMMSMAXMAMAMSSMSAMXXSXSASXXSAXAMAXMSXXASMMAAMMMMASMXMAXAMXXXAXMAMSXMAMMXMMSSSMSSSMSXSMAMAMAXMXMAXXAAXMSMSAMSMXMSAMXX
XMXAAMMASMMSMSSMMMAXAMAMAAXAMAMMAMAXAMXAXAMAMXMAMXMAXXAMAXMAAMMSSMXAMXXSAXAAMAMSAMXSXXASMSASMMSAAMAXMXAMAAMAMXXMAXXMXXMXMSSMMMXMXXMAXMAMMSMS
MASMMXSASAAAAAXAMSMSMSSSSMSXMASASMMSAMXMMXSAAAMMMAAMSSXMSSMMSAAMXMMMSAXMAMSMSSMXASAMXMSMAMAMAASMSMASXMMSMMMAMXMSSSXMAMMAMXAMXSAMSSSMSMMXAAAA
AMAAAAMASMMMMMSAMAAAAAAMAAMXXSXAMAXSXMMXAASMMMMASXSAAAAMAAXAAMAMASAAMMSMAMXAMXASAMMSAXAMMMSSMMMAMAAMAAAAASMXSXAAAAMXAXSAMSXMASASAAAXSAAMSMSM
SXSAMSMAMXXXMXXMSMSMMMSMMXMASAXXSXMMMAXXMMSAMXMAXSAMXSSMSSMSSMMSASMSMMMMASMAMMMMAMXSXSASAXXAMXMXMMSSSMMSSMAASMMMSMAMXMSAXXSMMSXMMSMMXMAXMAMX`
