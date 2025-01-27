// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// IProductService is an autogenerated mock type for the IProductService type
type IProductService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, product
func (_m *IProductService) Create(ctx context.Context, product entity.Product) (*entity.Product, error) {
	ret := _m.Called(ctx, product)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Product) (*entity.Product, error)); ok {
		return rf(ctx, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Product) *entity.Product); ok {
		r0 = rf(ctx, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Product) error); ok {
		r1 = rf(ctx, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx
func (_m *IProductService) Find(ctx context.Context) ([]*entity.Product, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []*entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Product, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Product); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *IProductService) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Product, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Product); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, id, product
func (_m *IProductService) UpdateByID(ctx context.Context, id string, product entity.Product) (*entity.Product, error) {
	ret := _m.Called(ctx, id, product)

	if len(ret) == 0 {
		panic("no return value specified for UpdateByID")
	}

	var r0 *entity.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Product) (*entity.Product, error)); ok {
		return rf(ctx, id, product)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Product) *entity.Product); ok {
		r0 = rf(ctx, id, product)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.Product) error); ok {
		r1 = rf(ctx, id, product)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIProductService creates a new instance of IProductService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIProductService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IProductService {
	mock := &IProductService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
