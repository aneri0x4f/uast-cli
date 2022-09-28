package utils

import "testing"

func TestHandleUnicode(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "/d/",
			output: "ḍ",
		},
		{
			input:  "d",
			output: "d",
		},
	}
	for _, tC := range testCases {
		t.Run("__"+tC.input+"__", func(t *testing.T) {
			if k, ok := Convertors["raw"]["iast"]; ok {
				for _, f := range k {
					tC.input = f(tC.input)
				}
			}

			if tC.input != tC.output {
				t.Fail()
			}
		})
	}
}

func TestUASTToIAST(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "/d/",
			output: "ḍa",
		},
		{
			input:  "/d/-",
			output: "ḍ",
		},
		{
			input:  "/d/h",
			output: "ḍha",
		},
		{
			input:  "/d/h-",
			output: "ḍh",
		},
		{
			input:  "/d/i",
			output: "ḍi",
		},
		{
			input:  "/d/hi",
			output: "ḍhi",
		},
		{
			input:  "d",
			output: "da",
		},
		{
			input:  "d-",
			output: "d",
		},
		{
			input:  "dh",
			output: "dha",
		},
		{
			input:  "dh-",
			output: "dh",
		},
		{
			input:  "l/i//au/",
			output: "līã",
		},
		{
			input:  "l-/au/",
			output: "lã",
		},
	}
	for _, tC := range testCases {
		t.Run("__"+tC.input+"__", func(t *testing.T) {
			if k, ok := Convertors["uast"]["iast"]; ok {
				for _, f := range k {
					tC.input = f(tC.input)
				}
			}

			if tC.input != tC.output {
				t.Fail()
			}
		})
	}
}
