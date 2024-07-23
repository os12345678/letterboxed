package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isWord bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord: false,
	}
}

type Trie struct {
	root *TrieNode
}

func newTrie() *Trie {
	return &Trie{
		root: newTrieNode(),
	}
}

func (t *Trie) insert(word string) {
    node := t.root
    for _, char := range word {
        if _, exists := node.children[char]; !exists {
            node.children[char] = newTrieNode()
        }
        node = node.children[char]
    }
    node.isWord = true
}

func (t *Trie) search(word string) bool {
    node := t.root
    for _, char := range word {
        if _, exists := node.children[char]; !exists {
            return false
        }
        node = node.children[char]
    }
    return node.isWord
}
