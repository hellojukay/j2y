package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var of = os.Stdout
var inf = os.Stdin

func init() {
	var output = flag.String("o", "", "output yaml file, default: stdout")
	var input = flag.String("i", "", "input json file, default: stdin")
	flag.Parse()
	if *output != "" {
		fh, err := os.Create(*output)
		if err != nil {
			fmt.Errorf("crate output file error, %s", err)
			os.Exit(1)
		}
		of = fh
	}
	if *input != "" {
		fh, err := os.Open(*input)
		if err != nil {
			fmt.Errorf("open input json file error, %s", err)
			os.Exit(1)
		}
		inf = fh
	}
}
func main() {
	var m interface{}
	bf, err := ioutil.ReadAll(inf)
	if err != nil {
		println("read json error")
		os.Exit(1)
	}
	err = json.Unmarshal(bf, &m)
	if err != nil {
		println("parse json error")
		os.Exit(1)
	}
	var buffer = bytes.Buffer{}
	var encoder = yaml.NewEncoder(&buffer)
	encoder.SetIndent(4)
	err = encoder.Encode(m)
	if err != nil {
		println("convert json to yaml error")
		os.Exit(1)
	}
	fmt.Fprintf(of, "%s", string(buffer.Bytes()))
}
