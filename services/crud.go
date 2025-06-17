// services/crud.go
package services

type CRUDService[T any] interface {
	List() []T
	Add(T) T
	Update(id int, t T) error
	Delete(id int) error
}
