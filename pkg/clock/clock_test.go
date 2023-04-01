package clock

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {

	wait := make(chan struct{})
	tac := 0
	c := Clock{
		Timeout:    time.Second * 5,
		TickerStep: time.Second,
		Tic: func(cc *Clock) {
			want := float64(5 - tac)
			get := cc.TimeLeft().Round(time.Second).Seconds()
			assert.Equal(t, want, get)
			tac++
		},
	}

	c.Start(func() {
		assert.True(t, tac == 6)
		close(wait)
	})

	<-wait
}
