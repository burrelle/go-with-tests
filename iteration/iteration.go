package iteration

const repeatCount = 5

// Repeat : Takes a character and and int and returns a string
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
