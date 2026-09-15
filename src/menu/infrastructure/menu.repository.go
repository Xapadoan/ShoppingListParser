package infrastructure

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/Xapadoan/shplsprsr/logger"
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

type DayMenu struct {
	Breakfast dom.MealMenu
	Lunch     dom.MealMenu
	Dinner    dom.MealMenu
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

func (weekMenu *WeekMenu) Adapter(id string) *dom.MenuCollection {
	menus := []*dom.MealMenu{}
	menus = append(menus, &weekMenu.Monday.Breakfast)
	menus = append(menus, &weekMenu.Monday.Lunch)
	menus = append(menus, &weekMenu.Monday.Dinner)
	menus = append(menus, &weekMenu.Tuesday.Breakfast)
	menus = append(menus, &weekMenu.Tuesday.Lunch)
	menus = append(menus, &weekMenu.Tuesday.Dinner)
	menus = append(menus, &weekMenu.Wednesday.Breakfast)
	menus = append(menus, &weekMenu.Wednesday.Lunch)
	menus = append(menus, &weekMenu.Wednesday.Dinner)
	menus = append(menus, &weekMenu.Thursday.Breakfast)
	menus = append(menus, &weekMenu.Thursday.Lunch)
	menus = append(menus, &weekMenu.Thursday.Dinner)
	menus = append(menus, &weekMenu.Friday.Breakfast)
	menus = append(menus, &weekMenu.Friday.Lunch)
	menus = append(menus, &weekMenu.Friday.Dinner)
	menus = append(menus, &weekMenu.Saturday.Breakfast)
	menus = append(menus, &weekMenu.Saturday.Lunch)
	menus = append(menus, &weekMenu.Saturday.Dinner)
	menus = append(menus, &weekMenu.Sunday.Breakfast)
	menus = append(menus, &weekMenu.Sunday.Lunch)
	menus = append(menus, &weekMenu.Sunday.Dinner)

	return &dom.MenuCollection{Id: id, Menus: menus}
}

type FileMenuRepository struct {
	assetsPath string
	logger     log.ILogger
}

func NewFileMenuRepository(assetsPath string, logger log.ILogger) *FileMenuRepository {
	return &FileMenuRepository{assetsPath, logger}
}

func (repo *FileMenuRepository) GetMenuCollection(id string) (*dom.MenuCollection, *dom.MenuError) {
	path := repo.assetsPath + "/" + id + ".json"
	repo.logger.Debug("Finding File: ", path)
	_, statErr := os.Stat(path)
	if statErr != nil && errors.Is(statErr, os.ErrNotExist) {
		repo.logger.Debug("File ", path, " not found")
		return nil, &dom.MenuError{Code: dom.NotFound}
	} else if statErr != nil {
		repo.logger.Warn("Failed to check file existence ", path)
		return nil, &dom.MenuError{Code: dom.FetchFailed}
	}

	data, readErr := os.ReadFile(path)
	if readErr != nil {
		repo.logger.Warn("Failed to read file ", path)
		return nil, &dom.MenuError{Code: dom.FetchFailed}
	}

	var collection dom.MenuCollection
	jsonError := json.Unmarshal(data, &collection)
	if jsonError != nil {
		repo.logger.Warn("Invalid data in file ", path)
		return nil, &dom.MenuError{Code: dom.InvalidData}
	}

	return &collection, nil
}
