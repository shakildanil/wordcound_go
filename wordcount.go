package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Tree struct {
	words string
	left  *Tree
	right *Tree
}

func (node *Tree) Insert(word string) error {
	if node == nil {
		return errors.New("Tree is null")
	}

	if node.words == word {
		return errors.New("Already in a tree")

	}

	if node.words > word {
		if node.left == nil {
			node.left = &Tree{words: word}
			return nil
		}
		return node.left.Insert(word)
	}

	if node.words < word {
		if node.right == nil {
			node.right = &Tree{words: word}
			return nil
		}
		return node.right.Insert(word)
	}
	return nil
}

func (node *Tree) PrintWords() {
	if node == nil {
		return
	}

	node.left.PrintWords()
	fmt.Println(node.words)
	node.right.PrintWords()
}

func main() {
	//ФОРМАТ ВВОДА: --<command> <filepath>
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	if len(strings.Split(text, " ")) > 2 {
		fmt.Println("Too many input arguments")
	} else if len(strings.Split(text, " ")) < 2 {
		fmt.Println("Not enough input arguments")
	} else if len(strings.Split(text, " ")) == 2 {
		command := strings.TrimSpace(strings.Split(text, " ")[0])
		path := strings.TrimSpace(strings.Split(text, " ")[1])
		if command == "--lines" || command == "--symbols" || command == "--words" || command == "--unique_words" {
			if command == "--lines" {
				count, err := countLines("./" + path)
				if err != nil {
					panic(err)
				}
				fmt.Println(count)
			}

			if command == "--symbols" {
				count, err := countSymbols("./" + path)
				if err != nil {
					panic(err)
				}
				fmt.Println(count)
			}

			if command == "--words" {
				words, err := scanWords("./" + path)
				if err != nil {
					panic(err)
				}

				fmt.Println(len(words))
			}

			if command == "--unique_words" {
				words, err := scanWords("./" + path)
				if err != nil {
					panic(err)
				}

				var a *Tree
				length := 0

				for _, word := range words {
					if a == nil {
						a = &Tree{words: word}
						length++
					} else {
						{
							err := a.Insert(word)
							if err == nil {
								length++
							}
						}
					}
				}
				fmt.Println(length)
			}

		} else {
			fmt.Println("Wrong command")
		}
	}
}

func countLines(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count++
	}
	return count, nil
}

func countSymbols(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}

	defer file.Close()
	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		count = count + len(scanner.Text())
	}
	return count, nil
}

func scanWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanWords)

	var words []string

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	return words, nil
}
