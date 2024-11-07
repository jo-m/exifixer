package main

import (
	"testing"
	"time"
)

func TestTsFromFileName(t *testing.T) {
	// Use the machine's local timezone
	loc := time.Now().Location()

	// Define test cases
	tests := []struct {
		fname    string
		expected time.Time
	}{
		{
			fname:    "IMG-20220330-WA0002.jpg",
			expected: time.Date(2022, 3, 30, 0, 0, 0, 0, loc),
		},
		{
			fname:    "IMG-20240922-WA0002.jpeg",
			expected: time.Date(2024, 9, 22, 0, 0, 0, 0, loc),
		},
		{
			fname:    "IMG_20240211_172433_486.jpg",
			expected: time.Date(2024, 2, 11, 17, 24, 33, 0, loc),
		},
		{
			fname:    "WhatsApp Image 2024-03-09 at 16.28.30.jpeg",
			expected: time.Date(2024, 3, 9, 16, 28, 30, 0, loc),
		},
		{
			fname:    "WhatsApp Image 2022-12-17 at 15.31.39(1).jpeg",
			expected: time.Date(2022, 12, 17, 15, 31, 39, 0, loc),
		},
		{
			fname:    "IMG-20240204-WA0004.jpg",
			expected: time.Date(2024, 2, 4, 0, 0, 0, 0, loc),
		},
		{
			fname:    "WhatsApp Image 2022-08-08 at 4.30.35 PM.jpeg",
			expected: time.Date(2022, 8, 8, 16, 30, 35, 0, loc),
		},
		{
			fname:    "WhatsApp Image 2022-08-08 at 4.30.35 PM(1).jpeg",
			expected: time.Date(2022, 8, 8, 16, 30, 35, 0, loc),
		},
	}

	for _, test := range tests {
		actual := tsFromFileName(test.fname)

		if actual == nil || !actual.Equal(test.expected) {
			t.Errorf("For filename %s, expected %v, got %v", test.fname, test.expected, actual)
		}
	}
}
