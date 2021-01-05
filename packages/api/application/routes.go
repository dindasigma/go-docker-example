package application

import (
	_ "github.com/dindasigma/go-docker-boilerplate/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/dindasigma/go-docker-boilerplate/packages/api/controllers"
	"github.com/dindasigma/go-docker-boilerplate/packages/api/middlewares"
	httpSwagger "github.com/swaggo/http-swagger"
)

func initializeRoutes() {
	// Home Route
	router.HandleFunc("/", middlewares.SetMiddlewareJSON(controllers.HomeController.Home)).Methods("GET")

	// Login Route
	router.HandleFunc("/login", middlewares.SetMiddlewareJSON(controllers.LoginController.Login)).Methods("POST")

	// Users routes
	/*s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")*/

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
