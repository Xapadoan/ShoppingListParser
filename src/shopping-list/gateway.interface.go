package shoppinglist

import (
	app "github.com/Xapadoan/shplsprsr/shopping-list/application"
	dom "github.com/Xapadoan/shplsprsr/shopping-list/domain"

	srv "github.com/Xapadoan/shplsprsr/server"
)

type ShoppingList = dom.ShoppingList

type ShoppingListError = dom.ShoppingListError

type CreateShoppingListParams = app.CreateShoppingListParams

type IShoppingListGateway interface {
	RegisterRoutes(server srv.IServeHttp)
}
