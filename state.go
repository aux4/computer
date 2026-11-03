package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

const logMaxBytesDefault = 10 * 1024 * 1024

// hostOS is runtime.GOOS unless AUX4_COMPUTER_OS overrides it (a test seam: it lets one machine exercise
// the provider choice and state-dir layout of the others).
func hostOS() string {
	if v := os.Getenv("AUX4_COMPUTER_OS"); v != "" {
		return v
	}
	return runtime.GOOS
}

func homeDir() string {
	if h, err := os.UserHomeDir(); err == nil {
		return h
	}
	return "."
}

// stateDirOverridden reports whether AUX4_COMPUTER_STATE_DIR is set. When it is, only that directory is
// used: the legacy computer-mac paths are ignored so tests never touch the real switch or log.
func stateDirOverridden() bool {
	return os.Getenv("AUX4_COMPUTER_STATE_DIR") != ""
}

// stateDir is the per-OS directory shared by the core and every provider: the STOP file and actions.jsonl.
func stateDir() string {
	if v := os.Getenv("AUX4_COMPUTER_STATE_DIR"); v != "" {
		return v
	}
	switch hostOS() {
	case "darwin":
		return filepath.Join(homeDir(), "Library", "Application Support", "aux4-computer")
	case "windows":
		if v := os.Getenv("LOCALAPPDATA"); v != "" {
			return filepath.Join(v, "aux4-computer")
		}
		return filepath.Join(homeDir(), "AppData", "Local", "aux4-computer")
	default:
		if v := os.Getenv("XDG_STATE_HOME"); v != "" {
			return filepath.Join(v, "aux4-computer")
		}
		return filepath.Join(homeDir(), ".local", "state", "aux4-computer")
	}
}

func stopFile() string { return filepath.Join(stateDir(), "STOP") }

func logFile() string { return filepath.Join(stateDir(), "actions.jsonl") }

// legacyMacStopFile is where aux4/computer-mac 0.6.x looks for its switch. `stop` writes it too so the
// switch reaches providers that predate the shared state dir.
func legacyMacStopFile() string {
	return filepath.Join(homeDir(), "Library", "Application Support", "aux4-computer-mac", "STOP")
}

func useLegacyMacPaths() bool {
	return hostOS() == "darwin" && !stateDirOverridden()
}

// stopFiles are every STOP file the core writes on `stop`, removes on `resume` and honours before a dispatch.
func stopFiles() []string {
	files := []string{stopFile()}
	if useLegacyMacPaths() {
		files = append(files, legacyMacStopFile())
	}
	return files
}

// logSources are every action log `log` reads, newest-rotation last. On macOS the provider may keep its log
// under ~/Library/Logs (contract v1 allows it), and computer-mac 0.6.x writes to its own legacy directory.
func logSources() []string {
	sources := []string{logFile()}
	if useLegacyMacPaths() {
		sources = append(sources,
			filepath.Join(homeDir(), "Library", "Logs", "aux4-computer", "actions.jsonl"),
			filepath.Join(homeDir(), "Library", "Logs", "aux4-computer-mac", "actions.jsonl"),
		)
	}
	return sources
}

// stopReason says why input is refused right now, or "" when it's allowed.
func stopReason() string {
	if os.Getenv("AUX4_COMPUTER_STOP") == "1" {
		return "AUX4_COMPUTER_STOP=1 is set"
	}
	if useLegacyMacPaths() && os.Getenv("AUX4_COMPUTER_MAC_STOP") == "1" {
		return "AUX4_COMPUTER_MAC_STOP=1 is set"
	}
	for _, f := range stopFiles() {
		if _, err := os.Stat(f); err == nil {
			return f + " exists"
		}
	}
	return ""
}

func nowTS() string { return time.Now().UTC().Format("2006-01-02T15:04:05.000Z") }

// truncateForLog keeps a short prefix of text that may be private, plus its length (same as computer-mac).
func truncateForLog(s string) string {
	r := []rune(s)
	if len(r) <= 16 {
		return s
	}
	return fmt.Sprintf("%s… (%d chars)", string(r[:16]), len(r))
}

func logMaxBytes() int64 {
	if n, err := strconv.ParseInt(os.Getenv("AUX4_COMPUTER_LOG_MAX_BYTES"), 10, 64); err == nil && n > 0 {
		return n
	}
	return logMaxBytesDefault
}

// appendActionLog writes one JSON line to the shared action log (0600, rotated to actions.jsonl.1 at 10MB).
// Logging never makes a command fail.
func appendActionLog(entry map[string]any) {
	path := logFile()
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return
	}
	if st, err := os.Stat(path); err == nil && st.Size() >= logMaxBytes() {
		_ = os.Rename(path, path+".1")
	}
	line := marshal(entry)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o600)
	if err != nil {
		return
	}
	defer f.Close()
	_ = f.Chmod(0o600)
	_, _ = f.Write(append(line, '\n'))
}

func cmdStop() {
	content := fmt.Sprintf("stopped at %s by 'aux4 computer stop'; 'aux4 computer resume' removes this file\n", nowTS())
	files := stopFiles()
	for _, f := range files {
		if err := os.MkdirAll(filepath.Dir(f), 0o700); err != nil {
			fail(map[string]any{"error": fmt.Sprintf("Could not create %s: %v", filepath.Dir(f), err)})
		}
		if err := os.WriteFile(f, []byte(content), 0o600); err != nil {
			fail(map[string]any{"error": fmt.Sprintf("Could not create the stop file at %s: %v", f, err)})
		}
	}
	appendActionLog(map[string]any{"ts": nowTS(), "provider": "core", "command": "stop", "args": map[string]any{}, "result": "ok"})
	printJSON(map[string]any{"stopped": true, "file": stopFile(), "files": files})
}

func cmdResume() {
	files := stopFiles()
	for _, f := range files {
		if err := os.Remove(f); err != nil && !os.IsNotExist(err) {
			fail(map[string]any{"error": fmt.Sprintf("Could not remove the stop file at %s: %v", f, err)})
		}
	}
	appendActionLog(map[string]any{"ts": nowTS(), "provider": "core", "command": "resume", "args": map[string]any{}, "result": "ok"})
	out := map[string]any{"stopped": false, "file": stopFile(), "files": files}
	// An environment switch can't be undone from here: say so instead of claiming input is allowed.
	if reason := stopReason(); reason != "" {
		out["stopped"] = true
		out["stopReason"] = reason + "; unset it to allow input"
	}
	printJSON(out)
}

type logEntry struct {
	ts    string
	order int
	value map[string]any
}

func readLogFile(path string, provider string, order *int, into *[]logEntry) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 4*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		var m map[string]any
		if json.Unmarshal([]byte(line), &m) != nil {
			continue
		}
		if _, ok := m["provider"]; !ok && provider != "" {
			m["provider"] = provider
		}
		ts, _ := m["ts"].(string)
		*order++
		*into = append(*into, logEntry{ts: ts, order: *order, value: m})
	}
}

func cmdLog(args []string) {
	limitText := "20"
	if len(args) > 0 && args[0] != "" {
		limitText = args[0]
	}
	limit, err := strconv.Atoi(limitText)
	if err != nil || limit <= 0 {
		fail(map[string]any{"error": fmt.Sprintf("Invalid --limit '%s': use a positive number", limitText)})
	}
	var entries []logEntry
	order := 0
	var paths []string
	for _, src := range logSources() {
		provider := ""
		if strings.Contains(src, "aux4-computer-mac") || strings.Contains(src, filepath.Join("Logs", "aux4-computer")) {
			provider = "mac"
		}
		for _, p := range []string{src + ".1", src} {
			if _, err := os.Stat(p); err == nil {
				paths = append(paths, p)
				readLogFile(p, provider, &order, &entries)
			}
		}
	}
	sort.SliceStable(entries, func(i, j int) bool {
		if entries[i].ts != entries[j].ts {
			return entries[i].ts < entries[j].ts
		}
		return entries[i].order < entries[j].order
	})
	if len(entries) > limit {
		entries = entries[len(entries)-limit:]
	}
	out := make([]map[string]any, 0, len(entries))
	for _, e := range entries {
		out = append(out, e.value)
	}
	if paths == nil {
		paths = []string{}
	}
	result := map[string]any{"entries": out, "path": logFile(), "paths": paths, "stopped": false}
	if reason := stopReason(); reason != "" {
		result["stopped"] = true
		result["stopReason"] = reason
	}
	printJSON(result)
}

// marshal encodes without HTML escaping, so "File > Save" stays readable (same as the providers).
func marshal(v any) []byte {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(v)
	return bytes.TrimRight(buf.Bytes(), "\n")
}

func printJSON(v any) {
	fmt.Println(string(marshal(v)))
}

// fail prints a JSON error on stderr and exits 1 (the shared error shape).
func fail(v map[string]any) {
	fmt.Fprintln(os.Stderr, string(marshal(v)))
	os.Exit(1)
}
