package domain

//Customer struct
type Customer struct {
	ID, Name, Email, Country string
	MobileNum                int
}

//CustomerStore interface to specify behaviours for CRUD on Customer
type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll() ([]Customer, error)
}

/*
Questions:
1. if we comment one of methods declaration from interface or dont have interface then also overall program would work?
Ans: mapstore implemented methods won't give any error but CustomerController will not work if its binded with Interface

2.
*/
