package controllers_test

import (
	"encoding/json"
	"gobase/assignment04/controllers"
	"gobase/assignment04/domain"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomerController", func() {
	var r *mux.Router
	var w *httptest.ResponseRecorder
	var store *FakeCustomerStore
	var controller controllers.CustomerController
	BeforeEach(func() {
		r = mux.NewRouter()
		store = newFakeCustomerStore()
		controller = controllers.CustomerController{
			Store: store,
		}
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

	Describe("Post a customer which does not exist", func() {
		Context("Provide a valid Customer data", func() {
			It("Should create a new Customer and get HTTP Status: 201", func() {
				r.Handle("/customers", controllers.ResponseHandler(controller.PostCustomer)).Methods("POST")
				customerJson := `{"ID":"CUST-301", "Email":"customer301@gmail.com", "Name":"Raj Sinha"}`
				req, err := http.NewRequest(
					"POST",
					"/customers",
					strings.NewReader(customerJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(201))
			})
		})

		Context("Provide a customer data that contains dupliate customer ID", func() {
			It("Should get HTTP Status: 400", func() {
				r.Handle("/customers", controllers.ResponseHandler(controller.PostCustomer)).Methods("POST")
				customerJson := `{"ID":"CUST-201", "Email":"customer201@gmail.com", "Name":"Rohan Singh"}`
				req, err := http.NewRequest(
					"POST",
					"/customers",
					strings.NewReader(customerJson),
				)
				Expect(err).NotTo(HaveOccurred())
				w = httptest.NewRecorder()
				r.ServeHTTP(w, req)
				Expect(w.Code).To(Equal(400))
			})
		})
	})

})

type FakeCustomerStore struct {
	customerStore []domain.Customer
}

func (custStore *FakeCustomerStore) GetCustomerById(Id string) (domain.Customer, error) {
	return domain.Customer{}, nil
}

func (custStore *FakeCustomerStore) GetAllCustomers() ([]domain.Customer, error) {
	return custStore.customerStore, nil
}

func (custStore *FakeCustomerStore) Create(customer domain.Customer) error {
	for _, u := range custStore.customerStore {
		if u.ID == customer.ID {
			return domain.ErrorIDExists
		}
	}
	custStore.customerStore = append(custStore.customerStore, customer)
	return nil
}

func (custStore *FakeCustomerStore) Delete(Id string) error {

	return nil
}

func (custStore *FakeCustomerStore) Update(Id string, customer domain.Customer) error {

	return nil
}

func newFakeCustomerStore() *FakeCustomerStore {
	store := &FakeCustomerStore{}
	store.Create(domain.Customer{
		ID:    "CUST-101",
		Email: "customer101@gmail.com",
		Name:  "Rahul Sharma",
	})
	store.Create(domain.Customer{
		ID:    "CUST-201",
		Email: "customer201@gmail.com",
		Name:  "Shanti Prashant",
	})
	return store
}
