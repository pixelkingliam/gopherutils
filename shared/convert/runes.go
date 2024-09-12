package convert

func RunifyString(str string) []rune {
	return []rune(str)
}
func RunifyStrings(strings []string) [][]rune {
	result := make([][]rune, len(strings))
	for i, s := range strings {
		result[i] = RunifyString(s)
	}
	return result
}
