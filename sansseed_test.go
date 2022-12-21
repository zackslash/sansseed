package sansseed_test

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/zackslash/sansseed"
	"github.com/zackslash/sansseed/entropy"
	"github.com/zackslash/sansseed/lengths"
	"github.com/zackslash/sansseed/wordlists/languages"
)

var testEnt []int

func getTestEntropy(t *testing.T) []int {
	if len(testEnt) <= 0 {
		ent, err := sansseed.New12WordEntropy()
		if err != nil {
			t.Errorf("SansSeed failed to generate entropy")
		}
		testEnt = ent
	}
	return testEnt
}

func hexToBits(s string) (string, []byte, error) {
	n, err := hex.DecodeString(s)
	if err != nil {
		return "", []byte{}, err
	}
	r := ""
	for _, b := range n {
		i := fmt.Sprintf("%08b", b)
		r += i
	}
	return r, n, nil

}

func TestReferenceInputOutput(t *testing.T) {
	wk := languages.BIP39English{}
	wrdlst := wk.GetList()

	refTests := map[string]string{
		"00000000000000000000000000000000":                                 "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon about",
		"7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f":                                 "legal winner thank year wave sausage worth useful legal winner thank yellow",
		"80808080808080808080808080808080":                                 "letter advice cage absurd amount doctor acoustic avoid letter advice cage above",
		"ffffffffffffffffffffffffffffffff":                                 "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo wrong",
		"000000000000000000000000000000000000000000000000":                 "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon agent",
		"7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f":                 "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal will",
		"808080808080808080808080808080808080808080808080":                 "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter always",
		"ffffffffffffffffffffffffffffffffffffffffffffffff":                 "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo when",
		"0000000000000000000000000000000000000000000000000000000000000000": "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art",
		"7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f7f": "legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth useful legal winner thank year wave sausage worth title",
		"8080808080808080808080808080808080808080808080808080808080808080": "letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic avoid letter advice cage absurd amount doctor acoustic bless",
		"ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff": "zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo zoo vote",
		"9e885d952ad362caeb4efe34a8e91bd2":                                 "ozone drill grab fiber curtain grace pudding thank cruise elder eight picnic",
		"6610b25967cdcca9d59875f5cb50b0ea75433311869e930b":                 "gravity machine north sort system female filter attitude volume fold club stay feature office ecology stable narrow fog",
		"68a79eaca2324873eacc50cb9c6eca8cc68ea5d936f98787c60c7ebc74e6ce7c": "hamster diagram private dutch cause delay private meat slide toddler razor book happy fancy gospel tennis maple dilemma loan word shrug inflict delay length",
		"c0ba5a8e914111210f2bd131f3d5e08d":                                 "scheme spot photo card baby mountain device kick cradle pact join borrow",
		"6d9be1ee6ebd27a258115aad99b7317b9c8d28b6d76431c3":                 "horn tenant knee talent sponsor spell gate clip pulse soap slush warm silver nephew swap uncle crack brave",
		"9f6a2878b2520799a44ef18bc7df394e7061a224d2c33cd015b157d746869863": "panda eyebrow bullet gorilla call smoke muffin taste mesh discover soft ostrich alcohol speed nation flash devote level hobby quick inner drive ghost inside",
		"23db8160a31d3e0dca3688ed941adbf3":                                 "cat swing flag economy stadium alone churn speed unique patch report train",
		"8197a4a47f0425faeaa69deebc05ca29c0a5b5cc76ceacc0":                 "light rule cinnamon wrap drastic word pride squirrel upgrade then income fatal apart sustain crack supply proud access",
		"066dca1a2bb7e8a1db2832148ce9933eea0f3ac9548d793112d9a95c9407efad": "all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform",
		"f30f8c1da665478f49b001d94c5fc452":                                 "vessel ladder alter error federal sibling chat ability sun glass valve picture",
		"c10ec20dc3cd9f652c7fac2f1230f7a3c828389a14392f05":                 "scissors invite lock maple supreme raw rapid void congress muscle digital elegant little brisk hair mango congress clump",
		"f585c11aec520db57dd353c69554b21a89b20fb0650966fa0a9d6f74fd989d8f": "void come effort suffer camp survey warrior heavy shoot primary clutch crush open amazing screen patrol group space point ten exist slush involve unfold",
	}

	for ent, expect := range refTests {
		s, b, err := hexToBits(ent)
		if err != nil {
			t.Error(err)
		}

		len, err := lengths.GetLengthTypeForMnemonic(expect)
		if err != nil {
			t.Error(err)
		}

		checklen := lengths.SeedBitLength(len)
		chk, err := entropy.CheckSumBits(b, checklen)
		if err != nil {
			t.Error(err)
		}

		i, err := entropy.BinaryStringToIntSlice(s + chk)
		if err != nil {
			t.Error(err)
		}

		res := []string{}
		for _, num := range i {
			res = append(res, wrdlst[num])
		}

		mnemon := strings.Join(res, " ")
		if mnemon != expect {
			t.Errorf("mnemonic output did not match reference output for entropy value %s\n Expected:%s\n got:%s\n", ent, expect, mnemon)
		}
	}

}

func TestChineaseSimplifiedSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseChineaseSimplified(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Chinease Simplified mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Chinease Simplified\n%s", res)
}

func TestChineaseTraditionalSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseChineaseTraditional(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Chinease Traditional mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Chinease Traditional\n%s", res)
}

func TestEnglishSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseEnglish(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate english mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: English\n%s", res)
}

func TestFrenchSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseFrench(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate French mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: French\n%s", res)
}

func TestItalianSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseItalian(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Italian mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Italian\n%s", res)
}

func TestJapaneseSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseJapanese(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Japanese mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Japanese\n%s", res)
}

func TestKoreanSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseKorean(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Korean mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Korean\n%s", res)
}

func TestSpanishSeedGeneration(t *testing.T) {
	r, err := sansseed.MnemonicPhraseSpanish(getTestEntropy(t), nil)
	if err != nil {
		t.Errorf("SansSeed failed to generate Spanish mnemonic")
	}

	res := ""
	for i, word := range r {
		res += fmt.Sprintf("%d.%s ", (i + 1), word)
	}
	t.Logf("Language: Spanish\n%s", res)
}

func TestNewMnemonicPhraseForLanguage(t *testing.T) {
	expectedLen := 24
	r, err := sansseed.New24WordMnemonicPhraseForLanguage("french")
	if err != nil {
		t.Errorf("SansSeed failed to generate a new mnemonic using NewMnemonicPhraseForLanguage: %s", err.Error())
	}

	st := strings.Split(r, " ")
	if len(st) != expectedLen {
		t.Errorf("NewMnemonicPhraseForLanguage mnemonic does not match. expected: %d got: %d", expectedLen, len(st))
	}

	t.Logf("Result: %s", r)
}

func TestGeneration(t *testing.T) {
	// Generate new random entropy
	ent, err := sansseed.NewWordEntropy(lengths.SeedBitLength(lengths.TwentyfourWordSeed))
	if err != nil {
		t.Errorf(err.Error())
	}

	// Generate an english mnemonic using the new entropy (Note this package currently supports 8 languages)
	res, err := sansseed.MnemonicPhraseForLanguage(ent, languages.BIP39English{})
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Printf("%s", res)
}

func TestGenerationZeroEnglish(t *testing.T) {
	expectedOut := "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon"

	// Generate seed all zeroes
	e, err := entropy.GetRandomEntropyBytesWithCheckSum(lengths.TwentyfourWordSeed)
	if err != nil {
		t.Errorf(err.Error())
	}

	ints, err := entropy.BinaryStringToIntSlice(e)
	if err != nil {
		t.Errorf(err.Error())
	}

	zeroEnt := []int{}
	for range ints {
		zeroEnt = append(zeroEnt, 0)
	}

	// Generate an english mnemonic using the new entropy (Note this package currently supports 8 languages)
	res, err := sansseed.MnemonicPhraseForLanguage(zeroEnt, languages.BIP39English{})
	if err != nil {
		t.Errorf(err.Error())
	}

	if expectedOut != strings.Join(res, " ") {
		t.Errorf("generation result did not match expected value")
	}
}
