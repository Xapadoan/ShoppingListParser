package menu

import (
	"strings"

	app "github.com/Xapadoan/shplsprsr/menu/application"
	dom "github.com/Xapadoan/shplsprsr/menu/domain"
	infra "github.com/Xapadoan/shplsprsr/menu/infrastructure"
	pres "github.com/Xapadoan/shplsprsr/menu/presentation"

	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type MenuGateway struct {
	logger log.ILogger
}

func NewMenuGateway(logger log.ILogger) *MenuGateway {
	return &MenuGateway{logger}
}

func (g *MenuGateway) RegisterRoutes(server srv.IServeHttp) {
	handler := pres.NewGetMenuCollectionHandler(g.logger)
	server.RegisterRoute(srv.Route{
		Method: srv.GET,
		Path:   "/menu/{id}",
		ParseURL: func(url string, req *srv.RouteRequest) *srv.ServerError {
			words := strings.Split(url, "/")
			req.Params = append(req.Params, words[2])

			return nil
		},
		HandleRequest: handler.HandleGetMenuCollection,
	})
}

func (g *MenuGateway) GetMenuCollection(id string) (*dom.MenuCollection, *dom.MenuError) {
	repo := infra.NewFileMenuRepository("../assets/menus", g.logger)
	useCase := app.NewGetMenuCollectionUsecase(repo.GetMenuCollection)
	menu, err := useCase.Exec(id)
	if err != nil {
		g.logger.Warn("Failed to get menu", id)
		return nil, err
	}

	return menu, nil
}
