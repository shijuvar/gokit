package controller

import "github.com/shijuvar/gokit/examples/testing/gomock/store"

type UserController struct {
	Store store.UserStore
}

func (c *UserController) Create(name string) error {
	return c.Store.AddUser(name)
}
