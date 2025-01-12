package iteration

func Repeat(ch string, count int) string {
	var result string

	for i := 0; i < count; i++ {
		result += ch
	}

	return result
}
