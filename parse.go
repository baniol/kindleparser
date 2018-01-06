package kindleparser

import (
	// "fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const lineDelimiter = "\r\n==========\r\n"

type Record struct {
	Title    string    `json:"title"`
	Author   string    `json:"author,omitempty"`
	Type     string    `json:"type"`
	Location string    `json:"loc,omitempty"`
	Page     string    `json:"page,omitempty"`
	Added    time.Time `json:"added"`
	Content  string    `json:"content"`
}

func ParseClippngs(filePath string) []*Record {
	text, err := getFileContents(filePath)
	if err != nil {
		panic("Error reading file")
	}

	lines := splitIntoRecords(text)

	var out []*Record

	// TODO: into goroutines (with limit ?) - what about ordering then ?
	for _, l := range lines {
		r := parseRecord(l)
		if len(r) == 4 {
			title, author := getTitleAndAuthor(r[0])
			if len(r) > 1 {
				kind, loc, page, added := parseSecondLine(r[1])
				record := &Record{
					Title:    title,
					Author:   author,
					Type:     kind,
					Added:    added,
					Location: loc,
					Page:     page,
					Content:  r[3],
				}
				out = append(out, record)
			}
		}
	}
	return out
}

func getTitleAndAuthor(line string) (string, string) {
	trimmed := strings.TrimSuffix(line, "\r")
	parts := strings.Split(trimmed, " (")

	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return parts[0], ""
	}

	title := parts[0]
	author := strings.TrimSuffix(parts[len(parts)-1], ")")

	return title, author
}

func getFileContents(clippingsFile string) (string, error) {
	text, err := ioutil.ReadFile(clippingsFile)
	if err != nil {
		return "", err
	}
	return string(text), nil
}

func splitIntoRecords(text string) []string {
	lines := strings.Split(text, lineDelimiter)
	return lines
}

func parseRecord(line string) []string {
	items := strings.Split(line, "\n")
	return items
}

func parseSecondLine(line string) (string, string, string, time.Time) {
	trimmed := strings.TrimPrefix(line, "- ")
	out := strings.SplitN(trimmed, " ", 4)

	kind := out[0]
	loc := out[2]
	addedStr := out[3]
	page := ""

	if loc == "Page" {
		rest := strings.Split(out[3], " | ")
		page = rest[0]
		if len(rest) == 3 {
			loc = strings.Trim(strings.Split(rest[1], " ")[1], " ")
			addedStr = rest[2]
		} else {
			loc = ""
		}
	}

	added := parseAdded(addedStr)

	return kind, loc, page, added
}

func parseAdded(added string) time.Time {

	p := strings.Split(added, ", ")

	year, _ := strconv.Atoi(p[2])

	ta := strings.Split(p[3], " ")
	tt := strings.Split(ta[0], ":")

	hour, _ := strconv.Atoi(strings.TrimPrefix(tt[0], "0"))
	minute, _ := strconv.Atoi(strings.TrimPrefix(tt[1], "0"))

	if ta[1] == "PM" || ta[1] == "PM\r" {
		hour += 12
	}

	d := strings.Split(p[1], " ")

	month := getMonth(d[0])
	day, _ := strconv.Atoi(d[1])

	return time.Date(year, month, day, hour, minute, 0, 0, time.Local)
}

func getMonth(name string) time.Month {
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	for i, m := range months {
		if name == m {
			mm := i + 1
			return time.Month(mm)
		}
	}
	return 0
}
