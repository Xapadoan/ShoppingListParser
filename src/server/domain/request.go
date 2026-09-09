package domain

type RouteParams []string
type RouteQuery map[string][]string
type RouteRequest struct {
	Path   string
	Params RouteParams
	Query  RouteQuery
	Body   string
}
