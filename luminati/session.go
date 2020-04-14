package luminati

import (
	"strconv"
	"sync"
)

var globalMx *sync.Mutex
var countID int

// MaxRequestBySession ...
var MaxRequestBySession = 10

// Session ...
type Session struct {
	id    int
	count int
	mx    *sync.Mutex
}

// CreateSession ...
func CreateSession() *Session {
	return &Session{id: uuid(), count: 0, mx: &sync.Mutex{}}
}

// Get ...
func (s *Session) Get() string {
	s.mx.Lock()
	if s.count >= MaxRequestBySession {
		s._reset()
	}
	s.count++
	s.mx.Unlock()

	return strconv.Itoa(s.id)
}

// Reset ...
func (s *Session) Reset() {
	s.mx.Lock()
	s._reset()
	s.mx.Unlock()
}

func (s *Session) _reset() {
	s.id = uuid()
	s.count = 0
}

func uuid() int {
	if globalMx == nil {
		globalMx = &sync.Mutex{}
	}

	globalMx.Lock()
	countID++
	globalMx.Unlock()

	return countID
}
