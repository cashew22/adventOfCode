package main

import (
	"fmt"
    "os"
    "bufio"
    "time"
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

func create_base_light(lines []string) {
    for h, line := range lines {
        for v, light := range line {
            if string(light) == "#" {
                base_light_init[h][v] = 1
            } else {
                base_light_init[h][v] = 0
            }
        }
    }
    base_light = base_light_init
}

func out_of_bound(v int, h int) (bool) {
	if v < 0 || v > 99 || h < 0 || h > 99 {
		return true
	}
	return false
}

func neighbors_on(v int, h int) (int) {
    on := 0
    for i:=v - 1; i<v+2; i++ {
    	for k:=h - 1; k<h+2; k++ {
    		if !out_of_bound(i,k) {
    			on += base_light[i][k]
    		}
    	}
    }
    on -= base_light[v][h]

    return on
}

func is_corner(x int, y int) (bool) {
	if x == 0 && y == 0 {return true}
	if x == 99 && y == 0 { return true}
	if x == 0 && y == 99 {return true}
	if x == 99 && y == 99 {return true}
	return false
}

func step(day2 bool) ([100][100]int) {
    var new_light [100][100]int

    for y, light_row := range base_light {
        for x,_ := range light_row {
            on := neighbors_on(x,y)
            if is_corner(x,y) && day2 {
            	new_light[x][y] = 1
            } else {
	            if base_light[x][y] == 1 {
	            	if on == 2 || on == 3 {
	            		new_light[x][y] = 1
	            	} else {
	            		new_light[x][y] = 0
	            	}
	            } else {
	            	if on == 3 {
	            		new_light[x][y] = 1
	            	} else {
	            		new_light[x][y] = 0
	            	}
	            }
	        }
        }
    }

    return new_light
}

func num_light_on() (int) {
	on := 0
	for _, light_row := range base_light {
        for _, light := range light_row {
        	on += light
        }
    }
    return on
}


var base_light_init [100][100]int
var base_light [100][100]int

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 18 GO !\n")


    //Read file
    data, err := readLines("input.txt")
    check(err)

    //create base light array
    create_base_light(data)

    for i:=0; i < 100 ; i++ {
        base_light = step(false)
    }

    fmt.Printf("Q1 answer = %d\n",num_light_on())

    //reset 
    base_light = base_light_init
    for i:=0; i < 100 ; i++ {
        base_light = step(true)
    }

    fmt.Printf("Q2 answer = %d\n",num_light_on())

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
