package main

import (
	"flag"
	"fmt"
	core "getsrv/core"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Lookup("logtostderr").Value.Set("true")
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()

	// Create DynamoDB client
	context := core.NewContext()
	server := core.NewServer(context)
	server.Start()
}
