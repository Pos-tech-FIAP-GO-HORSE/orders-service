// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// IOrderService is an autogenerated mock type for the IOrderService type
type IOrderService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, order
func (_m *IOrderService) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
	ret := _m.Called(ctx, order)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Order) (*entity.Order, error)); ok {
		return rf(ctx, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Order) *entity.Order); ok {
		r0 = rf(ctx, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Order) error); ok {
		r1 = rf(ctx, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Find provides a mock function with given fields: ctx
func (_m *IOrderService) Find(ctx context.Context) ([]*entity.Order, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Find")
	}

	var r0 []*entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*entity.Order, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*entity.Order); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Order)
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
func (_m *IOrderService) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Order, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Order); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ctx, id, order
func (_m *IOrderService) UpdateByID(ctx context.Context, id string, order entity.Order) (*entity.Order, error) {
	ret := _m.Called(ctx, id, order)

	if len(ret) == 0 {
		panic("no return value specified for UpdateByID")
	}

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Order) (*entity.Order, error)); ok {
		return rf(ctx, id, order)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, entity.Order) *entity.Order); ok {
		r0 = rf(ctx, id, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, entity.Order) error); ok {
		r1 = rf(ctx, id, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusByID provides a mock function with given fields: ctx, id, status
func (_m *IOrderService) UpdateStatusByID(ctx context.Context, id string, status string) (*entity.Order, error) {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateStatusByID")
	}

	var r0 *entity.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*entity.Order, error)); ok {
		return rf(ctx, id, status)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entity.Order); ok {
		r0 = rf(ctx, id, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIOrderService creates a new instance of IOrderService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIOrderService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IOrderService {
	mock := &IOrderService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
