package utils

import "strings"

type charMap = map[string]string

type langMap = struct {
	numbers    charMap
	vowels     charMap
	vowelSigns charMap
	consonants charMap
	misc       charMap
}

type langList string

const (
	gu langList = "gu"
	sa langList = "sa"
)

var unAspiratedConsonants = []string{
	"b",
	"c",
	"d",
	"g",
	"j",
	"k",
	"p",
	"t",
	"ḍ",
	"ṭ",
}

var devanagariNumMap = charMap{
	"0": "०",
	"1": "१",
	"2": "२",
	"3": "३",
	"4": "४",
	"5": "५",
	"6": "६",
	"7": "७",
	"8": "८",
	"9": "९",
}

var gujaratiNumMap = charMap{
	"0": "૦",
	"1": "૧",
	"2": "૨",
	"3": "૩",
	"4": "૪",
	"5": "૫",
	"6": "૬",
	"7": "૭",
	"8": "૮",
	"9": "૯",
}

var unicodeMap = charMap{
	"a":  "ā",
	"i":  "ī",
	"u":  "ū",
	"r":  "ṛ",
	"ru": "ṝ",
	"l":  "ḷ",
	"lu": "ḹ",
	"ll": "ḻ",
	"t":  "ṭ",
	"d":  "ḍ",
	"m":  "ṃ",
	"h":  "ḥ",
	"n":  "ñ",
	"nu": "ṅ",
	"nl": "ṇ",
	"su": "ś",
	"sl": "ṣ",
	"'":  "ऽ",
	".":  "।",
	"..": "॥",
	"om": "ॐ",
	"au": "ã",
}

var gujaratiCharDict = langMap{
	misc: charMap{
		"।": ".",
		"॥": "..",
		"ऽ": "'",
		"ॐ": "om",
	},
	numbers: charMap{
		"૦": "0",
		"૧": "1",
		"૨": "2",
		"૩": "3",
		"૪": "4",
		"૫": "5",
		"૬": "6",
		"૭": "7",
		"૮": "8",
		"૯": "9",
	},
	vowels: charMap{
		"a":  "અ",
		"ā":  "આ",
		"i":  "ઇ",
		"ī":  "ઈ",
		"u":  "ઉ",
		"ū":  "ઊ",
		"ṛ":  "ઋ",
		"e":  "એ",
		"ai": "ઐ",
		"o":  "ઓ",
		"au": "ઔ",
	},
	vowelSigns: charMap{
		"a":  "",
		"ā":  "ા",
		"i":  "િ",
		"ī":  "ી",
		"u":  "ુ",
		"ū":  "ૂ",
		"ṛ":  "ૃ",
		"e":  "ે",
		"ai": "ૈ",
		"o":  "ો",
		"au": "ૌ",
		"ṃ":  "ં",
		"ḥ":  "ઃ",
		"ã":  "ँ",
		"-":  "्",
	},
	consonants: charMap{
		"k":  "ક",
		"kh": "ખ",
		"g":  "ગ",
		"gh": "ઘ",
		"ṅ":  "ઙ",
		"c":  "ચ",
		"ch": "છ",
		"j":  "જ",
		"jh": "ઝ",
		"ñ":  "ઞ",
		"ṭ":  "ટ",
		"ṭh": "ઠ",
		"ḍ":  "ડ",
		"ḍh": "ઢ",
		"ṇ":  "ણ",
		"t":  "ત",
		"th": "થ",
		"d":  "દ",
		"dh": "ધ",
		"n":  "ન",
		"p":  "પ",
		"ph": "ફ",
		"b":  "બ",
		"bh": "ભ",
		"m":  "મ",
		"y":  "ય",
		"r":  "ર",
		"l":  "લ",
		"v":  "વ",
		"ś":  "શ",
		"ṣ":  "ષ",
		"s":  "સ",
		"h":  "હ",
		"ḻ":  "ળ",
	},
}

var devanagariCharDict = langMap{
	misc: charMap{
		"।": ".",
		"॥": "..",
		"ऽ": "'",
		"ॐ": "om",
	},
	numbers: charMap{
		"०": "0",
		"१": "1",
		"२": "2",
		"३": "3",
		"४": "4",
		"५": "5",
		"६": "6",
		"७": "7",
		"८": "8",
		"९": "9",
	},
	vowels: charMap{
		"a":  "अ",
		"ā":  "आ",
		"i":  "इ",
		"ī":  "ई",
		"u":  "उ",
		"ū":  "ऊ",
		"ṛ":  "ऋ",
		"ṝ":  "ॠ",
		"ḷ":  "ऌ",
		"ḹ":  "ॡ",
		"e":  "ए",
		"ai": "ऐ",
		"o":  "ओ",
		"au": "औ",
	},
	vowelSigns: charMap{
		"a":  "",
		"ā":  "ा",
		"i":  "ि",
		"ī":  "ी",
		"u":  "ु",
		"ū":  "ू",
		"ṛ":  "ृ",
		"ṝ":  "ॄ",
		"ḷ":  "ॢ",
		"ḹ":  "ॣ",
		"e":  "े",
		"ai": "ै",
		"o":  "ो",
		"au": "ौ",
		"ṃ":  "ं",
		"ḥ":  "ः",
		"ã":  "ँ",
		"-":  "्",
	},
	consonants: charMap{
		"k":  "क",
		"kh": "ख",
		"g":  "ग",
		"gh": "घ",
		"ṅ":  "ङ",
		"c":  "च",
		"ch": "छ",
		"j":  "ज",
		"jh": "झ",
		"ñ":  "ञ",
		"ṭ":  "ट",
		"ṭh": "ठ",
		"ḍ":  "ड",
		"ḍh": "ढ",
		"ṇ":  "ण",
		"t":  "त",
		"th": "थ",
		"d":  "द",
		"dh": "ध",
		"n":  "न",
		"p":  "प",
		"ph": "फ",
		"b":  "ब",
		"bh": "भ",
		"m":  "म",
		"y":  "य",
		"r":  "र",
		"l":  "ल",
		"v":  "व",
		"ś":  "श",
		"ṣ":  "ष",
		"s":  "स",
		"h":  "ह",
		"ḻ":  "ळ",
	},
}

var devanagariDataDict = charMap{
	"क": "k",
	"ख": "kh",
	"ग": "g",
	"घ": "gh",
	"ङ": "/nu/",
	"च": "c",
	"छ": "ch",
	"ज": "j",
	"झ": "jh",
	"ञ": "/n/",
	"ट": "/t/",
	"ठ": "/t/h",
	"ड": "/d/",
	"ढ": "/d/h",
	"ण": "/nl/",
	"त": "t",
	"थ": "th",
	"द": "d",
	"ध": "dh",
	"न": "n",
	"प": "p",
	"फ": "ph",
	"ब": "b",
	"भ": "bh",
	"म": "m",
	"य": "y",
	"र": "r",
	"ल": "l",
	"व": "v",
	"श": "/su/",
	"ष": "/sl/",
	"स": "s",
	"ह": "h",
	"ळ": "/ll/",
	"अ": "a",
	"आ": "/a/",
	"इ": "i",
	"ई": "/i/",
	"उ": "u",
	"ऊ": "/u/",
	"ऋ": "/r/",
	"ॠ": "/ru/",
	"ऌ": "/l/",
	"ॡ": "/lu/",
	"ए": "e",
	"ऐ": "ai",
	"ओ": "o",
	"औ": "au",
	"":  "a",
	"ा": "/a/",
	"ि": "i",
	"ी": "/i/",
	"ु": "u",
	"ू": "/u/",
	"ृ": "/r/",
	"ॄ": "/ru/",
	"ॢ": "/l/",
	"ॣ": "/lu/",
	"े": "e",
	"ै": "ai",
	"ो": "o",
	"ौ": "au",
	"ं": "/m/",
	"ः": "/h/",
	"ँ": "/au/",
	"्": "-",
	"ऽ": "\\/'/\\",
	"।": "\\/./\\",
	"॥": "\\/../\\",
	"ॐ": "\\/om/\\",
	"०": "\\/0/\\",
	"१": "\\/1/\\",
	"२": "\\/2/\\",
	"३": "\\/3/\\",
	"४": "\\/4/\\",
	"५": "\\/5/\\",
	"६": "\\/6/\\",
	"७": "\\/7/\\",
	"८": "\\/8/\\",
	"९": "\\/9/\\",
}

var iastDataDict = charMap{
	"०": "0",
	"१": "1",
	"२": "2",
	"३": "3",
	"४": "4",
	"५": "5",
	"६": "6",
	"७": "7",
	"८": "8",
	"९": "9",
	"ā": "a",
	"ī": "i",
	"ū": "u",
	"ṛ": "r",
	"ṝ": "ru",
	"ḷ": "l",
	"ḹ": "lu",
	"ḻ": "ll",
	"ṭ": "t",
	"ḍ": "d",
	"ṃ": "m",
	"ḥ": "h",
	"ñ": "n",
	"ṅ": "nu",
	"ṇ": "nl",
	"ś": "su",
	"ṣ": "sl",
	"ऽ": "'",
	"।": ".",
	"॥": "..",
	"ॐ": "om",
	"ã": "au",
}

var slpDataDict = charMap{
	"a": "a",
	"A": "ā",
	"i": "i",
	"I": "ī",
	"u": "u",
	"U": "ū",
	"e": "e",
	"E": "ai",
	"o": "o",
	"O": "au",
	"f": "ṛ",
	"F": "ṝ",
	"x": "ḷ",
	"X": "ḹ",
	"L": "ḻ",
	"|": "ḻh",
	"k": "k",
	"K": "kh",
	"g": "g",
	"G": "gh",
	"N": "ṅ",
	"c": "c",
	"C": "ch",
	"j": "j",
	"J": "jh",
	"Y": "ñ",
	"w": "ṭ",
	"W": "ṭh",
	"q": "ḍ",
	"Q": "ḍh",
	"R": "ṇ",
	"t": "t",
	"T": "th",
	"d": "d",
	"D": "dh",
	"n": "n",
	"p": "p",
	"P": "ph",
	"b": "b",
	"B": "bh",
	"m": "m",
	"M": "ṃ",
	"H": "ḥ",
	"y": "y",
	"r": "r",
	"l": "l",
	"v": "v",
	"S": "ś",
	"z": "ṣ",
	"s": "s",
	"h": "h",
	"'": "'",
	"~": "ã",
}

func in(a []string, x string) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}

	return false
}

// Function to map special characters to Unicode
func CreateHandleUnicode(lang langList) func(string) string {
	langDict := charMap{}

	for k, v := range unicodeMap {
		langDict[k] = v
	}

	switch lang {
	case gu:
		for k, v := range gujaratiNumMap {
			langDict[k] = v
		}
	default:
		for k, v := range devanagariNumMap {
			langDict[k] = v
		}
	}

	return func(uast string) string {
		uast = strings.Trim(strings.ToLower(uast), "\\")

		var str []string
		for _, v := range uast {
			str = append(str, string(v))
		}

		var arr []string

		for i := 0; i < len(str); {
			curr := str[i]

			if curr == "/" {
				var char []string

				for j := i + 1; j < len(str); j++ {
					curr := str[j]
					if curr == "/" {
						i = j
						break
					}

					if j == len(str)-1 {
						i = j
					}

					char = append(char, curr)
				}

				if v, ok := langDict[strings.Join(char, "")]; ok {
					arr = append(arr, v)
				}

				i++
				continue
			}

			arr = append(arr, curr)
			i++
		}

		return strings.Join(arr, "")
	}
}

// Function to create the function of parser
func CreateDataFunction(lang langList) func(string) string {
	obj := devanagariCharDict

	if lang == gu {
		obj = gujaratiCharDict
	}

	return func(data string) string {
		var ans []string

		for _, split := range strings.Split(data, "\\") {
			if _, ok := obj.misc[split]; ok {
				ans = append(ans, split)
				continue
			}

			if _, ok := obj.numbers[split]; ok {
				ans = append(ans, split)
				continue
			}

			if v, ok := obj.vowels[split]; ok {
				ans = append(ans, v)
				continue
			}

			var str []string
			for _, v := range split {
				str = append(str, string(v))
			}

			var arr []string
			for i := 0; i < len(str); {
				curr := str[i]
				if curr == "'" {
					arr = append(arr, "॑")
					i++
					continue
				}

				if curr == "`" {
					arr = append(arr, "॒")
					i++
					continue
				}

				if in(unAspiratedConsonants, curr) {
					var consonant string
					if i+1 < len(str) && str[i+1] == "h" {
						consonant = strings.Join(str[i:i+2], "")
						i += 2
					} else {
						consonant = curr
						i++
					}

					if v, ok := obj.consonants[consonant]; ok {
						arr = append(arr, v)
					}

					continue
				}

				if v, ok := obj.consonants[curr]; ok {
					arr = append(arr, v)
				}

				var vowel string
				if curr == "a" && (i+1 < len(str) &&
					(str[i+1] == "i" || str[i+1] == "u")) {
					vowel = strings.Join(str[i:i+2], "")
					i += 2
				} else {
					vowel = curr
					i++
				}

				if v, ok := obj.vowelSigns[vowel]; ok {
					arr = append(arr, v)
				}
			}

			ans = append(ans, strings.Join(arr, ""))
		}

		return strings.Join(ans, "")
	}
}

// Convert देवनागरी to UAST
func DevanagariToUAST(data string) string {
	var str []string
	for _, v := range data {
		str = append(str, string(v))
	}

	var arr []string

	for i := 0; i < len(str); i++ {
		curr := str[i]

		var next string
		if i+1 < len(str) {
			next = str[i+1]
		}

		if curr == "॑" {
			arr = append(arr, "\\'")
			continue
		}

		if curr == "॒" {
			arr = append(arr, "\\`")
			continue
		}

		var val string
		if v, ok := devanagariDataDict[curr]; ok {
			val = v
		} else {
			val = curr
		}

		var nextVal string
		if v, ok := devanagariDataDict[next]; ok {
			nextVal = v
		} else {
			nextVal = next
		}

		var checkVowel bool
		for _, v := range devanagariCharDict.vowels {
			if v == curr {
				checkVowel = true
				break
			}
		}

		var checkConsonant bool
		for _, v := range devanagariCharDict.consonants {
			if v == next {
				checkConsonant = true
				break
			}
		}

		if checkVowel && checkConsonant {
			arr = append(arr, val+"\\")
			continue
		}

		if in(unAspiratedConsonants, val) && nextVal == "h" {
			arr = append(arr, val+"a")
			continue
		}

		arr = append(arr, val)
	}

	return strings.Join(arr, "")
}

// Convert parsed UAST string to IAST
func DataToIAST(data string) string {
	data = strings.ReplaceAll(data, "\n", "")
	data = strings.ReplaceAll(data, "/'/", "/_/")
	data = strings.ReplaceAll(data, "/_/", "/'/")

	var ans []string

	for _, split := range strings.Split(data, "\\") {
		if split == "ॐ" {
			ans = append(ans, "oṃ")
			continue
		}

		if v, ok := devanagariCharDict.numbers[split]; ok {
			ans = append(ans, v)
			continue
		}

		if v, ok := devanagariCharDict.misc[split]; ok {
			ans = append(ans, v)
			continue
		}

		if split == "ḥ" || split == "ṃ" || split == "ã" {
			ans = append(ans, split)
			continue
		}

		var str []string
		for _, v := range split {
			str = append(str, string(v))
		}

		var arr []string
		for i := 0; i < len(str); {
			curr := str[i]

			if curr == "'" {
				// arr.push('॑');
				i++
				continue
			}

			if curr == "`" {
				// arr.push('॒');
				i++
				continue
			}

			var next string
			if i+1 < len(str) {
				next = str[i+1]
			}

			if next == "ḥ" || next == "ṃ" || next == "ã" {
				if _, ok := devanagariCharDict.consonants[curr]; ok {
					arr = append(arr, curr+"a"+next)
				} else {
					arr = append(arr, curr+next)
				}

				i += 2
				continue
			}

			if _, ok := devanagariCharDict.vowels[curr]; ok {
				arr = append(arr, curr)
				i++
				continue
			}

			if i == len(str)-1 {
				arr = append(arr, curr+"a")
				i++
				continue
			}

			if in(unAspiratedConsonants, curr) && next == "h" {
				var last string
				if i+2 < len(str) {
					last = str[i+2]
				}

				if _, ok := devanagariCharDict.vowelSigns[last]; !ok {
					arr = append(arr, curr+next+"a")
					i += 2
					continue
				}

				if last == "ḥ" || last == "ṃ" || last == "ã" {
					arr = append(arr, curr+next+"a"+last)
					i += 3
					continue
				}

				if last == "-" {
					i += 3
				} else {
					i += 2
				}
				arr = append(arr, curr+next)

				continue
			}

			if next == "-" {
				arr = append(arr, curr)
				i += 2
				continue
			}

			if _, ok := devanagariCharDict.vowelSigns[next]; ok {
				arr = append(arr, curr)
				i++
				continue
			}

			if curr == "ḥ" || curr == "ṃ" || curr == "ã" {
				arr = append(arr, curr)
				i++
				continue
			}

			arr = append(arr, curr+"a")
			i++
		}

		ans = append(ans, strings.Join(arr, ""))
	}

	return strings.Join(ans, "")
}

// Convert IAST to UAST
func IASTToUAST(data string) string {
	var str []string
	for _, v := range data {
		str = append(str, string(v))
	}

	var arr []string

	for i := 0; i < len(str); {
		curr := str[i]

		var next string
		if i+1 < len(str) {
			next = str[i+1]
		}

		if _, ok := devanagariCharDict.consonants[curr]; ok {
			if in(unAspiratedConsonants, curr) {
				if next == "a" && (i+2 < len(str) && str[i+2] == "h") {
					arr = append(arr, curr+"\\")
					i += 2
					continue
				}

				if next == "h" {
					var last string
					if i+2 < len(str) {
						last = str[i+2]
					}

					if _, ok := devanagariCharDict.consonants[last]; ok {
						arr = append(arr, curr+next+"-")
						i += 2
						continue
					}

					if last == "a" {
						if i+3 < len(str) {
							last = str[i+3]
						}
						if last == "i" || last == "u" {
							arr = append(arr, curr+next+"a"+last)
							i += 4
							continue
						}
						i += 3
					} else {
						i += 2
					}

					arr = append(arr, curr+next)

					continue
				}
			}

			if next == "a" {
				var last string
				if i+2 < len(str) {
					last = str[i+2]
				}

				if last == "i" || last == "u" {
					arr = append(arr, curr+"a"+last)
					i += 3
					continue
				}

				arr = append(arr, curr)
				i += 2
				continue
			}

			if _, ok := devanagariCharDict.consonants[next]; ok ||
				(next == "." || next == ".." || next == "'") ||
				i == len(str)-1 {
				arr = append(arr, curr+"-")
				i++
				continue
			}

			arr = append(arr, curr)
			i++
			continue
		}

		if curr == "a" && (next == "i" || next == "u") {
			arr = append(arr, curr+next+"\\")
			i += 2
			continue
		}

		if _, ok := devanagariCharDict.vowels[curr]; ok {
			if _, ok := devanagariCharDict.consonants[next]; ok {
				arr = append(arr, curr+"\\")
				i++
				continue
			}
		}

		arr = append(arr, curr)
		i++
	}

	var ans []string

	for k := 0; k < len(arr); k++ {
		curr := arr[k]

		var hasDash bool
		if strings.Contains(curr, "-") {
			hasDash = true
		}

		curr = strings.ReplaceAll(curr, "\\", "")
		curr = strings.ReplaceAll(curr, "-", "")

		for _, j := range []string{
			".",
			"'",
			"0",
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
			"9",
		} {
			if curr == "." && (k+1 < len(arr) && arr[k+1] == ".") {
				curr = strings.ReplaceAll(curr, curr, "\\/../\\")
				k++
				continue
			}

			curr = strings.ReplaceAll(curr, j, "\\/"+j+"/\\")
		}

		var val string
		if v, ok := iastDataDict[curr]; ok {
			val = "/" + v + "/"
		} else {
			val = curr
		}

		if in(unAspiratedConsonants, curr) &&
			k+1 < len(arr) && strings.Contains(arr[k+1], "h") {
			val += "a"
		}

		if hasDash {
			val += "-"
		}

		if _, ok := devanagariCharDict.vowels[curr]; ok {
			val += "\\"
		}

		ans = append(ans, val)
	}

	if len(ans) > 0 && len(str) > 0 {
		if _, ok := devanagariCharDict.consonants[ans[len(ans)-1]]; ok &&
			str[len(str)-1] != "a" {
			ans = append(ans, "-")
		}
	}

	var final []string

	for _, v := range strings.Join(ans, "") {
		l := string(v)
		if k, ok := iastDataDict[l]; ok {
			final = append(final, "/"+k+"/")
		} else {
			final = append(final, l)
		}
	}

	return strings.Join(final, "")
}

// Convert SLP1 to IAST
func SLPToIAST(data string) string {
	var str []string
	for _, v := range data {
		if c, ok := slpDataDict[string(v)]; ok {
			str = append(str, c)
		}
	}

	return strings.Join(str, "")
}

var Convertors = map[string](map[string]([]func(string) string)){
	"raw": {
		"iast": []func(string) string{
			CreateHandleUnicode(sa),
		},
	},
	"uast": {
		"devanagari": []func(string) string{
			CreateHandleUnicode(sa),
			CreateDataFunction(sa),
		},
		"iast": []func(string) string{
			CreateHandleUnicode(sa),
			DataToIAST,
		},
		"guj": []func(string) string{
			CreateHandleUnicode(gu),
			CreateDataFunction(gu),
		},
	},
	"devanagari": {
		"uast": []func(string) string{
			DevanagariToUAST,
		},
		"iast": []func(string) string{
			DevanagariToUAST,
			CreateHandleUnicode(sa),
			DataToIAST,
		},
	},
	"slp": {
		"iast": []func(string) string{
			SLPToIAST,
		},
		"uast": []func(string) string{
			SLPToIAST,
			IASTToUAST,
		},
		"devanagari": []func(string) string{
			SLPToIAST,
			IASTToUAST,
			CreateHandleUnicode(sa),
			CreateDataFunction(sa),
		},
	},
	"iast": {
		"uast": []func(string) string{
			IASTToUAST,
		},
		"devanagari": []func(string) string{
			IASTToUAST,
			CreateHandleUnicode(sa),
			CreateDataFunction(sa),
		},
	},
}
