### LetterBoxed solver

This project is a Go implementation of a solver for the LetterBoxed puzzle from the New York Times. The solver constructs a trie from a dictionary of valid words and finds all possible solutions that match the puzzle's constraints.

#### Features
- Trie-based Word Storage: Efficiently store and retrieve words using a trie.
- Puzzle Graph Construction: Build a graph of valid words that fit the puzzle's letter constraints.
- Solution Finding: Find all possible solutions, with options to print hints using external APIs.

#### Installation
Clone the repository:

``` 
git clone https://github.com/os12345678/letterboxed.git
cd letterboxed
```

#### Install dependencies:
```
go get ./...
```

#### Usage
Command-Line Arguments
```
--puzzle (default: mrf-sna-opu-gci): Puzzle input in abc-def-ghi-jkl format.
--dict (default: words.txt): Path to a newline-delimited text file of valid words.
--len (default: 3): Maximum length, in words, of solutions.
--hint (default: false): Print a hint using external vocabulary APIs.
```

Running the solver
```
go run cmd/letterboxed/main.go --puzzle="mrf-sna-opu-gci" --dict="words.txt" --len=3 --hint
```

#### Environment Variables for Hints
To use the hint feature, set the following environment variables with your API keys:

```
BigHugeLabsApiKey
WordnikApiKey
```
Example
```
export BigHugeLabsApiKey="your-bighugelabs-api-key"
export WordnikApiKey="your-wordnik-api-key"
```
