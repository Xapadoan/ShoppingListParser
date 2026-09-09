package menu

import (
	srv "github.com/Xapadoan/shplsprsr/server"

	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type IMenuGateway interface {
	RegisterRoutes(server srv.IServeHttp)
	GetWeekMenu(id string) (*dom.WeekMenu, *dom.MenuError)
}

type WeekMenu = dom.WeekMenu
type DayMenu = dom.DayMenu
type MealMenu = dom.MealMenu

type MenuError = dom.MenuError
type MenuErrorCode = dom.MenuErrorCode

const (
	NotFound    MenuErrorCode = dom.NotFound
	FetchFailed               = dom.FetchFailed
	InvalidData               = dom.InvalidData
)
