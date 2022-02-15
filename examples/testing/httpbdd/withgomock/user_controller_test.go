package withgomock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/shijuvar/gokit/examples/testing/httpbdd/controllers"
	"github.com/shijuvar/gokit/examples/testing/httpbdd/model"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {

	var (
		r         *mux.Router
		w         *httptest.ResponseRecorder
		handler   controllers.Handler
		mockCtrl  *gomock.Controller
		mockStore *MockUserStore
	)

	BeforeEach(func() {
		r = mux.NewRouter()
		mockCtrl = gomock.NewController(GinkgoT())
		mockStore = NewMockUserStore(mockCtrl)
		handler = controllers.Handler{Store: mockStore}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	// Specs for HTTP Get to "/users"
	Describe("Get list of Users", func() {
		Context("Get all Users from data store", func() {
			It("Should get list of Users", func() {
				mockUsers := getMockUsersList()
				mockStore.EXPECT().GetUsers().Return(mockUsers)
				r.HandleFunc("/users", handler.GetUsers).Methods("GET")
				req, err := http.NewRequest("GET", "/users", nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusOK))
				var users []model.User
				json.Unmarshal(w.Body.Bytes(), &users)
				// Verifying mocked data of 2 users
				Expect(len(users)).To(Equal(2))
			})
		})
	})

	// Specs for HTTP Post to "/users"
	Describe("Post a new User", func() {
		Context("Provide a valid User data", func() {
			It("Should create a new User and get HTTP Status: 201", func() {
				mockStore.EXPECT().AddUser(getMockUser()).Return(nil)
				r.HandleFunc("/users", handler.CreateUser).Methods("POST")
				userJson := `{"firstname": "Shiju", "lastname": "Varghese", "email": "shiju@xyz.com"}`

				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusCreated))
			})
		})
		Context("Provide a User data that contains duplicate email id", func() {
			It("Should get HTTP Status: 400", func() {
				mockStore.EXPECT().AddUser(getMockUser()).Return(model.ErrorEmailExists).Times(2)
				mockStore.AddUser(getMockUser())
				r.HandleFunc("/users", handler.CreateUser).Methods("POST")
				userJson := `{"firstname": "Shiju", "lastname": "Varghese", "email": "shiju@xyz.com"}`

				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

func getMockUsersList() []model.User {
	mockUsers := []model.User{
		model.User{
			FirstName: "Shiju",
			LastName:  "Varghese",
			Email:     "shiju@xyz.com",
		},
		model.User{
			FirstName: "Irene",
			LastName:  "Rose",
			Email:     "irene@xyz.com",
		},
	}
	return mockUsers
}

func getMockUser() model.User {
	return model.User{
		FirstName: "Shiju",
		LastName:  "Varghese",
		Email:     "shiju@xyz.com",
	}
}
