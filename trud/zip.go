package trud

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// WalkedFile is what is passed to a client WalkFunc in order
// to permit filtering and processing of a file from the filesystem
type WalkedFile struct {
	Path    string
	Info    os.FileInfo
	zipFile *zip.File
}

func (wf *WalkedFile) isZip() bool {
	return wf.Info != nil && !wf.Info.IsDir() && strings.HasSuffix(wf.Path, ".zip")
}

// Open opens the specified file for reading
func (wf *WalkedFile) Open() (io.ReadCloser, error) {
	if wf.zipFile != nil {
		return wf.zipFile.Open()
	}
	return os.Open(wf.Path)
}

// Walk walks a filesystem including uncompressing and exploring nested zip files
func Walk(root string, walkFunc func(wf *WalkedFile, err error) error) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		walkedFile := WalkedFile{}
		walkedFile.Path = root
		walkedFile.Info = info
		if err != nil {
			return walkFunc(&walkedFile, err)
		}
		return walk(&walkedFile, walkFunc)
	})
}

func walk(wf *WalkedFile, walkFunc func(wf *WalkedFile, err error) error) error {
	if wf.isZip() {
		var zr *zip.Reader
		if wf.zipFile != nil { // we're a zip nested inside a zip
			r, err := wf.Open()
			b, err := ioutil.ReadAll(r)
			if err != nil {
				return walkFunc(wf, err)
			}
			zr, err = zip.NewReader(bytes.NewReader(b), int64(len(b)))
			if err != nil {
				return walkFunc(wf, err)
			}
		} else { // we're just a regular zip file on the filesystem
			zrc, err := zip.OpenReader(wf.Path)
			if err != nil {
				return walkFunc(wf, err)
			}
			defer zrc.Close()
			zr = &zrc.Reader
		}
		for _, f := range zr.File {
			zippedFile := new(WalkedFile)
			zippedFile.Path = filepath.Join(wf.Path, f.Name)
			zippedFile.Info = f.FileInfo()
			zippedFile.zipFile = f
			if err := walk(zippedFile, walkFunc); err != nil {
				return err
			}
		}
		return nil
	}
	return walkFunc(wf, nil)
}
