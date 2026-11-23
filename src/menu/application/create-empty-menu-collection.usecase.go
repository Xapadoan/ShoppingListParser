package application

import (
	"uuid"

	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type CreateEmptyMenuCollectionUsecase struct {
	storeMenuCollection dom.IStoreMenuCollection
}

func (u *CreateEmptyMenuCollectionUsecase) Exec() (*dom.MenuCollection, *dom.MenuError) {
	id := uuid.NewV7()
	menu := &dom.MenuCollection{Id: id.String(), Menus: []*dom.MealMenu{}}
	err := u.storeMenuCollection(menu)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func NewCreateEmptyMenuCollectionUsecase(storeMenuCollection dom.IStoreMenuCollection) *CreateEmptyMenuCollectionUsecase {
	return &CreateEmptyMenuCollectionUsecase{storeMenuCollection}
}
