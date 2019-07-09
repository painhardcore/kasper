package file

import (
	"encoding/gob"
	"github.com/painhardcore/kasper/pkg/counter"
	"os"
	"strconv"
	"sync"
	"time"
)

// An store implements Counter with persistent file store.
// It is safe for concurrent use by multiple goroutines.
type store struct {
	sync.RWMutex
	dur      time.Duration // duration to count
	filename string        // path for file to store
	data     *Timecount
}

// implement Clocker interface, so we use current time
type realclock struct{}

// just return current time
func (realclock) Now() time.Time {
	return time.Now()
}

// Creates Counter which stores count for specified duration persistently to disk.
func New(filename string, dur time.Duration) (counter.Counter, error) {
	data := Timecount{cl: realclock{}}
	c := store{filename: filename, dur: dur, data: &data}
	err := c.load()
	if err != nil {
		return nil, err
	}
	return &c, err
}

// Implement Stringer interface to print count number.
func (s store) String() string {
	s.RLock()
	defer s.RUnlock()
	return strconv.FormatInt(s.data.Count, 10)
}

// Increase counter and flush changes to file.
// If time passed the deadline - counter will reset and next deadline will be set.
func (s *store) Inc() error {
	s.Lock()
	defer s.Unlock()
	if s.data.expired() {
		s.data.reset(s.dur)
	}
	s.data.Count += 1
	return s.write()
}

// For writing data to file.
func (s *store) write() error {
	file, err := os.Create(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewEncoder(file)
	err = decoder.Encode(s.data)

	return err
}

// Create initial store with due time
func (s *store) reset() error {
	s.data.reset(s.dur)
	return s.write()
}

// Loading data from disk at startup.
// If no file found - creates one.
func (s *store) load() error {
	s.Lock()
	defer s.Unlock()

	_, err := os.Stat(s.filename)
	if err != nil {
		return s.reset()
	}

	file, err := os.Open(s.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(s.data)
	if err != nil {
		return err
	}
	if s.data.expired() {
		s.data.reset(s.dur)
	}
	return nil
}
