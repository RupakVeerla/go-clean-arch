package main

import (
	"go-demo/controller"
	"go-demo/repository"
	"go-demo/router"
	"go-demo/service"
)

var (
	repo           repository.Repository     = repository.NewSqlRepo()
	userService    service.UserService       = service.NewUserService(repo)
	userController controller.UserController = controller.NewUserController(userService)
	muxrouter      router.Router             = router.NewMuxRouter()
)

const port string = ":4000"

func main() {
	muxrouter.Post("/user", userController.AddUser)
	muxrouter.Get("/user", userController.GetUser)

	muxrouter.Serve(port)
}
