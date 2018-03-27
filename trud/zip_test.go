package trud_test

import (
	"github.com/wardle/go-ods/trud"
	"strings"
	"testing"
)

/*
The test.zip file has the following contents:

├── f1
├── f2
├── f3
├── z1.zip
├── z2
│   └── z2.zip
└── z3.zip


z1 contains three files:
        0  03-24-2018 21:30   z1f1
        0  03-24-2018 21:30   z1f2
        0  03-24-2018 21:30   z1f3

z2 contains three files and a nested zip file, z4.zip
        0  03-24-2018 21:30   z2f1
        0  03-24-2018 21:30   z2f2
        0  03-24-2018 21:30   z2f3
	  430  03-24-2018 21:38   z4.zip

z3 contains three files in a subdirectory
        0  03-24-2018 21:30   z3/
        0  03-24-2018 21:30   z3/z3f3
        0  03-24-2018 21:30   z3/z3f2
		0  03-24-2018 21:30   z3/z3f1

z4 contains three files:
        0  03-24-2018 21:38   z4f1
        0  03-24-2018 21:38   z4f2
        0  03-24-2018 21:38   z4f3
*/

func TestWalkZip(t *testing.T) {
	found := false
	err := trud.Walk("test.zip", func(wf *trud.WalkedFile, err error) error {
		t.Logf("File: %s   -- %s (%d bytes)", wf.Path, wf.Info.Name(), wf.Info.Size())
		if err != nil {
			t.Error(err)
		}
		if strings.HasSuffix(wf.Path, "test.zip/z2/z2.zip/z4.zip/z4f1") {
			found = true
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("did not find file z4f1 in nested zip!")
	}
}
func TestWalkDirAndZip(t *testing.T) {
	found := false
	err := trud.Walk("./", func(wf *trud.WalkedFile, err error) error {
		if strings.HasSuffix(wf.Path, "test.zip/z2/z2.zip/z4.zip/z4f1") {
			found = true
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if !found {
		t.Error("did not find file z4f1 in nested zip!")
	}
}
