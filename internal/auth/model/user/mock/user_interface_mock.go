// Code generated by MockGen. DO NOT EDIT.
// Source: internal/auth/model/user/interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	user "github.com/mhdiiilham/segrato/internal/auth/model/user"
	mongo "go.mongodb.org/mongo-driver/mongo"
	options "go.mongodb.org/mongo-driver/mongo/options"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CheckUniqueness mocks base method.
func (m *MockRepository) CheckUniqueness(ctx context.Context, username, email string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUniqueness", ctx, username, email)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CheckUniqueness indicates an expected call of CheckUniqueness.
func (mr *MockRepositoryMockRecorder) CheckUniqueness(ctx, username, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUniqueness", reflect.TypeOf((*MockRepository)(nil).CheckUniqueness), ctx, username, email)
}

// Create mocks base method.
func (m *MockRepository) Create(arg0 context.Context, arg1 user.User) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), arg0, arg1)
}

// FindByID mocks base method.
func (m *MockRepository) FindByID(ctx context.Context, id string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, id)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockRepositoryMockRecorder) FindByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockRepository)(nil).FindByID), ctx, id)
}

// FindOne mocks base method.
func (m *MockRepository) FindOne(ctx context.Context, username string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, username)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockRepositoryMockRecorder) FindOne(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRepository)(nil).FindOne), ctx, username)
}

// PingMongoDB mocks base method.
func (m *MockRepository) PingMongoDB(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PingMongoDB", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// PingMongoDB indicates an expected call of PingMongoDB.
func (mr *MockRepositoryMockRecorder) PingMongoDB(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PingMongoDB", reflect.TypeOf((*MockRepository)(nil).PingMongoDB), ctx)
}

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, id string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, id)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, id)
}

// GetUserByAccessToken mocks base method.
func (m *MockService) GetUserByAccessToken(ctx context.Context, accessToken string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAccessToken", ctx, accessToken)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAccessToken indicates an expected call of GetUserByAccessToken.
func (mr *MockServiceMockRecorder) GetUserByAccessToken(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAccessToken", reflect.TypeOf((*MockService)(nil).GetUserByAccessToken), ctx, accessToken)
}

// Login mocks base method.
func (m *MockService) Login(ctx context.Context, username, password string) (user.User, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, username, password)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Login indicates an expected call of Login.
func (mr *MockServiceMockRecorder) Login(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), ctx, username, password)
}

// RegisterUser mocks base method.
func (m *MockService) RegisterUser(ctx context.Context, username, email, plainPassword string) (user.User, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", ctx, username, email, plainPassword)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockServiceMockRecorder) RegisterUser(ctx, username, email, plainPassword interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockService)(nil).RegisterUser), ctx, username, email, plainPassword)
}

// MockMongoCollection is a mock of MongoCollection interface.
type MockMongoCollection struct {
	ctrl     *gomock.Controller
	recorder *MockMongoCollectionMockRecorder
}

// MockMongoCollectionMockRecorder is the mock recorder for MockMongoCollection.
type MockMongoCollectionMockRecorder struct {
	mock *MockMongoCollection
}

// NewMockMongoCollection creates a new mock instance.
func NewMockMongoCollection(ctrl *gomock.Controller) *MockMongoCollection {
	mock := &MockMongoCollection{ctrl: ctrl}
	mock.recorder = &MockMongoCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMongoCollection) EXPECT() *MockMongoCollectionMockRecorder {
	return m.recorder
}

// FindOne mocks base method.
func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, filter}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOne", varargs...)
	ret0, _ := ret[0].(*mongo.SingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne.
func (mr *MockMongoCollectionMockRecorder) FindOne(ctx, filter interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, filter}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockMongoCollection)(nil).FindOne), varargs...)
}

// InsertOne mocks base method.
func (m *MockMongoCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, document}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InsertOne", varargs...)
	ret0, _ := ret[0].(*mongo.InsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockMongoCollectionMockRecorder) InsertOne(ctx, document interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, document}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockMongoCollection)(nil).InsertOne), varargs...)
}
