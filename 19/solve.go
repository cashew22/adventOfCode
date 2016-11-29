package main

import (
    "fmt"
    "os"
    "bufio"
    "time"
    "strings"
)

type operation struct {
    base string
    result string
}

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

func get_ops(input []string) ([]operation) {
    var ops []operation
    var opers operation
    for _, e := range input {
        splited := strings.Split(e, " ")
        opers = operation{base: splited[0], result: splited[2]}
        ops = append(ops, opers)
    }
    return ops
}

func replace_one_after_n(s string, old string, new_s string, n int) (string) {
	s_slice := s[:]
	index_target := 0
	//fmt.Println(s, n)
	for i:=0; i < n; i++ {
		index := strings.Index(string(s_slice), old)
		s_slice = s_slice[index+len(old):]
		index_target += index + len(old)
	}
	//fmt.Println(s[:index_target-1] + new_s + s[index_target:])
	return s[:index_target - len(old)] + new_s + s[index_target:]
}

func get_possibilities(ops operation, base_string string) ([]string) {
	var possibilities []string
	num_of_base := strings.Count(base_string, ops.base)
	for i:=0; i < num_of_base; i++ {
		possibilities = append(possibilities, replace_one_after_n(base_string, ops.base, ops.result, i+1))
	}

	fmt.Println(ops, num_of_base, len(possibilities))
	return possibilities
}

func add_new_molecules(molecule map[string]string, candidates []string) (map[string]string) {
	for _, candidate := range candidates {
		if molecule[candidate] == "" {
			molecule[candidate] = candidate
		}
	}

	return molecule
}

func main() {
    start := time.Now() //timestamp
    fmt.Printf("Solution for day 19 GO !\n")

    var oper []operation
	var molecule map[string]string
	molecule = make(map[string]string)
	var base_string = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"
	//var base_string = "SiCaCaCaCaSiCa"

    //Read file
    data, err := readLines("input.txt")
    check(err)

    oper = get_ops(data)

    for _,ops := range oper {
    	temp := get_possibilities(ops, base_string)
    	molecule = add_new_molecules(molecule, temp)
    }
    //fmt.Println(molecule)
    fmt.Printf("Q1 answer is : %d\n", len(molecule))

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
