package testifymock

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type Biz interface {
	A(a string) string
}

func Main(b Biz) (string, error) {
	result := b.A("Happy Chinese New Year!")
	return result, nil
}

// Start of test

type MockBiz struct {
	mock.Mock
}

func (m *MockBiz) A(a string) string {
	args := m.Called(a)
	return args.String(0)
}

type TestifyMockSuite struct {
	suite.Suite
}

func (s *TestifyMockSuite) TestA() {
	mockBiz := &MockBiz{}
	mockBiz.On("A", mock.Anything).Return("a string")
	Main(mockBiz)
	mockBiz.MethodCalled("A", "called")
}

func TestTestifyMock(t *testing.T) {
	suite.Run(t, &TestifyMockSuite{})
}
