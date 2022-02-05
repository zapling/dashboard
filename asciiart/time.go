package asciiart

import (
	"strings"
	"time"
)

func GetTime(t time.Time) string {
	str := t.Format("15:04:05")

	var asciiCharacters [][]string
	for _, char := range str {
		asciiCharacters = append(asciiCharacters, characters[char])
	}

	mergedCharacters := []string{"", "", "", "", ""}
	for char := 0; char < len(asciiCharacters); char++ {
		for row := 0; row < len(asciiCharacters[char]); row++ {
			mergedCharacters[row] = mergedCharacters[row] + asciiCharacters[char][row] + " "
		}
	}

	return strings.Join(mergedCharacters, "\n")
}
