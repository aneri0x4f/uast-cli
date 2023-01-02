// Made with ♥ in Bhāratavarṣa by Aneri Dalwadi and Dhruvil Dave
// भारतवर्षे अनेर्या अनिरुद्धेन च प्रणयात् एव निर्मित।

package main

import (
	"bufio"
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aneri0x4f/uast-cli/utils"
)

//go:embed LICENSE
var LICENSE string

//go:embed CITATIONS.md
var CITE string

func writeBuf(buf *bufio.ReadWriter, s string) {
	if _, err := buf.WriteString(s); err != nil {
		log.Fatal(err)
	}
}

func flushBuf(buf *bufio.ReadWriter) {
	if err := buf.Flush(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	const (
		UAST       string = "uast"
		IAST       string = "iast"
		RAW        string = "raw"
		SLP1       string = "slp"
		GUJARATI   string = "gu"
		TAMIL      string = "ta"
		KANNADA    string = "kn"
		ODIA       string = "or"
		TELUGU     string = "te"
		MALAYALAM  string = "ml"
		DEVANAGARI string = "devanagari"
	)

	from_schemes := []string{
		UAST,
		RAW,
		DEVANAGARI,
		IAST,
		SLP1,
		GUJARATI,
		ODIA,
		TAMIL,
		TELUGU,
		MALAYALAM,
		KANNADA,
	}
	to_schemes := []string{
		UAST,
		DEVANAGARI,
		IAST,
		GUJARATI,
		TAMIL,
		MALAYALAM,
		KANNADA,
		TELUGU,
		ODIA,
	}

	from := flag.String(
		"from",
		UAST,
		fmt.Sprintf(
			"from schema (%v)",
			from_schemes,
		),
	)
	to := flag.String(
		"to",
		DEVANAGARI,
		fmt.Sprintf(
			"to schema (%v)",
			to_schemes,
		),
	)

	input := flag.String("i", "", "Input file")
	output := flag.String("o", "", "Output file")

	flag.Parse()

	buf := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	switch *from {
	case
		UAST,
		RAW,
		DEVANAGARI,
		IAST,
		SLP1,
		GUJARATI,
		ODIA,
		TAMIL,
		TELUGU,
		MALAYALAM,
		KANNADA:
		writeBuf(buf, "`from`: "+*from+"\n")
	default:
		log.Printf("bad `from` value: %v: expected %v", *from, from_schemes)
	}

	switch *to {
	case
		UAST,
		DEVANAGARI,
		IAST,
		GUJARATI,
		TAMIL,
		MALAYALAM,
		KANNADA,
		TELUGU,
		ODIA:
		writeBuf(buf, "`to`: "+*to+"\n")
	default:
		log.Printf("bad `to` value: %v: expected %v", *to, to_schemes)
	}

	if *input != "" && *output != "" {
		f, err := os.ReadFile(*input)
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

			if os.WriteFile(*output, []byte(strings.Join(ans, "\n")), 0666) != nil {
				log.Fatal(err)
			}
		}

		return
	}

	if (*input != "" && *output == "") || (*input == "" && *output != "") {
		log.Fatalf("Either of `-i` or `-o` was missing")
	}

	var idx int
	for {
		writeBuf(
			buf,
			fmt.Sprintf(
				"\n\033[32mIn [\033[00m\033[01;32m%v\033[00m\033[32m]:\033[00m ",
				idx,
			),
		)

		flushBuf(buf)

		if s, err := buf.ReadString('\n'); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}

			writeBuf(buf, "\n")
			flushBuf(buf)
			return
		} else {
			var arr []string
			l := strings.Split(strings.TrimSpace(s), " ")

			if (len(l) == 2 || len(l) == 1) && strings.HasPrefix(l[0], "!") {
				switch l[0][1:] {
				case "from":
					{
						if len(l) == 1 {
							l = append(l, *from)
						}

						switch l[1] {
						case
							UAST,
							RAW,
							DEVANAGARI,
							IAST,
							SLP1,
							GUJARATI,
							ODIA,
							TAMIL,
							TELUGU,
							MALAYALAM,
							KANNADA:
							*from = l[1]
						default:
							log.Printf("bad `from` value: %v: expected %v", *from, from_schemes)
						}

						writeBuf(buf, "`from`: "+*from+"\n")
						writeBuf(buf, "`to`: "+*to+"\n")
						flushBuf(buf)
					}
				case "to":
					{
						if len(l) == 1 {
							l = append(l, *to)
						}

						switch l[1] {
						case
							UAST,
							DEVANAGARI,
							IAST,
							GUJARATI,
							TAMIL,
							MALAYALAM,
							KANNADA,
							TELUGU,
							ODIA:
							*to = l[1]
						default:
							log.Printf("bad `to` value: %v: expected %v", *to, to_schemes)
						}

						writeBuf(buf, "`from`: "+*from+"\n")
						writeBuf(buf, "`to`: "+*to+"\n")
						flushBuf(buf)
					}
				case "licence", "license":
					{
						writeBuf(buf, LICENSE)
						flushBuf(buf)
					}
				case "config":
					{
						writeBuf(buf, "`from`: "+*from+"\n")
						writeBuf(buf, "`to`: "+*to+"\n")
						flushBuf(buf)
					}
				case "citation":
					{
						writeBuf(buf, CITE)
						flushBuf(buf)
					}
				case "help":
					{
						log.Printf(
							"Available commands: %v",
							[]string{"from", "to", "license", "config", "help", "citation"},
						)
					}
				default:
					{
						log.Printf(
							"bad config value: %v: expected %v",
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

			writeBuf(
				buf,
				fmt.Sprintf(
					"\033[31mOut[\033[00m\033[01;31m%v\033[00m\033[31m]:\033[00m %v\n",
					idx, strings.Join(arr, " "),
				),
			)

			flushBuf(buf)
			idx++
		}
	}
}
