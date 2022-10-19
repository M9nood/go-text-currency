package converter

import (
	"strconv"
	"strings"
)

func getIntegerAndDecimalDigits(num float64) (string, string) {
	numStrSplit := strings.Split(strconv.FormatFloat(num, 'f', -1, 64), ".")
	if len(numStrSplit) < 2 {
		return numStrSplit[0], ""
	}
	return numStrSplit[0], numStrSplit[1]
}

func isZero(num float64) bool {
	return num == 0
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func getMillionUnitCount(s string) int {
	if len(s) <= MILLION_POSITION {
		return 0
	}
	return len(s) / MILLION_POSITION
}

func getUnitPosition(pos int) int {
	return pos % MILLION_POSITION
}

func isMillionPosition(pos int) bool {
	return pos == MILLION_POSITION
}

func isOverPrecisionDigit(digit string) bool {
	return len(digit) > DEFAULT_PRECISION_DIGIT
}

func isEmptyDigit(num string) bool {
	return num == "0" || num == ""
}
