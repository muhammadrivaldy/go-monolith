package util

func SearchIntInArray(object int, arr []int) (idx int, val int, res bool) {
	for h, i := range arr {
		if object == i {
			return h, i, true
		}
	}

	return 0, 0, false
}
