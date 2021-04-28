// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sapawarga/userpost-service/repository (interfaces: PostI)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/sapawarga/userpost-service/model"
	reflect "reflect"
)

// MockPostI is a mock of PostI interface
type MockPostI struct {
	ctrl     *gomock.Controller
	recorder *MockPostIMockRecorder
}

// MockPostIMockRecorder is the mock recorder for MockPostI
type MockPostIMockRecorder struct {
	mock *MockPostI
}

// NewMockPostI creates a new mock instance
func NewMockPostI(ctrl *gomock.Controller) *MockPostI {
	mock := &MockPostI{ctrl: ctrl}
	mock.recorder = &MockPostIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostI) EXPECT() *MockPostIMockRecorder {
	return m.recorder
}

// AddLikeOnPost mocks base method
func (m *MockPostI) AddLikeOnPost(arg0 context.Context, arg1 *model.AddOrRemoveLikeOnPostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLikeOnPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLikeOnPost indicates an expected call of AddLikeOnPost
func (mr *MockPostIMockRecorder) AddLikeOnPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLikeOnPost", reflect.TypeOf((*MockPostI)(nil).AddLikeOnPost), arg0, arg1)
}

// CheckIsExistLikeOnPostBy mocks base method
func (m *MockPostI) CheckIsExistLikeOnPostBy(arg0 context.Context, arg1 *model.AddOrRemoveLikeOnPostRequest) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsExistLikeOnPostBy", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckIsExistLikeOnPostBy indicates an expected call of CheckIsExistLikeOnPostBy
func (mr *MockPostIMockRecorder) CheckIsExistLikeOnPostBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsExistLikeOnPostBy", reflect.TypeOf((*MockPostI)(nil).CheckIsExistLikeOnPostBy), arg0, arg1)
}

// GetActor mocks base method
func (m *MockPostI) GetActor(arg0 context.Context, arg1 int64) (*model.UserResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActor", arg0, arg1)
	ret0, _ := ret[0].(*model.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActor indicates an expected call of GetActor
func (mr *MockPostIMockRecorder) GetActor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActor", reflect.TypeOf((*MockPostI)(nil).GetActor), arg0, arg1)
}

// GetDetailPost mocks base method
func (m *MockPostI) GetDetailPost(arg0 context.Context, arg1 int64) (*model.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDetailPost", arg0, arg1)
	ret0, _ := ret[0].(*model.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDetailPost indicates an expected call of GetDetailPost
func (mr *MockPostIMockRecorder) GetDetailPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDetailPost", reflect.TypeOf((*MockPostI)(nil).GetDetailPost), arg0, arg1)
}

// GetIsLikedByUser mocks base method
func (m *MockPostI) GetIsLikedByUser(arg0 context.Context, arg1 *model.IsLikedByUser) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIsLikedByUser", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIsLikedByUser indicates an expected call of GetIsLikedByUser
func (mr *MockPostIMockRecorder) GetIsLikedByUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsLikedByUser", reflect.TypeOf((*MockPostI)(nil).GetIsLikedByUser), arg0, arg1)
}

// GetListPost mocks base method
func (m *MockPostI) GetListPost(arg0 context.Context, arg1 *model.UserPostRequest) ([]*model.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPost", arg0, arg1)
	ret0, _ := ret[0].([]*model.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPost indicates an expected call of GetListPost
func (mr *MockPostIMockRecorder) GetListPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPost", reflect.TypeOf((*MockPostI)(nil).GetListPost), arg0, arg1)
}

// GetListPostByMe mocks base method
func (m *MockPostI) GetListPostByMe(arg0 context.Context, arg1 *model.UserPostByMeRequest) ([]*model.PostResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListPostByMe", arg0, arg1)
	ret0, _ := ret[0].([]*model.PostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListPostByMe indicates an expected call of GetListPostByMe
func (mr *MockPostIMockRecorder) GetListPostByMe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListPostByMe", reflect.TypeOf((*MockPostI)(nil).GetListPostByMe), arg0, arg1)
}

// GetMetadataPost mocks base method
func (m *MockPostI) GetMetadataPost(arg0 context.Context, arg1 *model.UserPostRequest) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataPost", arg0, arg1)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataPost indicates an expected call of GetMetadataPost
func (mr *MockPostIMockRecorder) GetMetadataPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataPost", reflect.TypeOf((*MockPostI)(nil).GetMetadataPost), arg0, arg1)
}

// GetMetadataPostByMe mocks base method
func (m *MockPostI) GetMetadataPostByMe(arg0 context.Context, arg1 *model.UserPostByMeRequest) (*int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataPostByMe", arg0, arg1)
	ret0, _ := ret[0].(*int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataPostByMe indicates an expected call of GetMetadataPostByMe
func (mr *MockPostIMockRecorder) GetMetadataPostByMe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataPostByMe", reflect.TypeOf((*MockPostI)(nil).GetMetadataPostByMe), arg0, arg1)
}

// InsertPost mocks base method
func (m *MockPostI) InsertPost(arg0 context.Context, arg1 *model.CreateNewPostRequestRepository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPost indicates an expected call of InsertPost
func (mr *MockPostIMockRecorder) InsertPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPost", reflect.TypeOf((*MockPostI)(nil).InsertPost), arg0, arg1)
}

// RemoveLikeOnPost mocks base method
func (m *MockPostI) RemoveLikeOnPost(arg0 context.Context, arg1 *model.AddOrRemoveLikeOnPostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLikeOnPost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveLikeOnPost indicates an expected call of RemoveLikeOnPost
func (mr *MockPostIMockRecorder) RemoveLikeOnPost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLikeOnPost", reflect.TypeOf((*MockPostI)(nil).RemoveLikeOnPost), arg0, arg1)
}

// UpdateStatusOrTitle mocks base method
func (m *MockPostI) UpdateStatusOrTitle(arg0 context.Context, arg1 *model.UpdatePostRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStatusOrTitle", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStatusOrTitle indicates an expected call of UpdateStatusOrTitle
func (mr *MockPostIMockRecorder) UpdateStatusOrTitle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatusOrTitle", reflect.TypeOf((*MockPostI)(nil).UpdateStatusOrTitle), arg0, arg1)
}
