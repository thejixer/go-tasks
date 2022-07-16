package knightmove

func FindIndex(x [8]string, v string) int {
	for index, element := range x {
		if element == v {
			return index
		}
	}
	return -1
}

func Contains(x [8]string, v string) bool {
	for _, s := range x {
		if v == s {
			return true
		}
	}
	return false
}
