package main

import (
	"bufio"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"uast/utils"
)

//go:embed LICENSE
var LICENSE string

//go:embed CITATIONS.md
var CITE string

func main() {
	const (
		UAST       string = "uast"
		IAST       string = "iast"
		RAW        string = "raw"
		DEVANAGARI string = "devanagari"
	)

	schemes := []string{UAST, RAW, DEVANAGARI, IAST}

	from := flag.String("from", UAST, fmt.Sprintf("from schema (%v)", schemes))
	to := flag.String("to", DEVANAGARI, fmt.Sprintf("to schema (%v)", schemes))

	input := flag.String("i", "", "Input file")
	output := flag.String("o", "", "Output file")

	flag.Parse()

	buf := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	switch *from {
	case RAW, DEVANAGARI, IAST, UAST:
		buf.WriteString("`from`: " + *from + "\n")
	default:
		log.Printf("bad `from` value: %v: expected %v", *from, schemes)
	}

	switch *to {
	case RAW, DEVANAGARI, IAST, UAST:
		buf.WriteString("`to`: " + *to + "\n")
	default:
		log.Printf("bad `to` value: %v: expected %v", *to, schemes)
	}

	if *input != "" && *output != "" {
		f, err := ioutil.ReadFile(*input)
		if err != nil {
			log.Fatal(err)
		}

		if k, ok := utils.Convertors[*from][*to]; ok {
			var ans []string

			for _, i := range strings.Split(string(f), "\n") {
				var arr []string

				for _, j := range strings.Split(i, " ") {
					for _, f := range k {
						j = f(j)
					}
					arr = append(arr, j)
				}

				ans = append(ans, strings.Join(arr, " "))
			}

			if ioutil.WriteFile(*output, []byte(strings.Join(ans, "\n")), 0666) != nil {
				log.Fatal(err)
			}
		}

		return
	}

	if (*input != "" && *output == "") || (*input == "" && *output != "") {
		log.Fatalf("Either of `-i` or `-o` was missing")
	}

	for {
		buf.WriteString(">>> ")
		buf.Flush()

		if s, err := buf.ReadString('\n'); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}

			buf.WriteString("\n")
			buf.Flush()
			return
		} else {
			var arr []string
			l := strings.Split(strings.TrimSpace(s), " ")

			if (len(l) == 2 || len(l) == 1) && strings.HasPrefix(l[0], "!") {
				switch l[0][1:] {
				case "from":
					{
						switch l[1] {
						case RAW, DEVANAGARI, IAST, UAST:
							*from = l[1]
						default:
							log.Printf("bad `from` value: %v: expected %v", *from, schemes)
						}

						buf.WriteString("`from`: " + *from + "\n")
						buf.WriteString("`to`: " + *to + "\n")
						buf.Flush()
					}
				case "to":
					{
						switch l[1] {
						case RAW, DEVANAGARI, IAST, UAST:
							*to = l[1]
						default:
							log.Printf("bad `from` value: %v: expected %v", *from, schemes)
						}

						buf.WriteString("`from`: " + *from + "\n")
						buf.WriteString("`to`: " + *to + "\n")
						buf.Flush()
					}
				case "licence", "license":
					{
						buf.WriteString(LICENSE)
						buf.Flush()
					}
				case "config":
					{
						buf.WriteString("`from`: " + *from + "\n")
						buf.WriteString("`to`: " + *to + "\n")
						buf.Flush()
					}
				case "citation":
					{
						buf.WriteString(CITE)
						buf.Flush()
					}
				case "help":
					{
						log.Printf("Available commands: %v",
							[]string{"from", "to", "license", "config", "help", "citation"},
						)
					}
				default:
					{
						log.Printf("bad config value: %v: expected %v",
							l[0],
							[]string{"from", "to", "license", "config", "help", "citation"},
						)
					}
				}

				continue
			}

			for _, v := range l {
				if k, ok := utils.Convertors[*from][*to]; ok {
					for _, f := range k {
						v = f(v)
					}
				}
				arr = append(arr, v)
			}

			buf.WriteString(strings.Join(arr, " ") + "\n")
			buf.Flush()
		}
	}
}
