package slack

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/broadinstitute/sherlock/go-shared/pkg/utils"
	"github.com/slack-go/slack"
)

func LinkHelper(url string, text string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}

const slackTextBlockLengthLimit = 3000

func chunkLinesToSectionMrkdwnBlocks(lines []string) []slack.Block {
	var chunks []string
	for _, line := range lines {
		if chunks == nil {
			chunks = make([]string, 0)
		}
		if len(chunks) > 0 && len(chunks[len(chunks)-1])+len("\n")+len(line) <= slackTextBlockLengthLimit {
			// If we can add the next line without going over the limit, do so
			chunks[len(chunks)-1] = chunks[len(chunks)-1] + "\n" + line
		} else {
			// Add the line in 3000 character chunks
			for slackTextBlockLengthLimit < len(line) {
				line, chunks = line[slackTextBlockLengthLimit:], append(chunks, line[:slackTextBlockLengthLimit])
			}
			chunks = append(chunks, line)
		}
	}
	return utils.Map(chunks, func(chunk string) slack.Block {
		return slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", chunk, false, true), nil, nil)
	})
}

func EscapeText(text string) string {
	// It's important that these be ordered, because we want ">" to be escaped as "&gt;" and not "&amp;gt;"
	for _, pair := range [][]string{
		{"&", "&amp;"},
		{"<", "&lt;"},
		{">", "&gt;"},
	} {
		text = strings.ReplaceAll(text, pair[0], pair[1])
	}
	return text
}

var markdownLinkRegex = regexp.MustCompile(`\[(.*?)]\((https://.*?)\)`)

func MarkdownLinksToSlackLinks(text string) string {
	// Slack doesn't support markdown links, so we have to convert them to slack links
	// https://api.slack.com/reference/surfaces/formatting#linking-urls
	return markdownLinkRegex.ReplaceAllStringFunc(text, func(match string) string {
		parts := markdownLinkRegex.FindStringSubmatch(match)
		return LinkHelper(parts[2], parts[1])
	})
}
