package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func genWords() []string {
	// Words to use
	wordList := [100]string{"apple", "ball", "cat", "dog", "egg", "fish", "grape", "hat", "ice", "jam", "kite", "leaf", "milk", "nose", "orange", "pen", "queen", "rain", "sun", "tree", "ant", "bird", "car", "door", "ear", "frog", "gift", "hand", "ink", "jug", "key", "lamp", "moon", "nest", "ocean", "pig", "quiz", "rose", "sock", "toy", "umbrella", "van", "water", "xray", "yarn", "zebra", "baby", "cake", "doll", "elbow", "fork", "goat", "hill", "iron", "jar", "king", "lion", "mask", "nut", "owl", "pear", "rope", "shoe", "tent", "up", "vest", "wall", "yard", "zipper", "antler", "beach", "corn", "dust", "edge", "fire", "gate", "hose", "inch", "joke", "knife", "lime", "mint", "nest", "piano", "star", "rain", "salt", "tent", "vine", "wave", "yawn", "zero", "wind", "lock", "bread", "cloud", "drum", "fast", "gold", "home"}

	// Slice to hold the result
	// NOTE: This has to be type slice to be compatible with strings.join
	randWords := make([]string, 8)

	// Generate a slice of words in random order
	for i := 0; i < 8; i++ {
		randNum := rand.IntN(100)
		//fmt.Println(randNum)
		randWords[i] = wordList[randNum-1]
	}

	// Return result
	return randWords
}

func main() {
	// Generate words
	randWords := genWords()

	// Copy word slice into a string
	result := strings.Join(randWords, " ")

	fmt.Println("Generated Passphrase Successfully:", result)
}
