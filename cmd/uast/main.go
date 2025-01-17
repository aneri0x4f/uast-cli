// Made with ♥ in Bhāratavarṣa by Aneri Dalwadi and Dhruvil Dave
// भारतवर्षे अनेर्या अनिरुद्धेन च प्रणयादेव निर्मितम्।

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"github.com/aneri0x4f/uast-cli/internal/utils"
	"golang.org/x/text/unicode/norm"
)

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
		UAST_IO    string = "uast-io"
		SLP1       string = "slp"
		GUJARATI   string = "gu"
		TAMIL      string = "ta"
		KANNADA    string = "kn"
		ODIA       string = "or"
		TELUGU     string = "te"
		MALAYALAM  string = "ml"
		DEVANĀGARĪ string = "devanāgarī"
	)

	from_schemes := []string{
		UAST,
		UAST_IO,
		DEVANĀGARĪ,
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
		DEVANĀGARĪ,
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
		UAST_IO,
		fmt.Sprintf(
			"from schema (%v)",
			from_schemes,
		),
	)
	to := flag.String(
		"to",
		DEVANĀGARĪ,
		fmt.Sprintf(
			"to schema (%v)",
			to_schemes,
		),
	)

	input := flag.String("i", "", "Input file")
	output := flag.String("o", "", "Output file")
	ver := flag.Bool("v", false, "version")

	flag.Parse()

	if *from == "devanagari" {
		*from = DEVANĀGARĪ
	}
	if *to == "devanagari" {
		*to = DEVANĀGARĪ
	}

	buf := bufio.NewReadWriter(
		bufio.NewReader(os.Stdin),
		bufio.NewWriter(os.Stdout),
	)

	if *ver {
		var commit string
		var buildtime string
		if info, ok := debug.ReadBuildInfo(); ok {
			for _, v := range info.Settings {
				if v.Key == "vcs.revision" {
					commit = v.Value
				} else if v.Key == "vcs.time" {
					buildtime = v.Value
				}
			}
		}
		if commit != "" && buildtime != "" {
			writeBuf(buf, "git commit hash: `"+commit+"`\n")
			writeBuf(buf, "git build datetime: `"+buildtime+"`\n")
		}
		writeBuf(buf, "For web version, visit `https://uast.dev`\n")
		writeBuf(buf, "For citations, visit `https://arxiv.org/abs/2203.14277`\n")
		flushBuf(buf)

		return
	}

	switch *from {
	case
		UAST,
		UAST_IO,
		DEVANĀGARĪ,
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
		log.Fatalf("bad `from` value: %v: expected %v", *from, from_schemes)
	}

	switch *to {
	case
		UAST,
		DEVANĀGARĪ,
		IAST,
		GUJARATI,
		TAMIL,
		MALAYALAM,
		KANNADA,
		TELUGU,
		ODIA:
		writeBuf(buf, "`to`: "+*to+"\n")
	default:
		log.Fatalf("bad `to` value: %v: expected %v", *to, to_schemes)
	}

	if *input != "" && *output != "" {
		f, err := os.ReadFile(*input)
		if err != nil {
			log.Fatal(err)
		}

		if k, ok := utils.Convertors[*from][*to]; ok {
			var ans []string

			for _, i := range strings.Split(norm.NFC.String(string(f)), "\n") {
				var arr []string

				for _, j := range strings.Split(i, " ") {
					for _, f := range k {
						j = f(j)
					}
					arr = append(arr, j)
				}

				ans = append(ans, norm.NFC.String(strings.Join(arr, " ")))
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

	flushBuf(buf)

	for {
		s, err := buf.ReadString('\n')
		if err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatal(err)
			}

			writeBuf(buf, "\n")
			flushBuf(buf)
			return
		}

		var arr []string
		l := strings.Split(norm.NFC.String(strings.TrimSpace(s)), " ")

		for _, v := range l {
			if k, ok := utils.Convertors[*from][*to]; ok {
				for _, f := range k {
					v = f(v)
				}
			}
			arr = append(arr, norm.NFC.String(v))
		}

		writeBuf(
			buf,
			fmt.Sprintf(
				"%v\n",
				strings.Join(arr, " "),
			),
		)

		flushBuf(buf)
	}
}
