package input

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

func FindAnswerGame() {
	var n int
	cnt := 1

	rand.Seed(time.Now().UnixNano())
	answer := rand.Intn(10)

	fmt.Println("Put the answer here")
	for {
		_, err := fmt.Scanln(&n)

		if err != nil {
			fmt.Println("please put correct number here", err)
			reader.ReadString('\n')
		} else {
			if n == answer {
				fmt.Println("You got an correct answer ", answer)
				fmt.Println("Count Tried", cnt)
				break
			} else {
				fmt.Println("retry", cnt)
			}
		}
		cnt++
	}
}

func Gamble() {
	limit := 5
	money := 1000
	answer := getRandomNum(limit)

	for {
		val, err := getInput()

		if err != nil {
			fmt.Println("Please number only here")
		} else {
			if val > limit {
				fmt.Println("it excceded")
			} else {
				if val == answer {
					money += 5000
					fmt.Println("Congratulation!!! Your current money ", money)
				} else {
					money -= 100
					fmt.Println("You got wrong!!! You current money", money)
				}
			}

			if hasReachedCondition(money, 0, 5000) {
				break
			}
		}
	}
}
func getRandomNum(cnt int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(cnt)
}

func getInput() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)

	if err != nil {
		reader.ReadString('\n')
	}

	return n, err
}

func hasReachedCondition(money int, minimum int, maximum int) bool {
	if money >= maximum || money <= minimum {
		return true
	} else {
		return false
	}
}

func switchTest(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("Hello %d\n", int(t))
	case string:
		fmt.Printf("Hello %s\n", string(t))
	}
}
