package ingredient

import (
	log "github.com/Xapadoan/shplsprsr/logger"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

type IngredientGateway struct {
	logger log.ILogger
}

func NewIngredientGateway(logger log.ILogger) *IngredientGateway {
	return &IngredientGateway{logger}
}

func (g *IngredientGateway) ParseIngredientQuantity(text string) (*IngredientQuantity, *IngredientError) {
	return dom.ParseIngredientQuantity(text)
}
