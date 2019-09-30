// Code generated by MockGen. DO NOT EDIT.
// Source: ./committee/committee.go

// Package mock_committee is a generated GoMock package.
package mock_committee

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	committee "github.com/iotexproject/iotex-election/committee"
	types "github.com/iotexproject/iotex-election/types"
	reflect "reflect"
	time "time"
)

// MockCommittee is a mock of Committee interface
type MockCommittee struct {
	ctrl     *gomock.Controller
	recorder *MockCommitteeMockRecorder
}

// MockCommitteeMockRecorder is the mock recorder for MockCommittee
type MockCommitteeMockRecorder struct {
	mock *MockCommittee
}

// NewMockCommittee creates a new mock instance
func NewMockCommittee(ctrl *gomock.Controller) *MockCommittee {
	mock := &MockCommittee{ctrl: ctrl}
	mock.recorder = &MockCommitteeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommittee) EXPECT() *MockCommitteeMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockCommittee) Start(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockCommitteeMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockCommittee)(nil).Start), arg0)
}

// Stop mocks base method
func (m *MockCommittee) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockCommitteeMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockCommittee)(nil).Stop), arg0)
}

// ResultByHeight mocks base method
func (m *MockCommittee) ResultByHeight(arg0 uint64) (*types.ElectionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResultByHeight", arg0)
	ret0, _ := ret[0].(*types.ElectionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResultByHeight indicates an expected call of ResultByHeight
func (mr *MockCommitteeMockRecorder) ResultByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResultByHeight", reflect.TypeOf((*MockCommittee)(nil).ResultByHeight), arg0)
}

// RawDataByHeight mocks base method
func (m *MockCommittee) RawDataByHeight(arg0 uint64) ([]*types.Bucket, []*types.Registration, time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawDataByHeight", arg0)
	ret0, _ := ret[0].([]*types.Bucket)
	ret1, _ := ret[1].([]*types.Registration)
	ret2, _ := ret[2].(time.Time)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// RawDataByHeight indicates an expected call of RawDataByHeight
func (mr *MockCommitteeMockRecorder) RawDataByHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawDataByHeight", reflect.TypeOf((*MockCommittee)(nil).RawDataByHeight), arg0)
}

// HeightByTime mocks base method
func (m *MockCommittee) HeightByTime(arg0 time.Time) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeightByTime", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeightByTime indicates an expected call of HeightByTime
func (mr *MockCommitteeMockRecorder) HeightByTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeightByTime", reflect.TypeOf((*MockCommittee)(nil).HeightByTime), arg0)
}

// LatestHeight mocks base method
func (m *MockCommittee) LatestHeight() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestHeight")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// LatestHeight indicates an expected call of LatestHeight
func (mr *MockCommitteeMockRecorder) LatestHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestHeight", reflect.TypeOf((*MockCommittee)(nil).LatestHeight))
}

// Status mocks base method
func (m *MockCommittee) Status() committee.STATUS {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(committee.STATUS)
	return ret0
}

// Status indicates an expected call of Status
func (mr *MockCommitteeMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCommittee)(nil).Status))
}

// PutNativePollByEpoch mocks base method
func (m *MockCommittee) PutNativePollByEpoch(arg0 uint64, arg1 time.Time, arg2 []*types.Bucket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutNativePollByEpoch", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutNativePollByEpoch indicates an expected call of PutNativePollByEpoch
func (mr *MockCommitteeMockRecorder) PutNativePollByEpoch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutNativePollByEpoch", reflect.TypeOf((*MockCommittee)(nil).PutNativePollByEpoch), arg0, arg1, arg2)
}

// NativeBucketsByEpoch mocks base method
func (m *MockCommittee) NativeBucketsByEpoch(arg0 uint64) ([]*types.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NativeBucketsByEpoch", arg0)
	ret0, _ := ret[0].([]*types.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NativeBucketsByEpoch indicates an expected call of NativeBucketsByEpoch
func (mr *MockCommitteeMockRecorder) NativeBucketsByEpoch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NativeBucketsByEpoch", reflect.TypeOf((*MockCommittee)(nil).NativeBucketsByEpoch), arg0)
}
