package ick0

import "testing"

func TestTokens(t *testing.T) {
	tks := MakeTokens("tk2 TK3 tk2")
	out := tks.String()
	if out != "tk2 TK3" {
		t.Fail()
	}

	if tks.Has("tk1") != false {
		t.Fail()
	}
	if tks.Has("tk2") != true {
		t.Fail()
	}
	if tks.Remove("TK3") != true {
		t.Fail()
	}
	if tks.Toggle("tk4"); tks.String() != "tk2 tk4" {
		t.Fail()
	}
	if tks.Toggle("tk2"); tks.String() != "tk4" {
		t.Fail()
	}
	if tks.Replace("tk4", "tk5"); tks.String() != "tk5" {
		t.Fail()
	}
}
