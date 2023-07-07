package clock

import "time"

// Clock provides a timer and a ticker with possibility to add callback functions at every tic and at the end of the timer.
type Clock struct {
	// The TargetID will be automatically removed after Timeout duration, if not zero.
	// The timer starts when the delete button is rendered.
	Timeout time.Duration

	// The optional ticker step, 1s by default
	// Ignored if 0 or greater than the timeout
	TickerStep time.Duration

	// The function to call at every ticker step. Set nil to ignore ticker.
	// TODO: clock - make tic a parameter of Start
	Tic func(*Clock)

	start  time.Time    // The countdown start time
	timer  *time.Timer  // overall timer
	ticker *time.Ticker // internal ticker to handle time left before closing
}

// Start starts the timer and the ticker, according to Timeout and TickerStep properties.
func (_clock *Clock) Start(_finished func()) {
	if _clock.Timeout == 0 {
		return
	}

	// Start the overall timer
	_clock.start = time.Now()
	_clock.timer = time.AfterFunc(_clock.Timeout, _finished)

	// Start the countdown
	tic := _clock.Tic
	if tic != nil && _clock.TickerStep <= _clock.Timeout {
		if _clock.TickerStep == 0 {
			_clock.TickerStep = 1 * time.Second
		}
		go func() {
			_clock.ticker = time.NewTicker(_clock.TickerStep)
			tic(_clock)
			for range _clock.ticker.C {
				tic(_clock)
			}
		}()
	}
}

// Stop stops the ticker and the timer.
func (_clock *Clock) Stop() {
	if _clock.ticker != nil {
		_clock.ticker.Stop()
	}
	if _clock.timer != nil {
		_clock.timer.Stop()
	}
}

// StartTime returns the last countdown start time. Can be a Zero time if the clock has never been started.
func (_clock *Clock) StartTime() time.Time {
	return _clock.start
}

// Timeleft returns the time left before the end of the timer, or zero if the clock is not started.
func (_clock *Clock) TimeLeft() time.Duration {
	if _clock.start.IsZero() {
		return 0
	}
	tl := time.Until(_clock.start.Add(_clock.Timeout))
	if tl < 0 {
		tl = 0
	}
	return tl
}
