package shoppinglist

import (
	log "github.com/Xapadoan/shplsprsr/logger"
	mnu "github.com/Xapadoan/shplsprsr/menu"
	rcp "github.com/Xapadoan/shplsprsr/recipe"
	srv "github.com/Xapadoan/shplsprsr/server"

	app "github.com/Xapadoan/shplsprsr/shopping-list/application"
	pres "github.com/Xapadoan/shplsprsr/shopping-list/presentation"
)

type ShoppingListGateway struct {
	logger        log.ILogger
	recipeGateway rcp.IRecipeGateway
	menuGateway   mnu.IMenuGateway
}

func NewShoppingListGateway(logger log.ILogger, recipeGateway rcp.IRecipeGateway, menuGateway mnu.IMenuGateway) *ShoppingListGateway {
	return &ShoppingListGateway{logger, recipeGateway, menuGateway}
}

func (g *ShoppingListGateway) RegisterRoutes(server srv.IServeHttp) {
	usecase := app.NewCreateShoppingListUsecase(g.recipeGateway, g.logger)
	getFromRecipeHandler := pres.NewCreateShoppingListHttpHandler(g.logger, usecase, g.recipeGateway)
	server.RegisterRoute(srv.Route{
		Method: srv.POST,
		Path:   "/shopping-list",
		ParseURL: func(url string, req *srv.RouteRequest) *srv.ServerError {
			return nil
		},
		HandleRequest: getFromRecipeHandler.HandleGetRecipeShoppingList,
	})
}
