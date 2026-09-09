package application

import (
	"strconv"

	log "github.com/Xapadoan/shplsprsr/logger"
	rcp "github.com/Xapadoan/shplsprsr/recipe"

	dom "github.com/Xapadoan/shplsprsr/shopping-list/domain"
)

type CreateShoppingListParams = []*struct {
	RecipeIds            []string
	NumberOfPeopleEating uint8
}

type CreateShoppingListUsecase struct {
	recipeGateway rcp.IRecipeGateway
	logger        log.ILogger
}

func NewCreateShoppingListUsecase(recipeGateway rcp.IRecipeGateway, logger log.ILogger) *CreateShoppingListUsecase {
	return &CreateShoppingListUsecase{recipeGateway, logger}
}

func (u *CreateShoppingListUsecase) Exec(params *CreateShoppingListParams) (*dom.ShoppingList, *dom.ShoppingListError) {
	list := &dom.ShoppingList{}

	for _, mealConfig := range *params {
		recipes := u.recipeGateway.FindRecipes(&rcp.FindRecipesParams{Ids: mealConfig.RecipeIds})
		if len(recipes) != len(mealConfig.RecipeIds) {
			u.logger.Warn("Find Recipes Partially failed (" + strconv.Itoa(len(recipes)) + "/" + strconv.Itoa(len(mealConfig.RecipeIds)) + ")")
		}

		for _, recipe := range recipes {
			recipe.AdaptQuantity(uint16(mealConfig.NumberOfPeopleEating))
			list.AddRecipeIngredients(recipe)
		}
	}

	return list, nil
}
