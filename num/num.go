package num

import (
	"bufio"
	"fmt"
	"os"
)

type colorType int

func printWeather(temp int, rain int) {
	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("It i/Users/semanticist/go/src/github.com/semanticist21/num.go hot and rainy")
		} else if rain >= 20 {
			fmt.Print("It is hot and wet")
		} else {
			fmt.Println("it is good to have a walk")
		}
	} else {
		if temp < 10 || rain >= 80 {
			fmt.Println("It is not good to go out")
		}
	}
}

const (
	Red colorType = iota
	Green
)

func colorToString(color colorType) string {
	switch color {
	case Red:
		return "Red"
	case Green:
		return "Green"
	default:
		return ""
	}
}

func getMyFavorColor() colorType {
	return Green
}

type Direction int

const (
	None Direction = iota
	North
	South
	West
	East
)

func getDirection(angle float64) Direction {
	switch {
	case angle >= 315, angle >= 0 && angle < 45:
		return North
	case 45 <= angle && angle < 135:
		return West
	case 135 <= angle && angle < 225:
		return South
	case 225 <= angle && angle < 315:
		return None
	default:
		return None
	}
}

func getInput() {
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("please direct put number here")
		var number int
		_, err := fmt.Scanln(&number)
		if err != nil {
			fmt.Println("")

			input, _ := stdin.ReadString('\n')
			fmt.Println("the error ", input, "was deleted")
			fmt.Println(err)
			continue
		}
		fmt.Printf("the number you input is %d. \n", number)
		if number%2 == 0 {
			break
		}
	}
}

func showMultipliers() {
	front := 2
	for {
		end := 1

		for {
			fmt.Printf("%d x %d, %d\n", front, end, front*end)

			if end == 9 {
				break
			}

			end++
		}

		if front == 9 {
			break
		}

		front++

	}
}

func findFortyFive() {
	front := 2
	limit := 10

OuterFor:
	for {
		end := 1

		for {
			if front*end == 45 {
				fmt.Println(front, end)
				break OuterFor
			}
			end++

			if end == limit {
				break
			}
		}

		if front == limit {
			break
		}
		front++
	}
}

func find45(num int) (int, bool) {
	for i := 0; i < 10; i++ {
		if num*i == 45 {
			return num * i, true
		}
	}
	return 0, false
}
func find(i int) (int, bool) {
	for i := 0; i < 10; i++ {
		result, success := find45(i)
		if success {
			return result, true
		}
	}
	return 0, false
}

func ExportTest() {

}
