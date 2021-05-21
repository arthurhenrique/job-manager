package repository

import "errors"

type Repository interface {
	Select(base string) error
	Create(base string) error
	Update(base string) error
}

type repositoryImpl struct{}

var instance = &repositoryImpl{}

func Get() Repository {
	return instance
}

func (r *repositoryImpl) Select(base string) error {
	return errors.New("not implemented")
}

func (r *repositoryImpl) Create(base string) error {
	return errors.New("not implemented")
}

func (r *repositoryImpl) Update(base string) error {
	return errors.New("not implemented")
}
