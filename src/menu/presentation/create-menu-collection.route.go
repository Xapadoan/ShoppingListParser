package presentation

import (
	"encoding/json"

	app "github.com/Xapadoan/shplsprsr/menu/application"
	infra "github.com/Xapadoan/shplsprsr/menu/infrastructure"

	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type CreateMenuCollectionHandler struct {
	logger log.ILogger
}

func NewCreateMenuCollectionHandler(logger log.ILogger) *CreateMenuCollectionHandler {
	return &CreateMenuCollectionHandler{logger}
}

func (h *CreateMenuCollectionHandler) HandleCreateMenuCollection(req *srv.RouteRequest, res *srv.RouteResponse) *srv.ServerError {
	h.logger.Debug("Creating new menu")

	params, parseBodyErr := parseCreateMenuCollectionParams(req.Body)
	if parseBodyErr != nil {
		h.logger.Debug("Bad Body:\n", req.Body, "\nParsing failed:\n", parseBodyErr.Error())
		return &srv.ServerError{Code: srv.BadRequest}
	}
	repo := infra.NewFileMenuRepository("../assets/menus", h.logger)
	usecase := app.NewCreateMenuCollectionUsecase(repo.StoreMenuCollection)
	collection, usecaseErr := usecase.Exec(params)
	if usecaseErr != nil {
		h.logger.Warn("Failed to create menu")
		return &srv.ServerError{Code: srv.Internal}
	}

	json, marshalError := json.Marshal(collection)
	if marshalError != nil {
		h.logger.Warn("Marshal JSON Failed:\n", marshalError.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	res.Status = srv.Created
	res.Body = json

	return nil
}

func parseCreateMenuCollectionParams(body string) (*app.CreateMenuCollectionParams, error) {
	var params app.CreateMenuCollectionParams
	err := json.Unmarshal([]byte(body), &params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}
