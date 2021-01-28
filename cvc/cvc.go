package cvc

import (
	"errors"
	"strings"
)

type (
	VigereChiper interface {
		Encode(string) (string, error)
		Decode(string) (string, error)
	}

	CVC struct {
		Key string
	}
)

// New for instance API reference
func New(k string) (v VigereChiper, err error) {
	cvc := &CVC{}
	k, err = sanitize(k)
	if err != nil {
		return v, err
	}
	cvc.Key = k

	return cvc, err
}

// Encode is func for encode
func (c *CVC) Encode(s string) (out string, err error) {
	s, err = sanitize(s)
	if err != nil {
		return out, err
	}
	return encrypt(s, c.Key), err
}

// Decode is func for decode
func (c *CVC) Decode(s string) (out string, err error) {
	s, err = sanitize(s)
	if err != nil {
		return out, err
	}
	return decrypt(s, c.Key), err
}

/*
	encrypt is func for encode plain text
	1. Get ASCII for key by modulos from message index
	2. Calculate with formula
	msg(i) =  (m1 + k1) mod len(alphabet)
*/
func encrypt(text, key string) (enc string) {

	var (
		encs = make([]rune, 0, len(text))
	)

	for i, v := range text {

		if v == 32 {
			encs = append(encs, v)
			continue
		}

		s := (((v - 65) + (rune(key[i%len(key)] - 65))) % 26) + 'A'
		encs = append(encs, s)
	}

	return string(encs)
}

/*
	encrypt is func for decrypt plain text
	1. Get ASCII for key by modulos from message index
	2. Calculate with formula
	msg(i) =  (m1 - k1) mod len(alphabet)
*/

func decrypt(text, key string) (enc string) {

	var (
		encs = make([]rune, 0, len(text))
	)

	for i, v := range text {

		if v == 32 {
			encs = append(encs, v)
			continue
		}

		s := ((((v - 65) - (rune(key[i%len(key)] - 65)) + 26) % 26) + 'A')
		encs = append(encs, s)
	}

	return string(encs)
}

/*
	sanitize is func for clean plain text dan keys
	For each plain text convert to ASCI for GET index
*/
func sanitize(msg string) (v string, err error) {

	plainText := []rune{}

	for _, v := range msg {
		// v is type data rune, rune is ASCII

		switch {
		case v >= 65 && v <= 90:
			// FOR ALPHABET A- Z
			plainText = append(plainText, v)
			continue
		case v >= 97 && v <= 122:
			// ALPHABET a-z convert to A-Z
			upper := toUpper(string(v))

			plainText = append(plainText, rune(upper[0]))
			continue
		case v == 32:
			plainText = append(plainText, v)
		default:
			return "", errors.New("invalid input")
		}

	}
	return string(plainText), err
}

// toUpper is func for convert to Upper
func toUpper(r string) string {
	return strings.ToUpper(r)
}
