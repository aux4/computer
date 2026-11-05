package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sort"
	"strings"
)

// ocrMatch is the subset of an ocr `matches[]` entry that --format text needs: the recognized text and
// the frame/center the provider already placed it at (click coordinates). The full entry carries more
// fields (confidence, match, windowCenter...); json.Unmarshal ignores what it doesn't ask for.
type ocrMatch struct {
	Text   string `json:"text"`
	Center struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"center"`
	Frame struct {
		X      float64 `json:"x"`
		Y      float64 `json:"y"`
		Width  float64 `json:"width"`
		Height float64 `json:"height"`
	} `json:"frame"`
}

type ocrResponse struct {
	Matches []ocrMatch `json:"matches"`
}

// execOcrText runs the provider's ocr command exactly like execPassthrough, except it captures stdout
// instead of streaming it, so the raw JSON never reaches the terminal. On success (exit 0) stdout is
// replaced with the --format text rendering; a failure is passed through unchanged (same stderr, same
// exit code) so error handling stays exactly as it is for the default json format.
func execOcrText(argv []string) int {
	cmd := exec.Command(aux4Binary(), argv...)
	cmd.Stdin = os.Stdin
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	// Let the child handle Ctrl-C itself; we only wait for it.
	signal.Ignore(os.Interrupt)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if code := exitErr.ExitCode(); code >= 0 {
				return code
			}
			return 1
		}
		fail(map[string]any{"error": fmt.Sprintf("Could not run aux4: %v", err)})
	}

	text, err := ocrJSONToText(stdout.Bytes())
	if err != nil {
		// Shouldn't happen (the provider returned exit 0 but not the documented shape) - fall back to the
		// raw JSON rather than hide the output.
		os.Stdout.Write(stdout.Bytes())
		return 0
	}
	fmt.Println(text)
	return 0
}

// ocrJSONToText parses an ocr success response and renders its matches as plain text lines in reading
// order. When --text was given to ocr, matches is already filtered to the lines that matched, so the
// text output only contains those lines too - --format text does not change what --text selects, only
// how it's printed.
func ocrJSONToText(data []byte) (string, error) {
	var resp ocrResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return "", err
	}
	return ocrRowsToText(groupOcrRows(resp.Matches)), nil
}

// groupOcrRows groups OCR matches into visual rows by vertical position and sorts each row left-to-right.
// Matches are first ordered by vertical center; a match joins the current row when its center is within
// tolerance of the row's running average center (tolerance = 60% of the taller of the row's and the
// match's frame height, so same-row jitter of a few pixels doesn't split a row, but a genuinely separate
// line below/above starts a new one). Each returned row is itself sorted by horizontal center.
func groupOcrRows(matches []ocrMatch) [][]ocrMatch {
	if len(matches) == 0 {
		return nil
	}
	sorted := make([]ocrMatch, len(matches))
	copy(sorted, matches)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Center.Y < sorted[j].Center.Y
	})

	var rows [][]ocrMatch
	var row []ocrMatch
	var rowY, rowHeight float64

	flush := func() {
		if len(row) > 0 {
			rows = append(rows, row)
		}
	}

	for _, m := range sorted {
		h := m.Frame.Height
		if h <= 0 {
			h = 1
		}
		if len(row) == 0 {
			row = []ocrMatch{m}
			rowY = m.Center.Y
			rowHeight = h
			continue
		}
		tolerance := rowHeight
		if h > tolerance {
			tolerance = h
		}
		tolerance *= 0.6
		if diff(m.Center.Y, rowY) <= tolerance {
			row = append(row, m)
			n := float64(len(row))
			rowY = (rowY*(n-1) + m.Center.Y) / n
			if h > rowHeight {
				rowHeight = h
			}
		} else {
			flush()
			row = []ocrMatch{m}
			rowY = m.Center.Y
			rowHeight = h
		}
	}
	flush()

	for _, r := range rows {
		sort.SliceStable(r, func(i, j int) bool {
			return r[i].Center.X < r[j].Center.X
		})
	}
	return rows
}

// ocrRowsToText joins each row's text left-to-right with two spaces, one row per line.
func ocrRowsToText(rows [][]ocrMatch) string {
	lines := make([]string, 0, len(rows))
	for _, row := range rows {
		texts := make([]string, 0, len(row))
		for _, m := range row {
			texts = append(texts, m.Text)
		}
		lines = append(lines, strings.Join(texts, "  "))
	}
	return strings.Join(lines, "\n")
}

func diff(a, b float64) float64 {
	if a < b {
		return b - a
	}
	return a - b
}
