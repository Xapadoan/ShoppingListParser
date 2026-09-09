package application

import (
	log "github.com/Xapadoan/shplsprsr/logger"

	dom "github.com/Xapadoan/shplsprsr/ingredient/domain"
)

type ParseIngredientQuantityUsecase struct {
	logger  log.ILogger
	parsers []dom.IParseIngredientQuantity
}

func NewParseIngredientQuantityUsecase(logger log.ILogger, parsers []dom.IParseIngredientQuantity) *ParseIngredientQuantityUsecase {
	return &ParseIngredientQuantityUsecase{logger, parsers}
}

func (u *ParseIngredientQuantityUsecase) Exec(text string) (*dom.IngredientQuantity, *dom.IngredientError) {
	for _, parse := range u.parsers {
		i, err := parse(text)
		if err == nil {
			return i, nil
		}
	}

	return nil, &dom.IngredientError{Code: dom.ParsingFailed}
}
