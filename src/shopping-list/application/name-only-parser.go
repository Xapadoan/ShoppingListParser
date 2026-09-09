package application

import (
	dom "github.com/Xapadoan/shplsprsr/shopping-list/domain"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
	log "github.com/Xapadoan/shplsprsr/logger"
)

type NameOnlyParser struct {
	logger log.ILogger
}

func NewNameOnlyParser(logger log.ILogger) *NameOnlyParser {
	return &NameOnlyParser{logger}
}

func (p *NameOnlyParser) Parse(text string) (*ing.IngredientQuantity, *dom.ShoppingListError) {
	p.logger.Debug("Name Only parser success")
	return &ing.IngredientQuantity{Name: text}, nil
}
