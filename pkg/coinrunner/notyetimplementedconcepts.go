package coinrunner

import "time"

// check how many requests came from this ip in the last 20minutes
func CheckVelocity(token Token) int {
	count := 0
	for _, t := range Memory {
		// check if its the same ip
		if t.SenderIp == token.SenderIp {
			// calculate delta and check if it is +-10min
			timeDelta := t.Timestamp.Sub(token.Timestamp)
			if timeDelta < 0 {
				timeDelta = -timeDelta
			}
			if timeDelta < 10*time.Minute {
				count++
			}
		}
	}

	return count
}

func RetryAction(token Token, maxRetries int) {
	delay := time.Second * 1
	var err error

	// for a specific number of retries
	for range maxRetries {
		// call to action that can fail
		err = nil // err = ActionToRetry(token)
		//if action completed return
		if err == nil {
			return
		}
		// if action failed, sleep and try again
		time.Sleep(delay)
		// increase delay after each retry
		delay = delay * 2
	}
}

type RateLimiter struct {
	Last     time.Time
	Interval time.Duration
}

func NewRateLimiter(interval time.Duration) RateLimiter {
	return RateLimiter{
		Interval: interval,
	}
}

// check when the last action happened and decide if we allow or not
func (rl RateLimiter) Allow() bool {
	now := time.Now()
	if now.Sub(rl.Last) >= rl.Interval {
		rl.Last = now
		return true
	}

	return false
}
