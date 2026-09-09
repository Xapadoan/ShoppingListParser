package application

import (
	log "github.com/Xapadoan/shplsprsr/logger"
	dom "github.com/Xapadoan/shplsprsr/recipe/domain"
)

type GetRecipeUseCase struct {
	logger    log.ILogger
	getRecipe dom.IGetRecipe
}

func (u GetRecipeUseCase) Exec(id string) (*dom.Recipe, *dom.RecipeError) {
	recipe, err := u.getRecipe(id)
	if err != nil {
		u.logger.Warn("Failed to get recipe with id ", id, ":\n", err.Error())
		return &dom.Recipe{}, err
	}

	return recipe, nil
}

func NewGetRecipeUseCase(getRecipe dom.IGetRecipe, logger log.ILogger) *GetRecipeUseCase {
	return &GetRecipeUseCase{logger, getRecipe}
}
