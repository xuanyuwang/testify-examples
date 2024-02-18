package unexpectedcleanuporder_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

var externalResource = 1

type TestUnexpectedCleanupOrder struct {
	suite.Suite
}

func (s *TestUnexpectedCleanupOrder) TearDownTest() {
	fmt.Println("TearDownTest")
}

func (s *TestUnexpectedCleanupOrder) SetupSubTest() {
	externalResource = 2
	fmt.Printf("\tSetupSubTest. The externalResoure is: %d\n", externalResource)
	s.T().Cleanup(func() {
		externalResource = 1
		fmt.Printf("\tCleanup subTest. The externalResoure is: %d\n", externalResource)
	})
}

// func (s *TestUnexpectedCleanupOrder) TearDownSubTest() {
// 	externalResource = 1
// 	fmt.Printf("\tTeardown subTest. The externalResoure is: %d\n", externalResource)
// }

// Test case
func (s *TestUnexpectedCleanupOrder) TestCaseA() {
	fmt.Printf("Start TestCase A. The externalResoure is: %d\n", externalResource)

	s.Run("Sub-TestCase A", func() {
		fmt.Printf("\tFinish Sub-TestCase A. The externalResoure is: %d\n", externalResource)
	})
	fmt.Printf("Back to TestCase A. The externalResource is: %d\n", externalResource)
}

// Entry point of the whole test suite
func TestUnexpectedCleanupOrderSuite(t *testing.T) {
	suite.Run(t, &TestUnexpectedCleanupOrder{})
}
