package main

import (
	"bufio"
	"fmt"
	"os"

	"trie"
)

func loadDictionary(t *trie.Trie, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        t.Insert(word)
    }

    return scanner.Err()
}

func main() {
    t := trie.NewTrie()
    err := loadDictionary(t, "words.txt")
    if err != nil {
        fmt.Println("Error loading dictionary:", err)
        return
    }

    fmt.Println("Dictionary loaded successfully")
}
