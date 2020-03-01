package protocol

import (
	"math/rand"
)

func GenerateRandomKey(length uint64) string {
	var tritSet = []rune{'1', '2', '3'}
	var lim = uint64(0)

	var key []rune

	for lim < length {
		chosenIndex := rand.Intn(3)
		key = append(key, tritSet[chosenIndex])
		lim++
	}
	return string(key)
}

func GenerateRandomBasis(length uint64) string {
	var basisSet = []rune{'A', 'B', 'C', 'D'}
	var chosenBasis []rune
	var lim = uint64(0)

	for lim < length {
		chosenIndex := rand.Intn(3)
		chosenBasis = append(chosenBasis, basisSet[chosenIndex])
		lim++
	}

	return string(chosenBasis)
}

func TritsToQutrits(key string, basis string) string {
	length := len(key)
	var qutritSet map[string]rune
	qutritSet = make(map[string]rune)
	qutritSet = GetQutrits()
	var qutrits []rune

	for index:=0; index < length; index++ {
		mKey := string([]rune{rune(key[index]), rune(basis[index])})
		qutrits = append(qutrits, qutritSet[mKey])
	}

	return string(qutrits)
}

func QutritsToTrits(key string, basis string) string {
	length := len(key)
	var tritSet map[string]rune
	tritSet = make(map[string]rune)
	tritSet = GetTrits()
	var trits []rune

	for index:=0; index < length; index++ {
		mKey := string([]rune{rune(key[index]), rune(basis[index])})
		trits = append(trits, tritSet[mKey])
	}

	return string(trits)
}


// Utilities

func GetQutrits() map[string]rune {
	var m map[string]rune
	m = make(map[string]rune)

	m["1A"] = 'a'
	m["2A"] = 'b'
	m["3A"] = 'c'
	m["1B"] = 'd'
	m["2B"] = 'e'
	m["3B"] = 'f'
	m["1C"] = 'g'
	m["2C"] = 'h'
	m["3C"] = 'i'
	m["1D"] = 'j'
	m["2D"] = 'k'
	m["3D"] = 'l'

	return m
}

func GetTrits() map[string]rune {
	var m map[string]rune
	m = make(map[string]rune)

	m["aA"] = '1'
	m["bA"] = '2'
	m["cA"] = '3'
	m["dB"] = '1'
	m["eB"] = '2'
	m["fB"] = '3'
	m["gC"] = '1'
	m["hC"] = '2'
	m["iC"] = '3'
	m["jD"] = '1'
	m["kD"] = '2'
	m["lD"] = '3'

	return m
}

