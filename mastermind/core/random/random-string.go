package random

import (
	"crypto/rand"
)

type RandType uint8

const (
	ALPHANUM RandType = iota + 1
	ALPHA
	NUMBER
	ASCII
)

func randStr(strSize int, randType RandType) string {

	var dictionary string

	switch randType {
	case ALPHANUM:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case ALPHA:
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case NUMBER:
		dictionary = "0123456789"
	case ASCII:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz" + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
	}

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

func RandAlphanum(strSize int) string {
	return randStr(strSize, ALPHANUM)
}

func RandAlpha(strSize int) string {
	return randStr(strSize, ALPHA)
}

func RandNumber(strSize int) string {
	return randStr(strSize, NUMBER)
}

func RandAscii(strSize int) string {
	return randStr(strSize, ASCII)
}
