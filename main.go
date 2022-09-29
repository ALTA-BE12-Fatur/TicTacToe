package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	array := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	var playing = true
	var player = 1
	for playing == true {

		papan(array, player)



		var inputX, inputY int
		fmt.Println()
		fmt.Println("INPUT KOORDINAT")
		fmt.Print("Koordinat X : ")
		fmt.Scanln(&inputX)
		fmt.Print("Koordinat Y : ")
		fmt.Scanln(&inputY)

		player++
		player %= 2

		// var res = inputX * inputY
		// fmt.Println(res)

		var change = true
		fmt.Println(change)

		for y := 1; y <= 3; y++ {
			for x := 1; x <= 3; x++ {
				if (y == 1 && x == 1) && (inputX == 1 && inputY == 1) {
					if array[0][0] == 1 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[0][0] = 1
					}
				}
				if (y == 1 && x == 2) && (inputX == 2 && inputY == 1) {
					if array[0][1] == 2 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[0][1] = 2
					}
				}
				if (y == 1 && x == 3) && (inputX == 3 && inputY == 1) {
					if array[0][2] == 3 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[0][2] = 3
					}
				}
				if (y == 2 && x == 1) && (inputX == 1 && inputY == 2) {
					if array[1][0] == 4 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[1][0] = 4
					}
				}
				if (y == 2 && x == 2) && (inputX == 2 && inputY == 2) {
					if array[1][1] == 5 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[1][1] = 5
					}
				}
				if (y == 2 && x == 3) && (inputX == 3 && inputY == 2) {
					if array[1][2] == 6 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[1][2] = 6
					}
				}
				if (y == 3 && x == 1) && (inputX == 1 && inputY == 3) {
					if array[2][0] == 7 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[2][0] = 7
					}
				}
				if (y == 3 && x == 2) && (inputX == 2 && inputY == 3) {
					if array[2][1] == 8 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[2][1] = 8
					}
				}
				if (y == 3 && x == 3) && (inputX == 3 && inputY == 3) {
					if array[2][2] == 9 {
						fmt.Println("++++ Kolom sudah terisi ++++")
						change = false
					} else {
						array[2][2] = 9
					}
				}
			}
		}
		time.Sleep(2 * time.Second)
		clear()
		// playing = false
	}
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func papan(array [][]int, player int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				if array[i][2] == 0 {
					fmt.Println("_")
				} else if player == 1 {
					fmt.Println("X")
				} else if player == 0 {
					fmt.Println("O")
				}
			} else if i == 1 && j == 2 {
				if array[i][2] == 0 {
					fmt.Println("_")
				} else if player == 1 {
					fmt.Println("X")
				} else if player == 0 {
					fmt.Println("O")
				}
			} else if i == 2 && j == 2 {
				if array[i][2] == 0 {
					fmt.Println("_")
				} else if player == 1 {
					fmt.Println("X")
				} else if player == 0 {
					fmt.Println("O")
				}
			} else {
				if array[i][j] == 0 {
					fmt.Print("_")
					fmt.Print(" | ")
				} else if player == 1 {
					fmt.Print("X")
					fmt.Print(" | ")
				} else if player == 0 {
					fmt.Print("O")
					fmt.Print(" | ")
				}
			}
		}
	}
}

func checkForWin(array [][]int) int {
	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}