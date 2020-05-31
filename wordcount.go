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
		path := "./" + strings.TrimSpace(strings.Split(text, " ")[1])

		if command == "--lines" || command == "--symbols" || command == "--words" || command == "--unique_words" {
			res, _ := readFromFile(path) //получаем массив строк из файла

			if command == "--lines" {
				fmt.Println(countLines(res))
			}

			if command == "--symbols" {
				fmt.Println(countSymbols(res))
			}

			if command == "--words" {
				fmt.Println(countWords(res))
			}

			if command == "--unique_words" {
				var words []string
				for _, line := range res {
					for _, word := range strings.Split(line, " ") {
						words = append(words, word)
					}
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

func readFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	var res []string
	rd := bufio.NewScanner(file)
	for rd.Scan() {
		res = append(res, rd.Text())
	}
	return res, nil
}

func countLines(lines []string) int {
	return len(lines)
}

func countSymbols(lines []string) int {
	count := 0
	for _, line := range lines {
		count = count + len(line)
	}
	return count
}

func countWords(lines []string) int {
	count := 0
	for _, line := range lines {
		count = count + len(strings.Split(line, " "))
	}
	return count
}
