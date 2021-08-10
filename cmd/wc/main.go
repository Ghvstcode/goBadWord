package main

import (
	"github.com/Ghvstcode/wc/wordcount"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(wordcount.NewAnalyzer())
}
