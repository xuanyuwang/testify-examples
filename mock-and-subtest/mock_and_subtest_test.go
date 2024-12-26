package mockandsubtest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Prod code
type ExternalService interface {
	Work()
}

type Server struct {
	externalService ExternalService
}

func NewServer(externalService ExternalService) *Server {
	return &Server{
		externalService: externalService,
	}
}

// Test code
type ServerSuite struct {
	suite.Suite
	ExternalService *MockExternalService
	Server
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, &ServerSuite{})
}

// Run before all test cases
func (s *ServerSuite) SetupSuite() {
	s.ExternalService = NewMockExternalService(s.T())
	s.Server = Server{externalService: s.ExternalService}
}

func (s *ServerSuite) TearDownTest() {
	// s.ExternalService.AssertExpectations(s.T())
}

func (s *ServerSuite) TestA() {
	fmt.Println("TestA is running")
	s.ExternalService.EXPECT().Work().Times(1)
}

func (s *ServerSuite) TestB() {
	fmt.Println("TestB is running")
	s.Server.externalService.Work()
}
