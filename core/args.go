package core

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	FileType   string
	InputFile  string
	OutputFile string
	Verbose    bool
	OtherArgs  []string
}

func ParseArgs() Args {
	var args Args

	flag.StringVar(&args.FileType, "type", "txt", "file type (csv, json, txt)")
	flag.StringVar(&args.InputFile, "input", "", "input file path")
	flag.StringVar(&args.OutputFile, "output", "", "output file path")
	flag.BoolVar(&args.Verbose, "verbose", false, "enable verbose mode")

	flag.Parse()

	args.OtherArgs = flag.Args()

	return args
}

func PrintUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [options] [additional arguments]\n", os.Args[0])
	flag.PrintDefaults()
}
