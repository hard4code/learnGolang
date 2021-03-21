package base

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func FnFile() () {
	pa := "/Users/mac/Documents/内容"
	files, err := ioutil.ReadDir(pa)
	for _, tmp := range files {
		re, _ := regexp.Compile("(detail).*png")
		name := re.FindString(tmp.Name())
		if len(name) != 0 {
			fmt.Println(path.Join(pa, tmp.Name()), path.Join(pa, re.FindString(tmp.Name())))
			err = os.Rename(path.Join(pa, tmp.Name()), path.Join(pa, name))
			if nil != err {
				fmt.Println(err)
			}
		}
	}
	filepath.IsAbs(pa)
	filepath.Clean(pa)
}
