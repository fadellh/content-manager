package mocks

import (
	"content/business/blog"

	mock "github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m Repository) FindContentById(id int) (*blog.Blog, error) {
	ret := _m.Called(id)

	var r0 *blog.Blog
	if rf, ok := ret.Get(0).(func(id int) *blog.Blog); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*blog.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(id int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
