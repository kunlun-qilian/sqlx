package builder_test

import (
	"testing"

	. "github.com/kunlun-qilian/sqlx/v3/builder"
	. "github.com/kunlun-qilian/sqlx/v3/builder/buidertestingutils"
	"github.com/onsi/gomega"
)

func TestFunc(t *testing.T) {
	t.Run("invalid", func(t *testing.T) {
		gomega.NewWithT(t).Expect(Func("")).To(BeExpr(""))
	})
	t.Run("count", func(t *testing.T) {
		gomega.NewWithT(t).Expect(Count()).To(BeExpr("COUNT(1)"))
	})
	t.Run("AVG", func(t *testing.T) {
		gomega.NewWithT(t).Expect(Avg()).To(BeExpr("AVG(*)"))
	})
}
