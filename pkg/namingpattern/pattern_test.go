package namingpattern

import "testing"

func TestIsValidHTMLName(t *testing.T) {

	tests := []struct {
		In     string
		Target bool
	}{
		{In: "", Target: false},
		{In: " ", Target: false},
		{In: ".", Target: false},
		{In: ">", Target: false},
		{In: "1", Target: false},

		{In: "_", Target: true},
		{In: "a", Target: true},
		{In: "À", Target: true},

		{In: "name", Target: true},
		{In: "name1", Target: true},
		{In: "ÀÀÀ", Target: true},
		{In: "_n.ame", Target: true},

		{In: "n>", Target: false},
		{In: "na>", Target: false},
		{In: "na>e", Target: false},
		{In: "name2 n", Target: false},
		{In: "nam>e", Target: false},
	}

	for _, tst := range tests {
		out := IsValidName(tst.In)
		if out != tst.Target {
			t.Errorf("%q failed. target: %v --> out: %v", tst.In, tst.Target, out)
		}
	}

}

func TestCompilechartset(t *testing.T) {

	tests := []struct {
		in     string
		wanted int
	}{
		{in: "", wanted: 0},
		{in: "A", wanted: 1},
		{in: "A|B", wanted: 2},
		{in: `[\xC0-\xD6]`, wanted: 1},
		{in: `[\xC0-\xD6]|[\x00F8-\x02FF]`, wanted: 2},
	}

	for _, tst := range tests {
		out, _ := compileCharset(tst.in)
		if len(out) != tst.wanted {
			t.Errorf("%q failed: out: %v -> wanted: %v", tst.in, len(out), tst.wanted)
		}
	}

}

func TestGetrune(t *testing.T) {
	tests := []struct {
		in     string
		wanted rune
	}{
		{in: " ", wanted: 32},
		{in: "A", wanted: 65},
		{in: "☺", wanted: 9786},
		{in: `\xC0`, wanted: 192},
		{in: `\xC0D`, wanted: 3085},
		{in: `\x02FF`, wanted: 767},
		{in: "", wanted: -1},
		{in: "AA", wanted: -1},
		{in: `\xGHIJ`, wanted: -1},
	}

	for _, tst := range tests {
		out, err := parseRune(tst.in)
		if tst.wanted == -1 {
			if err == nil {
				t.Errorf("%q failed: out: %v -> wanted: %v", tst.in, err, tst.wanted)
			}
		} else {
			if out != tst.wanted {
				t.Errorf("%q failed: out: %v -> wanted: %v", tst.in, out, tst.wanted)
			}
		}
	}
}
