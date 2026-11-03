package main

import (
	"encoding/json"
	"os"
	"testing"
)

type aux4File struct {
	Profiles []struct {
		Name     string `json:"name"`
		Commands []struct {
			Name string `json:"name"`
		} `json:"commands"`
	} `json:"profiles"`
}

// The reserved-name list must match the core's `computer` profile exactly.
func TestCoreCommandsMatchPackage(t *testing.T) {
	data, err := os.ReadFile("package/.aux4")
	if err != nil {
		t.Fatal(err)
	}
	var f aux4File
	if err := json.Unmarshal(data, &f); err != nil {
		t.Fatal(err)
	}
	found := map[string]bool{}
	for _, p := range f.Profiles {
		if p.Name == "computer" {
			for _, c := range p.Commands {
				found[c.Name] = true
				if !coreCommands[c.Name] {
					t.Errorf("command %q is in package/.aux4 but not in coreCommands", c.Name)
				}
			}
		}
	}
	for name := range coreCommands {
		if !found[name] {
			t.Errorf("coreCommands has %q but package/.aux4 does not", name)
		}
	}
	for _, p := range knownProviders {
		if coreCommands[p.name] {
			t.Errorf("provider name %q collides with a core command", p.name)
		}
	}
}

func TestParseFlags(t *testing.T) {
	flags, err := parseFlags([]string{"--app", "Calculator", "--text", "it's $(x)"})
	if err != nil || flags["app"] != "Calculator" || flags["text"] != "it's $(x)" {
		t.Fatalf("unexpected %v %v", flags, err)
	}
	if _, err := parseFlags([]string{"stray"}); err == nil {
		t.Fatal("expected an error for a stray argument")
	}
	if _, err := parseFlags([]string{"--app"}); err == nil {
		t.Fatal("expected an error for a flag without a value")
	}
}

func TestIsStateChanging(t *testing.T) {
	cases := []struct {
		command string
		flags   map[string]string
		want    bool
	}{
		{"click", nil, true},
		{"screenshot", nil, false},
		{"menu", map[string]string{}, false},
		{"menu", map[string]string{"path": "File > Save"}, true},
		{"dialog", map[string]string{}, false},
		{"dialog", map[string]string{"press": "OK"}, true},
		{"window list", nil, false},
		{"window close", nil, true},
		{"clipboard read", nil, false},
		{"clipboard paste", nil, true},
	}
	for _, c := range cases {
		if got := isStateChanging(c.command, c.flags); got != c.want {
			t.Errorf("isStateChanging(%q, %v) = %v, want %v", c.command, c.flags, got, c.want)
		}
	}
}

func TestTruncateForLog(t *testing.T) {
	if got := truncateForLog("short"); got != "short" {
		t.Fatal(got)
	}
	if got := truncateForLog("a long and private sentence"); got != "a long and priva… (27 chars)" {
		t.Fatal(got)
	}
}

func TestStateDirPerOS(t *testing.T) {
	t.Setenv("AUX4_COMPUTER_STATE_DIR", "")
	t.Setenv("HOME", "/home/sally")
	t.Setenv("XDG_STATE_HOME", "")
	t.Setenv("AUX4_COMPUTER_OS", "linux")
	if got := stateDir(); got != "/home/sally/.local/state/aux4-computer" {
		t.Fatal(got)
	}
	t.Setenv("XDG_STATE_HOME", "/xdg")
	if got := stateDir(); got != "/xdg/aux4-computer" {
		t.Fatal(got)
	}
	t.Setenv("AUX4_COMPUTER_OS", "darwin")
	if got := stateDir(); got != "/home/sally/Library/Application Support/aux4-computer" {
		t.Fatal(got)
	}
	t.Setenv("AUX4_COMPUTER_STATE_DIR", "/custom")
	if got := stateDir(); got != "/custom" || useLegacyMacPaths() {
		t.Fatal(got)
	}
}
