package presentation

import (
	"encoding/json"

	log "github.com/Xapadoan/shplsprsr/logger"
	app "github.com/Xapadoan/shplsprsr/menu/application"

	infra "github.com/Xapadoan/shplsprsr/menu/infrastructure"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type UpdateMenuCollectionHandler struct {
	logger log.ILogger
}

func NewUpdateMenuCollectionHandler(logger log.ILogger) *UpdateMenuCollectionHandler {
	return &UpdateMenuCollectionHandler{logger}
}

func (h *UpdateMenuCollectionHandler) HandleUpdateMenuCollection(req *srv.RouteRequest, res *srv.RouteResponse) *srv.ServerError {
	h.logger.Debug("Updating menu: ", req.Params[0])

	params, parseBodyErr := parseUpdateMenuCollectionParams(req.Body)
	if parseBodyErr != nil {
		h.logger.Debug("Bad Body:\n", req.Body, "\nParsing failed:\n", parseBodyErr.Error())
		return &srv.ServerError{Code: srv.BadRequest}
	}
	repo := infra.NewFileMenuRepository("../assets/menus", h.logger)
	usecase := app.NewUpdateCollectionUsecase(repo.EditMenuCollection)
	collection, usecaseErr := usecase.Exec(req.Params[0], params)
	if usecaseErr != nil {
		h.logger.Warn("Failed to update menu with id:", req.Params[0])
		return &srv.ServerError{Code: srv.Internal}
	}

	json, marshalError := json.Marshal(collection)
	if marshalError != nil {
		h.logger.Warn("Marshal JSON Failed:\n", marshalError.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	res.Status = srv.Ok
	res.Body = json

	return nil
}

func parseUpdateMenuCollectionParams(body string) (*app.UpdateMenuCollectionParams, error) {
	var params app.UpdateMenuCollectionParams
	err := json.Unmarshal([]byte(body), &params)
	if err != nil {
		return nil, err
	}

	return &params, nil
}
