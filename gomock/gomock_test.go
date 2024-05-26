package gomocktest

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type GoMockSuite struct {
	suite.Suite
	foo        *MockFoo
	controller *gomock.Controller
}

func (s *GoMockSuite) SetupSuite() {
	s.controller = gomock.NewController(s.T())
	s.foo = NewMockFoo(s.controller)
}

// gomock has to be cleaned after each test.
// In this example, behaviours set up in TestA are used in TestB if we don't call s.controller.Finish after each test
// Uncomment s.controller.Finish() will expose errors in TestA
func (s *GoMockSuite) TearDownTest() {
	// s.controller.Finish()
}

func (s *GoMockSuite) TestA() {
	s.foo.EXPECT().A(gomock.Any()).Return("hello").Times(2)
	res := s.foo.A(3)
	s.Assert().Equal(res, "hello")
}

func (s *GoMockSuite) TestB() {
	res := s.foo.A(3)
	s.Assert().Equal(res, "hello")
}

func TestTestifyMock(t *testing.T) {
	suite.Run(t, &GoMockSuite{})
}
