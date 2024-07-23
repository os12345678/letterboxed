package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	trie "github.com/os12345678/letterboxed/pkg"
)

func parseLetterbox(letterbox string) []string {
	return strings.Split(letterbox, "-")
}

func loadDictionary(t *trie.Trie, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        t.Insert(strings.ToLower(word))
    }

    return scanner.Err()
}

func getInnerWords(node *trie.TrieNode, lastSide string, sides []string) []*trie.TrieNode {
    validNodes := []*trie.TrieNode{}
	if node.IsWord {
		validNodes = append(validNodes, node)
	}
    if node.Children != nil {
		for _, nextSide := range sides {
			if nextSide != lastSide {
				for _, nextLetter := range nextSide {
					if nextNode, exists := node.Children[rune(nextLetter)]; exists {
						validNodes = append(validNodes, getInnerWords(nextNode, nextSide, sides)...)
					}
				}
			}
		}
	}
	return validNodes
}

func getAllWords(t *trie.Trie, sides []string) []string {
    allValidNodes := []*trie.TrieNode{}

    for _, startingSide := range sides {
        for _, startingLetter := range startingSide {
            if node, exists := t.Root.Children[rune(startingLetter)]; exists {
                allValidNodes = append(allValidNodes, getInnerWords(node, startingSide, sides)...)

            }
        }
    }
    words := []string{}
	for _, node := range allValidNodes {
		words = append(words, node.GetWord())
	}
	return words    
}



func main() {
    t := trie.NewTrie()
    err := loadDictionary(t, "words.txt")
    if err != nil {
        fmt.Println("Error loading dictionary:", err)
        return
    }

    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)
    input = strings.ReplaceAll(input, " ", "")
    sides := parseLetterbox(input)
    // letters := strings.Join(sides, "")
    words := getAllWords(t, sides)
    fmt.Println(words)

    


    
}
