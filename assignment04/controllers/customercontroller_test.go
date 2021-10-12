package controllers_test

import (
	"encoding/json"
	"gobase/assignment04/controllers"
	"gobase/assignment04/domain"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomerController", func() {
	var r *mux.Router
	var w *httptest.ResponseRecorder
	//var store *FakeCustomerStore
	var controller controllers.CustomerController
	BeforeEach(func() {
		r = mux.NewRouter()
		// controller = controllers.CustomerController{
		// 	Store: store,
		// }
	})

	Describe("Get list of Customers", func() {
		Context("Get all Customers from data store", func() {
			It("Should get list of Customers", func() {
				r.Handle("/customers", controllers.ResponseHandler(controller.GetAllCustomers)).Methods("GET")
				req, err := http.NewRequest("GET", "/customers", nil)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(200))
				var customers []domain.Customer
				json.Unmarshal(w.Body.Bytes(), &customers)
				Expect(len(customers)).To(Equal(0))
			})
		})
	})
})

type FakeCustomerStore struct {
	customerStore []domain.Customer
}

func (custStore *FakeCustomerStore) GetCustomers() []domain.Customer {
	return custStore.customerStore
}
