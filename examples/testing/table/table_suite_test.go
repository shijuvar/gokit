package table_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestTable(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	//RunSpecs(t, "Table Suite")
	RunSpecsWithDefaultAndCustomReporters(t, "Table Suite", []Reporter{junitReporter})

}
