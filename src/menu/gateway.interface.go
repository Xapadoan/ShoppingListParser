package menu

import (
	srv "github.com/Xapadoan/shplsprsr/server"

	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type IMenuGateway interface {
	RegisterRoutes(server srv.IServeHttp)
	GetWeekMenu(id string) (*dom.MenuCollection, *dom.MenuError)
}

type MealMenu = dom.MealMenu
type MenuCollection = dom.MenuCollection

type MenuError = dom.MenuError
type MenuErrorCode = dom.MenuErrorCode

const (
	NotFound    MenuErrorCode = dom.NotFound
	FetchFailed               = dom.FetchFailed
	InvalidData               = dom.InvalidData
)
