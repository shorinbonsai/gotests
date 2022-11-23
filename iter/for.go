package iter

const loopCount = 5

func Loopy(character string, count int) string {
	var looped string
	for i := 0; i < count; i++ {
		looped += character
	}
	return looped
}
