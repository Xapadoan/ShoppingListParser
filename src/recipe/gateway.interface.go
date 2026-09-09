package recipe

import (
	dom "github.com/Xapadoan/shplsprsr/recipe/domain"

	srv "github.com/Xapadoan/shplsprsr/server"
)

type IRecipeGateway interface {
	RegisterRoutes(server srv.IServeHttp)
	GetRecipe(id string) (*dom.Recipe, *dom.RecipeError)
	FindRecipes(params *dom.FindRecipesParams) []*dom.Recipe
}

type Recipe = dom.Recipe

type RecipeError = dom.RecipeError
type RecipeErrorCode = dom.RecipeErrorCode

const (
	NotFound RecipeErrorCode = dom.NotFound
)

type FindRecipesParams = dom.FindRecipesParams
