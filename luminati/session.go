package luminati

import (
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

var globalMx *sync.Mutex
var countID int

// MaxRequestBySession ...
var MaxRequestBySession = 10

// Session ...
type Session struct {
	ID    string
	count int
	mx    *sync.Mutex
}

// CreateSession ...
func CreateSession() *Session {
	return &Session{ID: uuid(), count: 0, mx: &sync.Mutex{}}
}

// Get ...
func (s *Session) Get() string {
	s.mx.Lock()
	if s.count >= MaxRequestBySession {
		s._reset()
	}
	s.count++
	s.mx.Unlock()

	return s.ID
}

// Reset ...
func (s *Session) Reset() {
	s.mx.Lock()
	s._reset()
	s.mx.Unlock()
}

func (s *Session) _reset() {
	s.ID = uuid()
	s.count = 0
}

func uuid() string {
	if globalMx == nil {
		globalMx = &sync.Mutex{}
	}

	globalMx.Lock()
	countID++
	uuid := generation(countID)
	globalMx.Unlock()

	return uuid
}

func generation(seed int) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}

	uuid := fmt.Sprintf("%x%x%x%x%x%d", b[0:4], b[4:6], b[6:8], b[8:10], b[10:], seed)

	return uuid
}
