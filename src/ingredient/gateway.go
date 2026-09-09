package ingredient

import (
	log "github.com/Xapadoan/shplsprsr/logger"

	app "github.com/Xapadoan/shplsprsr/ingredient/application"
)

type IngredientGateway struct {
	logger log.ILogger
}

func NewIngredientGateway(logger log.ILogger) *IngredientGateway {
	return &IngredientGateway{logger}
}

func (g *IngredientGateway) ParseIngredientQuantity(text string, parsers []IParseIngredientQuantity) (*IngredientQuantity, *IngredientError) {
	usecase := app.NewParseIngredientQuantityUsecase(g.logger, parsers)
	return usecase.Exec(text)
}
