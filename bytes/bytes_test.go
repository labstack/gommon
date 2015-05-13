package bytes

import "testing"

func TestFormat(t *testing.T) {
	// B
	f := Format(515)
	if f != "515 B" {
		t.Errorf("formatted bytes should be 515 B, found %s", f)
	}

	// MB
	f = Format(13231323)
	if f != "13.23 MB" {
		t.Errorf("formatted bytes should be 13.23 MB, found %s", f)
	}

	// Exact
	f = Format(1000 * 1000 * 1000)
	if f != "1.00 GB" {
		t.Errorf("formatted bytes should be 1.00 GB, found %s xxx", f)
	}
}

func TestFormatB(t *testing.T) {
	f := FormatB(1323)
	if f != "1.29 KiB" {
		t.Errorf("formatted bytes should be 1.29 KiB, found %s xxx", f)
	}
}
