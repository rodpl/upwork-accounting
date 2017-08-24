package nbp

import (
	. "gopkg.in/check.v1"
	"testing"
	"time"
)

func Test(t *testing.T) { TestingT(t) }

type NbpSuite struct{}

var _ = Suite(&NbpSuite{})

func (s *NbpSuite) TestUSDollacOnDat(c *C) {
	var date, _ = time.Parse("2006-01-02", "2018-08-14")
	var rate = GetExchangeRate("USD", date)
	c.Assert(rate, Equals, 3.6278)
}

func (s *NbpSuite) IgnaroTestUSDollacOnDateWhereAreNoRatesExampleOnSunday(c *C) {
	var date, _ = time.Parse("2006-01-02", "2018-08-15")
	var rate = GetExchangeRate("USD", date)
	c.Assert(rate, Equals, 3.6278)
}
