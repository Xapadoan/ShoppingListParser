package domain

type MealMenu struct {
	Title     string
	RecipeIds []string
}

type MenuCollection struct {
	Id    string
	Menus []*MealMenu
}

