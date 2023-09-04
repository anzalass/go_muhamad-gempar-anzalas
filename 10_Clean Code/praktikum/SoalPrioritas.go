package main

type User struct {
	Id       int
	Username string
	Password int
}

type UserService struct {
	t []User
}

func (user UserService) getallusers() []User {

	return user.t

}

func (user UserService) getuserbyid(id int) User {

	for _, r := range user.t {

		if id == r.Id {

			return r

		}

	}

	return User{}

}
