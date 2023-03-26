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

func IntToRupiah(i int) string {
	result := humanize.Comma(int64(i))
	return "Rp. " + strings.ReplaceAll(result, ",", ".")
}
