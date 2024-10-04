package service

import (
	"sql-tx-abstraction/repository"
)

type ExampleService interface {
	Do() error
}

type exampleService struct {
	exampleRepo repository.ExampleRepo
	txProvider  repository.TransactionProvider
}

func NewExampleService(
	exampleRepo repository.ExampleRepo,
	txProvider repository.TransactionProvider,
) ExampleService {
	return exampleService{
		exampleRepo: exampleRepo,
		txProvider:  txProvider,
	}
}

// func (s exampleService) Do() error {
// 	fmt.Println("Without RunInTx")
// 	s.exampleRepo.Do()

// 	fmt.Println("")

// 	fmt.Println("With RunInTx")
// 	s.txProvider.RunInTx(func(txRepos repository.Repositories) error {
// 		return txRepos.ExampleRepo.Do()
// 	})

// 	return nil
// }

func (s exampleService) Do() error {
	err := s.txProvider.RunInTx(func(txRepos repository.Repositories) error {
		return txRepos.ExampleRepo.Do()
	})
	if err != nil {
		return err
	}

	return nil
}
