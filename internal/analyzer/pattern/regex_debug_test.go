package pattern

import "testing"
import "regexp"
import "fmt"

func TestRegexDebugging(t *testing.T) {
	pattern := `(?i)password\s*(:=|=)\s*".+"`

	re, err := regexp.Compile(pattern)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Matched?", re.MatchString("password := \"supersecret123\""))
	fmt.Println("Matched line?", re.MatchString("	password := \"supersecret123\" // vulnerable"))
}
