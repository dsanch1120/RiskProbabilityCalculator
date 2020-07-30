/*
	Created by Daniel Sanchez
	July 29th, 2020
	Runs simulation of risk battle given user-entered information
 */
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

//Global Variables
var clear map[string]func()

//Allows the program to clear the terminal in Windows and Unix operating systems
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

//Clears the Screen
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func getStrength(option int) int {
	for {
		if option == 1 {
			fmt.Println("Enter attack strength")
		} else {
			fmt.Println("Enter defense strength")
		}
		reader := bufio.NewReader(os.Stdin)
		//Gets user input
		temp, _ := reader.ReadString('\n')
		temp = strings.Replace(temp, "\n", "", -1)
		strength, err := strconv.Atoi(temp)

		if err == nil && (0 < strength) {
			CallClear()
			return strength
		} else {
			CallClear()
			fmt.Println("Invalid input")
			fmt.Println()
		}
	}
}

func battle(attackPower int, defensePower int) {
	//Variables to be output
	attackWins := 0
	defenseWins := 0
	//Generates random seed
	rand.Seed(time.Now().UnixNano())

	//Iterates until attacker or defender runs out of troops
	for i := 0; i < 1000; i++ {
		attackStrength := attackPower
		defenseStrength := defensePower
		for {
			var attackDice int
			var defenseDice int
			var attackRoll []int
			var defenseRoll []int

			//Determines number of dice each side gets
			if attackStrength >= 3 {
				attackDice = 3
			} else {
				attackDice = attackStrength
			}
			if defenseStrength >= 2 {
				defenseDice = 2
			} else {
				defenseDice = defenseStrength
			}

			//Rolls dice for attacker and defender
			for i := 0; i < attackDice; i++ {
				a := rand.Intn(6) + 1
				attackRoll = append(attackRoll, a)
			}
			for i := 0; i < defenseDice; i++ {
				d := rand.Intn(6) + 1
				defenseRoll = append(defenseRoll, d)
			}

			//Sorts the slices in descending order
			sort.Sort(sort.Reverse(sort.IntSlice(attackRoll)))
			sort.Sort(sort.Reverse(sort.IntSlice(defenseRoll)))

			if (len(attackRoll) >= 1) && (len(defenseRoll) >= 1) {
				if attackRoll[0] > defenseRoll[0] {
					defenseStrength--
				} else {
					attackStrength--
				}
			}
			if defenseStrength <= 0 {
				attackWins++
				break
			}
			if attackStrength <= 0 {
				defenseWins++
				break
			}
			if (len(attackRoll) >= 2) && (len(defenseRoll) >= 2) {
				if attackRoll[1] > defenseRoll[1] {
					defenseStrength--
				} else {
					attackStrength--
				}
			}
			if defenseStrength <= 0 {
				attackWins++
				break
			}
			if attackStrength <= 0 {
				defenseWins++
				break
			}
		}
	}
	fmt.Println("Attacker Wins " + strconv.Itoa(attackWins) + "/1000 times")
	fmt.Println("Defender Wins " + strconv.Itoa(defenseWins) + "/1000 times")
}

func menu() {
	//Displays welcome message to user
	CallClear()
	fmt.Println("Welcome to Risk Battle Probability Calculator")
	fmt.Println()

	//Gets the attack strength (number of dice)
	attackStrength := getStrength(1)
	defenseStrength := getStrength(2)


	//Guesses outcome of a battle based upon attack and defenseStrength
	battle(attackStrength, defenseStrength)

}

func main() {
	menu()
}