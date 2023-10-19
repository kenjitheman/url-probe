package core

import "flag"

type Args struct {
	Source string
	File   string
	URLs   []string
}

func ParseArgs() (Args, error) {
	var args Args

	flag.StringVar(&args.Source, "source", "args", "Source of URLs (json, csv, text, args)")
	flag.StringVar(&args.File, "file", "", "File containing URLs (for json, csv, and text sources)")

	flag.Parse()

	args.URLs = flag.Args()

	return args, nil
}
