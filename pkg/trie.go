package trie

type TrieNode struct {
	Value    rune
	Parent   *TrieNode
	Children map[rune]*TrieNode
	IsWord   bool
}

func NewTrieNode(value rune, parent *TrieNode) *TrieNode {
	return &TrieNode{
		Value:    value,
		Parent:   parent,
		Children: make(map[rune]*TrieNode),
		IsWord:   false,
	}
}

type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(0, nil),
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewTrieNode(char, node)
		}
		node = node.Children[char]
	}
	node.IsWord = true
}
func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			return false
		}
		node = node.Children[char]
	}
	return node.IsWord
}

func (n *TrieNode) GetWord() string {
	if n.Parent != nil {
		return n.Parent.GetWord() + string(n.Value)
	}
	return ""
}