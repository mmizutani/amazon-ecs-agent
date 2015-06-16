// Copyright 2015 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/amazon-ecs-agent/agent/wsclient (interfaces: ClientServer)

package mock_wsclient

import (
	wsclient "github.com/aws/amazon-ecs-agent/agent/wsclient"
	gomock "github.com/golang/mock/gomock"
)

// Mock of ClientServer interface
type MockClientServer struct {
	ctrl     *gomock.Controller
	recorder *_MockClientServerRecorder
}

// Recorder for MockClientServer (not exported)
type _MockClientServerRecorder struct {
	mock *MockClientServer
}

func NewMockClientServer(ctrl *gomock.Controller) *MockClientServer {
	mock := &MockClientServer{ctrl: ctrl}
	mock.recorder = &_MockClientServerRecorder{mock}
	return mock
}

func (_m *MockClientServer) EXPECT() *_MockClientServerRecorder {
	return _m.recorder
}

func (_m *MockClientServer) AddRequestHandler(_param0 wsclient.RequestHandler) {
	_m.ctrl.Call(_m, "AddRequestHandler", _param0)
}

func (_mr *_MockClientServerRecorder) AddRequestHandler(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddRequestHandler", arg0)
}

func (_m *MockClientServer) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockClientServer) Connect() error {
	ret := _m.ctrl.Call(_m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Connect() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Connect")
}

func (_m *MockClientServer) MakeRequest(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "MakeRequest", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) MakeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MakeRequest", arg0)
}

func (_m *MockClientServer) Serve() error {
	ret := _m.ctrl.Call(_m, "Serve")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockClientServerRecorder) Serve() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Serve")
}
