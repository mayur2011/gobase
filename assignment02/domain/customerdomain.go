package domain

//Customer struct
type Customer struct {
	ID, Name, Email, Country string
	MobileNum                int
}

//CustomerStore interface to specify behaviours for CRUD on Customer
type CustomerStore interface {
	Create(Customer) error
	//Update(string, Customer) error
	// Delete(string) error
	// GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}
