package service

import "fmt"

type Storage interface {
	Get(id int) string
}

type service struct {
	storage Storage
}

func New(storage Storage) *service {
	return &service{storage: storage}
}

func (s *service) Mul2(a int) int {
	fmt.Println(s.storage.Get(a))
	return a * 2
}
