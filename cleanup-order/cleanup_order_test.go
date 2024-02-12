package cleanuporder_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

var tearDownSubTest []string

type TestCleanupOrder struct {
	suite.Suite
}

func (s *TestCleanupOrder) SetupSuite() {
	fmt.Println("Run Interface SetupAllSuite: SetupSuite")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface SetupAllSuite: SetupSuite")
	})
}

func (s *TestCleanupOrder) SetupTest() {
	fmt.Println("Run Interface SetupTestSuite: SetupTest")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface SetupTestSuite: SetupTest")
	})
}

func (s *TestCleanupOrder) TeatDownSuite() {
	fmt.Println("Run Interface TearDownAllSuite: TearDownSuite")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface TearDownAllSuite: TearDownSuite")
	})
}

func (s *TestCleanupOrder) TearDownTest() {
	fmt.Println("Run Interface TearDownTestSuite: TearDownTest")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface TearDownTestSuite: TearDownTest")
	})
}

func (s *TestCleanupOrder) BeforeTest(suiteName, testName string) {
	fmt.Printf("Run Interface BeforeTest: BeforeTest for suite %s - test %s\n", suiteName, testName)
	s.T().Cleanup(func() {
		fmt.Printf("Cleanup Interface BeforeTest: BeforeTest for suite %s - test %s\n", suiteName, testName)
	})
}

func (s *TestCleanupOrder) AfterTest(suiteName, testName string) {
	fmt.Printf("Run Interface AfterTest: AfterTest for suite %s - test %s\n", suiteName, testName)
	s.T().Cleanup(func() {
		fmt.Printf("Cleanup Interface AfterTest: AfterTest for suite %s - test %s\n", suiteName, testName)
	})
}

func (s *TestCleanupOrder) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	fmt.Printf("Run Interface WithStats: HandleStats for suite %s\n", suiteName)
	s.T().Cleanup(func() {
		fmt.Printf("Cleanup Interface WithStats: HandleStats for suite %s\n", suiteName)
	})
}

func (s *TestCleanupOrder) SetupSubTest() {
	fmt.Println("Run Interface SetupSubTest: SetupSubTest")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface SetupSubTest: SetupSubTest")
	})
}

func (s *TestCleanupOrder) TearDownSubTest() {
	fmt.Println("Run Interface TearDownSubTest: TearDownSubTest")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup Interface TearDownSubTest: TearDownSubTest")
		fmt.Printf("\tTear down sub test %v\n", tearDownSubTest)
	})
}

// Test case
func (s *TestCleanupOrder) TestCaseA() {
	fmt.Println("Run TestCase A")
	s.T().Cleanup(func() {
		fmt.Println("Cleanup TestCase A")
	})

	s.Run("Sub-TestCase A", func() {
		fmt.Println("Run Sub-TestCase A")
		s.T().Cleanup(func() {
			fmt.Println("Cleanup Sub-TestCase A")
		})
		tearDownSubTest = append(tearDownSubTest, "Sub-TestCase A")

		s.Run("Sub-Sub-TestCase A", func() {
			s.T().Cleanup(func() {
				fmt.Println("Cleanup Sub-Sub-TestCase A")
			})
			fmt.Println("Run Sub-Sub-TestCase A")
			tearDownSubTest = append(tearDownSubTest, "Sub-Sub-TestCase A")
			fmt.Println("Finish Sub-Sub-TestCase A")
		})
		s.Run("Sub-Sub-TestCase B", func() {
			s.T().Cleanup(func() {
				fmt.Println("Cleanup Sub-Sub-TestCase B")
			})
			fmt.Println("Run Sub-Sub-TestCase B")
			tearDownSubTest = append(tearDownSubTest, "Sub-Sub-TestCase B")
			fmt.Println("Finish Sub-Sub-TestCase B")
		})

		fmt.Println("Finish Sub-TestCase A")
	})
	fmt.Println("Finish TestCase A")
}

// Entry point of the whole test suite
func TestCleanupOrderSuite(t *testing.T) {
	suite.Run(t, &TestCleanupOrder{})
}
