package utils

import (
	"strings"
	"testing"
)

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
			if k, ok := Convertors["uast-io"]["iast"]; ok {
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

func TestDevanāgarīToIAST(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{
			input:  "मङ्गलं भगवान्विष्णुर्मङ्गलं गरुडध्वजः। मङ्गलं पुण्डरीकाक्षो मङ्गलायतनं हरिः॥",
			output: "maṅgalaṃ bhagavānviṣṇurmaṅgalaṃ garuḍadhvajaḥ. maṅgalaṃ puṇḍarīkākṣo maṅgalāyatanaṃ hariḥ..",
		},
		{
			input:  "अग्निमीळे पुरोहितं यज्ञस्य देवमृत्विजम्। होतारं रत्नधातमम्॥ अग्निः पूर्वेभिरृषिभिरीड्यो नूतनैरूत। स देवाँ एह वक्षति॥ अग्निना रयिमश्नवत्पोषमेव दिवेदिवे। यशसं वीरवत्तमम्॥ अग्ने यं यज्ञमध्वरं विश्वतः परिभूरसि। स इद्देवेषु गच्छति॥ अग्निर्होता कविक्रतुः सत्यश्चित्रश्रवस्तमः। देवो देवेभिरा गमत्॥ यदङ्ग दाशुषे त्वमग्ने भद्रं करिष्यसि। तवेत्तत्सत्यमङ्गिरः॥ उप त्वाग्ने दिवेदिवे दोषावस्तर्धिया वयम्। नमो भरन्त एमसि॥ राजन्तमध्वराणां गोपामृतस्य दीदिविम्। वर्धमानं स्वे दमे॥ स नः पितेव सूनवेऽग्ने सूपायनो भव। सचस्वा नः स्वस्तये॥",
			output: "agnimīḻe purohitaṃ yajñasya devamṛtvijam. hotāraṃ ratnadhātamam.. agniḥ pūrvebhirṛṣibhirīḍyo nūtanairūta. sa devāã eha vakṣati.. agninā rayimaśnavatpoṣameva divedive. yaśasaṃ vīravattamam.. agne yaṃ yajñamadhvaraṃ viśvataḥ paribhūrasi. sa iddeveṣu gacchati.. agnirhotā kavikratuḥ satyaścitraśravastamaḥ. devo devebhirā gamat.. yadaṅga dāśuṣe tvamagne bhadraṃ kariṣyasi. tavettatsatyamaṅgiraḥ.. upa tvāgne divedive doṣāvastardhiyā vayam. namo bharanta emasi.. rājantamadhvarāṇāṃ gopāmṛtasya dīdivim. vardhamānaṃ sve dame.. sa naḥ piteva sūnave'gne sūpāyano bhava. sacasvā naḥ svastaye..",
		},
		{
			input:  "ॐ भूर्भुवः स्वः तत्सवितुर्वरेण्यं भर्गो देवस्य धीमहि। धियो यो नः प्रचोदयात्॥",
			output: "om bhūrbhuvaḥ svaḥ tatsaviturvareṇyaṃ bhargo devasya dhīmahi. dhiyo yo naḥ pracodayāt..",
		},
	}

	for _, tC := range testCases {
		t.Run("__"+tC.input+"__", func(t *testing.T) {
			l := strings.Split(tC.input, " ")
			m := Convertors["devanāgarī"]["iast"]
			o := []string{}

			for _, i := range l {
				for _, f := range m {
					i = f(i)
				}

				o = append(o, i)
			}

			if strings.Join(o, " ") != tC.output {
				t.Fail()
			}
		})
	}
}
