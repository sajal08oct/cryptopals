package Set1

import (
	"fmt"
)

// Challenge #1
func TestChallenge1() {
	trial := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	dec, err := decodeHex(trial)
	if err != nil {
		fmt.Println("Error while decoding hex: s", trial)
		return
	}
	f := base64Encode(dec)
	fmt.Printf("s", f)
}

// Challenge 1# ends here

// Challenge #2 starts here

func TestChallenge2() {
	inputHex1 := []byte("1c0111001f010100061a024b53535009181c")
	inputHex2 := []byte("686974207468652062756c6c277320657965")

	input1, _ := decodeHex(inputHex1)
	input2, _ := decodeHex(inputHex2)
	tResult, _ := xorFixed(input1, input2)
	result := encodeHex(tResult)

	fmt.Printf("s", result)
}

// challenge #2 ends here

// challenge #3 starts here

func TestChallenge3() {
	temp := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	ans := singleByteXorCipher(temp)

	fmt.Printf("%s", ans)
}
