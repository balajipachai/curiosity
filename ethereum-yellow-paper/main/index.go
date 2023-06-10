package main

import (
	"fmt"

	"example.com/trie"
)

func main() {
	fmt.Println("********************Creating a new Trie with `balaji`, `ball`, `apple`, `app` & `oppo`********************")
	trieDS := trie.NewTrie()

	trieDS.Insert("balaji")
	trieDS.Insert("ball")
	trieDS.Insert("apple")
	trieDS.Insert("app")
	trieDS.Insert("oppo")

	fmt.Println("\n********************Printing the Trie After Insertion********************")
	trieDS.Print(trieDS.Root, "")

	fmt.Println("********************Deleting from the Trie********************")
	fmt.Println("Deletes `app`, `ball`")
	trieDS.Delete("app")
	trieDS.Delete("ball")

	fmt.Println("\n********************Printing the Trie After Deletion********************")
	trieDS.Print(trieDS.Root, "")

	fmt.Println("********************Search `app` & `ball` in the Trie returns `false`********************")
	fmt.Println("\ttrie.Search(app) = ", trieDS.Search("app"))
	fmt.Println("\ttrie.Search(ball) = ", trieDS.Search("ball"))

	fmt.Println("\n********************Search `balaji` in the Trie returns `true`********************")
	fmt.Println("\ttrie.Search(balaji) = ", trieDS.Search("balaji"))
}
