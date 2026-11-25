package application

import (
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type UpdateMenuCollectionParams struct {
	Menus []*dom.MealMenu
}

type UpdateMenuCollectionUsecase struct {
	editMenuCollection dom.IEditMenuCollection
}

func NewUpdateCollectionUsecase(editMenuCollection dom.IEditMenuCollection) *UpdateMenuCollectionUsecase {
	return &UpdateMenuCollectionUsecase{editMenuCollection}
}

func (u *UpdateMenuCollectionUsecase) Exec(id string, params *UpdateMenuCollectionParams) (*dom.MenuCollection, *dom.MenuError) {
	menu := &dom.MenuCollection{Id: id, Menus: params.Menus}
	return menu, u.editMenuCollection(menu)
}
