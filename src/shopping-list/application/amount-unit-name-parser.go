package application

import (
	"regexp"
	"strconv"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
	log "github.com/Xapadoan/shplsprsr/logger"

	dom "github.com/Xapadoan/shplsprsr/shopping-list/domain"
)

type AmountUnitNameParser struct {
	logger log.ILogger
}

func NewAmountUnitParser(logger log.ILogger) *AmountUnitNameParser {
	return &AmountUnitNameParser{logger}
}

func (p *AmountUnitNameParser) Parse(text string) (*ing.IngredientQuantity, *dom.ShoppingListError) {
	re := regexp.MustCompile("([0-9]+) ?(g|u|cas|cac|ml|mL|cl|cL)? (.*)")
	matchs := re.FindStringSubmatch(text)
	if len(matchs) < 1 {
		msg := "String" + text + "does not match regexp" + re.String()
		p.logger.Warn(msg)
		return &ing.IngredientQuantity{}, &dom.ShoppingListError{Code: dom.ParsingFailed, Message: msg}
	}

	quantity, quantityErr := strconv.ParseUint(matchs[1], 10, 64)
	if quantityErr != nil {
		msg := "Failed to parse" + matchs[1] + "as integer:\n" + quantityErr.Error()
		p.logger.Warn(msg)
		return &ing.IngredientQuantity{}, &dom.ShoppingListError{Code: dom.ParsingFailed, Message: msg}
	}

	unit, ratio := parseUnitAndRatio(matchs[2])

	p.logger.Debug("AmountUnitName parser success")
	return &ing.IngredientQuantity{Name: matchs[3], Amount: float32(quantity) * ratio, Unit: unit}, nil
}

func parseUnitAndRatio(text string) (ing.IngredientUnit, float32) {
	switch text {
	case "g":
		return ing.Gram, 1
	case "ml", "mL":
		return ing.Milliliter, 1
	case "cl", "cL":
		return ing.Milliliter, 10
	case "cas", "Cas":
		return ing.Cas, 1
	case "cac", "Cac":
		return ing.Cac, 1
	default:
		return ing.Unit, 1
	}
}
