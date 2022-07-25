package main

import "fmt"

func main() {
	dictionary := "CSOITEUIWUIZNSROCNKFD"
	word := "GOLANG"
	message := ""
	keyIndex := 0

	for i := 0; i < len(dictionary); i++ {

		c := dictionary[i] - 'A'
		k := word[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		message += string(c)

		keyIndex++
		keyIndex %= len(word)
	}

	fmt.Println(message)

}
