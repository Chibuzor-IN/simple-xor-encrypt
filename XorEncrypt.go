package main

import (
	"bufio"
	"time"
	"math/rand"
	"os"
)

var key = generateKey()

func main() {	
	print("Enter input : ")
	var inputText = getInput()

	// var key int = generateKey()
	encryptedString := encryptString(key, inputText)
	decryptedString := decryptString(key, encryptedString)

	println("INPUT STRING IS :", inputText, "\n")
	println(encryptedString)
	println(decryptedString)
}

func getInput() (inputText string){
	reader := bufio.NewReader(os.Stdin)
	inputText, _ = reader.ReadString('\n')

	return
}

func generateKey() (key int){
	// Pseudorandomly generates an ascii character
	rand.Seed(time.Now().UnixNano())
	key = rand.Intn(127)

	return
}

func encryptString(key int, text string) (encryptedText string){
	key = generateKey()
	for _, element := range text {		
		replacementChar := int(element) ^ key		
		encryptedText += string(replacementChar)
	}

	return
}

func decryptString(key int, encryptedText string) (text string){
	for _, element := range encryptedText {		
		replacementChar := int(element) ^ key		
		text += string(replacementChar)
	}

	return
}
