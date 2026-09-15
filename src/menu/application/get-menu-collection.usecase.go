package application

import (
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type GetMenuCollectionUsecase struct {
	getMenuCollection dom.IGetMenuCollection
}

func (u *GetMenuCollectionUsecase) Exec(id string) (*dom.MenuCollection, *dom.MenuError) {
	menu, err := u.getMenuCollection(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func NewGetMenuCollectionUsecase(getMenuCollection dom.IGetMenuCollection) *GetMenuCollectionUsecase {
	return &GetMenuCollectionUsecase{getMenuCollection}
}
