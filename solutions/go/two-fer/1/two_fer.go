package twofer

import "fmt"

func ShareWith(name string) string {
	template := "One for %s, one for me."
    knownNames := []string{"Alice", "Bohdan", "Zaphod", "Bob"}
    for _, n := range knownNames {
        if n == name {
            return fmt.Sprintf(template, name)
        }
    }
    return fmt.Sprintf(template, "you")
}
