package checkup_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCheckup(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "checkup Suite")
}
