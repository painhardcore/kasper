package file

import "time"

type Clocker interface {
	Now() time.Time
}

type Timecount struct {
	cl    Clocker
	Count int64
	Due   time.Time
}

func (t *Timecount) reset(dur time.Duration) {
	t.Due = t.cl.Now().Add(dur)
	t.Count = 0
}
func (t *Timecount) expired() bool {
	return t.cl.Now().After(t.Due)
}
