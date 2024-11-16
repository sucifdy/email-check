package email

import (
	"regexp"
)

func IsValidEmail(email string) bool {
	// Regular expression
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
