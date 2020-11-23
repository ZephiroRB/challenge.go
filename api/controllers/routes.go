package controllers

import "github.com/ZephiroRB/challenge.go/api/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	s.Router.HandleFunc("/api/user/register", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

	s.Router.HandleFunc("/api/user/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	s.Router.HandleFunc("/api/skills/register", middlewares.SetMiddlewareJSON(s.CreateSkill)).Methods("POST")

	s.Router.HandleFunc("/api/user/skills", middlewares.SetMiddlewareJSON(s.CreateUserSkill)).Methods("POST")

}
