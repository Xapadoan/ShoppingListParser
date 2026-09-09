package recipe

import (
	"strings"

	app "github.com/Xapadoan/shplsprsr/recipe/application"
	dom "github.com/Xapadoan/shplsprsr/recipe/domain"
	infra "github.com/Xapadoan/shplsprsr/recipe/infrastructure"
	pres "github.com/Xapadoan/shplsprsr/recipe/presentation"

	ing "github.com/Xapadoan/shplsprsr/ingredient"
	log "github.com/Xapadoan/shplsprsr/logger"
	srv "github.com/Xapadoan/shplsprsr/server"
)

type RecipeGateway struct {
	ingredientGateway ing.IIngredientGateway
	logger            log.ILogger
}

func NewRecipeGateway(logger log.ILogger, ingredientGateway ing.IIngredientGateway) *RecipeGateway {
	return &RecipeGateway{ingredientGateway, logger}
}

func (g *RecipeGateway) GetRecipe(id string) (*Recipe, *RecipeError) {
	repo := infra.NewFileRecipeRepository("../assets/recipes", g.logger)
	useCase := app.NewGetRecipeUseCase(repo.Get, g.logger)

	return useCase.Exec(id)
}

func (g *RecipeGateway) FindRecipes(params *dom.FindRecipesParams) []*dom.Recipe {
	repo := infra.NewFileRecipeRepository("../assets/recipes", g.logger)
	useCase := app.NewFindRecipesUseCase(g.logger, repo)

	return useCase.Exec(params)
}

func (g *RecipeGateway) RegisterRoutes(server srv.IServeHttp) {
	handler := pres.NewGetRecipeHandler(g.logger, g.GetRecipe)
	server.RegisterRoute(srv.Route{
		Method: srv.GET,
		Path:   "/recipe/{id}",
		ParseURL: func(url string, req *srv.RouteRequest) *srv.ServerError {
			words := strings.Split(url, "/")
			req.Params = append(req.Params, words[2])

			return nil
		},
		HandleRequest: handler.HandleGetRecipe,
	})
}
