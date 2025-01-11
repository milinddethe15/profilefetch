package utility

func CleanString(input string, suffix string) string {
	if len(input) >= len(suffix) && input[len(input)-len(suffix):] == suffix {
		return input[:len(input)-len(suffix)]
	}
	return input
}
