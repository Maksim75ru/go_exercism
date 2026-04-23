package airportrobot

import (
    "fmt"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
    LanguageName() string
    Greet(a string) string
}

func SayHello(name string, g Greeter) string {
    return fmt.Sprintf("I can speak %s: %s!", g.LanguageName(), g.Greet(name))
}


type Italian struct {
    languageName string
    greet string
}

func (i Italian) LanguageName() string {
    i.languageName = "Italian"
    return i.languageName

}

func (i Italian) Greet(a string) string {
    i.greet = "Ciao " + a
    return i.greet
}



type Portuguese struct {
    languageName string
    greet string
}

func (p Portuguese) LanguageName() string {
    p.languageName = "Portuguese"
    return p.languageName
}

func (p Portuguese) Greet(a string) string {
    p.greet = "Olá " + a
    return p.greet
}