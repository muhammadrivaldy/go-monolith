package util

import (
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

func StringToInt(s string) int {
	res, _ := strconv.ParseInt(s, 10, 64)
	return int(res)
}

func StringToInt64(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func Int64ToRupiah(i int64) string {
	result := humanize.Comma(int64(i))
	return "Rp. " + strings.ReplaceAll(result, ",", ".")
}

func StringsToInts(s []string) (res []int) {
	for _, i := range s {
		res = append(res, StringToInt(i))
	}
	return
}
