package greeter

import "github.com/sotnikoff/greeter/internal/words"

func Hello(name string) string {
	return words.Greeting() + ", " + shout(name)
}

func shout(name string) string {
	return name + "!"
}
