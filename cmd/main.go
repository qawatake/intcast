package main

import (
	"github.com/qawatake/intcast"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(intcast.Analyzer) }
