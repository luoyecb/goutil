package ufile

import (
	"bytes"
	"io"
	"os"
)

// test:TestUfile

func FileExists(fpath string) bool {
	_, err := os.Stat(fpath)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func EnsureDir(dir string) error {
	if !FileExists(dir) {
		return os.MkdirAll(dir, 0775)
	}
	return nil
}

func JoinPath(paths ...string) string {
	var buf bytes.Buffer
	for _, p := range paths {
		if p == "" {
			continue
		}
		buf.WriteString(p)
		buf.WriteByte('/')
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}

// CopyFile
// src must exists
// dst must not exists
func CopyFile(src, dst string) error {
	inFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	return err
}

func WalkDir(dir string, fn func(name string, isDir bool), ignores ...string) error {
	entrys, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	if fn == nil {
		return nil
	}

	for _, entry := range entrys {
		fname := entry.Name()
		if fname == "" || fname[0] == '.' { // ignore "." ".."
			continue
		}
		for _, ignore := range ignores { // ignore ignores
			if fname == ignore {
				continue
			}
		}
		fn(fname, entry.IsDir())
	}
	return nil
}
