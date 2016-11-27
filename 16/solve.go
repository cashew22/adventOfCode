package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
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

func remove_stuff(sue []string, stuff string, stuff_num int) ([]string) {
	var keep []string
	for _, e := range sue {
		if strings.Contains(e, stuff) {
			// Split on comma.
	    	result := strings.Split(e, ",")

	    	for _, i := range result {
	    		if strings.Contains(i, stuff) {
	    			// Use : to find number
	    			index := strings.LastIndex(i, ":") + 1
	    			s_num, err := strconv.Atoi(i[index + 1:])
	    			check(err)
	    			//Look if sue is not sue
	    			switch stuff {
	    			case "cats":
	    				if s_num > stuff_num {
	    				keep = append(keep, e)
	    			}
	    			case "trees":
	    				if s_num > stuff_num {
	    				keep = append(keep, e)
	    			}
	    			case "pomeranians":
	    				if s_num < stuff_num {
	    				keep = append(keep, e)
	    			}
	    			case "goldfish":
	    				if s_num < stuff_num {
	    				keep = append(keep, e)
	    			}
	    			default:
	    				if s_num == stuff_num {
	    				keep = append(keep, e)
	    			}
	    			}
	    		}
	    	}
	    } else {
	    	keep = append(keep,e)
	    }
	}

	return keep
}

//Sue Stat
const children = 3
const cats = 7
const samoyeds = 2
const pomeranians = 3
const akitas = 0
const vizslas = 0
const goldfish = 5
const trees = 3
const cars = 2
const perfumes = 1

func main() {
	fmt.Printf("Solution for day 16 GO !\n")

	//Read input
	sue, err := readLines("input.txt")
	check(err)
	fmt.Printf("There are %d sue in the list\n", len(sue))

	//children
	sue = remove_stuff(sue, "children", children)
	fmt.Printf("Get rid of the kid and then there are %d sue in the list\n", len(sue))

	//cats
	sue = remove_stuff(sue, "cats", cats)
	fmt.Printf("Get rid of the cats and then there are %d sue in the list\n", len(sue))

	//samoyeds
	sue = remove_stuff(sue, "samoyeds", samoyeds)
	fmt.Printf("Get rid of the samoyeds and then there are %d sue in the list\n", len(sue))

	//pomeranians
	sue = remove_stuff(sue, "pomeranians", pomeranians)
	fmt.Printf("Get rid of the pomeranians and then there are %d sue in the list\n", len(sue))

	//akitas
	sue = remove_stuff(sue, "akitas", akitas)
	fmt.Printf("Get rid of the akitas and then there are %d sue in the list\n", len(sue))

	//vizslas
	sue = remove_stuff(sue, "vizslas", vizslas)
	fmt.Printf("Get rid of the vizslas and then there are %d sue in the list\n", len(sue))

	//goldfish
	sue = remove_stuff(sue, "goldfish", goldfish)
	fmt.Printf("Get rid of the goldfish and then there are %d sue in the list\n", len(sue))

	//trees
	sue = remove_stuff(sue, "trees", trees)
	fmt.Printf("Get rid of the trees and then there are %d sue in the list\n", len(sue))

	//cars
	sue = remove_stuff(sue, "cars", cars)
	fmt.Printf("Get rid of the cars and then there are %d sue in the list\n", len(sue))

	//perfumes
	sue = remove_stuff(sue, "perfumes", perfumes)
	fmt.Printf("Get rid of the perfumes and then there are %d sue in the list\n", len(sue))

	fmt.Println(sue)
}