package main

import (
"fmt"
"regexp"
"strconv"
"strings"
"bufio"
"os"
)

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func get_operation(input string) (oper string) {
	return input[0:3]
}

func jio(input string, input_num int) (int) {
	if input_num == 1 {
		re := regexp.MustCompile("-?[0-9]+")
   		match := re.FindAllString(input, -1)
   		number, err := strconv.Atoi(match[0])
   		check(err)
   		return number
   	} else {
   		return 1
   	}
}

func main() {
	fmt.Printf("Solution for day 13 GO !\n")

	//var a = 0
	var a = 1
	var b = 0
	var pointer = 0
	prog, err := readLines("input.txt")
	check(err)

	var oper string
	for {
		oper = get_operation(string(prog[pointer]))
		//fmt.Printf("Next operation is : %s\n", string(prog[pointer]))

		switch oper {
			case "hlf":
				if strings.Contains(string(prog[pointer]), "a") {
					a = a/2
				} else {
					b = a/2
				}
				pointer++
			case "tpl":
				if strings.Contains(string(prog[pointer]), "a") {
					a = a * 3
				} else {
					b = b * 3
				}
				pointer++
			case "inc":
				if strings.Contains(string(prog[pointer]), "a") {
					a++
				} else {
					b++
				}
				pointer++
			case "jmp":
				offset := jio(string(prog[pointer]), 1)
				pointer = pointer + offset
			case "jie":
				offset := 0
				if strings.Contains(string(prog[pointer]), "a") {
					if (a % 2) > 0 {
						offset = jio(string(prog[pointer]), 0)
					} else {
						offset = jio(string(prog[pointer]), 1)
					}
				} else {
					if (b % 2) > 0 {
						offset = jio(string(prog[pointer]), 0)
					} else {
						offset = jio(string(prog[pointer]), 1)
					}
				}
				pointer = pointer + offset
			case "jio":
				offset := 0
				if strings.Contains(string(prog[pointer]), "a") {
					offset = jio(string(prog[pointer]), a)
				} else {
					offset = jio(string(prog[pointer]), b)
				}
				pointer = pointer + offset
		}

		if pointer < 0 || pointer > 46 {
			fmt.Printf("Final = a = %d, b = %d\n", a, b)
			os.Exit(0)
		}

		//fmt.Printf("a = %d, b = %d\n", a, b)
	}


}