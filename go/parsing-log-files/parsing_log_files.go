package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[-~=*]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	for _, line := range lines {
		line = strings.ToLower(line)
		re := regexp.MustCompile(`".*password.*"`)
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	tagged := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+(\S+)`)

	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			username := match[1]
			tagged[i] = fmt.Sprintf("[USR] %s %s", username, line)
		} else {
			tagged[i] = line
		}
	}

	return tagged
}
