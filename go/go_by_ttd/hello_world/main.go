package main

import "fmt"

const (
	englishGreeting = "hello"
)

func HelloWorld(language, name string) string {
	var prefix string
	if name == "" {
		name = "world"
	}
	switch language {
	case "Bangla":
		prefix = englishGreeting

	default:
		prefix = englishGreeting
	}
	return prefix + " " + name

}
func main() {
	fmt.Println(HelloWorld("Bangla", "eyanat"))
}
