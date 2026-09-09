package presentation

import (
	"encoding/json"

	rcp "github.com/Xapadoan/shplsprsr/recipe"
	app "github.com/Xapadoan/shplsprsr/shopping-list/application"

	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type CreateShoppingListHttpHandler struct {
	logger        log.ILogger
	useCase       app.CreateShoppingListUsecase
	recipeGateway rcp.IRecipeGateway
}

func NewCreateShoppingListHttpHandler(
	logger log.ILogger,
	useCase *app.CreateShoppingListUsecase,
	recipeGateway rcp.IRecipeGateway,
) *CreateShoppingListHttpHandler {
	return &CreateShoppingListHttpHandler{logger, *useCase, recipeGateway}
}

func (h *CreateShoppingListHttpHandler) HandleGetRecipeShoppingList(req *srv.RouteRequest, res *srv.RouteResponse) *srv.ServerError {
	body, bodyErr := parseBody(req.Body)
	if bodyErr != nil {
		h.logger.Debug("Bad Body")
		return &srv.ServerError{Code: srv.BadRequest}
	}

	list, listErr := h.useCase.Exec(body)
	if listErr != nil {
		h.logger.Warn("Failed to create shopping list for recipe", req.Params[0])
		return &srv.ServerError{Code: srv.Internal}
	}

	json, jsonErr := json.Marshal(list)
	if jsonErr != nil {
		h.logger.Warn("Failed to Marshal JSON:\n", jsonErr.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	res.Status = srv.Ok
	res.Body = json

	return nil
}

func parseBody(body string) (*app.CreateShoppingListParams, error) {
	var params app.CreateShoppingListParams
	err := json.Unmarshal([]byte(body), &params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}
