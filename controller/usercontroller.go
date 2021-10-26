package controller

import (
	"encoding/json"
	"go-demo/entity"
	"go-demo/service"
	"net/http"
)

type UserController interface {
	AddUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
}

type userController struct{}

var userService service.UserService

func NewUserController(service service.UserService) UserController {
	userService = service
	return &userController{}
}

func (*userController) AddUser(w http.ResponseWriter, r *http.Request) {
	user := entity.User{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error encoding user Request"))
		return
	}
	err = userService.AddUser(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (*userController) GetUser(w http.ResponseWriter, r *http.Request) {
	users, err := userService.GetUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
