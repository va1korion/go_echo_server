package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.ConfigureLogger(); err != nil {
		return err
	}

	s.logger.Infof("Starting up...")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *APIServer) ConfigureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)

	if err != nil {
		return err
	}

	s.logger.SetLevel(level)
	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.router.HandleFunc("/hello", s.HandleHello())
}

func (s *APIServer) HandleHello() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if _, err := io.WriteString(writer, "Hello there!"); err != nil {
			s.logger.Warnf("Something wrong with hello handler")
		}
	}
}
