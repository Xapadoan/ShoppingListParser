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

func (repo *FileMenuRepository) StoreMenuCollection(collection *dom.MenuCollection) *dom.MenuError {
	path := repo.assetsPath + "/" + collection.Id + ".json"
	repo.logger.Debug("Storing File:", path)

	_, statErr := os.Stat(path)
	if statErr == nil || !errors.Is(statErr, os.ErrNotExist) {
		repo.logger.Warn("Menu with id", collection.Id, "already exists")
		return &dom.MenuError{Code: dom.AlreadyExists}
	}

	fileContent, marshalError := json.Marshal(collection)
	if marshalError != nil {
		repo.logger.Error(marshalError, "Failed to marshal collection", collection.Id)
		return &dom.MenuError{Code: dom.InvalidData}
	}

	writeErr := os.WriteFile(path, fileContent, 0666)
	if writeErr != nil {
		repo.logger.Error(writeErr, "Failed write file", path)
		return &dom.MenuError{Code: dom.StoreFailed}
	}

	return nil
}

func (repo *FileMenuRepository) EditMenuCollection(collection *dom.MenuCollection) *dom.MenuError {
	path := repo.assetsPath + "/" + collection.Id + ".json"
	repo.logger.Debug("Storing File:", path)

	_, statErr := os.Stat(path)
	if statErr != nil && errors.Is(statErr, os.ErrNotExist) {
		repo.logger.Warn("Menu with id", collection.Id, "not found")
		return &dom.MenuError{Code: dom.NotFound}
	}
	if statErr != nil {
		repo.logger.Warn("Stat error for menu with id:", collection.Id)
		return &dom.MenuError{Code: dom.StoreFailed}
	}

	fileContent, marshalError := json.Marshal(collection)
	if marshalError != nil {
		repo.logger.Error(marshalError, "Failed to marshal collection", collection.Id)
		return &dom.MenuError{Code: dom.InvalidData}
	}

	writeErr := os.WriteFile(path, fileContent, 0666)
	if writeErr != nil {
		repo.logger.Error(writeErr, "Failed write file", path)
		return &dom.MenuError{Code: dom.StoreFailed}
	}

	return nil
}
