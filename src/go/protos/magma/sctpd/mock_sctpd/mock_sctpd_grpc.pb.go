// Code generated by MockGen. DO NOT EDIT.
// Source: magma/sctpd/sctpd_grpc.pb.go

// Package mock_sctpd is a generated GoMock package.
package mock_sctpd

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sctpd "github.com/magma/magma/src/go/protos/magma/sctpd"
	grpc "google.golang.org/grpc"
)

// MockSctpdDownlinkClient is a mock of SctpdDownlinkClient interface.
type MockSctpdDownlinkClient struct {
	ctrl     *gomock.Controller
	recorder *MockSctpdDownlinkClientMockRecorder
}

// MockSctpdDownlinkClientMockRecorder is the mock recorder for MockSctpdDownlinkClient.
type MockSctpdDownlinkClientMockRecorder struct {
	mock *MockSctpdDownlinkClient
}

// NewMockSctpdDownlinkClient creates a new mock instance.
func NewMockSctpdDownlinkClient(ctrl *gomock.Controller) *MockSctpdDownlinkClient {
	mock := &MockSctpdDownlinkClient{ctrl: ctrl}
	mock.recorder = &MockSctpdDownlinkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSctpdDownlinkClient) EXPECT() *MockSctpdDownlinkClientMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockSctpdDownlinkClient) Init(ctx context.Context, in *sctpd.InitReq, opts ...grpc.CallOption) (*sctpd.InitRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Init", varargs...)
	ret0, _ := ret[0].(*sctpd.InitRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init.
func (mr *MockSctpdDownlinkClientMockRecorder) Init(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSctpdDownlinkClient)(nil).Init), varargs...)
}

// SendDl mocks base method.
func (m *MockSctpdDownlinkClient) SendDl(ctx context.Context, in *sctpd.SendDlReq, opts ...grpc.CallOption) (*sctpd.SendDlRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendDl", varargs...)
	ret0, _ := ret[0].(*sctpd.SendDlRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDl indicates an expected call of SendDl.
func (mr *MockSctpdDownlinkClientMockRecorder) SendDl(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDl", reflect.TypeOf((*MockSctpdDownlinkClient)(nil).SendDl), varargs...)
}

// MockSctpdDownlinkServer is a mock of SctpdDownlinkServer interface.
type MockSctpdDownlinkServer struct {
	ctrl     *gomock.Controller
	recorder *MockSctpdDownlinkServerMockRecorder
}

// MockSctpdDownlinkServerMockRecorder is the mock recorder for MockSctpdDownlinkServer.
type MockSctpdDownlinkServerMockRecorder struct {
	mock *MockSctpdDownlinkServer
}

// NewMockSctpdDownlinkServer creates a new mock instance.
func NewMockSctpdDownlinkServer(ctrl *gomock.Controller) *MockSctpdDownlinkServer {
	mock := &MockSctpdDownlinkServer{ctrl: ctrl}
	mock.recorder = &MockSctpdDownlinkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSctpdDownlinkServer) EXPECT() *MockSctpdDownlinkServerMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockSctpdDownlinkServer) Init(arg0 context.Context, arg1 *sctpd.InitReq) (*sctpd.InitRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0, arg1)
	ret0, _ := ret[0].(*sctpd.InitRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Init indicates an expected call of Init.
func (mr *MockSctpdDownlinkServerMockRecorder) Init(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockSctpdDownlinkServer)(nil).Init), arg0, arg1)
}

// SendDl mocks base method.
func (m *MockSctpdDownlinkServer) SendDl(arg0 context.Context, arg1 *sctpd.SendDlReq) (*sctpd.SendDlRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendDl", arg0, arg1)
	ret0, _ := ret[0].(*sctpd.SendDlRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendDl indicates an expected call of SendDl.
func (mr *MockSctpdDownlinkServerMockRecorder) SendDl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendDl", reflect.TypeOf((*MockSctpdDownlinkServer)(nil).SendDl), arg0, arg1)
}

// mustEmbedUnimplementedSctpdDownlinkServer mocks base method.
func (m *MockSctpdDownlinkServer) mustEmbedUnimplementedSctpdDownlinkServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSctpdDownlinkServer")
}

// mustEmbedUnimplementedSctpdDownlinkServer indicates an expected call of mustEmbedUnimplementedSctpdDownlinkServer.
func (mr *MockSctpdDownlinkServerMockRecorder) mustEmbedUnimplementedSctpdDownlinkServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSctpdDownlinkServer", reflect.TypeOf((*MockSctpdDownlinkServer)(nil).mustEmbedUnimplementedSctpdDownlinkServer))
}

// MockUnsafeSctpdDownlinkServer is a mock of UnsafeSctpdDownlinkServer interface.
type MockUnsafeSctpdDownlinkServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSctpdDownlinkServerMockRecorder
}

// MockUnsafeSctpdDownlinkServerMockRecorder is the mock recorder for MockUnsafeSctpdDownlinkServer.
type MockUnsafeSctpdDownlinkServerMockRecorder struct {
	mock *MockUnsafeSctpdDownlinkServer
}

// NewMockUnsafeSctpdDownlinkServer creates a new mock instance.
func NewMockUnsafeSctpdDownlinkServer(ctrl *gomock.Controller) *MockUnsafeSctpdDownlinkServer {
	mock := &MockUnsafeSctpdDownlinkServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSctpdDownlinkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSctpdDownlinkServer) EXPECT() *MockUnsafeSctpdDownlinkServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSctpdDownlinkServer mocks base method.
func (m *MockUnsafeSctpdDownlinkServer) mustEmbedUnimplementedSctpdDownlinkServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSctpdDownlinkServer")
}

// mustEmbedUnimplementedSctpdDownlinkServer indicates an expected call of mustEmbedUnimplementedSctpdDownlinkServer.
func (mr *MockUnsafeSctpdDownlinkServerMockRecorder) mustEmbedUnimplementedSctpdDownlinkServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSctpdDownlinkServer", reflect.TypeOf((*MockUnsafeSctpdDownlinkServer)(nil).mustEmbedUnimplementedSctpdDownlinkServer))
}

// MockSctpdUplinkClient is a mock of SctpdUplinkClient interface.
type MockSctpdUplinkClient struct {
	ctrl     *gomock.Controller
	recorder *MockSctpdUplinkClientMockRecorder
}

// MockSctpdUplinkClientMockRecorder is the mock recorder for MockSctpdUplinkClient.
type MockSctpdUplinkClientMockRecorder struct {
	mock *MockSctpdUplinkClient
}

// NewMockSctpdUplinkClient creates a new mock instance.
func NewMockSctpdUplinkClient(ctrl *gomock.Controller) *MockSctpdUplinkClient {
	mock := &MockSctpdUplinkClient{ctrl: ctrl}
	mock.recorder = &MockSctpdUplinkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSctpdUplinkClient) EXPECT() *MockSctpdUplinkClientMockRecorder {
	return m.recorder
}

// CloseAssoc mocks base method.
func (m *MockSctpdUplinkClient) CloseAssoc(ctx context.Context, in *sctpd.CloseAssocReq, opts ...grpc.CallOption) (*sctpd.CloseAssocRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloseAssoc", varargs...)
	ret0, _ := ret[0].(*sctpd.CloseAssocRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAssoc indicates an expected call of CloseAssoc.
func (mr *MockSctpdUplinkClientMockRecorder) CloseAssoc(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAssoc", reflect.TypeOf((*MockSctpdUplinkClient)(nil).CloseAssoc), varargs...)
}

// NewAssoc mocks base method.
func (m *MockSctpdUplinkClient) NewAssoc(ctx context.Context, in *sctpd.NewAssocReq, opts ...grpc.CallOption) (*sctpd.NewAssocRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewAssoc", varargs...)
	ret0, _ := ret[0].(*sctpd.NewAssocRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAssoc indicates an expected call of NewAssoc.
func (mr *MockSctpdUplinkClientMockRecorder) NewAssoc(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAssoc", reflect.TypeOf((*MockSctpdUplinkClient)(nil).NewAssoc), varargs...)
}

// SendUl mocks base method.
func (m *MockSctpdUplinkClient) SendUl(ctx context.Context, in *sctpd.SendUlReq, opts ...grpc.CallOption) (*sctpd.SendUlRes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendUl", varargs...)
	ret0, _ := ret[0].(*sctpd.SendUlRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendUl indicates an expected call of SendUl.
func (mr *MockSctpdUplinkClientMockRecorder) SendUl(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUl", reflect.TypeOf((*MockSctpdUplinkClient)(nil).SendUl), varargs...)
}

// MockSctpdUplinkServer is a mock of SctpdUplinkServer interface.
type MockSctpdUplinkServer struct {
	ctrl     *gomock.Controller
	recorder *MockSctpdUplinkServerMockRecorder
}

// MockSctpdUplinkServerMockRecorder is the mock recorder for MockSctpdUplinkServer.
type MockSctpdUplinkServerMockRecorder struct {
	mock *MockSctpdUplinkServer
}

// NewMockSctpdUplinkServer creates a new mock instance.
func NewMockSctpdUplinkServer(ctrl *gomock.Controller) *MockSctpdUplinkServer {
	mock := &MockSctpdUplinkServer{ctrl: ctrl}
	mock.recorder = &MockSctpdUplinkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSctpdUplinkServer) EXPECT() *MockSctpdUplinkServerMockRecorder {
	return m.recorder
}

// CloseAssoc mocks base method.
func (m *MockSctpdUplinkServer) CloseAssoc(arg0 context.Context, arg1 *sctpd.CloseAssocReq) (*sctpd.CloseAssocRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseAssoc", arg0, arg1)
	ret0, _ := ret[0].(*sctpd.CloseAssocRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CloseAssoc indicates an expected call of CloseAssoc.
func (mr *MockSctpdUplinkServerMockRecorder) CloseAssoc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseAssoc", reflect.TypeOf((*MockSctpdUplinkServer)(nil).CloseAssoc), arg0, arg1)
}

// NewAssoc mocks base method.
func (m *MockSctpdUplinkServer) NewAssoc(arg0 context.Context, arg1 *sctpd.NewAssocReq) (*sctpd.NewAssocRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAssoc", arg0, arg1)
	ret0, _ := ret[0].(*sctpd.NewAssocRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewAssoc indicates an expected call of NewAssoc.
func (mr *MockSctpdUplinkServerMockRecorder) NewAssoc(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAssoc", reflect.TypeOf((*MockSctpdUplinkServer)(nil).NewAssoc), arg0, arg1)
}

// SendUl mocks base method.
func (m *MockSctpdUplinkServer) SendUl(arg0 context.Context, arg1 *sctpd.SendUlReq) (*sctpd.SendUlRes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendUl", arg0, arg1)
	ret0, _ := ret[0].(*sctpd.SendUlRes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendUl indicates an expected call of SendUl.
func (mr *MockSctpdUplinkServerMockRecorder) SendUl(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendUl", reflect.TypeOf((*MockSctpdUplinkServer)(nil).SendUl), arg0, arg1)
}

// mustEmbedUnimplementedSctpdUplinkServer mocks base method.
func (m *MockSctpdUplinkServer) mustEmbedUnimplementedSctpdUplinkServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSctpdUplinkServer")
}

// mustEmbedUnimplementedSctpdUplinkServer indicates an expected call of mustEmbedUnimplementedSctpdUplinkServer.
func (mr *MockSctpdUplinkServerMockRecorder) mustEmbedUnimplementedSctpdUplinkServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSctpdUplinkServer", reflect.TypeOf((*MockSctpdUplinkServer)(nil).mustEmbedUnimplementedSctpdUplinkServer))
}

// MockUnsafeSctpdUplinkServer is a mock of UnsafeSctpdUplinkServer interface.
type MockUnsafeSctpdUplinkServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeSctpdUplinkServerMockRecorder
}

// MockUnsafeSctpdUplinkServerMockRecorder is the mock recorder for MockUnsafeSctpdUplinkServer.
type MockUnsafeSctpdUplinkServerMockRecorder struct {
	mock *MockUnsafeSctpdUplinkServer
}

// NewMockUnsafeSctpdUplinkServer creates a new mock instance.
func NewMockUnsafeSctpdUplinkServer(ctrl *gomock.Controller) *MockUnsafeSctpdUplinkServer {
	mock := &MockUnsafeSctpdUplinkServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeSctpdUplinkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeSctpdUplinkServer) EXPECT() *MockUnsafeSctpdUplinkServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedSctpdUplinkServer mocks base method.
func (m *MockUnsafeSctpdUplinkServer) mustEmbedUnimplementedSctpdUplinkServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedSctpdUplinkServer")
}

// mustEmbedUnimplementedSctpdUplinkServer indicates an expected call of mustEmbedUnimplementedSctpdUplinkServer.
func (mr *MockUnsafeSctpdUplinkServerMockRecorder) mustEmbedUnimplementedSctpdUplinkServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedSctpdUplinkServer", reflect.TypeOf((*MockUnsafeSctpdUplinkServer)(nil).mustEmbedUnimplementedSctpdUplinkServer))
}
