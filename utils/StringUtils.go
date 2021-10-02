package utils

import (
	"math/rand"
	"time"
)

func GenerarCadena(tamano int) string {
	var letras = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return generarValor(letras, tamano)
}

func GenerarNumero() int {
	return rand.Intn(99) + 1
}

func generarValor(r []rune, n int) string {
	rand.Seed(time.Now().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = r[rand.Intn(len(r))]
	}
	return string(s)
}
