package iteration

const repeatCount = 5

func repeat(character string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return
}

func main() {
}
