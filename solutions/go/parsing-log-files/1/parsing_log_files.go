package parsinglogfiles

import (
	"regexp"
	"strings"
)

// Task 1: check valid log prefix
func IsValidLine(text string) bool {
	validPrefixes := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, prefix := range validPrefixes {
		if strings.HasPrefix(text, prefix) {
			return true
		}
	}
	return false
}

// Task 2: split logs by special delimiters like <*>, <==>, <~~~>, etc.
func SplitLogLine(text string) []string {
	// regex: < [~-*=]* >   (zero or more allowed chars between brackets)
	re := regexp.MustCompile(`<[-~*=]*>`)
	return re.Split(text, -1)
}

// Task 3: count quoted "password" mentions
func CountQuotedPasswords(lines []string) int {
	count := 0
	re := regexp.MustCompile(`"[^"]*password[^"]*"`) // case-insensitive check later
	for _, line := range lines {
		low := strings.ToLower(line)
		if re.MatchString(low) {
			count++
		}
	}
	return count
}

// Task 4: remove all "end-of-line123" fragments
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// Task 5: tag user mentions
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	result := []string{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			user := match[1]
			line = "[USR] " + user + " " + line
		}
		result = append(result, line)
	}
	return result
}
