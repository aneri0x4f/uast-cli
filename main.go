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
	"uast/utils"
)

//go:embed LICENSE
var LICENSE string

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
				default:
					{
						log.Printf("bad config value: %v: expected %v",
							l[0],
							[]string{"from", "to", "license", "config"},
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
