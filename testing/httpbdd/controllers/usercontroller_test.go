package controllers_test

import (
	"encoding/json"
	"gobase/testing/httpbdd/controllers"
	"gobase/testing/httpbdd/domain"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserController", func() {
	var r *mux.Router
	var w *httptest.ResponseRecorder
	var store *FakeUserStore
	var controller controllers.UserController
	BeforeEach(func() {
		r = mux.NewRouter()
		store = newFakeUserStore()
		controller = controllers.UserController{
			//Injecting a test stub
			Store: store,
		}
	})

	//Specs for HTTP Get to "/users"
	Describe("Get list of Users", func() {
		Context("Get all Users from data store", func() {
			It("Should get list of Users", func() {
				r.HandleFunc("/users", controller.GetUsers).Methods("GET")
				req, err := http.NewRequest("GET", "/users", nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(200))
				var users []domain.User
				json.Unmarshal(w.Body.Bytes(), &users)
				//Verifying mocked data of 2 users
				Expect(len(users)).To(Equal(2))
			})
		})
	})

	// Specs for HTTP Post to "/users"
	Describe("Post a new User", func() {
		Context("Provide a valid User data", func() {
			It("Should create a new User and get HTTP Status: 201", func() {
				r.HandleFunc("/users", controller.CreateUser).Methods("POST")
				userJson := `{"firstName" : "Jonn", "lastName": "Pod", "email": "jonn_pod@gmail.com"}`
				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(201))
			})
		})

		Context("Provide a User data that contains duplicate email id", func() {
			It("Should get HTTP Status: 400", func() {
				r.HandleFunc("/users", controller.CreateUser).Methods("POST")
				userJson := `{"firstName" : "Pruthvi", "lastName": "Singh", "email": "ps@gmail.com"}`
				req, err := http.NewRequest(
					"POST",
					"/users",
					strings.NewReader(userJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(400))
			})
		})
	})
})

//FakeUserStore provides a mocked implementation of interface domain.UserStore
type FakeUserStore struct {
	userStore []domain.User
}

func (store *FakeUserStore) GetUsers() []domain.User {
	return store.userStore
}

func (store *FakeUserStore) AddUser(user domain.User) error {
	//Check whether email exists
	for _, u := range store.userStore {
		if u.Email == user.Email {
			return domain.ErrorEmailsExists
		}
	}
	store.userStore = append(store.userStore, user)
	return nil
}

func newFakeUserStore() *FakeUserStore {
	store := &FakeUserStore{}
	store.AddUser(domain.User{
		FirstName: "Pruthvi",
		LastName:  "Singh",
		Email:     "ps@gmail.com",
	})
	store.AddUser(domain.User{
		FirstName: "Praful",
		LastName:  "Sharma",
		Email:     "ps@yahoo.com",
	})
	return store
}
