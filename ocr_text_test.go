package main

import "testing"

func makeMatch(text string, x, y, width, height float64) ocrMatch {
	m := ocrMatch{Text: text}
	m.Center.X = x
	m.Center.Y = y
	m.Frame.X = x - width/2
	m.Frame.Y = y - height/2
	m.Frame.Width = width
	m.Frame.Height = height
	return m
}

func TestGroupOcrRowsEmpty(t *testing.T) {
	if rows := groupOcrRows(nil); rows != nil {
		t.Fatalf("expected nil rows for no matches, got %v", rows)
	}
	if text := ocrRowsToText(groupOcrRows(nil)); text != "" {
		t.Fatalf("expected empty text for no matches, got %q", text)
	}
}

// Same row, given out of left-to-right order and with a couple of pixels of y jitter (as real OCR boxes
// have) - they must land on one row, sorted by x.
func TestGroupOcrRowsSameRowWithJitterAndOutOfOrder(t *testing.T) {
	matches := []ocrMatch{
		makeMatch("75", 300, 101, 40, 30),
		makeMatch("42+", 100, 100, 50, 30),
		makeMatch("x100", 500, 99, 45, 30),
	}
	rows := groupOcrRows(matches)
	if len(rows) != 1 {
		t.Fatalf("expected 1 row, got %d: %v", len(rows), rows)
	}
	if got := ocrRowsToText(rows); got != "42+  75  x100" {
		t.Fatalf("unexpected text: %q", got)
	}
}

// Multiple rows (e.g. Calculator's display above its button grid) must come back top-to-bottom, each
// sorted left-to-right, regardless of the input order.
func TestGroupOcrRowsMultipleRows(t *testing.T) {
	matches := []ocrMatch{
		// bottom row (buttons), given out of order
		makeMatch("%", 220, 500, 40, 40),
		makeMatch("AC", 100, 500, 40, 40),
		makeMatch("÷", 340, 500, 40, 40),
		// top row (display), given first
		makeMatch("7,542", 300, 100, 80, 30),
	}
	rows := groupOcrRows(matches)
	if len(rows) != 2 {
		t.Fatalf("expected 2 rows, got %d: %v", len(rows), rows)
	}
	if got := ocrRowsToText(rows); got != "7,542\nAC  %  ÷" {
		t.Fatalf("unexpected text: %q", got)
	}
}

// Rows separated by a gap close to a button's own height must still split - tolerance scales with the
// height of the taller entry being compared, not with an absolute pixel count.
func TestGroupOcrRowsDistinctRowsNotMerged(t *testing.T) {
	matches := []ocrMatch{
		makeMatch("row1", 100, 100, 60, 30),
		makeMatch("row2", 100, 160, 60, 30),
	}
	rows := groupOcrRows(matches)
	if len(rows) != 2 {
		t.Fatalf("expected 2 distinct rows, got %d: %v", len(rows), rows)
	}
}

func TestOcrJSONToTextParsesMatchesOnly(t *testing.T) {
	data := []byte(`{
		"app": "Calculator",
		"matches": [
			{"text": "AC", "center": {"x": 100, "y": 500}, "frame": {"x": 80, "y": 480, "width": 40, "height": 40}},
			{"text": "7,542", "center": {"x": 300, "y": 100}, "frame": {"x": 260, "y": 85, "width": 80, "height": 30}}
		]
	}`)
	text, err := ocrJSONToText(data)
	if err != nil {
		t.Fatal(err)
	}
	if text != "7,542\nAC" {
		t.Fatalf("unexpected text: %q", text)
	}
}

func TestOcrJSONToTextNoMatches(t *testing.T) {
	text, err := ocrJSONToText([]byte(`{"matches": []}`))
	if err != nil {
		t.Fatal(err)
	}
	if text != "" {
		t.Fatalf("expected empty text, got %q", text)
	}
}
