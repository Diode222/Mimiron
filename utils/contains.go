package utils

func StringContains(strSlice []string, str string) bool {
	b := false
	if strSlice == nil {
		return b
	}
	for _, s := range strSlice {
		if s == str {
			return true
		}
	}
	return false
}
