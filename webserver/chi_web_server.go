package webserver

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

type ChiWebServer struct {
	engine *chi.Mux
	srv    *http.Server
	wg     *sync.WaitGroup
}

func Chi() Server {
	return &ChiWebServer{}
}

func (s *ChiWebServer) Initialize() {
	s.engine = chi.NewRouter()
	s.wg = observer.GetWaitGroup()
}

func (s *ChiWebServer) Shutdown() error {
	return s.srv.Close()
}

func (s *ChiWebServer) InjectMiddlewares() {
	s.engine.Use(middleware.Recoverer)
	s.engine.Use(accessControlMiddleware)
}

func (s *ChiWebServer) InjectCustomMiddlewares() {
	for _, srvMiddleware := range ServerMiddlewares {
		s.registerCustomMiddleware(srvMiddleware)
	}
}

func (s *ChiWebServer) InjectRoutes() {

	for _, route := range ServerRoutes {
		s.engine.MethodFunc(route.HttpMethod.String(), route.Path, func(w http.ResponseWriter, r *http.Request) {
			s.wg.Add(1)
			defer s.wg.Done()
			webContext := &chiWebContext{writer: w, request: r}

			route.Function(webContext)
		})

		logging.Info("Registered route [%7s] %s", route.HttpMethod, route.Path)
	}
}

func (s *ChiWebServer) ListenAndServe() error {
	serverPortStr := os.Getenv("SDKOPEN_SERVER_PORT")
	if serverPortStr == "" {
		serverPortStr = "8080"
	}

	serverPort, err := strconv.Atoi(serverPortStr)
	if err != nil {
		serverPort = 8080
	}

	s.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: s.engine,
	}
	return s.srv.ListenAndServe()
}

func (s *ChiWebServer) registerCustomMiddleware(m IMiddleware) {
	s.engine.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			webCtx := &chiWebContext{writer: w, request: r}
			if err := m.Apply(webCtx); err != nil {
				w.WriteHeader(400)
				if err := json.NewEncoder(w).Encode(err); err != nil {
					logging.Error("%s", err.Error())
				}
				return
			}
			next.ServeHTTP(w, r.WithContext(webCtx.Context()))
		})
	})
}
