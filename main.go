package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/procfs-status-parser/parser"
)

func main() {
	p := os.Getpagesize()
	rss := parser.Parse()

	rssString := string(rss)

	i, _ := strconv.Atoi(rssString)

	fmt.Println(i * p)
}
