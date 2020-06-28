package main

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"log"
	"os"
	"strconv"
	"strings"
)

const _amountOfInput = 2 // expecting '1d12' and '5' (dice and atk bonus)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var attack Attack
	fmt.Print(generateExpectedInputQuestion(attack))

	for scanner.Scan() {
		input := scanner.Text()

		empty := attack.findEmpty()

		//		fmt.Printf("response: %v", empty)

		switch empty {
		case "Attack Bonus":
			atkBonus, err := strconv.Atoi(input)
			if err != nil {
				fmt.Print(err)
				continue
			}
			attack.AttackBonus = &atkBonus

		case "Damage Dice":
			attack.DamageDice = &input

		case "Damage Bonus":
			dmgBonus, err := strconv.Atoi(input)
			if err != nil {
				fmt.Print(err)
				continue
			}
			attack.DamageBonus = &dmgBonus

		}

		if attack.findEmpty() == "" {
			// Done, print answer
			avgDmg := calcAvgDmg(*attack.DamageDice, *attack.DamageBonus)
			fmt.Printf("Average Damage: %v\n\n", avgDmg)

			maxAC := calcMaxAC(*attack.AttackBonus, avgDmg)
			fmt.Printf("Max AC: %v\n\n", maxAC)
			attack.reset()
		}
		fmt.Print(generateExpectedInputQuestion(attack))
	}
}

func generateExpectedInputQuestion(atk Attack) string {
	return atk.findEmpty() + ": "
}

//func expectedInputIsComplete(attack Attack) bool {
//	return attack.findEmpty == ""
//}

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

func calcMaxAC(atkBonus int, avgDmg decimal.Decimal) decimal.Decimal {
	_atkBonus := decimal.NewFromInt(int64(atkBonus))
	_halfDmg := avgDmg.Div(decimal.NewFromInt(2))
	_sixTeen := decimal.NewFromInt(16)

	return _atkBonus.Sub(_halfDmg).Add(_sixTeen)
}
