package gross

const (
	QuarterOfADozen = 3
	HalfOfADozen    = 6
	Dozen           = 12
	SmallGross      = 120
	Gross           = 144
	GreatGross      = 1728
)

// Units stores the GrossStore unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": QuarterOfADozen,
		"half_of_a_dozen":    HalfOfADozen,
		"dozen":              Dozen,
		"small_gross":        SmallGross,
		"gross":              Gross,
		"great_gross":        GreatGross,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	unitValue, existsUnits := units[unit]

	if !existsUnits {
		return false
	}

	bill[item] += unitValue

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	currentQuantity, itemExists := bill[item]
	if !itemExists {
		return false
	}

	unitValue, unitExists := units[unit]
	if !unitExists {
		return false
	}

	newQuantity := currentQuantity - unitValue

	switch {
	case newQuantity > 0:
		bill[item] = newQuantity
	case newQuantity == 0:
		delete(bill, item)
	default:
		return false
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	if !exists {
		return 0, false
	}
	return quantity, true
}
