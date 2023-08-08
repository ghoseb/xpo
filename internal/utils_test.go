package internal

import (
	"testing"
)

func TestParseExplicit0(t *testing.T) {
	opt := []string{"5,0,0"}
	res := ParseExplicit(opt)

	if res != nil {
		t.Fatalf("%s <-- should not get parsed.", opt)
	}
}

func TestParseExplicit1(t *testing.T) {
	opt := []string{"5,0,0:ERR"}
	res := ParseExplicit(opt)

	if len(res) != 1 {
		t.Fatalf("Expecting 1 match, but got %d.", len(res))
	}

	if res["ERR"] != rgbCode(5, 0, 0) {
		t.Fatalf("Parsing of %s failed!", opt)
	}
}

func TestParseExplicit2(t *testing.T) {
	opt := []string{"5,0,0:ERR", "5,4,0:WARN"}
	res := ParseExplicit(opt)

	if len(res) != 2 {
		t.Fatalf("Expecting 2 matches, but got %d.", len(res))
	}

	if res["WARN"] != rgbCode(5, 4, 0) {
		t.Fatalf("Parsing of %s failed!", opt)
	}
}

func TestParseExplicit3(t *testing.T) {
	opt := []string{"5,0,0:ERR", "5,4,0:WARN", "2,2,5:INFO"}
	res := ParseExplicit(opt)

	if len(res) != 3 {
		t.Fatalf("Expecting 3 matches, but got %d.", len(res))
	}

	if res["ERR"] != rgbCode(5, 0, 0) && res["WARN"] != rgbCode(5, 4, 0) && res["INFO"] != rgbCode(2, 2, 5) {
		t.Fatalf("Parsing of %s failed!", opt)
	}
}
