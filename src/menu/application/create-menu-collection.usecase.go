package application

import (
	"uuid"

	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type CreateMenuCollectionParams struct {
	Menus []*dom.MealMenu
}

type CreateMenuCollectionUsecase struct {
	storeMenuCollection dom.IStoreMenuCollection
}

func (u *CreateMenuCollectionUsecase) Exec(params *CreateMenuCollectionParams) (*dom.MenuCollection, *dom.MenuError) {
	id := uuid.NewV7()
	menu := &dom.MenuCollection{Id: id.String(), Menus: params.Menus}
	err := u.storeMenuCollection(menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func NewCreateMenuCollectionUsecase(storeMenuCollection dom.IStoreMenuCollection) *CreateMenuCollectionUsecase {
	return &CreateMenuCollectionUsecase{storeMenuCollection}
}
