package task

import (
	"testing"
	"time"

	"github.com/prashantv/gostub"
)

func Test_DateTime(t *testing.T) {
	stubs := gostub.Stub(&timeNow, func() time.Time {
		return time.Date(2015, 6, 1, 0, 0, 0, 0, time.UTC)
	})
	defer stubs.Reset()
	t.Log(GetUnixTime())
}
