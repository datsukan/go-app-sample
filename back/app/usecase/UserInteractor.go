package usecase

import (
	"portfolio/app/domain"
)

type UserInteractor struct {
	DB   DBRepository
	User UserRepository
}

func (interactor *UserInteractor) GetAll() (users []domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Usersの取得
	foundUsers, _ := interactor.User.Find(db)
	for _, v := range foundUsers {
		users = append(users, v.BuildForGet())
	}
	return users, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Get(id int) (user domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Userの取得
	foundUser, err := interactor.User.FindByID(db, id)
	if err != nil {
		return domain.UsersForGet{}, NewResultStatus(404, err)
	}
	user = foundUser.BuildForGet()
	return user, NewResultStatus(200, nil)
}

func (interactor *UserInteractor) Create(user domain.Users) (createdUser domain.UsersForGet, resultStatus *ResultStatus) {
	db := interactor.DB.Connect()
	// Userの登録
	err := interactor.User.Create(db, &user)
	if err != nil {
		return domain.UsersForGet{}, NewResultStatus(500, err)
	}
	createdUser = user.BuildForGet()
	return createdUser, NewResultStatus(200, nil)
}
