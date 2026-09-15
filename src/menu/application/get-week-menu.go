package application

import (
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type GetWeekMenuUseCase struct {
	getMenuCollection dom.IGetMenuCollection
}

func (u *GetWeekMenuUseCase) Exec(id string) (*dom.MenuCollection, *dom.MenuError) {
	menu, err := u.getMenuCollection(id)
	if err != nil {
		return nil, err
	}

	return menu, nil
}

func NewGetWeekMenuUseCase(getMenuCollection dom.IGetMenuCollection) *GetWeekMenuUseCase {
	return &GetWeekMenuUseCase{getMenuCollection}
}
