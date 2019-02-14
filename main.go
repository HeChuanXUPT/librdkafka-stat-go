package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var version = "V0.1"
var showVersion bool
var filename string

func init() {
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.StringVar(&filename, "f", "", "filename")
}

func printVersion() {
	fmt.Println(version)
	os.Exit(0)
}

func parseFile(filename string) Stat {
	var data Stat
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(buf, &data); err != nil {
		log.Fatalln(err)
	}
	return data
}

func main() {
	flag.Parse()

	if showVersion {
		printVersion()
	}

	if len(filename) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	stat := parseFile(filename)
	fmt.Println(stat)
}
