package presentation

import (
	"encoding/json"

	app "github.com/Xapadoan/shplsprsr/menu/application"
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
	infra "github.com/Xapadoan/shplsprsr/menu/infrastructure"

	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type GetMenuCollectionHandler struct {
	logger log.ILogger
}

func NewGetMenuCollectionHandler(logger log.ILogger) *GetMenuCollectionHandler {
	return &GetMenuCollectionHandler{logger}
}

func (h *GetMenuCollectionHandler) HandleGetMenuCollection(req *srv.RouteRequest, res *srv.RouteResponse) *srv.ServerError {
	if len(req.Params) != 1 {
		h.logger.Debug("Invalid Parameters")
		return &srv.ServerError{Code: srv.BadRequest}
	}

	repo := infra.NewFileMenuRepository("../assets/menus", h.logger)
	usecase := app.NewGetMenuCollectionUsecase(repo.GetMenuCollection)
	menu, usecaseErr := usecase.Exec(req.Params[0])
	if usecaseErr != nil && usecaseErr.Code == dom.NotFound {
		h.logger.Debug("Menu with id", req.Params[0], "not found")
		return &srv.ServerError{Code: srv.NotFound}
	}
	if usecaseErr != nil {
		h.logger.Warn("Failed to get menu with id", req.Params[0], ":\n", usecaseErr.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	json, marshalError := json.Marshal(menu)
	if marshalError != nil {
		h.logger.Warn("Marshal JSON Failed:\n", marshalError.Error())
		return &srv.ServerError{Code: srv.Internal}
	}

	res.Status = srv.Ok
	res.Body = json

	return nil
}
