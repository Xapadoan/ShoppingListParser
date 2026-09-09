package main

import (
	ing "github.com/Xapadoan/shplsprsr/ingredient"
	log "github.com/Xapadoan/shplsprsr/logger"
	mnu "github.com/Xapadoan/shplsprsr/menu"
	rcp "github.com/Xapadoan/shplsprsr/recipe"
	srv "github.com/Xapadoan/shplsprsr/server"
	shp "github.com/Xapadoan/shplsprsr/shopping-list"
)

func main() {
	loggerGateway := log.LoggerGateway{}
	serverGateway := srv.NewServerGateway(&loggerGateway)
	server := serverGateway.NewHttpServer()

	menuGateway := mnu.NewMenuGateway(loggerGateway.NewLogger("[MenuGateway]"))
	menuGateway.RegisterRoutes(server)

	ingredientGateway := ing.NewIngredientGateway(loggerGateway.NewLogger("[IngredientGateway]"))

	recipeGateway := rcp.NewRecipeGateway(loggerGateway.NewLogger("[RecipeGateway]"), ingredientGateway)
	recipeGateway.RegisterRoutes(server)

	shoppingListGateway := shp.NewShoppingListGateway(
		loggerGateway.NewLogger("[ShoppingListGateway]"),
		recipeGateway,
		menuGateway,
	)
	shoppingListGateway.RegisterRoutes(server)
	server.Start(8040)
}
