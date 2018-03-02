package controller_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/shijuvar/gokit/examples/testing/gomock/controller"
	"github.com/shijuvar/gokit/examples/testing/gomock/mocks"
)

var _ = Describe("Controller", func() {

	var (
		mockCtrl       *gomock.Controller
		mockStore      *mocks.MockUserStore
		userController *controller.UserController
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockStore = mocks.NewMockUserStore(mockCtrl)
		userController = &controller.UserController{Store: mockStore}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("Add User", func() {
		mockStore.EXPECT().AddUser("shijuvar").Return(nil).Times(1)
		err := userController.Create("shijuvar")
		Ω(err).Should(BeNil())

	})
})
