package application

import (
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type GetWeekMenuUseCase struct {
	repo dom.IGetWeekMenu
}

func (u *GetWeekMenuUseCase) Exec(id string) (*dom.WeekMenu, *dom.MenuError) {
	menu, err := u.repo.GetWeekMenu(id)
	if err != nil {
		return &dom.WeekMenu{}, err
	}

	return menu, nil
}

func NewGetWeekMenuUseCase(repo dom.IGetWeekMenu) *GetWeekMenuUseCase {
	return &GetWeekMenuUseCase{repo: repo}
}
