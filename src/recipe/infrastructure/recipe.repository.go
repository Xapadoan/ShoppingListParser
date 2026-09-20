package infrastructure

import (
	"encoding/json"
	"errors"
	"os"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
	log "github.com/Xapadoan/shplsprsr/logger"

	dom "github.com/Xapadoan/shplsprsr/recipe/domain"
)

type FileRecipeRepository struct {
	assetsPath              string
	parseIngredientQuantity ing.IParseIngredientQuantity
	logger                  log.ILogger
}

func NewFileRecipeRepository(assetsPath string, parseIngredientQuantity ing.IParseIngredientQuantity, logger log.ILogger) *FileRecipeRepository {
	return &FileRecipeRepository{assetsPath, parseIngredientQuantity, logger}
}

func (repo *FileRecipeRepository) Get(id string) (*dom.Recipe, *dom.RecipeError) {
	path := repo.assetsPath + "/" + id + ".json"
	repo.logger.Debug("Finding File ", path)
	_, statErr := os.Stat(path)
	if statErr != nil && errors.Is(statErr, os.ErrNotExist) {
		repo.logger.Debug("File ", path, " not found")
		return nil, dom.NewRecipeError(dom.NotFound, "File not found")
	} else if statErr != nil {
		repo.logger.Warn("Failed to check file existence ", path)
		return nil, dom.NewRecipeError(dom.FetchFailed, statErr.Error())
	}
	data, readError := os.ReadFile(path)
	if readError != nil {
		repo.logger.Warn("Failed to read file ", path)
		return nil, dom.NewRecipeError(dom.FetchFailed, readError.Error())
	}

	var jsonRecipe JsonRecipe
	jsonError := json.Unmarshal(data, &jsonRecipe)
	if jsonError != nil {
		repo.logger.Warn("Invalid data in file", path, ":\n", jsonError.Error())
		return nil, dom.NewRecipeError(dom.InvalidData, jsonError.Error())
	}

	recipe, adapterError := jsonRecipe.DomainAdapter(repo.parseIngredientQuantity)
	if adapterError != nil {
		repo.logger.Warn("Failed to adapt json recipe")
		return nil, dom.NewRecipeError(dom.InvalidData, adapterError.Error())
	}

	return recipe, nil
}

func (repo *FileRecipeRepository) Find(params *dom.FindRecipesParams) []*dom.Recipe {
	var recipes []*dom.Recipe
	for _, id := range params.Ids {
		recipe, err := repo.Get(id)
		if err != nil {
			repo.logger.Warn("Failed to get recipe", id)
		} else {
			recipes = append(recipes, recipe)
		}
	}

	return recipes
}
