package fiputil

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func IsHidden(name string) bool {
	//todo: windows
	if len(name) != 0 && name[0] == '.' {
		return true
	} else {
		return false
	}
}

func CopyDir(src string, dest string, filter func(string) bool) (err error) {
	if err := os.MkdirAll(dest, 0755); err != nil {
		return err
	}

	children, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}

	for _, child := range children {
		name := child.Name()
		if filter != nil && !filter(name) {
			continue
		}
		childSrc := filepath.Join(src, name)
		childDest := filepath.Join(dest, name)
		if child.IsDir() {
			err = CopyDir(childSrc, childDest, filter)
		} else {
			err = CopyFile(childSrc, childDest)
		}
		if err != nil {
			return
		}
	}
	return
}

func CopyFile(src string, dest string) (err error) {
	//log.Println("copy file,src:",src,"dest:",dest)
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755) //os.Create(dest)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return
}

func ReadFile(path string) (content []byte, len uint32, err error) {
	file, err := os.Open(path)

	if err != nil {
		return
	}
	defer file.Close()

	bwCon := new(bytes.Buffer)
	l, err := io.Copy(bwCon, file)
	content = bwCon.Bytes()
	len = uint32(l)
	return
}

func WriteFile(path string, content []byte) (err error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer file.Close()

	file.Write(content)
	return
}

func GetExtension(path string) string {
	idx := strings.LastIndex(path, ".")
	if idx == -1 {
		return ""
	} else {
		return path[idx+1:]
	}
}

func RemoveExtension(path string) string {
	idx := strings.LastIndex(path, ".")
	if idx == -1 {
		return path
	} else {
		return path[:idx]
	}
}

func CombinePath(base string, extra ...string) {
	/*for i:=1;i<len(extra);i++ {
		if os.IsPathSeparator(base[len(base)-1]) {
			if os.IsPathSeparator(extra[i][0]) {
				base += extra[1:]
			} else {
				base += extra
			}
		} else {
			if os.IsPathSeparator(extra[i][0]) {
				base += extra
			} else {
				base += os.PathSeparator + extra
			}
		}
	}*/
}
