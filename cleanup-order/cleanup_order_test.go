package cleanuporder_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestCleanupOrder struct {
	suite.Suite
}

func (s *TestCleanupOrder) SetupSuite() {
	fmt.Println("Run Interface SetupAllSuite: SetupSuite")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface SetupAllSuite: SetupSuite")
	})
}

func (s *TestCleanupOrder) SetupTest() {
	fmt.Println("Run Interface SetupTestSuite: SetupTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface SetupTestSuite: SetupTest")
	})
}

func (s *TestCleanupOrder) TeatDownSuite() {
	fmt.Println("Run Interface TearDownAllSuite: TearDownSuite")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface TearDownAllSuite: TearDownSuite")
	})
}

func (s *TestCleanupOrder) TearDownTest() {
	fmt.Println("Run Interface TearDownTestSuite: TearDownTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface TearDownTestSuite: TearDownTest")
	})
}

func (s *TestCleanupOrder) BeforeTest(suiteName, testName string) {
	fmt.Println("Run Interface BeforeTest: BeforeTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface BeforeTest: BeforeTest")
	})
}

func (s *TestCleanupOrder) AfterTest(suiteName, testName string) {
	fmt.Println("Run Interface AfterTest: AfterTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface AfterTest: AfterTest")
	})
}

func (s *TestCleanupOrder) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Println("Run Interface WithStats: HandleStats")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface WithStats: HandleStats")
	})
}

func (s *TestCleanupOrder) SetupSubTest() {
	fmt.Println("Run Interface SetupSubTest: SetupSubTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface SetupSubTest: SetupSubTest")
	})
}

func (s *TestCleanupOrder) TearDownSubTest() {
	fmt.Println("Run Interface TearDownSubTest: TearDownSubTest")
	s.T().Cleanup(func() {
		fmt.Println("Registered Interface TearDownSubTest: TearDownSubTest")
	})
}

// Test case
func (s *TestCleanupOrder) TestCaseA() {
	fmt.Println("Run TestCaseA - Top level")
	s.T().Cleanup(func() {
		fmt.Println("Registered at TestCaseA - Top level")
	})

	s.Run("TestCaseA - Subtest A", func() {
		s.T().Cleanup(func() {
			fmt.Println("Registered at SubTestCaseA")
		})
		fmt.Println("Run Sub-TestCaseA")
		s.Run("TestCaseA - Sub-subtest A", func() {
			s.T().Cleanup(func() {
				fmt.Println("Registered at Sub-sub-TestCaseA")
			})

			fmt.Println("Run Sub-sub-TestCaseA")
		})
	})
}

// Entry point of the whole test suite
func TestCleanupOrderSuite(t *testing.T) {
	suite.Run(t, &TestCleanupOrder{})
}
