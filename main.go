package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"uast_cli/utils"
)

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
		log.Fatalf("bad `from` value: %v: expected %v", *from, schemes)
	}

	switch *to {
	case RAW, DEVANAGARI, IAST, UAST:
		buf.WriteString("`to`: " + *to + "\n")
	default:
		log.Fatalf("bad `to` value: %v: expected %v", *to, schemes)
	}

	for {
		buf.WriteString(">>> ")
		buf.Flush()

		if s, err := buf.ReadString('\n'); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}

			return
		} else {
			var arr []string
			for _, v := range strings.Split(strings.TrimSpace(s), " ") {
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
