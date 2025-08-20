package airportrobot

import "fmt"

// Greeter interface defines how all language greeters should behave.
type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

// SayHello takes a visitor name and a Greeter, then builds the message.
func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

// Italian implements the Greeter interface for Italian.
type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

// Portuguese implements the Greeter interface for Portuguese.
type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
