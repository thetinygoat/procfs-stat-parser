package parser

import (
	"bytes"
	"io/ioutil"
	"os"
)

func Parse() []byte {
	f, err := os.OpenFile("/proc/self/statm", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	d := bytes.Split(b, []byte(" "))
	return d[1]
}
