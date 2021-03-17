// Package view all view related structs and functions
// Generated code - do not modify it will be overwritten!!
// Time: 17.03.2021 09:33:01.755
package view

import (
	. "/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	*echo.Echo
	env *Env
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewServer(env *Env) *echo.Echo {
	s := new(Server)
	s.env = env
	s.Echo = echo.New()

	// assets will be loaded from /static directory as /assets/*
	s.Static("/static", "assets")

	// Instantiate a template registry with an array of template set
	// Ref: https://gist.github.com/rand99/808e6e9702c00ce64803d94abff65678
	renderer := &Template{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = renderer

	// Middleware
	s.Use(middleware.Logger())
	s.Use(middleware.Recover())

	// Routes
	s.GET("/", s.DashboardHandler) // Opens Dashboard

	// routes for Project
	s.GET("/projects", s.ListProjects)
	s.GET("/projects/:id", s.GetProject)
	s.GET("/projects/new", s.NewProject)
	s.POST("/projects", s.CreateProject)
	s.POST("/projects/:id", s.UpdateProject)
	s.POST("/projects/:id/delete", s.DeleteProject)

	return s
}

/*
type Handler interface {
	Get(c echo.Context) error
	List(c echo.Context) error
	New(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	List(c echo.Context) error
}
*/
