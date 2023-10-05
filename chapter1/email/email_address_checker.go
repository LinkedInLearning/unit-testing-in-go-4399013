package linkedinlearning

import "strings"

// IsLinkedInEmployee checks if a given email address ends
// with linkedin.com
func IsLinkedInEmployee(address string) bool {
	return strings.HasSuffix(address, "@linkedin.com")
}
