package controllers

import (
	"strconv"
	"time"

	"portfolio/app/domain"
	"portfolio/app/interfaces/database"
	"portfolio/app/usecase"
)

type UsersController struct {
	Interactor usecase.UserInteractor
}

type StoreJsonRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	ImageUrl string `json:"image_url"`
}

func NewUsersController(db database.DB) *UsersController {
	return &UsersController{
		Interactor: usecase.UserInteractor{
			DB:   &database.DBRepository{DB: db},
			User: &database.UserRepository{},
		},
	}
}

func (controller *UsersController) Index(c Context) {

	users, res := controller.Interactor.GetAll()
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", users))
}

func (controller *UsersController) Show(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UsersController) Store(c Context) {

	var json StoreJsonRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, NewH(err.Error(), nil))
		return
	}

	user := domain.Users{}
	user.Username = json.Username
	user.Password = json.Password
	user.Email = json.Email
	birthday, err := time.Parse("2006-01-02", json.Birthday)

	if err != nil {
		c.JSON(422, err)
		return
	}

	user.Birthday = birthday
	user.ImageUrl = json.ImageUrl
	createdUser, res := controller.Interactor.Create(user)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", createdUser))
}
