package infrastructure

import (
	"encoding/json"
	"errors"
	"os"

	log "github.com/Xapadoan/shplsprsr/logger"
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
)

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
