package main

import "fmt"

type ExampleService interface {
	ExecExample() error
}

type exampleService struct {
	rm RepositoryManager
}

func NewExampleService(rm RepositoryManager) ExampleService {
	return &exampleService{rm: rm}
}

func (s *exampleService) ExecExample() error {
	fmt.Println("Running without TX")
	s.rm.ExampleRepo().ExecExample()

	fmt.Println("Running with TX")
	s.rm.RunAtomic(func(atomicRepositoryManager RepositoryManager) error {
		return atomicRepositoryManager.ExampleRepo().ExecExample()
	})

	return nil
}
