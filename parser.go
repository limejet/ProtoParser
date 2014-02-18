package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"text/template"
)

const (
	DEFAULTS_FILE = "defaults.json"
	TEMPLATE_FILE = "mall.tex"
	OUT_SUFFIX    = ".tex"
)

type person struct {
	Namn  string
	Titel string
}

type punkt struct {
	Titel    string
	Brödtext string
	Beslut   string
}

func k(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specify a file name.")
		return
	}
	infile := os.Args[1]

	t := template.Must(template.New("mall.tex").Delims("#(", ")#").ParseFiles(TEMPLATE_FILE))

	file, err := os.Open(DEFAULTS_FILE)
	k(err)

	var defaults map[string]interface{}

	d := json.NewDecoder(file)
	err = d.Decode(&defaults)
	k(err)
	file.Close()

	out, err := os.Create(infile + OUT_SUFFIX)
	k(err)
	defer out.Close()

	k(t.Execute(out, defaults))

}

func ParseProtocol(filename string) (out map[string]interface{}) {
	o, err := os.Open(filename)
	k(err)
	defer o.Close()
	s := bufio.NewScanner(o)

	for s.Scan() {
		s.Text()

	}

}
