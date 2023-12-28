package timer

import (
	"strings"
	"time"
)

func GenerateQuestTime(text string) time.Duration {
	// insanelly I found 600 wpm to be a good reading time for most players

	return time.Duration(int(
		len(strings.Fields(text)) * int(time.Minute) / 600),
	)
}
