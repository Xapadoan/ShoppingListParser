package presentation

import (
	"encoding/json"

	dom "github.com/Xapadoan/shplsprsr/recipe/domain"

	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server/domain"
)

type GetRecipeHandler struct {
	logger    log.ILogger
	getRecipe dom.IGetRecipe
}

func NewGetRecipeHandler(logger log.ILogger, getRecipe dom.IGetRecipe) *GetRecipeHandler {
	return &GetRecipeHandler{logger, getRecipe}
}

func (h *GetRecipeHandler) HandleGetRecipe(req *srv.RouteRequest, res *srv.RouteResponse) *srv.ServerError {
	if len(req.Params) != 1 {
		return &srv.ServerError{Code: srv.BadRequest}
	}

	recipe, getErr := h.getRecipe(req.Params[0])
	if getErr != nil && getErr.IsNotFound() {
		h.logger.Debug("Recipe with id", req.Params[0], "not found")
		return &srv.ServerError{Code: srv.NotFound}
	}
	if getErr != nil {
		h.logger.Warn("Failed to get recipe with id", req.Params[0], ":\n", getErr.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	json, marshalError := json.Marshal(recipe)
	if marshalError != nil {
		h.logger.Warn("Marshal JSON Failed: ", marshalError.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	res.Status = srv.Ok
	res.Body = json

	return nil
}
