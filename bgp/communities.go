package bgp

import (
	"fmt"
	"strconv"
	"strings"
)

func createRange(num int) string {
	if num > 9 {
		panic(fmt.Sprintf("%d is higher than 9", num))
	}
	result := ""
	if num == 0 {
		result = "[1-9]"
	} else if num == 9 {
		result = "[0-8]"
	} else {
		less := num - 1
		more := num + 1
		result = fmt.Sprintf("[0-%d%d-9]", less, more)
	}
	return result
}

func intToSlice(num int) []int {
	parts := make([]int, 0, 3)
	filled := fmt.Sprintf("%03d", num)
	for i := 0; i < len(filled); i++ {
		num64, err := strconv.ParseInt(string(filled[i]), 10, 64)
		if err != nil {
			panic(err)
		}
		parts = append(parts, int(num64))
	}
	return parts
}

func InverseNumberMatch(input int) string {
	chars := intToSlice(input)
	finalParts := []string{}
	if input < 10 {
		// e.g. 1
		r := createRange(input)
		finalParts = append(finalParts,
			fmt.Sprintf("[1-9]0%s", r),
			fmt.Sprintf("00%s", r),
			fmt.Sprintf("0[1-9]%d", input),
			fmt.Sprintf("0[1-9]%s", r),
			fmt.Sprintf("[1-9][1-9]%s", r),
			fmt.Sprintf("[1-9][1-9]%d", input),
			fmt.Sprintf("[1-9]0%d", input),
		)

	} else if input < 100 {
		// e.g. 012
		r0 := createRange(chars[1])
		r1 := createRange(chars[2])
		finalParts = append(finalParts,
			"00[1-9]",
			"[1-9][0-9][0-9]",
			fmt.Sprintf("0%s%d", r0, chars[1]),
			fmt.Sprintf("0%d%s", chars[0], r1),
			fmt.Sprintf("0%s[0-9]", r0),
			fmt.Sprintf("0[0-9]%s", r1),
		)
	} else if input < 1000 {
		// e.g. 345
		r0 := createRange(chars[0])
		r1 := createRange(chars[1])
		r2 := createRange(chars[2])
		finalParts = append(finalParts,
			// 1 digit
			"00[1-9]",
			// 2 digit
			"0[0-9][1-9]",
			"0[1-9][0-9]",
			// 3 digit
			fmt.Sprintf("%d%d%s", chars[0], chars[1], r2),
			fmt.Sprintf("%d%s%d", chars[0], r1, chars[2]),
			fmt.Sprintf("%s%d%d", r0, chars[1], chars[2]),
			fmt.Sprintf("%s[0-9][0-9]", r0),
			fmt.Sprintf("[1-9]%s[0-9]", r1),
			fmt.Sprintf("[1-9][0-9]%s", r2),
		)
	}
	return fmt.Sprintf("(%s)", strings.Join(finalParts, "|"))
}
