package testutils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func (s *ExampleTestSuite) TestDocker() {
	s.compose("down")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
