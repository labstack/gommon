package bytes

import (
	"testing"
	"github.com/labstack/gommon/bytes"
	"fmt"
)

func TestFormat(t *testing.T) {
	fmt.Println(bytes.Format(1323))
	// Zero
	f := Format(0)
	if f != "--" {
		t.Errorf("formatted bytes should be --, found %s", f)
	}

	// B
	f = Format(515)
	if f != "515 B" {
		t.Errorf("formatted bytes should be 515 B, found %s", f)
	}

	// MB
	f = Format(13231323)
	if f != "12.62 MB" {
		t.Errorf("formatted bytes should be 12.62 MB, found %s", f)
	}

	// Exact
	f = Format(1024 * 1024 * 1024)
	if f != "1.00 GB" {
		t.Errorf("formatted bytes should be 1.00 GB, found %s", f)
	}
}

func TestBinaryPrefix(t *testing.T) {
	BinaryPrefix(true)
	f := Format(1323)
	if f != "1.29 KiB" {
		t.Errorf("formatted bytes should be 1.29 KiB, found %s", f)
	}
}
