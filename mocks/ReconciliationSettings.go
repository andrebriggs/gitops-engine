// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	"github.com/argoproj/argo-cd/engine/resource"
	"github.com/stretchr/testify/mock"

	"github.com/argoproj/argo-cd/engine/pkg/apis/application/v1alpha1"
)

// ReconciliationSettings is an autogenerated mock type for the ReconciliationSettings type
type ReconciliationSettings struct {
	mock.Mock
}

// GetAppInstanceLabelKey provides a mock function with given fields:
func (_m *ReconciliationSettings) GetAppInstanceLabelKey() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigManagementPlugins provides a mock function with given fields:
func (_m *ReconciliationSettings) GetConfigManagementPlugins() ([]v1alpha1.ConfigManagementPlugin, error) {
	ret := _m.Called()

	var r0 []v1alpha1.ConfigManagementPlugin
	if rf, ok := ret.Get(0).(func() []v1alpha1.ConfigManagementPlugin); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]v1alpha1.ConfigManagementPlugin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKustomizeBuildOptions provides a mock function with given fields:
func (_m *ReconciliationSettings) GetKustomizeBuildOptions() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceOverrides provides a mock function with given fields:
func (_m *ReconciliationSettings) GetResourceOverrides() (map[string]v1alpha1.ResourceOverride, error) {
	ret := _m.Called()

	var r0 map[string]v1alpha1.ResourceOverride
	if rf, ok := ret.Get(0).(func() map[string]v1alpha1.ResourceOverride); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]v1alpha1.ResourceOverride)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcesFilter provides a mock function with given fields:
func (_m *ReconciliationSettings) GetResourcesFilter() (*resource.ResourcesFilter, error) {
	ret := _m.Called()

	var r0 *resource.ResourcesFilter
	if rf, ok := ret.Get(0).(func() *resource.ResourcesFilter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resource.ResourcesFilter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Subscribe provides a mock function with given fields: subCh
func (_m *ReconciliationSettings) Subscribe(subCh chan<- bool) {
	_m.Called(subCh)
}

// Unsubscribe provides a mock function with given fields: subCh
func (_m *ReconciliationSettings) Unsubscribe(subCh chan<- bool) {
	_m.Called(subCh)
}
