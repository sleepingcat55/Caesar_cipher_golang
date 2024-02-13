package main

import (
	"fmt"
	"strings"
)

var (
	alphabet = [27]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", " "}
)

func chiff_cesar(mot string, deca int) string {
	var res string
	dec := 0
	for i:=0; i<len(mot); i++ {
		for j:=0; j<len(alphabet); j++ {
			if string(mot[i]) == alphabet[j] {
				if j+deca > len(alphabet)-1 {
					// dec = len(alphabet) + deca
					// dec = dec - len(alphabet)
					dec = (j+deca)-(len(alphabet)-1)
					res += alphabet[dec-1]
					
				} else {
					res += alphabet[j+deca]
				}
			}
		}
	}
	return res
}

func dechiff_cesar(chiffre string, dec int) string {
	resu := ""
	decaa := 0
	for i:=0; i<len(chiffre); i++ {
		for j:=0; j<len(alphabet); j++ {
			if string(chiffre[i]) == alphabet[j] {
				if j-dec < 0 {
					decaa = len(alphabet) - (dec-j)
					resu += alphabet[decaa]
					// resu += "1"
				} else {
					resu += alphabet[j-dec]
				}
			}
		}
	}
	return resu
}

func main() {
	decalage := 15

	// entree := "abcdefghijklmn opqrstuvwxyz"
	entree := "Hello how are you today"
	mot_a_chiffrer := strings.ToLower(entree)
	fmt.Printf("%s\n", mot_a_chiffrer)

	fmt.Println("Chiffrement : ")
	chiffrement := chiff_cesar(mot_a_chiffrer, decalage)
	fmt.Printf("%s", chiffrement)
	fmt.Println()

	fmt.Println("Dechiffrement : ")
	dechiffrement := dechiff_cesar(chiffrement, decalage)
	fmt.Printf("%s\n", dechiffrement)
	fmt.Println()
}