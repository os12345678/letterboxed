package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/karan/vocabulary"
	trie "github.com/os12345678/letterboxed/pkg"
)

/*
LetterBox
  ├── inputString: "mrf-sna-opu-gci"
  ├── sides: {"mrf", "sna", "opu", "gci"}
  ├── letters: {'m', 'r', 'f', 's', 'n', 'a', 'o', 'p', 'u', 'g', 'c', 'i'}
  ├── lenThreshold: 3
  ├── trie: Trie
  ├── words: ["some", "valid", "words"]
  └── graph: map[rune]map[rune]map[string][]string

*/
type LetterBox struct {
    inputString string
    sides map[string]struct{}
    letters map[rune]struct{}
    lenThreshold int
    trie *trie.Trie
    words []string
    graph map[rune]map[rune]map[string][]string
}

/*
NewLetterBox
  ├── Initialize LetterBox
  ├── Parse sides and letters
  ├── Load dictionary into trie
  ├── Get all valid words
  └── Build puzzle graph
*/
func NewLetterBoxed(inputString, dictionary string, lenThreshold int) *LetterBox {
    start := time.Now()
	defer func() {
		fmt.Printf("LetterBoxed initialization took %v\n", time.Since(start))
	}()

    lb := &LetterBox{
		inputString:  strings.ToLower(inputString),
		sides:        make(map[string]struct{}),
		letters:      make(map[rune]struct{}),
		lenThreshold: lenThreshold,
		trie:         trie.NewTrie(),
		graph:        make(map[rune]map[rune]map[string][]string),
	}

    for _, side := range strings.Split(lb.inputString, "-") {
		lb.sides[side] = struct{}{}
		for _, letter := range side {
			lb.letters[letter] = struct{}{}
		}
	}

    file, err := os.Open(dictionary)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        lb.trie.Insert(strings.ToLower(word))
    }
    
    if err := scanner.Err(); err != nil {
		panic(err)
	}

    lb.words = lb.getAllWords()
    lb.buildPuzzleGraph()

    return lb
}

/*
getAllWords
  ├── Iterate over sides and letters
  └── Collect words from trie
*/
func (lb *LetterBox) getAllWords() []string {
    start := time.Now()
	defer func() {
		fmt.Printf("getPuzzleWords took %v\n", time.Since(start))
	}()

	var allValidNodes []*trie.TrieNode
	for side := range lb.sides {
		for _, letter := range side {
			if lb.trie.Root.Children[letter] != nil {
				allValidNodes = append(allValidNodes, lb.getInnerWords(lb.trie.Root.Children[letter], side)...)
			}
		}
	}

	var words []string
	for _, node := range allValidNodes {
		words = append(words, node.GetWord())
	}
	return words
}

/*
collectWords
  ├── Check if node is terminal
  ├── Collect current word
  └── Recursively collect words from children
*/
func (lb *LetterBox) getInnerWords(node *trie.TrieNode, lastSide string) []*trie.TrieNode {
	var validNodes []*trie.TrieNode
	if node.IsWord {
		validNodes = append(validNodes, node)
	}
	if node.Children != nil {
		for nextSide := range lb.sides {
			if nextSide != lastSide {
				for _, nextLetter := range nextSide {
					if node.Children[nextLetter] != nil {
						validNodes = append(validNodes, lb.getInnerWords(node.Children[nextLetter], nextSide)...)
					}
				}
			}
		}
	}
	return validNodes
}

/*
buildPuzzleGraph
  ├── Iterate over words
  ├── Create letter sets and keys
  └── Populate graph with words
*/
func (lb *LetterBox) buildPuzzleGraph() {
	for _, word := range lb.words {
		start := rune(word[0])
		end := rune(word[len(word)-1])
		letterSet := make(map[rune]struct{})
		for _, letter := range word {
			letterSet[letter] = struct{}{}
		}
		letterKey := makeKey(letterSet)

		if lb.graph[start] == nil {
			lb.graph[start] = make(map[rune]map[string][]string)
		}
		if lb.graph[start][end] == nil {
			lb.graph[start][end] = make(map[string][]string)
		}
		lb.graph[start][end][letterKey] = append(lb.graph[start][end][letterKey], word)
	}
}

/*
makeKey
  └── Concatenate letters into a string
*/
func makeKey(letterSet map[rune]struct{}) string {
	var key strings.Builder
	for letter := range letterSet {
		key.WriteRune(letter)
	}
	return key.String()
}


/*
findAllSolutions
  ├── Iterate over letters
  └── Find solutions for each combination
*/
func (lb *LetterBox) findAllSolutions() [][]string {
	start := time.Now()
	defer func() {
		fmt.Printf("findAllSolutions took %v\n", time.Since(start))
	}()

	var allSolutions [][]string
	for firstLetter := range lb.letters {
		for lastLetter := range lb.letters {
			for letterEdge, edgeWords := range lb.graph[firstLetter][lastLetter] {
				letterSet := make(map[rune]struct{})
				for _, letter := range letterEdge {
					letterSet[letter] = struct{}{}
				}
				allSolutions = append(allSolutions, lb.findSolutionsInner([][]string{edgeWords}, letterSet, lastLetter)...)
			}
		}
	}
	return allSolutions
}

/*
findSolutionsInner
  ├── Check for complete solution
  ├── Check length threshold
  ├── Iterate over graph edges
  └── Recursively find solutions
*/
func (lb *LetterBox) findSolutionsInner(pathWords [][]string, letters map[rune]struct{}, nextLetter rune) [][]string {
	if len(letters) == 12 {
		var solution []string
		for _, words := range pathWords {
			solution = append(solution, words...)
		}
		return [][]string{solution}
	} else if len(pathWords) == lb.lenThreshold {
		return nil
	}

	var solutions [][]string
	for lastLetter := range lb.graph[nextLetter] {
		for letterEdge, edgeWords := range lb.graph[nextLetter][lastLetter] {
			newLetters := make(map[rune]struct{})
			for letter := range letters {
				newLetters[letter] = struct{}{}
			}
			for _, letter := range letterEdge {
				newLetters[letter] = struct{}{}
			}
			if len(newLetters) > len(letters) {
				solutions = append(solutions, lb.findSolutionsInner(append(pathWords, edgeWords), newLetters, lastLetter)...)
			}
		}
	}
	return solutions
}


func main() {
	puzzle := flag.String("puzzle", "mrf-sna-opu-gci", "puzzle input in abd-def-ghi-jkl format")
	dict := flag.String("dict", "words.txt", "path to newline-delimited text file of valid words")
	lenThreshold := flag.Int("len", 3, "maximum length, in words, of solutions")
    hint := flag.Bool("hint", false, "print hint")
	flag.Parse()

	fmt.Println("solving puzzle", *puzzle)
	p := NewLetterBoxed(*puzzle, *dict, *lenThreshold)
	fmt.Println(len(p.words), "valid words found")
	metaSolutions := p.findAllSolutions()
    if *hint {
        c := &vocabulary.Config{BigHugeLabsApiKey: os.Getenv("BigHugeLabsApiKey"), WordnikApiKey: os.Getenv("WordnikApiKey")}
        v, err := vocabulary.New(c)
        if err != nil {
            panic(err)
        }
        word, err := v.Word(metaSolutions[0][0])
        if err != nil {
            panic(err)
        }
        fmt.Printf("word.Word = %s \n", word.Word)
        fmt.Printf("word.Meanings = %s \n", word.Meanings)
        fmt.Printf("word.Synonyms = %s \n", word.Synonyms)
        fmt.Printf("word.UsageExample = %s \n", word.UsageExample)
	} else {
        for _, solution := range metaSolutions {
            fmt.Println(solution)
        }
    }
}

