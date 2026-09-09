package domain

type MealMenu struct {
	Title     string
	RecipeIds []string
}

type DayMenu struct {
	Breakfast MealMenu
	Lunch     MealMenu
	Dinner    MealMenu
}

type WeekMenu struct {
	Monday    DayMenu
	Tuesday   DayMenu
	Wednesday DayMenu
	Thursday  DayMenu
	Friday    DayMenu
	Saturday  DayMenu
	Sunday    DayMenu
}
