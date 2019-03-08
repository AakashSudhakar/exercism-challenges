// STATUS: Pending.

package proverb

import "fmt"

// Proverb() takes a list of nouns and returns a thematic proverb with each noun referenced.
func Proverb(objects []string) []string {
	proverb := []string{}
	if len(objects) <= 0 {
		return proverb
	}

	for iterator := 0; iterator+1 <= len(objects); iterator++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", objects[iterator], objects[iterator+1])
		proverb = append(proverb, line)
	}
	proverb = append(proverb, fmt.Sprintf("And all for the want of a %s.", objects[0]))
	return proverb
}
