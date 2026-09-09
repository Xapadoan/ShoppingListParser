package application

import (
	log "github.com/Xapadoan/shplsprsr/logger"

	dom "github.com/Xapadoan/shplsprsr/recipe/domain"
)

type FindRecipesUseCase struct {
	logger log.ILogger
	repo   dom.IFindRecipe
}

func NewFindRecipesUseCase(logger log.ILogger, repo dom.IFindRecipe) *FindRecipesUseCase {
	return &FindRecipesUseCase{logger, repo}
}

func (u *FindRecipesUseCase) Exec(params *dom.FindRecipesParams) []*dom.Recipe {
	recipes := u.repo.Find(params)
	return recipes
}
