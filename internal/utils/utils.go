// सखेति मत्वा प्रसभं यदुक्तं हे कृष्ण हे यादव हे सखेति।
// अजानता महिमानं तवेदं मया प्रमादात्प्रणयेन वापि॥
// यच्चावहासार्थमसत्कृतोऽसि विहारशय्यासनभोजनेषु।
// एकोऽथवाप्यच्युत तत्समक्षं तत्क्षामये त्वामहमप्रमेयम्॥

package utils

import (
	"maps"
	"regexp"
	"slices"
	"strings"
)

type charMap = map[string]string

type langMap = struct {
	numbers    charMap
	vowels     charMap
	vowelSigns charMap
	consonants charMap
	misc       charMap
}

type langList = string

const (
	gu langList = "gu"
	sa langList = "sa"
	ml langList = "ml"
	or langList = "or"
	te langList = "te"
	kn langList = "kn"
	ta langList = "ta"
)

var charDict = map[langList]langMap{
	gu: {
		misc: charMap{
			"।": ".",
			"॥": "..",
			"ઽ": "'",
			"ૐ": "om",
		},
		numbers: charMap{
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
		},
		vowels: charMap{
			"a":  "અ",
			"ā":  "આ",
			"i":  "ઇ",
			"ī":  "ઈ",
			"u":  "ઉ",
			"ū":  "ઊ",
			"ṛ":  "ઋ",
			"ṝ":  "ૠ",
			"ḷ":  "ઌ",
			"ḹ":  "ૡ",
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
			"ṝ":  "ૄ",
			"ḷ":  "ૢ",
			"ḹ":  "ૣ",
			"e":  "ે",
			"ai": "ૈ",
			"o":  "ો",
			"au": "ૌ",
			"ṃ":  "ં",
			"ḥ":  "ઃ",
			"ã":  "ઁ",
			"-":  "્",
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
	},
	ta: {
		misc: charMap{
			"।": ".",
			"॥": "..",
			"𑌽": "'",
			"𑍐": "om",
		},
		numbers: charMap{
			"0": "௦",
			"1": "௧",
			"2": "௨",
			"3": "௩",
			"4": "௪",
			"5": "௫",
			"6": "௬",
			"7": "௭",
			"8": "௮",
			"9": "௯",
		},
		vowels: charMap{
			"a":  "𑌅",
			"ā":  "𑌆",
			"i":  "𑌇",
			"ī":  "𑌈",
			"u":  "𑌉",
			"ū":  "𑌊",
			"ṛ":  "𑌋",
			"ṝ":  "𑍠",
			"ḷ":  "𑌌",
			"ḹ":  "𑍡",
			"e":  "𑌏",
			"ai": "𑌐",
			"o":  "𑌓",
			"au": "𑌔",
		},
		vowelSigns: charMap{
			"a":  "",
			"ā":  "𑌾",
			"i":  "𑌿",
			"ī":  "𑍀",
			"u":  "𑍁",
			"ū":  "𑍂",
			"ṛ":  "𑍃",
			"ṝ":  "𑍄",
			"ḷ":  "𑍢",
			"ḹ":  "𑍣",
			"e":  "𑍇",
			"ai": "𑍈",
			"o":  "𑍋",
			"au": "𑍗",
			"ṃ":  "𑌂",
			"ḥ":  "𑌃",
			"ã":  "𑌁",
			"-":  "𑍍",
		},
		consonants: charMap{
			"k":  "𑌕",
			"kh": "𑌖",
			"g":  "𑌗",
			"gh": "𑌘",
			"ṅ":  "𑌙",
			"c":  "𑌚",
			"ch": "𑌛",
			"j":  "𑌜",
			"jh": "𑌝",
			"ñ":  "𑌞",
			"ṭ":  "𑌟",
			"ṭh": "𑌠",
			"ḍ":  "𑌡",
			"ḍh": "𑌢",
			"ṇ":  "𑌣",
			"t":  "𑌤",
			"th": "𑌥",
			"d":  "𑌦",
			"dh": "𑌧",
			"n":  "𑌨",
			"p":  "𑌪",
			"ph": "𑌫",
			"b":  "𑌬",
			"bh": "𑌭",
			"m":  "𑌮",
			"y":  "𑌯",
			"r":  "𑌰",
			"l":  "𑌲",
			"v":  "𑌵",
			"ś":  "𑌶",
			"ṣ":  "𑌷",
			"s":  "𑌸",
			"h":  "𑌹",
			"ḻ":  "𑌳",
		},
	},
	ml: {
		misc: charMap{
			"।":  ".",
			"॥":  "..",
			"ഽ":  "'",
			"ഓം": "om",
		},
		numbers: charMap{
			"0": "൦",
			"1": "൧",
			"2": "൨",
			"3": "൩",
			"4": "൪",
			"5": "൫",
			"6": "൬",
			"7": "൭",
			"8": "൮",
			"9": "൯",
		},
		vowels: charMap{
			"a":  "അ",
			"ā":  "ആ",
			"i":  "ഇ",
			"ī":  "ഈ",
			"u":  "ഉ",
			"ū":  "ഊ",
			"ṛ":  "ഋ",
			"ṝ":  "ൠ",
			"ḷ":  "ഌ",
			"ḹ":  "ൡ",
			"e":  "എ",
			"ai": "ഐ",
			"o":  "ഒ",
			"au": "ഔ",
		},
		vowelSigns: charMap{
			"a":  "",
			"ā":  "ാ",
			"i":  "ി",
			"ī":  "ീ",
			"u":  "ു",
			"ū":  "ൂ",
			"ṛ":  "ൃ",
			"ṝ":  "ൄ",
			"ḷ":  "ൢ",
			"ḹ":  "ൣ",
			"e":  "െ",
			"ai": "ൈ",
			"o":  "ൊ",
			"au": "ൗ",
			"ṃ":  "ം",
			"ḥ":  "ഃ",
			"ã":  "ഁ",
			"-":  "്",
		},
		consonants: charMap{
			"k":  "ക",
			"kh": "ഖ",
			"g":  "ഗ",
			"gh": "ഘ",
			"ṅ":  "ങ",
			"c":  "ച",
			"ch": "ഛ",
			"j":  "ജ",
			"jh": "ഝ",
			"ñ":  "ഞ",
			"ṭ":  "ട",
			"ṭh": "ഠ",
			"ḍ":  "ഡ",
			"ḍh": "ഢ",
			"ṇ":  "ണ",
			"t":  "ത",
			"th": "ഥ",
			"d":  "ദ",
			"dh": "ധ",
			"n":  "ന",
			"p":  "പ",
			"ph": "ഫ",
			"b":  "ബ",
			"bh": "ഭ",
			"m":  "മ",
			"y":  "യ",
			"r":  "ര",
			"l":  "ല",
			"v":  "വ",
			"ś":  "ശ",
			"ṣ":  "ഷ",
			"s":  "സ",
			"h":  "ഹ",
			"ḻ":  "ള",
		},
	},
	te: {
		misc: charMap{
			"।":  ".",
			"॥":  "..",
			"ఽ":  "'",
			"ఓం": "om",
		},
		numbers: charMap{
			"0": "౦",
			"1": "౧",
			"2": "౨",
			"3": "౩",
			"4": "౪",
			"5": "౫",
			"6": "౬",
			"7": "౭",
			"8": "౮",
			"9": "౯",
		},
		vowels: charMap{
			"a":  "అ",
			"ā":  "ఆ",
			"i":  "ఇ",
			"ī":  "ఈ",
			"u":  "ఉ",
			"ū":  "ఊ",
			"ṛ":  "ఋ",
			"ṝ":  "ౠ",
			"ḷ":  "ఌ",
			"ḹ":  "ౡ",
			"e":  "ఎ",
			"ai": "ఐ",
			"o":  "ఒ",
			"au": "ఔ",
		},
		vowelSigns: charMap{
			"a":  "",
			"ā":  "ా",
			"i":  "ి",
			"ī":  "ీ",
			"u":  "ు",
			"ū":  "ూ",
			"ṛ":  "ృ",
			"ṝ":  "ౄ",
			"ḷ":  "ౢ",
			"ḹ":  "ౣ",
			"e":  "ె",
			"ai": "ై",
			"o":  "ొ",
			"au": "ౌ",
			"ṃ":  "ం",
			"ḥ":  "ః",
			"ã":  "ఁ",
			"-":  "్",
		},
		consonants: charMap{
			"k":  "క",
			"kh": "ఖ",
			"g":  "గ",
			"gh": "ఘ",
			"ṅ":  "ఙ",
			"c":  "చ",
			"ch": "ఛ",
			"j":  "జ",
			"jh": "ఝ",
			"ñ":  "ఞ",
			"ṭ":  "ట",
			"ṭh": "ఠ",
			"ḍ":  "డ",
			"ḍh": "ఢ",
			"ṇ":  "ణ",
			"t":  "త",
			"th": "థ",
			"d":  "ద",
			"dh": "ధ",
			"n":  "న",
			"p":  "ప",
			"ph": "ఫ",
			"b":  "బ",
			"bh": "భ",
			"m":  "మ",
			"y":  "య",
			"r":  "ర",
			"l":  "ల",
			"v":  "వ",
			"ś":  "శ",
			"ṣ":  "ష",
			"s":  "స",
			"h":  "హ",
			"ḻ":  "ళ",
		},
	},
	kn: {
		misc: charMap{
			"।":  ".",
			"॥":  "..",
			"ಽ":  "'",
			"ಓಂ": "om",
		},
		numbers: charMap{
			"0": "೦",
			"1": "೧",
			"2": "೨",
			"3": "೩",
			"4": "೪",
			"5": "೫",
			"6": "೬",
			"7": "೭",
			"8": "೮",
			"9": "೯",
		},
		vowels: charMap{
			"a":  "ಅ",
			"ā":  "ಆ",
			"i":  "ಇ",
			"ī":  "ಈ",
			"u":  "ಉ",
			"ū":  "ಊ",
			"ṛ":  "ಋ",
			"ṝ":  "ೠ",
			"ḷ":  "ಌ",
			"ḹ":  "ೡ",
			"e":  "ಎ",
			"ai": "ಐ",
			"o":  "ಒ",
			"au": "ಔ",
		},
		vowelSigns: charMap{
			"a":  "",
			"ā":  "ಾ",
			"i":  "ಿ",
			"ī":  "ೀ",
			"u":  "ು",
			"ū":  "ೂ",
			"ṛ":  "ೃ",
			"ṝ":  "ೄ",
			"ḷ":  "ೢ",
			"ḹ":  "ೣ",
			"e":  "ೆ",
			"ai": "ೈ",
			"o":  "ೊ",
			"au": "ೌ",
			"ṃ":  "ಂ",
			"ḥ":  "ಃ",
			"ã":  "ಁ",
			"-":  "್",
		},
		consonants: charMap{
			"k":  "ಕ",
			"kh": "ಖ",
			"g":  "ಗ",
			"gh": "ಘ",
			"ṅ":  "ಙ",
			"c":  "ಚ",
			"ch": "ಛ",
			"j":  "ಜ",
			"jh": "ಝ",
			"ñ":  "ಞ",
			"ṭ":  "ಟ",
			"ṭh": "ಠ",
			"ḍ":  "ಡ",
			"ḍh": "ಢ",
			"ṇ":  "ಣ",
			"t":  "ತ",
			"th": "ಥ",
			"d":  "ದ",
			"dh": "ಧ",
			"n":  "ನ",
			"p":  "ಪ",
			"ph": "ಫ",
			"b":  "ಬ",
			"bh": "ಭ",
			"m":  "ಮ",
			"y":  "ಯ",
			"r":  "ರ",
			"l":  "ಲ",
			"v":  "ವ",
			"ś":  "ಶ",
			"ṣ":  "ಷ",
			"s":  "ಸ",
			"h":  "ಹ",
			"ḻ":  "ಳ",
		},
	},
	or: {
		misc: charMap{
			"।":  ".",
			"॥":  "..",
			"ଽ":  "'",
			"ଓଁ": "om",
		},
		numbers: charMap{
			"0": "୦",
			"1": "୧",
			"2": "୨",
			"3": "୩",
			"4": "୪",
			"5": "୫",
			"6": "୬",
			"7": "୭",
			"8": "୮",
			"9": "୯",
		},
		vowels: charMap{
			"a":  "ଅ",
			"ā":  "ଆ",
			"i":  "ଇ",
			"ī":  "ଈ",
			"u":  "ଉ",
			"ū":  "ଊ",
			"ṛ":  "ଋ",
			"ṝ":  "ୠ",
			"ḷ":  "ଌ",
			"ḹ":  "ୡ",
			"e":  "ଏ",
			"ai": "ଐ",
			"o":  "ଓ",
			"au": "ଔ",
		},
		vowelSigns: charMap{
			"a":  "",
			"ā":  "ା",
			"i":  "ି",
			"ī":  "ୀ",
			"u":  "ୁ",
			"ū":  "ୂ",
			"ṛ":  "ୃ",
			"ṝ":  "ୄ",
			"ḷ":  "ୢ",
			"ḹ":  "ୣ",
			"e":  "େ",
			"ai": "ୈ",
			"o":  "ୋ",
			"au": "ୌ",
			"ṃ":  "ଂ",
			"ḥ":  "ଃ",
			"ã":  "ଁ",
			"-":  "୍",
		},
		consonants: charMap{
			"k":  "କ",
			"kh": "ଖ",
			"g":  "ଗ",
			"gh": "ଘ",
			"ṅ":  "ଙ",
			"c":  "ଚ",
			"ch": "ଛ",
			"j":  "ଜ",
			"jh": "ଝ",
			"ñ":  "ଞ",
			"ṭ":  "ଟ",
			"ṭh": "ଠ",
			"ḍ":  "ଡ",
			"ḍh": "ଢ",
			"ṇ":  "ଣ",
			"t":  "ତ",
			"th": "ଥ",
			"d":  "ଦ",
			"dh": "ଧ",
			"n":  "ନ",
			"p":  "ପ",
			"ph": "ଫ",
			"b":  "ବ",
			"bh": "ଭ",
			"m":  "ମ",
			"y":  "ୟ",
			"r":  "ର",
			"l":  "ଲ",
			"v":  "ୱ",
			"ś":  "ଶ",
			"ṣ":  "ଷ",
			"s":  "ସ",
			"h":  "ହ",
			"ḻ":  "ଳ",
		},
	},
	sa: {
		misc: charMap{
			"।": ".",
			"॥": "..",
			"ऽ": "'",
			"ॐ": "om",
		},
		numbers: charMap{
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

var devanāgarīDataDict = charMap{
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
	"०": "\\0\\",
	"१": "\\1\\",
	"२": "\\2\\",
	"३": "\\3\\",
	"४": "\\4\\",
	"५": "\\5\\",
	"६": "\\6\\",
	"७": "\\7\\",
	"८": "\\8\\",
	"९": "\\9\\",
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

var allowedSymbols = []string{
	"!",
	"\"",
	"(",
	")",
	",",
	"-",
	":",
	"=",
	"?",
	"|",
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

var iastAllowed = []string{
	"-",
	"a",
	"ai",
	"au",
	"b",
	"bh",
	"c",
	"ch",
	"d",
	"dh",
	"e",
	"g",
	"gh",
	"h",
	"i",
	"j",
	"jh",
	"k",
	"kh",
	"l",
	"m",
	"n",
	"o",
	"p",
	"ph",
	"r",
	"s",
	"t",
	"th",
	"u",
	"v",
	"y",
	"ã",
	"ñ",
	"ā",
	"ī",
	"ś",
	"ū",
	"ऽ",
	"ॐ",
	"।",
	"॥",
	"०",
	"१",
	"२",
	"३",
	"४",
	"५",
	"६",
	"७",
	"८",
	"९",
	"ḍ",
	"ḍh",
	"ḥ",
	"ḷ",
	"ḹ",
	"ḻ",
	"ṃ",
	"ṅ",
	"ṇ",
	"ṛ",
	"ṝ",
	"ṣ",
	"ṭ",
	"ṭh",
}

// Function to map special characters to Unicode
func createHandleUnicode(lang langList) func(string) string {
	langDict := charMap{}

	maps.Copy(langDict, unicodeMap)

	switch lang {
	case gu:
		maps.Copy(
			langDict,
			charMap{
				"om": "ૐ",
				"'":  "ઽ",
			},
		)
	case sa:
		maps.Copy(
			langDict,
			charMap{
				"om": "ॐ",
				"'":  "ऽ",
			},
		)
	case or:
		maps.Copy(
			langDict,
			charMap{
				"om": "ଓଁ",
				"'":  "ଽ",
			},
		)
	case kn:
		maps.Copy(
			langDict,
			charMap{
				"om": "ಓಂ",
				"'":  "ಽ",
			},
		)
	case te:
		maps.Copy(
			langDict,
			charMap{
				"'":  "ఽ",
				"om": "ఓం",
			},
		)
	case ml:
		maps.Copy(
			langDict,
			charMap{
				"'":  "ഽ",
				"om": "ഓം",
			},
		)
	case ta:
		maps.Copy(
			langDict,
			charMap{
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

		for i := 0; i < len(str); i++ {
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

				continue
			}

			arr = append(arr, curr)
		}

		return strings.Join(arr, "")
	}
}

// script to devanāgarī
var devanāgarīScriptDict = map[langList]charMap{
	gu: {
		"।": "।",
		"॥": "॥",
		"ઽ": "ऽ",
		"ૐ": "ॐ",
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
	or: {
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
	kn: {
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
	te: {
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
	ml: {
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
	ta: {
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
	},
	sa: {},
}

// devanāgarī to script
var reverseDevanāgarīScriptDict = (func() map[langList]charMap {
	d := map[langList]charMap{}

	for ll, v := range devanāgarīScriptDict {
		m := charMap{}
		for k, v2 := range v {
			m[v2] = k
		}

		d[ll] = m
	}

	return d
})()

func createDevanāgarīToScript(lang langList) func(string) string {
	obj := reverseDevanāgarīScriptDict[lang]
	return func(data string) string {
		var ans []string

		for _, v := range data {
			c := string(v)
			if k, ok := obj[c]; ok {
				ans = append(ans, k)
			}
		}

		return strings.Join(ans, "")
	}
}

func createScriptFunction(lang langList) func(string) string {
	obj := devanāgarīScriptDict[lang]

	return func(s string) string {
		var arr []string

		for _, v := range s {
			l := string(v)

			if k, ok := obj[l]; ok {
				arr = append(arr, k)
				continue
			}

			if _, ok := slices.BinarySearch(
				allowedSymbols,
				l,
			); ok {
				arr = append(arr, l)
			}
		}

		return strings.Join(arr, "")
	}
}

// Convert parsed UAST string to IAST
func dataToIAST(data string) string {
	data = string(
		regexp.
			MustCompile(`[\[\]{}^~@#$%&*_;.<>\n\v\t\r\f]`).
			ReplaceAll([]byte(data), []byte("")),
	)

	var ans []string

	for split := range strings.SplitSeq(data, "\\") {
		if _, ok := charDict[sa].numbers[split]; ok {
			ans = append(ans, split)
			continue
		}

		if v, ok := charDict[sa].misc[split]; ok {
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

			if _, ok := slices.BinarySearch(
				allowedSymbols,
				curr,
			); ok {
				arr = append(arr, curr)
				i++
				continue
			}

			var next string
			if i+1 < len(str) {
				next = str[i+1]
			}

			if next == "ḥ" || next == "ṃ" || next == "ã" {
				if _, ok := charDict[sa].consonants[curr]; ok {
					arr = append(arr, curr+"a"+next)
				} else {
					arr = append(arr, curr+next)
				}

				i += 2
				continue
			}

			if _, ok := charDict[sa].vowels[curr]; ok {
				arr = append(arr, curr)
				i++
				continue
			}

			if i == len(str)-1 {
				if curr == "ḥ" || curr == "ṃ" || curr == "ã" {
					arr = append(arr, curr)
				} else {
					arr = append(arr, curr+"a")
				}

				i++
				continue
			}

			if _, ok := slices.BinarySearch(unAspiratedConsonants, curr); ok && next == "h" {
				var last string
				if i+2 < len(str) {
					last = str[i+2]
				}

				if _, ok := charDict[sa].vowelSigns[last]; !ok {
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

			if _, ok := charDict[sa].vowelSigns[next]; ok {
				arr = append(arr, curr)
				i++
				continue
			}

			if curr == "ḥ" || curr == "ṃ" || curr == "ã" {
				arr = append(arr, curr)
				i++
				continue
			}

			if _, ok := slices.BinarySearch(iastAllowed, curr); !ok {
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
func iastToUAST(data string) string {
	var str []string
	for _, v := range string(
		regexp.
			MustCompile(`[\[\]{}^~@#$%&*\-_;<>]`).
			ReplaceAll([]byte(data), []byte("")),
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

		if _, ok := charDict[sa].consonants[curr]; ok {
			if _, ok := slices.BinarySearch(unAspiratedConsonants, curr); ok {
				if next == "a" && (i+2 < len(str) && str[i+2] == "h") {
					arr = append(arr, curr+"a\\")
					i += 2
					continue
				}

				if next == "h" {
					var last string
					if i+2 < len(str) {
						last = str[i+2]
					}

					if _, ok := charDict[sa].consonants[last]; ok {
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

			if _, ok := charDict[sa].consonants[next]; ok ||
				(next == "." || next == ".." || next == "'") ||
				i == len(str)-1 {
				arr = append(arr, curr+"-")
				i++
				continue
			}

			if next == "ã" {
				arr = append(arr, curr+"-", "ã")
				i += 2
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

		if _, ok := charDict[sa].vowels[curr]; ok {
			if _, ok := charDict[sa].consonants[next]; ok {
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

		hasDash := strings.Contains(curr, "-")

		curr = strings.ReplaceAll(curr, "\\", "")
		curr = strings.ReplaceAll(curr, "-", "")

		for _, j := range []string{
			".",
			"'",
		} {
			if curr == "." && (k+1 < len(arr) && arr[k+1] == ".") {
				curr = strings.ReplaceAll(curr, curr, "\\/../\\")
				k++
				continue
			}

			curr = strings.ReplaceAll(curr, j, "\\/"+j+"/\\")
		}

		for j := range charDict[sa].numbers {
			curr = strings.ReplaceAll(curr, j, "\\"+j+"\\")
		}

		val := curr

		if _, ok := slices.BinarySearch(unAspiratedConsonants, curr); ok &&
			k+1 < len(arr) && arr[k+1] == "h" {
			val += "a"
		}

		if hasDash {
			val += "-"
		}

		if _, ok := charDict[sa].vowels[curr]; ok {
			val += "\\"
		}

		ans = append(ans, val)
	}

	if len(ans) > 0 && len(str) > 0 {
		if _, ok := charDict[sa].consonants[ans[len(ans)-1]]; ok &&
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

// Function to create the function of parser
func createDataFunction(lang langList) func(string) string {
	obj := charDict[lang]

	return func(data string) string {
		var ans []string

		for split := range strings.SplitSeq(data, "\\") {
			if _, ok := obj.misc[split]; ok {
				ans = append(ans, split)
				continue
			}

			if v, ok := obj.numbers[split]; ok {
				ans = append(ans, v)
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

				if _, ok := slices.BinarySearch(
					[]string{"!", "\"", "(", ")", ",", ":", "=", "?"},
					curr,
				); ok {
					arr = append(arr, curr)
					i++
					continue
				}

				if _, ok := slices.BinarySearch(unAspiratedConsonants, curr); ok {
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
func devanāgarīToUAST(data string) string {
	var str []string
	for _, v := range data {
		str = append(str, string(v))
	}

	var arr []string

	for i := range str {
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
		if v, ok := devanāgarīDataDict[curr]; ok {
			val = v
		} else {
			val = curr
		}

		var nextVal string
		if v, ok := devanāgarīDataDict[next]; ok {
			nextVal = v
		} else {
			nextVal = next
		}

		var checkVowel bool
		for _, v := range charDict[sa].vowels {
			if v == curr {
				checkVowel = true
				break
			}
		}

		var checkConsonant bool
		for _, v := range charDict[sa].consonants {
			if v == next {
				checkConsonant = true
				break
			}
		}

		if checkVowel && checkConsonant {
			arr = append(arr, val+"\\")
			continue
		}

		if _, ok := slices.BinarySearch(unAspiratedConsonants, val); ok && nextVal == "h" {
			arr = append(arr, val+"a")
			continue
		}

		arr = append(arr, val)
	}

	return strings.Join(arr, "")
}

// Convert SLP1 to IAST
func slpToIAST(data string) string {
	var str []string
	for _, v := range data {
		if c, ok := slpDataDict[string(v)]; ok {
			str = append(str, c)
		}
	}

	return strings.Join(str, "")
}

type funcList string

const (
	hu funcList = "handleUnicode"
	df funcList = "dataFunction"
	sd funcList = "scriptTodevanāgarī"
	ds funcList = "devanāgarīToScript"
)

type builder = map[langList](map[funcList](func(string) string))

var builderFuncs = func() builder {
	m := make(builder)

	for _, v := range []langList{
		gu,
		sa,
		ml,
		or,
		te,
		kn,
		ta,
	} {
		m[v] = map[funcList]func(string) string{
			df: createDataFunction(v),
			hu: createHandleUnicode(v),
			sd: createScriptFunction(v),
			ds: createDevanāgarīToScript(v),
		}
	}

	return m
}()

var Convertors = map[string](map[string]([]func(string) string)){
	"uast-io": {
		"iast": []func(string) string{
			builderFuncs[sa][hu],
		},
		"devanāgarī": []func(string) string{
			builderFuncs[sa][hu],
			iastToUAST,
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"uast": []func(string) string{
			builderFuncs[sa][hu],
			iastToUAST,
		},
		"gu": []func(string) string{
			builderFuncs[gu][hu],
			iastToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
		"or": []func(string) string{
			builderFuncs[or][hu],
			iastToUAST,
			builderFuncs[or][hu],
			builderFuncs[or][df],
		},
		"kn": []func(string) string{
			builderFuncs[kn][hu],
			iastToUAST,
			builderFuncs[kn][hu],
			builderFuncs[kn][df],
		},
		"ml": []func(string) string{
			builderFuncs[ml][hu],
			iastToUAST,
			builderFuncs[ml][hu],
			builderFuncs[ml][df],
		},
		"ta": []func(string) string{
			builderFuncs[ta][hu],
			iastToUAST,
			builderFuncs[ta][hu],
			builderFuncs[ta][df],
		},
		"te": []func(string) string{
			builderFuncs[te][hu],
			iastToUAST,
			builderFuncs[te][hu],
			builderFuncs[te][df],
		},
	},
	"uast": {
		"devanāgarī": []func(string) string{
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"iast": []func(string) string{
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"gu": []func(string) string{
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
		"or": []func(string) string{
			builderFuncs[or][hu],
			builderFuncs[or][df],
		},
		"ta": []func(string) string{
			builderFuncs[ta][hu],
			builderFuncs[ta][df],
		},
		"te": []func(string) string{
			builderFuncs[te][hu],
			builderFuncs[te][df],
		},
		"kn": []func(string) string{
			builderFuncs[kn][hu],
			builderFuncs[kn][df],
		},
		"ml": []func(string) string{
			builderFuncs[ml][hu],
			builderFuncs[ml][df],
		},
	},
	"devanāgarī": {
		"uast": []func(string) string{
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"gu": []func(string) string{
			builderFuncs[gu][ds],
		},
		"or": []func(string) string{
			builderFuncs[or][ds],
		},
		"kn": []func(string) string{
			builderFuncs[kn][ds],
		},
		"te": []func(string) string{
			builderFuncs[te][ds],
		},
		"ta": []func(string) string{
			builderFuncs[ta][ds],
		},
		"ml": []func(string) string{
			builderFuncs[ml][ds],
		},
	},
	"slp": {
		"iast": []func(string) string{
			slpToIAST,
		},
		"uast": []func(string) string{
			slpToIAST,
			iastToUAST,
		},
		"devanāgarī": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"gu": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
		"or": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[or][hu],
			builderFuncs[or][df],
		},
		"kn": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[kn][hu],
			builderFuncs[kn][df],
		},
		"ta": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[ta][hu],
			builderFuncs[ta][df],
		},
		"te": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[te][hu],
			builderFuncs[te][df],
		},
		"ml": []func(string) string{
			slpToIAST,
			iastToUAST,
			builderFuncs[ml][hu],
			builderFuncs[ml][df],
		},
	},
	"iast": {
		"uast": []func(string) string{
			iastToUAST,
		},
		"devanāgarī": []func(string) string{
			iastToUAST,
			builderFuncs[sa][hu],
			builderFuncs[sa][df],
		},
		"gu": []func(string) string{
			iastToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
		"or": []func(string) string{
			iastToUAST,
			builderFuncs[or][hu],
			builderFuncs[or][df],
		},
		"kn": []func(string) string{
			iastToUAST,
			builderFuncs[kn][hu],
			builderFuncs[kn][df],
		},
		"ta": []func(string) string{
			iastToUAST,
			builderFuncs[ta][hu],
			builderFuncs[ta][df],
		},
		"te": []func(string) string{
			iastToUAST,
			builderFuncs[gu][hu],
			builderFuncs[gu][df],
		},
		"ml": []func(string) string{
			iastToUAST,
			builderFuncs[ml][hu],
			builderFuncs[ml][df],
		},
	},
	"gu": {
		"devanāgarī": []func(string) string{
			builderFuncs[gu][sd],
		},
		"uast": []func(string) string{
			builderFuncs[gu][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[gu][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"or": []func(string) string{
			builderFuncs[gu][sd],
			builderFuncs[or][ds],
		},
		"kn": []func(string) string{
			builderFuncs[gu][sd],
			builderFuncs[kn][ds],
		},
		"ta": []func(string) string{
			builderFuncs[gu][sd],
			builderFuncs[ta][ds],
		},
		"te": []func(string) string{
			builderFuncs[gu][sd],
			builderFuncs[te][ds],
		},
		"ml": []func(string) string{
			builderFuncs[gu][sd],
			builderFuncs[ml][ds],
		},
	},
	"or": {
		"devanāgarī": []func(string) string{
			builderFuncs[or][sd],
		},
		"uast": []func(string) string{
			builderFuncs[or][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[or][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"gu": []func(string) string{
			builderFuncs[or][sd],
			builderFuncs[gu][ds],
		},
		"kn": []func(string) string{
			builderFuncs[or][sd],
			builderFuncs[kn][sd],
		},
		"ta": []func(string) string{
			builderFuncs[or][sd],
			builderFuncs[ta][ds],
		},
		"te": []func(string) string{
			builderFuncs[or][sd],
			builderFuncs[te][ds],
		},
		"ml": []func(string) string{
			builderFuncs[or][sd],
			builderFuncs[ml][ds],
		},
	},
	"kn": {
		"devanāgarī": []func(string) string{
			builderFuncs[kn][sd],
		},
		"uast": []func(string) string{
			builderFuncs[kn][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[kn][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"or": []func(string) string{
			builderFuncs[kn][sd],
			builderFuncs[or][ds],
		},
		"gu": []func(string) string{
			builderFuncs[kn][sd],
			builderFuncs[gu][ds],
		},
		"ta": []func(string) string{
			builderFuncs[kn][sd],
			builderFuncs[ta][ds],
		},
		"te": []func(string) string{
			builderFuncs[kn][sd],
			builderFuncs[te][ds],
		},
		"ml": []func(string) string{
			builderFuncs[kn][sd],
			builderFuncs[ml][ds],
		},
	},
	"te": {
		"devanāgarī": []func(string) string{
			builderFuncs[te][sd],
		},
		"uast": []func(string) string{
			builderFuncs[te][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[te][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"or": []func(string) string{
			builderFuncs[te][sd],
			builderFuncs[or][ds],
		},
		"kn": []func(string) string{
			builderFuncs[te][sd],
			builderFuncs[kn][ds],
		},
		"ta": []func(string) string{
			builderFuncs[te][sd],
			builderFuncs[ta][ds],
		},
		"gu": []func(string) string{
			builderFuncs[te][sd],
			builderFuncs[gu][ds],
		},
		"ml": []func(string) string{
			builderFuncs[te][sd],
			builderFuncs[ml][ds],
		},
	},
	"ta": {
		"devanāgarī": []func(string) string{
			builderFuncs[ta][sd],
		},
		"uast": []func(string) string{
			builderFuncs[ta][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[ta][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"or": []func(string) string{
			builderFuncs[ta][sd],
			builderFuncs[or][ds],
		},
		"kn": []func(string) string{
			builderFuncs[ta][sd],
			builderFuncs[kn][ds],
		},
		"gu": []func(string) string{
			builderFuncs[ta][sd],
			builderFuncs[gu][ds],
		},
		"te": []func(string) string{
			builderFuncs[ta][sd],
			builderFuncs[te][ds],
		},
		"ml": []func(string) string{
			builderFuncs[ta][sd],
			builderFuncs[ml][ds],
		},
	},
	"ml": {
		"devanāgarī": []func(string) string{
			builderFuncs[ml][sd],
		},
		"uast": []func(string) string{
			builderFuncs[ml][sd],
			devanāgarīToUAST,
		},
		"iast": []func(string) string{
			builderFuncs[ml][sd],
			devanāgarīToUAST,
			builderFuncs[sa][hu],
			dataToIAST,
		},
		"or": []func(string) string{
			builderFuncs[ml][sd],
			builderFuncs[or][ds],
		},
		"kn": []func(string) string{
			builderFuncs[ml][sd],
			builderFuncs[kn][ds],
		},
		"ta": []func(string) string{
			builderFuncs[ml][sd],
			builderFuncs[ta][ds],
		},
		"te": []func(string) string{
			builderFuncs[ml][sd],
			builderFuncs[te][ds],
		},
		"gu": []func(string) string{
			builderFuncs[ml][sd],
			builderFuncs[gu][ds],
		},
	},
}
