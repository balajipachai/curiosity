package trie

import "fmt"

/*
* The TrieNode type represents a node in a Trie data structure with 26 children nodes and a boolean value indicating if it marks the end of a word.
* @property children - The `children` property is an array of pointers to `TrieNode` objects.
* Each index in the array represents a letter of the alphabet (a-z), and the corresponding pointer points to the child node that represents that letter.
* This allows for efficient traversal of the trie when searching for words.
* @property {bool} isEnd - The `isEnd` property is a boolean value that indicates whether the current node represents the end of a word in the Trie data structure.
* If `isEnd` is true, it means that the node represents the last character of a word that has been inserted into the Trie.
 */
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

/*
* The Trie type represents a trie data structure with a root node.
* @property root - The root node of the Trie data structure. It is the starting point of the Trie and does not contain any character value.
* All the actual characters are stored in the child nodes of the root node.
 */
type Trie struct {
	Root *TrieNode // Since we are accessing it from the main go module, thus root is marked with a beginning Capital letter
}

// The function creates a new instance of a Trie data structure.
func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{},
	}
}

/*
* The `Insert` function is a method of the `Trie` struct that takes a string `word` as input and inserts it into the Trie data structure.
* It does this by iterating over each character in the `word` and checking if the corresponding child node exists in the Trie.
* If the child node does not exist, it creates a new `TrieNode` and sets it as the child node.
* It then moves to the next node and repeats the process until it reaches the end of the `word`.
* Finally, it marks the last node as the end of the `word` by setting the `isEnd` property to `true`.
 */
func (trie *Trie) Insert(word string) {
	node := trie.Root

	for _, char := range word {
		/*
		* `index := char - 'a'` is converting the character `char` to an index value between 0 and 25.
		* This is done by subtracting the ASCII value of the character 'a' (97) from the ASCII value of the character `char`.
		* For example, if `char` is 'c', then `index` will be 2 (99 - 97).
		* This index value is used to access the corresponding child node in the `children` array of the current node.
		 */
		index := char - 'a'
		// Checks if node.children at index is nil or not
		// if nil then adds a Trie Node, if not, then
		// moves the node to the next one
		if node.children[index] == nil {
			node.children[index] = &TrieNode{}
		}
		node = node.children[index]
	}
	// Marks the end of the word
	node.isEnd = true
}

/*
* The `Search` function is a method of the `Trie` struct that takes a string `word` as input and searches for it in the Trie data structure.
* It does this by iterating over each character in the `word` and checking if the corresponding child node exists in the Trie.
* If the child node does not exist, it returns `false` indicating that the `word` does not exist in the Trie.
* If the child node exists, it moves to the next node and repeats the process until it reaches the end of the `word`.
* Finally, it checks if the last node is marked as the end of a `word` by checking the `isEnd` property.
* If `isEnd` is `true`, it returns `true` indicating that the `word` exists in the Trie, otherwise it returns `false`.
 */
func (trie *Trie) Search(word string) bool {
	node := trie.Root

	for _, char := range word {
		/*
		* If node.children[index] != nil, implies the Trie contains a character, in this case continue to the next char.
		* If node.children[index] == nil, it implies the Trie does not contain that character.
		* At last if the loop runs till the end i.e. node.children[index] != nil for all char's then check the isEnd property.
		* If isEnd == true, return true indicating that the word exists in the Trie else return false
		 */

		/*
		* `index := char - 'a'` is converting the character `char` to an index value between 0 and 25.
		* This is done by subtracting the ASCII value of the character 'a' (97) from the ASCII value of the character `char`.
		* For example, if `char` is 'c', then `index` will be 2 (99 - 97).
		* This index value is used to access the corresponding child node in the `children` array of the current node.
		 */
		index := char - 'a'

		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}

	return node.isEnd
}

/*
* The `Delete` function is a method of the `Trie` struct that takes a string `word` as input and deletes it from the Trie data structure.
* It does this by iterating over each character in the `word` and checking if the corresponding child node exists in the Trie.
* If the child node does not exist, it returns a message indicating that the `word` does not exist in the Trie.
* If the child node exists, it moves to the next node and repeats the process until it reaches the end of the `word`.
* Finally, it checks if the last node is marked as the end of a `word` by checking the `isEnd` property.
* If `isEnd` is `true`, it deletes the node by setting the corresponding child node in the parent node to `nil`.
* If `isEnd` is `false`, it returns a message indicating that the `word` does not exist in the Trie.
 */
func (trie *Trie) Delete(word string) {
	node := trie.Root
	parent := trie.Root
	parentIndex := 0

	for _, ch := range word {
		index := ch - 'a'
		if node.children[index] == nil {
			fmt.Println("Word doesn't exist in the trie")
			return
		}
		parent = node
		parentIndex = int(index)
		node = node.children[index]
	}

	if !node.isEnd {
		fmt.Println("Word doesn't exist in the trie")
		return
	}

	// Deletes the node
	parent.children[parentIndex] = nil

}

/*
* The `Print` function is a method of the `Trie` struct that takes a `child` node pointer and a `prefix` string as input and recursively prints all the words in the Trie data structure.
* It starts by setting the `node` variable to the root node of the Trie.
* If the `node` is marked as the end of a word (`isEnd` is `true`), it prints the `prefix` string.
* Then, it iterates over each child node of the `node` and recursively calls the `Print` function with the child node and the `prefix` string concatenated with the index of the child node as a string.
* This process continues until all the nodes in the Trie have been traversed and all the words have been printed.
 */
func (trie *Trie) Print(node *TrieNode, prefix string) {
	if node.isEnd {
		fmt.Println("\t", prefix)
	}

	for ch, child := range node.children {
		if child != nil {
			trie.Print(child, prefix+string(ch+'a'))
		}
	}
}
