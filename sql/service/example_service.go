package service

import (
	"fmt"
	"sql-playground/repository"
)

type ExampleService interface {
	Do() error
}

type exampleService struct {
	rm repository.RepositoryManager
}

func NewExampleService(rm repository.RepositoryManager) ExampleService {
	return exampleService{rm: rm}
}

// Use this one when running the main.go
func (s exampleService) Do() error {
	fmt.Println("Running without TX")
	s.rm.ExampleRepo().Do()

	fmt.Println("")

	fmt.Println("Running with TX")
	s.rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
		return atomicRM.ExampleRepo().Do()
	})

	return nil
}

// // Use this one when running tests
// func (s exampleService) Do() error {
// 	err := s.rm.RunAtomic(func(atomicRM repository.RepositoryManager) error {
// 		return atomicRM.ExampleRepo().Do()
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
