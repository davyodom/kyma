// Code generated by mockery v1.0.0
package mocks

import apperrors "github.com/kyma-project/kyma/components/application-connector/internal/apperrors"
import mock "github.com/stretchr/testify/mock"
import remoteenv "github.com/kyma-project/kyma/components/application-connector/internal/metadata/remoteenv"

// ServiceRepository is an autogenerated mock type for the ServiceRepository type
type ServiceRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: remoteEnvironment, service
func (_m *ServiceRepository) Create(remoteEnvironment string, service remoteenv.Service) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, service)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, remoteenv.Service) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceRepository) Delete(remoteEnvironment string, id string) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, id)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, string) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// Get provides a mock function with given fields: remoteEnvironment, id
func (_m *ServiceRepository) Get(remoteEnvironment string, id string) (remoteenv.Service, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment, id)

	var r0 remoteenv.Service
	if rf, ok := ret.Get(0).(func(string, string) remoteenv.Service); ok {
		r0 = rf(remoteEnvironment, id)
	} else {
		r0 = ret.Get(0).(remoteenv.Service)
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string, string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: remoteEnvironment
func (_m *ServiceRepository) GetAll(remoteEnvironment string) ([]remoteenv.Service, apperrors.AppError) {
	ret := _m.Called(remoteEnvironment)

	var r0 []remoteenv.Service
	if rf, ok := ret.Get(0).(func(string) []remoteenv.Service); ok {
		r0 = rf(remoteEnvironment)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]remoteenv.Service)
		}
	}

	var r1 apperrors.AppError
	if rf, ok := ret.Get(1).(func(string) apperrors.AppError); ok {
		r1 = rf(remoteEnvironment)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(apperrors.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: remoteEnvironment, service
func (_m *ServiceRepository) Update(remoteEnvironment string, service remoteenv.Service) apperrors.AppError {
	ret := _m.Called(remoteEnvironment, service)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, remoteenv.Service) apperrors.AppError); ok {
		r0 = rf(remoteEnvironment, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}