package mockandsubtest

import "github.com/stretchr/testify/suite"

type BaseTestSuite struct {
	suite.Suite
	ExternalService
	Server
}

// Run before all test cases
func (s *BaseTestSuite) SetupSuite() {
	s.ExternalService = NewMockExternalService(s.T())
	s.Server = Server{s.ExternalService}
}
