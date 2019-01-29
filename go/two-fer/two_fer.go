// STATUS: Passed submission.

package twofer

import "fmt"

// Function that returns two-fer string expression with variable declaration.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
