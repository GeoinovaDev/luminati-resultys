package luminati

import (
	"sync"
	"time"
)

// Pool ...
type Pool struct {
	urls []*URL
	seek int
	mx   *sync.Mutex
}

// CreatePool ...
func CreatePool() *Pool {
	return &Pool{seek: 0, mx: &sync.Mutex{}}
}

// Add ...
func (p *Pool) Add(url *URL) *Pool {
	p.urls = append(p.urls, url)

	return p
}

// PopAndPush ...
func (p *Pool) PopAndPush(url *URL) *Pool {
	p.mx.Lock()
	for i, _url := range p.urls {
		if url.Equal(_url) {
			p.urls = append(p.urls[:i], p.urls[i+1:]...)
			p.urls = append(p.urls, _url)
			break
		}
	}
	p.mx.Unlock()
	return p
}

// RemoveTemporary
func (p *Pool) RemoveTemporary(url *URL, sec int) *Pool {
	var urlRemoved *URL = nil

	p.mx.Lock()
	for i, _url := range p.urls {
		if url.Equal(_url) {
			urlRemoved = _url
			p.urls = append(p.urls[:i], p.urls[i+1:]...)
			break
		}
	}
	p.mx.Unlock()

	go func() {
		time.Sleep(time.Duration(sec) * time.Second)
		if urlRemoved != nil {
			p.mx.Lock()
			p.urls = append(p.urls, urlRemoved)
			p.mx.Unlock()
		}
	}()

	return p
}

// Get ...
func (p *Pool) Get() *URL {
	p.mx.Lock()
	url := p.urls[p.seek%len(p.urls)]
	p.seek++
	p.mx.Unlock()

	return url
}

// Clear ...
func (p *Pool) Clear() {
	p.mx.Lock()
	p.seek = 0
	p.urls = []*URL{}
	p.mx.Unlock()
}
