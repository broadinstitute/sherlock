package slack

import "fmt"

func LinkHelper(url string, text string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}
