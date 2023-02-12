package stringFields

// Вся эта функция заменяется одним strings.Fields()
func StringToArray(str string) []string {
	result := []string{}
	var temp string
	for k, v := range str {
		switch {
		case v == 32:
			result = append(result, temp)
			temp = ""
		case k == len(str)-1:
			temp += string(v)
			result = append(result, temp)
			//break
		default:
			temp += string(v)
		}
	}
	return result
}
