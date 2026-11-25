package infrastructure

import (
	"io"
	"net/http"
	"strconv"

	dom "github.com/Xapadoan/shplsprsr/server/domain"

	log "github.com/Xapadoan/shplsprsr/logger"
)

type HttpMuxServer struct {
	http   *http.ServeMux
	logger log.ILogger
}

func NewHttpMuxServer(logger log.ILogger) *HttpMuxServer {
	return &HttpMuxServer{http: http.NewServeMux(), logger: logger}
}

func (server *HttpMuxServer) RegisterRoute(route dom.Route) {
	server.logger.Debug("Registering route", route.Method.String(), route.Path)
	server.http.HandleFunc(route.Method.String()+" "+route.Path, func(res http.ResponseWriter, req *http.Request) {
		server.logger.Debug(req.Method, req.URL.Path)
		request := dom.RouteRequest{}
		parseUrlErr := route.ParseURL(req.URL.Path, &request)
		if parseUrlErr != nil {
			res.WriteHeader(int(dom.BadRequestError))
			res.Write([]byte("BadRequest"))
			return
		}
		body, bodyErr := io.ReadAll(req.Body)
		if bodyErr != nil {
			res.WriteHeader(int(dom.BadRequestError))
			res.Write([]byte("BadBody"))
			return
		}
		request.Body = string(body)

		response := dom.RouteResponse{}

		err := route.HandleRequest(&request, &response)
		if err != nil {
			res.WriteHeader(int(dom.StatusCodeFromServerErrorCode(err.Code)))
			res.Write([]byte(err.String()))
			return
		}

		res.WriteHeader((int)(response.Status))
		res.Write(response.Body)
	})
}

func addCorsHeaders(res http.ResponseWriter) {
	headers := res.Header()
	headers.Add("Access-Control-Allow-Origin", "http://localhost:4200")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept")
	headers.Add("Access-Control-Allow-Methods", "GET,POST,PUT,OPTIONS")
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addCorsHeaders(w)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (server *HttpMuxServer) Start(port int) {
	server.logger.Debug("Starting HttpMuxServer")

	portAsString := strconv.Itoa(port)
	server.logger.Info("Server Listening on http://localhost:" + portAsString)
	http.ListenAndServe(":"+portAsString, corsMiddleware(server.http))
}
