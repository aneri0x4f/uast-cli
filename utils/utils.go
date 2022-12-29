// सखेति मत्वा प्रसभं यदुक्तं
// हे कृष्ण हे यादव हे सखेति।
// अजानता महिमानं तवेदं
// मया प्रमादात्प्रणयेन वापि॥
// यच्चावहासार्थमसत्कृतोऽसि
// विहारशय्यासनभोजनेषु।
// एकोऽथवाप्यच्युत तत्समक्षं
// तत्क्षामये त्वामहमप्रमेयम्॥

package utils

import (
	"regexp"
	"strings"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"golang.org/x/text/unicode/norm"
)

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
	ml langList = "ml"
	or langList = "or"
	te langList = "te"
	kn langList = "kn"
	ta langList = "ta"
)

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
		"ã":  "ઁ",
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
	".":  "।",
	"..": "॥",
	"au": "ã",
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

// Function to map special characters to Unicode
func CreateHandleUnicode(lang langList) func(string) string {
	langDict := charMap{}

	maps.Copy(langDict, unicodeMap)

	switch lang {
	case gu:
		maps.Copy(langDict,
			charMap{
				"0":  "૦",
				"1":  "૧",
				"2":  "૨",
				"3":  "૩",
				"4":  "૪",
				"5":  "૫",
				"6":  "૬",
				"7":  "૭",
				"8":  "૮",
				"9":  "૯",
				"om": "ॐ",
				"'":  "ઽ",
			},
		)
	case sa:
		maps.Copy(langDict,
			charMap{
				"0":  "०",
				"1":  "१",
				"2":  "२",
				"3":  "३",
				"4":  "४",
				"5":  "५",
				"6":  "६",
				"7":  "७",
				"8":  "८",
				"9":  "९",
				"om": "ॐ",
				"'":  "ऽ",
			},
		)
	case or:
		maps.Copy(langDict,
			charMap{
				"0":  "୦",
				"1":  "୧",
				"2":  "୨",
				"3":  "୩",
				"4":  "୪",
				"5":  "୫",
				"6":  "୬",
				"7":  "୭",
				"8":  "୮",
				"9":  "୯",
				"om": "ଓଁ",
				"'":  "ଽ",
			},
		)
	case kn:
		maps.Copy(langDict,
			charMap{
				"0":  "೦",
				"1":  "೧",
				"2":  "೨",
				"3":  "೩",
				"4":  "೪",
				"5":  "೫",
				"6":  "೬",
				"7":  "೭",
				"8":  "೮",
				"9":  "೯",
				"om": "ಓಂ",
				"'":  "ಽ",
			},
		)
	case te:
		maps.Copy(langDict,
			charMap{
				"0":  "౦",
				"1":  "౧",
				"2":  "౨",
				"3":  "౩",
				"4":  "౪",
				"5":  "౫",
				"6":  "౬",
				"7":  "౭",
				"8":  "౮",
				"9":  "౯",
				"'":  "ఽ",
				"om": "ఓం",
			},
		)
	case ml:
		maps.Copy(langDict,
			charMap{
				"0":  "൦",
				"1":  "൧",
				"2":  "൨",
				"3":  "൩",
				"4":  "൪",
				"5":  "൫",
				"6":  "൬",
				"7":  "൭",
				"8":  "൮",
				"9":  "൯",
				"'":  "ഽ",
				"om": "ഓം",
			},
		)
	case ta:
		maps.Copy(langDict,
			charMap{
				"0":  "௦",
				"1":  "௧",
				"2":  "௨",
				"3":  "௩",
				"4":  "௪",
				"5":  "௫",
				"6":  "௬",
				"7":  "௭",
				"8":  "௮",
				"9":  "௯",
				"'":  "𑌽",
				"om": "𑍐",
			},
		)
	default:
		panic("Unhandled case")
	}

	return func(uast string) string {
		var str []string

		for _, v := range strings.Trim(strings.ToLower(uast), "\\") {
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

		return norm.NFC.String(strings.Join(arr, ""))
	}
}

func CreateScriptFunction(lang langList) func(string) string {
	obj := charMap{}

	switch lang {
	case gu:
		maps.Copy(obj,
			charMap{
				"।": "।",
				"॥": "॥",
				"ઽ": "ऽ",
				"ॐ": "ॐ",
				"૦": "०",
				"૧": "१",
				"૨": "२",
				"૩": "३",
				"૪": "४",
				"૫": "५",
				"૬": "६",
				"૭": "७",
				"૮": "८",
				"૯": "९",
				"અ": "अ",
				"આ": "आ",
				"ઇ": "इ",
				"ઈ": "ई",
				"ઉ": "उ",
				"ઊ": "ऊ",
				"ઋ": "ऋ",
				"ૠ": "ॠ",
				"ઌ": "ऌ",
				"ૡ": "ॡ",
				"એ": "ए",
				"ઐ": "ऐ",
				"ઓ": "ओ",
				"ઔ": "औ",
				"ા": "ा",
				"િ": "ि",
				"ી": "ी",
				"ુ": "ु",
				"ૂ": "ू",
				"ૃ": "ृ",
				"ૄ": "ॄ",
				"ૢ": "ॢ",
				"ૣ": "ॣ",
				"ે": "े",
				"ૈ": "ै",
				"ો": "ो",
				"ૌ": "ौ",
				"ં": "ं",
				"ઃ": "ः",
				"ઁ": "ँ",
				"્": "्",
				"ક": "क",
				"ખ": "ख",
				"ગ": "ग",
				"ઘ": "घ",
				"ઙ": "ङ",
				"ચ": "च",
				"છ": "छ",
				"જ": "ज",
				"ઝ": "झ",
				"ઞ": "ञ",
				"ટ": "ट",
				"ઠ": "ठ",
				"ડ": "ड",
				"ઢ": "ढ",
				"ણ": "ण",
				"ત": "त",
				"થ": "थ",
				"દ": "द",
				"ધ": "ध",
				"ન": "न",
				"પ": "प",
				"ફ": "फ",
				"બ": "ब",
				"ભ": "भ",
				"મ": "म",
				"ય": "य",
				"ર": "र",
				"લ": "ल",
				"વ": "व",
				"શ": "श",
				"ષ": "ष",
				"સ": "स",
				"હ": "ह",
				"ળ": "ळ",
			},
		)
	case or:
		maps.Copy(obj,
			charMap{
				"।":  "।",
				"॥":  "॥",
				"ଽ":  "ऽ",
				"ଓଁ": "ॐ",
				"୦":  "०",
				"୧":  "१",
				"୨":  "२",
				"୩":  "३",
				"୪":  "४",
				"୫":  "५",
				"୬":  "६",
				"୭":  "७",
				"୮":  "८",
				"୯":  "९",
				"ଅ":  "अ",
				"ଆ":  "आ",
				"ଇ":  "इ",
				"ଈ":  "ई",
				"ଉ":  "उ",
				"ଊ":  "ऊ",
				"ଋ":  "ऋ",
				"ୠ":  "ॠ",
				"ଌ":  "ऌ",
				"ୡ":  "ॡ",
				"ଏ":  "ए",
				"ଐ":  "ऐ",
				"ଓ":  "ओ",
				"ଔ":  "औ",
				"ା":  "ा",
				"ି":  "ि",
				"ୀ":  "ी",
				"ୁ":  "ु",
				"ୂ":  "ू",
				"ୃ":  "ृ",
				"ୄ":  "ॄ",
				"ୢ":  "ॢ",
				"ୣ":  "ॣ",
				"େ":  "े",
				"ୈ":  "ै",
				"ୋ":  "ो",
				"ୌ":  "ौ",
				"ଂ":  "ं",
				"ଃ":  "ः",
				"ଁ":  "ँ",
				"୍":  "्",
				"କ":  "क",
				"ଖ":  "ख",
				"ଗ":  "ग",
				"ଘ":  "घ",
				"ଙ":  "ङ",
				"ଚ":  "च",
				"ଛ":  "छ",
				"ଜ":  "ज",
				"ଝ":  "झ",
				"ଞ":  "ञ",
				"ଟ":  "ट",
				"ଠ":  "ठ",
				"ଡ":  "ड",
				"ଢ":  "ढ",
				"ଣ":  "ण",
				"ତ":  "त",
				"ଥ":  "थ",
				"ଦ":  "द",
				"ଧ":  "ध",
				"ନ":  "न",
				"ପ":  "प",
				"ଫ":  "फ",
				"ବ":  "ब",
				"ଭ":  "भ",
				"ମ":  "म",
				"ୟ":  "य",
				"ର":  "र",
				"ଲ":  "ल",
				"ୱ":  "व",
				"ଶ":  "श",
				"ଷ":  "ष",
				"ସ":  "स",
				"ହ":  "ह",
				"ଳ":  "ळ",
			},
		)
	case "kn":
		maps.Copy(obj,
			charMap{
				"।":  "।",
				"॥":  "॥",
				"ಽ":  "ऽ",
				"ಓಂ": "ॐ",
				"೦":  "०",
				"೧":  "१",
				"೨":  "२",
				"೩":  "३",
				"೪":  "४",
				"೫":  "५",
				"೬":  "६",
				"೭":  "७",
				"೮":  "८",
				"೯":  "९",
				"ಅ":  "अ",
				"ಆ":  "आ",
				"ಇ":  "इ",
				"ಈ":  "ई",
				"ಉ":  "उ",
				"ಊ":  "ऊ",
				"ಋ":  "ऋ",
				"ೠ":  "ॠ",
				"ಌ":  "ऌ",
				"ೡ":  "ॡ",
				"ಎ":  "ए",
				"ಐ":  "ऐ",
				"ಒ":  "ओ",
				"ಔ":  "औ",
				"ಾ":  "ा",
				"ಿ":  "ि",
				"ೀ":  "ी",
				"ು":  "ु",
				"ೂ":  "ू",
				"ೃ":  "ृ",
				"ೄ":  "ॄ",
				"ೢ":  "ॢ",
				"ೣ":  "ॣ",
				"ೆ":  "े",
				"ೈ":  "ै",
				"ೊ":  "ो",
				"ೌ":  "ौ",
				"ಂ":  "ं",
				"ಃ":  "ः",
				"ಁ":  "ँ",
				"್":  "्",
				"ಕ":  "क",
				"ಖ":  "ख",
				"ಗ":  "ग",
				"ಘ":  "घ",
				"ಙ":  "ङ",
				"ಚ":  "च",
				"ಛ":  "छ",
				"ಜ":  "ज",
				"ಝ":  "झ",
				"ಞ":  "ञ",
				"ಟ":  "ट",
				"ಠ":  "ठ",
				"ಡ":  "ड",
				"ಢ":  "ढ",
				"ಣ":  "ण",
				"ತ":  "त",
				"ಥ":  "थ",
				"ದ":  "द",
				"ಧ":  "ध",
				"ನ":  "न",
				"ಪ":  "प",
				"ಫ":  "फ",
				"ಬ":  "ब",
				"ಭ":  "भ",
				"ಮ":  "म",
				"ಯ":  "य",
				"ರ":  "र",
				"ಲ":  "ल",
				"ವ":  "व",
				"ಶ":  "श",
				"ಷ":  "ष",
				"ಸ":  "स",
				"ಹ":  "ह",
				"ಳ":  "ळ",
			},
		)
	case te:
		maps.Copy(obj,
			charMap{
				"।":  "।",
				"॥":  "॥",
				"ఽ":  "ऽ",
				"ఓం": "ॐ",
				"౦":  "०",
				"౧":  "१",
				"౨":  "२",
				"౩":  "३",
				"౪":  "४",
				"౫":  "५",
				"౬":  "६",
				"౭":  "७",
				"౮":  "८",
				"౯":  "९",
				"అ":  "अ",
				"ఆ":  "आ",
				"ఇ":  "इ",
				"ఈ":  "ई",
				"ఉ":  "उ",
				"ఊ":  "ऊ",
				"ఋ":  "ऋ",
				"ౠ":  "ॠ",
				"ఌ":  "ऌ",
				"ౡ":  "ॡ",
				"ఎ":  "ए",
				"ఐ":  "ऐ",
				"ఒ":  "ओ",
				"ఔ":  "औ",
				"ా":  "ा",
				"ి":  "ि",
				"ీ":  "ी",
				"ు":  "ु",
				"ూ":  "ू",
				"ృ":  "ृ",
				"ౄ":  "ॄ",
				"ౢ":  "ॢ",
				"ౣ":  "ॣ",
				"ె":  "े",
				"ై":  "ै",
				"ొ":  "ो",
				"ౌ":  "ौ",
				"ం":  "ं",
				"ః":  "ः",
				"ఁ":  "ँ",
				"్":  "्",
				"క":  "क",
				"ఖ":  "ख",
				"గ":  "ग",
				"ఘ":  "घ",
				"ఙ":  "ङ",
				"చ":  "च",
				"ఛ":  "छ",
				"జ":  "ज",
				"ఝ":  "झ",
				"ఞ":  "ञ",
				"ట":  "ट",
				"ఠ":  "ठ",
				"డ":  "ड",
				"ఢ":  "ढ",
				"ణ":  "ण",
				"త":  "त",
				"థ":  "थ",
				"ద":  "द",
				"ధ":  "ध",
				"న":  "न",
				"ప":  "प",
				"ఫ":  "फ",
				"బ":  "ब",
				"భ":  "भ",
				"మ":  "म",
				"య":  "य",
				"ర":  "र",
				"ల":  "ल",
				"వ":  "व",
				"శ":  "श",
				"ష":  "ष",
				"స":  "स",
				"హ":  "ह",
				"ళ":  "ळ",
			},
		)
	case ml:
		maps.Copy(obj,
			charMap{
				"।":  "।",
				"॥":  "॥",
				"ഽ":  "ऽ",
				"ഓം": "ॐ",
				"൦":  "०",
				"൧":  "१",
				"൨":  "२",
				"൩":  "३",
				"൪":  "४",
				"൫":  "५",
				"൬":  "६",
				"൭":  "७",
				"൮":  "८",
				"൯":  "९",
				"അ":  "अ",
				"ആ":  "आ",
				"ഇ":  "इ",
				"ഈ":  "ई",
				"ഉ":  "उ",
				"ഊ":  "ऊ",
				"ഋ":  "ऋ",
				"ൠ":  "ॠ",
				"ഌ":  "ऌ",
				"ൡ":  "ॡ",
				"എ":  "ए",
				"ഐ":  "ऐ",
				"ഒ":  "ओ",
				"ഔ":  "औ",
				"ാ":  "ा",
				"ി":  "ि",
				"ീ":  "ी",
				"ു":  "ु",
				"ൂ":  "ू",
				"ൃ":  "ृ",
				"ൄ":  "ॄ",
				"ൢ":  "ॢ",
				"ൣ":  "ॣ",
				"െ":  "े",
				"ൈ":  "ै",
				"ൊ":  "ो",
				"ൗ":  "ौ",
				"ം":  "ं",
				"ഃ":  "ः",
				"ഁ":  "ँ",
				"്":  "्",
				"ക":  "क",
				"ഖ":  "ख",
				"ഗ":  "ग",
				"ഘ":  "घ",
				"ങ":  "ङ",
				"ച":  "च",
				"ഛ":  "छ",
				"ജ":  "ज",
				"ഝ":  "झ",
				"ഞ":  "ञ",
				"ട":  "ट",
				"ഠ":  "ठ",
				"ഡ":  "ड",
				"ഢ":  "ढ",
				"ണ":  "ण",
				"ത":  "त",
				"ഥ":  "थ",
				"ദ":  "द",
				"ധ":  "ध",
				"ന":  "न",
				"പ":  "प",
				"ഫ":  "फ",
				"ബ":  "ब",
				"ഭ":  "भ",
				"മ":  "म",
				"യ":  "य",
				"ര":  "र",
				"ല":  "ल",
				"വ":  "व",
				"ശ":  "श",
				"ഷ":  "ष",
				"സ":  "स",
				"ഹ":  "ह",
				"ള":  "ळ",
			},
		)
	case ta:
		maps.Copy(obj,
			charMap{
				"।": "।",
				"॥": "॥",
				"𑌽": "ऽ",
				"𑍐": "ॐ",
				"௦": "०",
				"௧": "१",
				"௨": "२",
				"௩": "३",
				"௪": "४",
				"௫": "५",
				"௬": "६",
				"௭": "७",
				"௮": "८",
				"௯": "९",
				"𑌅": "अ",
				"𑌆": "आ",
				"𑌇": "इ",
				"𑌈": "ई",
				"𑌉": "उ",
				"𑌊": "ऊ",
				"𑌋": "ऋ",
				"𑍠": "ॠ",
				"𑌌": "ऌ",
				"𑍡": "ॡ",
				"𑌏": "ए",
				"𑌐": "ऐ",
				"𑌓": "ओ",
				"𑌔": "औ",
				"𑌾": "ा",
				"𑌿": "ि",
				"𑍀": "ी",
				"𑍁": "ु",
				"𑍂": "ू",
				"𑍃": "ृ",
				"𑍄": "ॄ",
				"𑍢": "ॢ",
				"𑍣": "ॣ",
				"𑍇": "े",
				"𑍈": "ै",
				"𑍋": "ो",
				"𑍗": "ौ",
				"𑌂": "ं",
				"𑌃": "ः",
				"𑌁": "ँ",
				"𑍍": "्",
				"𑌕": "क",
				"𑌖": "ख",
				"𑌗": "ग",
				"𑌘": "घ",
				"𑌙": "ङ",
				"𑌚": "च",
				"𑌛": "छ",
				"𑌜": "ज",
				"𑌝": "झ",
				"𑌞": "ञ",
				"𑌟": "ट",
				"𑌠": "ठ",
				"𑌡": "ड",
				"𑌢": "ढ",
				"𑌣": "ण",
				"𑌤": "त",
				"𑌥": "थ",
				"𑌦": "द",
				"𑌧": "ध",
				"𑌨": "न",
				"𑌪": "प",
				"𑌫": "फ",
				"𑌬": "ब",
				"𑌭": "भ",
				"𑌮": "म",
				"𑌯": "य",
				"𑌰": "र",
				"𑌲": "ल",
				"𑌵": "व",
				"𑌶": "श",
				"𑌷": "ष",
				"𑌸": "स",
				"𑌹": "ह",
				"𑌳": "ळ",
			})
	case sa:
		maps.Copy(obj, charMap{})
	default:
		panic("Unhandled case")
	}

	return func(s string) string {
		var arr []string

		for _, v := range s {
			l := string(v)

			if v, ok := obj[l]; ok {
				arr = append(arr, v)
			}
		}

		return norm.NFC.String(strings.Join(arr, ""))
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

				if lang == sa {
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
				}

				if slices.Contains(
					[]string{",", "?", "!", "\"", ":", "(", ")", "="},
					curr,
				) {
					arr = append(arr, curr)
					i++
					continue
				}

				if slices.Contains(unAspiratedConsonants, curr) {
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

		return norm.NFC.String(strings.Join(ans, ""))
	}
}

// Convert देवनागरी to UAST
func DevanagariToUAST(data string) string {
	var str []string
	for _, v := range norm.NFC.String(data) {
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

		if slices.Contains(unAspiratedConsonants, val) && nextVal == "h" {
			arr = append(arr, val+"a")
			continue
		}

		arr = append(arr, val)
	}

	return norm.NFC.String(strings.Join(arr, ""))
}

// Convert parsed UAST string to IAST
func DataToIAST(data string) string {
	data = string(
		regexp.
			MustCompile(`[\[\]^~@#$%&*_;\n\v\t\r\f]`).
			ReplaceAll([]byte(norm.NFC.String(data)), []byte("")),
	)

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
				// arr = append(arr, "॑")
				i++
				continue
			}

			if curr == "`" {
				// arr = append(arr, "॒")
				i++
				continue
			}

			if slices.Contains(
				[]string{",", "?", "!", "\"", "-", ":", "(", ")", "="},
				curr,
			) {
				arr = append(arr, curr)
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
				if curr == "ḥ" || curr == "ṃ" || curr == "ã" {
					arr = append(arr, curr)
					i++
					continue
				}

				arr = append(arr, curr+"a")
				i++
				continue
			}

			if slices.Contains(unAspiratedConsonants, curr) && next == "h" {
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

	return norm.NFC.String(strings.Join(ans, ""))
}

// Convert IAST to UAST
func IASTToUAST(data string) string {
	var str []string
	for _, v := range string(
		regexp.
			MustCompile(`[\[\]^~@#$%&*\-_;]`).
			ReplaceAll([]byte(norm.NFC.String(data)), []byte("")),
	) {
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
			if slices.Contains(unAspiratedConsonants, curr) {
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

		val := curr

		if slices.Contains(unAspiratedConsonants, curr) &&
			k+1 < len(arr) && arr[k+1] == "h" {
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

	return norm.NFC.String(strings.Join(final, ""))
}

// Convert SLP1 to IAST
func SLPToIAST(data string) string {
	var str []string
	for _, v := range data {
		if c, ok := slpDataDict[string(v)]; ok {
			str = append(str, c)
		}
	}

	return norm.NFC.String(strings.Join(str, ""))
}

type funcList string

const (
	hu funcList = "handleUnicode"
	df funcList = "dataFunction"
)

var builderFuncs = map[langList](map[funcList](func(string) string)){
	gu: {
		hu: CreateHandleUnicode(gu),
		df: CreateDataFunction(gu),
	},
	sa: {
		hu: CreateHandleUnicode(sa),
		df: CreateDataFunction(sa),
	},
}

var Convertors = map[string](map[string]([]func(string) string)){
	"raw": {
		"iast": []func(string) string{
			builderFuncs[sa][hu],
		},
		"devanagari": []func(string) string{
			builderFuncs[sa][hu],
			IASTToUAST,
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"uast": []func(string) string{
			builderFuncs[sa][hu],
			IASTToUAST,
		},
		"guj": []func(string) string{
			builderFuncs[gu][hu],
			IASTToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
	},
	"uast": {
		"devanagari": []func(string) string{
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"iast": []func(string) string{
			builderFuncs[sa][hu],
			DataToIAST,
		},
		"guj": []func(string) string{
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
	},
	"devanagari": {
		"uast": []func(string) string{
			DevanagariToUAST,
		},
		"iast": []func(string) string{
			DevanagariToUAST,
			builderFuncs[sa][hu],
			DataToIAST,
		},
		"guj": []func(string) string{
			DevanagariToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
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
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"guj": []func(string) string{
			SLPToIAST,
			IASTToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
	},
	"iast": {
		"uast": []func(string) string{
			IASTToUAST,
		},
		"devanagari": []func(string) string{
			IASTToUAST,
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"guj": []func(string) string{
			IASTToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
	},
}
