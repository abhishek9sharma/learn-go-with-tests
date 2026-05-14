package hello

const greetEnglish = "Hello, "
const greetHindi = "Namaste, "

func Hello(name string, language string) string {
    if name == "" {
        name = "world"
    }

    var greet string
    if language == "Hindi" {
        greet = greetHindi
    } else {
        greet = greetEnglish
    }
    
    return greet + name
}