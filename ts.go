package main

import (
	"log"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type FileType int

const (
	JPEG FileType = iota + 1
)

var patterns = []struct {
	re     *regexp.Regexp
	layout string
}{
	{
		re:     regexp.MustCompile(`(?:IMG|VID)-(\d{8})-WA\d+`),
		layout: "20060102",
	},
	{
		re:     regexp.MustCompile(`(?:IMG|VID)_(\d{8})_(\d{6})_\d+`),
		layout: "20060102  150405",
	},
	{
		re:     regexp.MustCompile(`WhatsApp (?:Image|Video) ([0-9\-]{10}) at ([0-9.]{8}).*`),
		layout: "2006-01-02  15.04.05",
	},
	{
		re:     regexp.MustCompile(`WhatsApp (?:Image|Video) ([0-9\-]{10}) at ([0-9.APM ]{10,11})(?:\(\d+\))?`),
		layout: "2006-01-02  3.04.05 PM",
	},
}

func tsFromFileName(fname string) *time.Time {
	// We always assume local timezone.
	loc := time.Now().Location()

	// Trim file extension.
	fname = strings.TrimSuffix(fname, filepath.Ext(fname))

	for _, p := range patterns {
		match := p.re.FindStringSubmatch(fname)
		if match == nil {
			continue
		}
		toParse := strings.Join(match[1:], "  ")
		ts, err := time.ParseInLocation(p.layout, toParse, loc)
		if err != nil {
			log.Fatalf("Failed to parse '%s' as '%s': %s", toParse, p.layout, err)
		}
		ts.In(loc)
		return &ts
	}

	return nil
}
