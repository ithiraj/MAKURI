package main

import (
	"fmt"
	"sort"
)

func main() {
	var str string
	var iteration_stage1, value1, iteration4value1 int

	var stor [92]string

	stor[0] = "_"
	stor[1] = "a"
	stor[2] = "b"
	stor[3] = "c"
	stor[4] = "d"
	stor[5] = "e"
	stor[6] = "f"
	stor[7] = "g"
	stor[8] = "h"
	stor[9] = "i"
	stor[10] = "j"
	stor[11] = "k"
	stor[12] = "l"
	stor[13] = "m"
	stor[14] = "n"
	stor[15] = "o"
	stor[16] = "p"
	stor[17] = "q"
	stor[18] = "r"
	stor[19] = "s"
	stor[20] = "t"
	stor[21] = "u"
	stor[22] = "v"
	stor[23] = "w"
	stor[24] = "x"
	stor[25] = "y"
	stor[26] = "z"
	stor[27] = "A"
	stor[28] = "B"
	stor[29] = "C"
	stor[30] = "D"
	stor[31] = "E"
	stor[32] = "F"
	stor[33] = "G"
	stor[34] = "H"
	stor[35] = "I"
	stor[36] = "J"
	stor[37] = "K"
	stor[38] = "L"
	stor[39] = "M"
	stor[40] = "N"
	stor[41] = "O"
	stor[42] = "P"
	stor[43] = "Q"
	stor[44] = "R"
	stor[45] = "S"
	stor[46] = "T"
	stor[47] = "U"
	stor[48] = "V"
	stor[49] = "W"
	stor[50] = "X"
	stor[51] = "Y"
	stor[52] = "Z"
	stor[53] = "-"
	stor[54] = "$"
	stor[55] = "%"
	stor[56] = "^"
	stor[57] = "&"
	stor[58] = "*"
	stor[59] = "("
	stor[60] = ")"
	stor[61] = "!"
	stor[62] = "@"
	stor[63] = "#"
	stor[64] = "="
	stor[65] = "+"
	stor[66] = "`"
	stor[67] = "~"
	stor[68] = "["
	stor[69] = "]"
	stor[70] = "{"
	stor[71] = "}"
	stor[72] = ";"
	stor[73] = "'"
	stor[74] = "|"
	stor[75] = ":"
	stor[76] = ","
	stor[77] = "."
	stor[78] = "/"
	stor[79] = "<"
	stor[80] = ">"
	stor[81] = "?"
	stor[82] = "0"
	stor[83] = "1"
	stor[84] = "2"
	stor[85] = "3"
	stor[86] = "4"
	stor[87] = "5"
	stor[88] = "6"
	stor[89] = "7"
	stor[90] = "8"
	stor[91] = "9"

	var stor1 [10]string

	stor1[0] = "0"
	stor1[1] = "1"
	stor1[2] = "2"
	stor1[3] = "3"
	stor1[4] = "4"
	stor1[5] = "5"
	stor1[6] = "6"
	stor1[7] = "7"
	stor1[8] = "8"
	stor1[9] = "9"

	map_stage1 := map[int]string{
		0: "Z",
		1: "O",
		2: "t",
		3: "T",
		4: "R",
		5: "F",
		6: "s",
		7: "S",
		8: "E",
		9: "N",
	}

	map_stage2 := map[int]string{
		0:  "0_",
		1:  "1_",
		2:  "2_",
		3:  "3_",
		4:  "4_",
		5:  "5_",
		6:  "6_",
		7:  "7_",
		8:  "8_",
		9:  "9_",
		10: "0a",
		11: "1a",
		12: "2a",
		13: "3a",
		14: "4a",
		15: "5a",
		16: "6a",
		17: "7a",
		18: "8a",
		19: "9a",
		20: "0b",
		21: "1b",
		22: "2b",
		23: "3b",
		24: "4b",
		25: "5b",
		26: "6b",
		27: "7b",
		28: "8b",
		29: "9b",
		30: "0c",
		31: "1c",
		32: "2c",
		33: "3c",
		34: "4c",
		35: "5c",
		36: "6c",
		37: "7c",
		38: "8c",
		39: "9c",
		40: "0d",
		41: "1d",
		42: "2d",
		43: "3d",
		44: "4d",
		45: "5d",
		46: "6d",
		47: "7d",
		48: "8d",
		49: "9d",
		50: "0e",
		51: "1e",
		52: "2e",
		53: "3e",
		54: "4e",
		55: "5e",
		56: "6e",
		57: "7e",
		58: "8e",
		59: "9e",
		60: "0f",
		61: "1f",
		62: "2f",
		63: "3f",
		64: "4f",
		65: "5f",
		66: "6f",
		67: "7f",
		68: "8f",
		69: "9f",
		70: "0g",
		71: "1g",
		72: "2g",
		73: "3g",
		74: "4g",
		75: "5g",
		76: "6g",
		77: "7g",
		78: "8g",
		79: "9g",
		80: "0h",
		81: "1h",
		82: "2h",
		83: "3h",
		84: "4h",
		85: "5h",
		86: "6h",
		87: "7h",
		88: "8h",
		89: "9h",
		90: "0i",
		91: "1i",
	}

	fmt.Println("ENTER THE CIPHER")
	fmt.Scan(&str)

	slice_stage1 := make([]string, len(str))

	var slice2_stage1 []string
	flag := 0
	pos := 0
	iteration2_stage1 := 0

	for iteration_stage1 = 0; iteration_stage1 < len(str); iteration_stage1++ {

		for iteration2_stage1 = 0; iteration2_stage1 < len(map_stage1); iteration2_stage1++ {

			if string(str[iteration_stage1]) == map_stage1[iteration2_stage1] {
				flag = 1
				pos = iteration2_stage1

			}

		}
		if flag == 0 {

			slice_stage1[iteration_stage1] = string(str[iteration_stage1])
		} else {

			slice_stage1[iteration_stage1] = stor1[pos]
			flag = 0
		}

	}

	num_sli := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	sort.Strings(num_sli)

	for i := 0; i < len(slice_stage1); {
		var res2 int

		flag1 := 0

		res1 := sort.SearchStrings(num_sli, slice_stage1[i])
		if i == len(slice_stage1)-1 {
			flag1 = 1

		} else {

			res2 = sort.SearchStrings(num_sli, slice_stage1[i+1])
		}

		if res1 != 10 && res2 != 10 {
			res := slice_stage1[i] + "_"
			slice2_stage1 = append(slice2_stage1, res)
			i += 1
		} else if flag1 == 1 {
			res := slice_stage1[i] + "_"
			slice2_stage1 = append(slice2_stage1, res)
			i += 1
		} else {
			res := slice_stage1[i] + slice_stage1[i+1]
			slice2_stage1 = append(slice2_stage1, res)
			i += 2
		}

	}

	slice1_stage2 := make([]string, len(slice2_stage1))

	for iteration1_stage2 := 0; iteration1_stage2 < len(slice2_stage1); iteration1_stage2++ {
		for iteration2_stage2 := 0; iteration2_stage2 < len(stor); iteration2_stage2++ {

			if slice2_stage1[iteration1_stage2] == map_stage2[iteration2_stage2] {
				slice1_stage2[iteration1_stage2] = stor[iteration2_stage2]

			}
		}
	}
	fmt.Println(slice1_stage2)

	var outstring string

	for iteration4value1 = 0; iteration4value1 < len(slice1_stage2); iteration4value1++ {

		if slice1_stage2[iteration4value1] != "" {
			value1 = iteration4value1 + 1

		}

	}

	for iteration4outstring := 0; iteration4outstring < value1; iteration4outstring++ {
		outstring = outstring + string(slice1_stage2[iteration4outstring])
	}
	fmt.Println("----", outstring)

}
