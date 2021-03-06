// Code generated by MockGen. DO NOT EDIT.
// Source: engine.go

// Package mocks is a generated GoMock package.
package mocks

import (
	. "github.com/drone/drone-runtime/engine"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Setup mocks base method
func (m *MockEngine) Setup(arg0 *Config) error {
	ret := m.ctrl.Call(m, "Setup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup
func (mr *MockEngineMockRecorder) Setup(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockEngine)(nil).Setup), arg0)
}

// Create mocks base method
func (m *MockEngine) Create(arg0 *Step) error {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockEngineMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockEngine)(nil).Create), arg0)
}

// Start mocks base method
func (m *MockEngine) Start(arg0 *Step) error {
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockEngineMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockEngine)(nil).Start), arg0)
}

// Wait mocks base method
func (m *MockEngine) Wait(arg0 *Step) (*State, error) {
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(*State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Wait indicates an expected call of Wait
func (mr *MockEngineMockRecorder) Wait(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockEngine)(nil).Wait), arg0)
}

// Tail mocks base method
func (m *MockEngine) Tail(arg0 *Step) (io.ReadCloser, error) {
	ret := m.ctrl.Call(m, "Tail", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tail indicates an expected call of Tail
func (mr *MockEngineMockRecorder) Tail(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*MockEngine)(nil).Tail), arg0)
}

// Upload mocks base method
func (m *MockEngine) Upload(arg0 *Step, arg1 string, arg2 io.Reader) error {
	ret := m.ctrl.Call(m, "Upload", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload
func (mr *MockEngineMockRecorder) Upload(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockEngine)(nil).Upload), arg0, arg1, arg2)
}

// Download mocks base method
func (m *MockEngine) Download(arg0 *Step, arg1 string) (io.ReadCloser, *FileInfo, error) {
	ret := m.ctrl.Call(m, "Download", arg0, arg1)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(*FileInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Download indicates an expected call of Download
func (mr *MockEngineMockRecorder) Download(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockEngine)(nil).Download), arg0, arg1)
}

// Destroy mocks base method
func (m *MockEngine) Destroy(arg0 *Config) error {
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockEngineMockRecorder) Destroy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockEngine)(nil).Destroy), arg0)
}
