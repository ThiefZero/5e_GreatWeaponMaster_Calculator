package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

const _amountOfInput = 2 // expecting '1d12' and '5' (dice and atk bonus)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var expectedInput []string
	fmt.Print(generateExpectedInputQuestion(len(expectedInput)))

	for scanner.Scan() {
		input := scanner.Text()
		expectedInput = append(expectedInput, input)
		fmt.Print(generateExpectedInputQuestion(len(expectedInput)))

		if expectedInputIsComplete(expectedInput) {
			// check all inputs
			damageDice := expectedInput[0]

			// check damageBonus
			damageBonus, err := strconv.Atoi(expectedInput[1])
			if err != nil {
				fmt.Printf("Invalid damageBonus: %v %v\n\n", expectedInput[1], err)
				expectedInput = nil
				fmt.Print(generateExpectedInputQuestion(len(expectedInput)))
				continue
			}

			fmt.Printf("Average Damage: %v\n\n", calcAvgDmg(damageDice, damageBonus))
			expectedInput = nil
			fmt.Print(generateExpectedInputQuestion(len(expectedInput)))

		}
	}
}

func generateExpectedInputQuestion(amountInputted int) string {
	switch amountInputted {
	case 0:
		return "Damage dice used: "
	case 1:
		return "Damage bonus: "
	}

	return ""
}

func expectedInputIsComplete(input []string) bool {
	return len(input) == _amountOfInput
}

func calcAvgDmg(damageDice string, damageBonus int) decimal.Decimal {
	var dice []decimal.Decimal
	var avgDmgOfDice decimal.Decimal

	diceStrings := strings.Split(damageDice, "d") // '1d12' -> [1 12]

	// convert to int and put in dice array
	for _, element := range diceStrings {
		converted, err := strconv.Atoi(element)

		if err != nil {
			log.Print(err)
		}

		dice = append(dice, decimal.NewFromInt(int64(converted)))
	}

	maxValue := decimal.NewFromInt(1).Add(dice[1]) // ex: 1 + 12 = 13
	average := maxValue.Div(decimal.NewFromInt(2)) // 13 / 2 = 6.5 for 1d12 die
	avgDmgOfDice = average.Mul(dice[0])            // 6.5 * 2 for 2 dice

	return decimal.NewFromInt(int64(damageBonus)).Add(avgDmgOfDice)
}
