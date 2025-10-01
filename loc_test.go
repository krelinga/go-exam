package exam

import (
	"strings"
	"testing"
)

func TestLocString(t *testing.T) {
	tests := []struct {
		name     string
		loc      Loc
		expected string
	}{
		{
			name:     "basic location",
			loc:      Loc{File: "/path/to/file.go", Line: 42},
			expected: "/path/to/file.go:42",
		},
		{
			name:     "relative path",
			loc:      Loc{File: "file.go", Line: 1},
			expected: "file.go:1",
		},
		{
			name:     "zero line",
			loc:      Loc{File: "test.go", Line: 0},
			expected: "test.go:0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.loc.String()
			if result != tt.expected {
				t.Errorf("String() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestHere(t *testing.T) {
	// Call Here() and verify it captures the correct file and line
	loc := Here()

	// Check that the file contains "loc_test.go"
	if !strings.HasSuffix(loc.File, "loc_test.go") {
		t.Errorf("Here() captured file %q, want file ending with 'loc_test.go'", loc.File)
	}

	// The line should be 40 (where Here() was called)
	// We can't hardcode the exact line, but we can verify it's reasonable
	if loc.Line <= 0 {
		t.Errorf("Here() captured line %d, want positive line number", loc.Line)
	}

	// Verify String() works with the captured location
	str := loc.String()
	if !strings.Contains(str, "loc_test.go") {
		t.Errorf("String() = %q, want string containing 'loc_test.go'", str)
	}
	if !strings.Contains(str, ":") {
		t.Errorf("String() = %q, want string containing ':'", str)
	}
}

func TestHereLineNumber(t *testing.T) {
	// Test that consecutive calls to Here() return consecutive line numbers
	loc1 := Here()
	loc2 := Here()

	if loc1.File != loc2.File {
		t.Errorf("Here() calls returned different files: %q vs %q", loc1.File, loc2.File)
	}

	// loc2 should be exactly 1 line after loc1
	if loc2.Line != loc1.Line+1 {
		t.Errorf("Here() line numbers: loc1=%d, loc2=%d, want consecutive lines", loc1.Line, loc2.Line)
	}
}

func helperFunction() Loc {
	return Here()
}

func TestHereFromHelper(t *testing.T) {
	// Test that Here() captures the location where it's called, not where the function is defined
	loc := helperFunction()

	// Should capture the line inside helperFunction, not the line where helperFunction is called
	if !strings.HasSuffix(loc.File, "loc_test.go") {
		t.Errorf("Here() from helper captured file %q, want file ending with 'loc_test.go'", loc.File)
	}

	// The line should be inside helperFunction (around line 73)
	if loc.Line <= 0 {
		t.Errorf("Here() from helper captured line %d, want positive line number", loc.Line)
	}
}
