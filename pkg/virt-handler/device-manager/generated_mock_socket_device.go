// Automatically generated by MockGen. DO NOT EDIT!
// Source: socket_device.go

package device_manager

import (
	gomock "github.com/golang/mock/gomock"

	safepath "kubevirt.io/kubevirt/pkg/safepath"
)

// Mock of PermissionManager interface
type MockPermissionManager struct {
	ctrl     *gomock.Controller
	recorder *_MockPermissionManagerRecorder
}

// Recorder for MockPermissionManager (not exported)
type _MockPermissionManagerRecorder struct {
	mock *MockPermissionManager
}

func NewMockPermissionManager(ctrl *gomock.Controller) *MockPermissionManager {
	mock := &MockPermissionManager{ctrl: ctrl}
	mock.recorder = &_MockPermissionManagerRecorder{mock}
	return mock
}

func (_m *MockPermissionManager) EXPECT() *_MockPermissionManagerRecorder {
	return _m.recorder
}

func (_m *MockPermissionManager) ChownAtNoFollow(path *safepath.Path, uid int, gid int) error {
	ret := _m.ctrl.Call(_m, "ChownAtNoFollow", path, uid, gid)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockPermissionManagerRecorder) ChownAtNoFollow(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ChownAtNoFollow", arg0, arg1, arg2)
}
