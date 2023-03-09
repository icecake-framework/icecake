package clock

import "time"

type Clock struct {
	// The TargetID will be automatically removed after Timeout duration, if not zero.
	// The timer starts when the delete button is rendered.
	Timeout time.Duration

	// The optional ticker step, 1s by default.
	// Ignored if greater than the timeout
	TickerStep time.Duration

	// the function to call at every ticker step
	Tic func(*Clock)

	start  time.Time    // The countdown start time
	timer  *time.Timer  // overall timer
	ticker *time.Ticker // internal ticker to handle time left before closing
}

func (_clock *Clock) Start(_finished func()) {
	if _clock.Timeout == 0 {
		return
	}

	if _clock.TickerStep == 0 {
		_clock.TickerStep = 1 * time.Second
	}

	// Start the overall timer
	_clock.start = time.Now()
	_clock.timer = time.AfterFunc(_clock.Timeout, _finished)

	// Start the countdown
	if _clock.Tic != nil && _clock.TickerStep <= _clock.Timeout {
		go func() {
			_clock.ticker = time.NewTicker(_clock.TickerStep)
			_clock.Tic(_clock)
			for _ = range _clock.ticker.C {
				_clock.Tic(_clock)
			}
		}()
	}
}

func (_clock *Clock) Stop() {
	if _clock.ticker != nil {
		_clock.ticker.Stop()
	}
	if _clock.timer != nil {
		_clock.timer.Stop()
	}
}

func (_clock *Clock) StartTime() time.Time {
	return _clock.start
}

func (_clock *Clock) TimeLeft() time.Duration {
	tl := time.Until(_clock.start.Add(_clock.Timeout))
	if tl < 0 {
		tl = 0
	}
	return tl
}
