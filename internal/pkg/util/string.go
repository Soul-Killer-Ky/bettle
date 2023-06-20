package util

func InArrayString(item string, arr []string) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
