package main

import (
	"fmt"
)

type Player struct {
	hitpoint int
	armor int
	damage int
}

type Item struct {
	name string
	price int
	armor int
	damage int
}
	var empty = Item{name: "empty", price: 0, armor: 0, damage: 0}

	var Dagger = Item{name: "Dagger", price: 8, armor: 0, damage: 4}
	var Shortsword = Item{name: "Shortsword", price: 10, armor: 0, damage: 5}
	var Warhammer = Item{name: "Warhammer", price: 25, armor: 0, damage: 6}
	var Longsword = Item{name: "Longsword", price: 40, armor: 0, damage: 7}
	var Greataxe = Item{name: "Greataxe", price: 74, armor: 0, damage: 8}
	var weapon = []Item{empty, Dagger, Shortsword, Warhammer, Longsword, Greataxe}

	var Leather = Item{name: "Leather", price: 13, armor: 1, damage: 0}
	var Chainmail = Item{name: "Chainmail", price: 31, armor: 2, damage: 0}
	var Splintmail = Item{name: "Splintmail", price: 53, armor: 3, damage: 0}
	var Bandedmail = Item{name: "Bandedmail", price: 75, armor: 4, damage: 0}
	var Platemail = Item{name: "Platemail", price: 102, armor: 5, damage: 0}
	var armor = []Item{empty, Leather, Chainmail, Splintmail, Bandedmail, Platemail}

	var dam1 = Item{name: "dam1", price: 25, armor: 0, damage: 1}
	var dam2 = Item{name: "dam2", price: 50, armor: 0, damage: 2}
	var dam3 = Item{name: "dam3", price: 100, armor: 0, damage: 3}
	var def1 = Item{name: "def1", price: 20, armor: 1, damage: 0}
	var def2 = Item{name: "def2", price: 40, armor: 2, damage: 0}
	var def3 = Item{name: "def3", price: 80, armor: 3, damage: 0}
	var ring = []Item{empty, dam1, dam2, dam3, def1, def2, def3}

	var gold = 0
	var w = 0
	var a = 0
	var r1 = 0
	var r2 = 0

func Battle(boss Player, hero Player) (int) {
	boss_hitpoint := boss.hitpoint
	hero_hitpoint := hero.hitpoint
	damage := 0

	for {
		// player strike
		damage = hero.damage - boss.armor
		if damage <= 0 { damage = 1 }
		boss_hitpoint -= damage
		//tcheck hitpoint
		fmt.Printf("Hero hit with %d damage, boss hitpoint = %d\n", damage, boss_hitpoint)
		if boss_hitpoint <= 0 { return 1 }
		//boss strike
		damage = boss.damage - hero.armor
		if damage <= 0 { damage = 1 }
		hero_hitpoint -= damage
		//tcheck hitpoint
		fmt.Printf("Boss hit with %d damage, hero hitpoint = %d\n", damage, hero_hitpoint)
		if hero_hitpoint <= 0 { return 0 }
	}
}

func shop(hero *Player) (int) {
	hero.damage = weapon[w].damage
	fmt.Printf("Hero choose %s as a weapon \n", weapon[w].name)
	gold = weapon[w].price

	hero.armor = armor[a].armor
	fmt.Printf("Hero choose %s as an armor \n", armor[a].name)
	gold += armor[a].price

	hero.damage += ring[r1].damage
	hero.armor += ring[r1].armor
	fmt.Printf("Hero choose %s as ring \n", ring[r1].name)
	gold += ring[r1].price

	hero.damage += ring[r2].damage
	hero.armor += ring[r2].armor
	fmt.Printf("Hero choose %s as ring \n", ring[r2].name)
	gold += ring[r2].price

	if w == 5 && a == 5 && r1 == 6 && r2 == 6 {
		fmt.Printf("We use all possibility\n")
		return -1
	}

	// increment our "number"
	w++
	if w == 6 {
		w = 1
		a++
		if a == 6 {
			a = 0
			r1++
			if r1 == 7 {
				r1 = 0
				r2 ++
			}
		}
	}

	return gold
}

func printmin (win []int) {
	m := 500000
	for _, e := range win {
    	if e < m {
        	m = e
    	}
	}
	fmt.Printf("The minimal cost is: %d\n",m)
}

func printmax (fail []int) {
	m := 0
	fmt.Println(fail)
	for _, e := range fail {
    	if e > m {
        	m = e
    	}
	}
	fmt.Printf("The maximal cost is: %d\n",m)
}

func main() {

	boss := Player{hitpoint: 108, armor: 2, damage: 8}
	hero := Player{hitpoint: 100, armor: 0, damage: 0}
	var win []int
	var fail []int

	for {
		winner := Battle(boss, hero)

		if winner == 1 {
			fmt.Printf("Hero win !!! with %d gold spend \n", gold)
			if r1 == r2 {
				fmt.Println("We cheat on this one :(")
			} else {win = append(win, gold)}
		} else {
			fmt.Printf("Boss win with %d gold\n", gold)
			if r1 == r2 {
				fmt.Println("We cheat on this one :(")
			} else {fail = append(fail, gold)}
		}

		gold = shop(&hero)
		if gold == -1 {
			printmin(win)
			printmax(fail)
			return
		}
	}

}