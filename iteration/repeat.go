package iteration

const repeatCount = 5

// Repeat receives a char and returns a string
func Repeat(character string) string {
	var repeated string

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
