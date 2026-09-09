package domain

type IGetWeekMenu interface {
	GetWeekMenu(id string) (*WeekMenu, *MenuError)
}
