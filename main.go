package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	Array := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	var playing = false
	var playerCheck = 1


	for playing != true {
		fmt.Println()
		papan(Array)

		player := playerCheck%2+1
		if player == 1 {
			fmt.Println()
			fmt.Println("++ Player 1 Jalan ++")
		} else {
			fmt.Println()
			fmt.Println("++ Player 2 Jalan ++")
		}

		move := move()

		play := jalan(move, player, Array)


		result := winCheck(play)
		if result > 0 {
			fmt.Printf("Player %d wins!\n\n", result)
			playing = true
		} else if playerCheck == 9 {
			// Tie game example: 0 2 1 3 4 7 5 8 6
			fmt.Printf("!!++++ Pertandingan Seri! ++++!!\n\n")
			playing = true
		} else {
			playerCheck++
		}


	}
}

func move () int{
	var inputX, inputY int
		fmt.Println()
		fmt.Println("~ INPUT KOORDINAT ~")
		fmt.Println()
		fmt.Print("Koordinat X : ")
		fmt.Scanln(&inputX)
		fmt.Print("Koordinat Y : ")
		fmt.Scanln(&inputY)
		arr := koordinat(inputX, inputY)
		return arr
}

func jalan(point int, player int, b []int) []int {
	if b[point] != 0 {
		fmt.Println("Kolom Sudah Diisi")
		point = move()
		b = jalan(point, player, b)
	} else {
		if player == 1 {
			b[point] = 1
		} else if player == 2 {
			b[point] = 10
		}
	}

	if player == 1 {
		b[point] = 1
	} else if player == 2 {
		b[point] = 10
	}
	return b
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func papan(array []int) {
	for i, val := range array{
		if val == 0 {
			fmt.Print("_")
		} else if val == 1{
			fmt.Printf("X")
		} else if val == 10 {
			fmt.Printf("O")
		}

		if i > 0 && (i+1)%3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf(" | ")
		}
	}
}

func koordinat(inputX, inputY int) int{
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {
			if (y == 1 && x == 1) && (inputX == 1 && inputY == 1) {
					return 0 
			}
			if (y == 1 && x == 2) && (inputX == 2 && inputY == 1) {
					return 1 
			}
			if (y == 1 && x == 3) && (inputX == 3 && inputY == 1) {
					return 2 
			}
			if (y == 2 && x == 1) && (inputX == 1 && inputY == 2) {
					return 3 
			}
			if (y == 2 && x == 2) && (inputX == 2 && inputY == 2) {
					return 4 
			}
			if (y == 2 && x == 3) && (inputX == 3 && inputY == 2) {
					return 5 
			}
			if (y == 3 && x == 1) && (inputX == 1 && inputY == 3) {
					return 6 
			}
			if (y == 3 && x == 2) && (inputX == 2 && inputY == 3) {
					return 7 
			}
			if (y == 3 && x == 3) && (inputX == 3 && inputY == 3) {
					return 8 
			}
		}
	}
	time.Sleep(2 * time.Second)
	clear()
	return 0
}

func winCheck(b []int) int {
	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	sums[0] = b[2] + b[4] + b[6]
	sums[1] = b[0] + b[3] + b[6]
	sums[2] = b[1] + b[4] + b[7]
	sums[3] = b[2] + b[5] + b[8]
	sums[4] = b[0] + b[4] + b[8]
	sums[5] = b[6] + b[7] + b[8]
	sums[6] = b[3] + b[4] + b[5]
	sums[7] = b[0] + b[1] + b[2]
	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30 {
			return 2
		}
	}
	return 0
}