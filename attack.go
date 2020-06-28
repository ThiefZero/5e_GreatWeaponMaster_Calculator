package main

// Attack represents a character's attack
type Attack struct {
	AttackBonus *int
	DamageDice  *string
	DamageBonus *int
}

func (atk *Attack) findEmpty() string {
	if atk.AttackBonus == nil {
		return "Attack Bonus"
	} else if atk.DamageDice == nil {
		return "Damage Dice"
	} else if atk.DamageBonus == nil {
		return "Damage Bonus"
	}
	return ""

}

// Reset resets the attack properties to nil
func (atk *Attack) reset() {
	atk.AttackBonus = nil
	atk.DamageDice = nil
	atk.DamageBonus = nil
}
