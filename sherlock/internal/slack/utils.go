package slack

import (
	"fmt"
	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/slack-go/slack"
	"strings"
)

func LinkHelper(url string, text string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}

const slackTextBlockLengthLimit = 3000

func chunkLinesToSectionMrkdwnBlocks(lines []string) []slack.Block {
	var chunks []string
	for _, line := range lines {
		if len(chunks) == 0 {
			// If no chunks so far, begin with the first line
			chunks = []string{line}
		} else if len(chunks[len(chunks)-1])+len("\n")+len(line) < slackTextBlockLengthLimit {
			// If we can add the next line without going over the limit, do so
			chunks[len(chunks)-1] = chunks[len(chunks)-1] + "\n" + line
		} else {
			// Otherwise, split to a new chunk
			chunks = append(chunks, line)
		}
	}
	return utils.Map(chunks, func(chunk string) slack.Block {
		return slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", chunk, false, true), nil, nil)
	})
}

func EscapeText(text string) string {
	replacements := map[string]string{
		"&": "&amp;",
		"<": "&lt;",
		">": "&gt;",
	}
	for k, v := range replacements {
		text = strings.ReplaceAll(text, k, v)
	}
	return text
}
