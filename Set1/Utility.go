package Set1

import (
	"encoding/base64"
	"encoding/hex"
	"errors"
	"sort"
)

func decodeHex(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)
	return eb
}

func encodeHex(input []byte) []byte {
	res := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(res, input)
	return res
}
func xorFixed(input1, input2 []byte) ([]byte, error) {
	len1 := len(input1)
	len2 := len(input2)
	if len1 != len2 {
		return nil, errors.New("len of inputs are not equal")
	}
	res := make([]byte, len1)
	for i := 0; i < len1; i++ {
		res[i] = input1[i] ^ input2[i]
	}
	return res, nil
}

var LETTER_FREQUENCY = map[string]float64{
	"a": 8.497,
	"b": 1.492,
	"c": 2.202,
	"d": 4.253,
	"e": 11.162,
	"f": 2.228,
	"g": 2.015,
	"h": 6.094,
	"i": 7.546,
	"j": 0.153,
	"k": 1.292,
	"l": 4.025,
	"m": 2.406,
	"n": 6.749,
	"o": 7.507,
	"p": 1.929,
	"q": 0.095,
	"r": 7.587,
	"s": 6.327,
	"t": 9.356,
	"u": 2.758,
	"v": 0.978,
	"w": 2.560,
	"x": 0.150,
	"y": 1.994,
	"z": 0.077,
	" ": 15.00,
}

func getFrequencyScore(text []byte) float64 {
	score := 0.0
	for i := 0; i < len(text); i++ {
		score = score + LETTER_FREQUENCY[string(text[i])]
	}
	return score
}

func singleByteXorCipher(input []byte) []byte {
	hexDecoded, _ := decodeHex(input)
	m := make(map[float64][]byte)
	var keys []float64
	for i := 0; i < 256; i++ {
		re := make([]byte, len(hexDecoded))
		for j := 0; j < len(hexDecoded); j++ {
			re[j] = hexDecoded[j] ^ byte(i)
		}
		score := getFrequencyScore(re)
		m[score] = re
		keys = append(keys, score)
	}
	sort.Float64s(keys)
	if keys != nil {
		return m[keys[len(keys)-1]]
	}
	return nil
}
