package hello

const greet = "Namaste, "

func Hello(name string) string {
	if name == "" {
		return greet + "world"
	}
	return greet + name
}