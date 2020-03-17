// Code generated by mockery v1.0.0. DO NOT EDIT.

package tg

import mock "github.com/stretchr/testify/mock"

// MockNameTagGenerator is an autogenerated mock type for the NameTagGenerator type
type MockNameTagGenerator struct {
	mock.Mock
}

// NameTG provides a mock function with given fields: namespace, ingressName, serviceName, servicePort, targetType, protocol
func (_m *MockNameTagGenerator) NameTG(namespace string, ingressName string, serviceName string, servicePort string, targetType string, protocol string) string {
	ret := _m.Called(namespace, ingressName, serviceName, servicePort, targetType, protocol)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string, string, string, string, string) string); ok {
		r0 = rf(namespace, ingressName, serviceName, servicePort, targetType, protocol)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// TagTG provides a mock function with given fields: namespace, ingressName, serviceName, servicePort
func (_m *MockNameTagGenerator) TagTG(namespace string, ingressName string, serviceName string, servicePort string) map[string]string {
	ret := _m.Called(namespace, ingressName, serviceName, servicePort)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string, string, string) map[string]string); ok {
		r0 = rf(namespace, ingressName, serviceName, servicePort)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}

// TagTGGroup provides a mock function with given fields: namespace, ingressName
func (_m *MockNameTagGenerator) TagTGGroup(namespace string, ingressName string) map[string]string {
	ret := _m.Called(namespace, ingressName)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string, string) map[string]string); ok {
		r0 = rf(namespace, ingressName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}
