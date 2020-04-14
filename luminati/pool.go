package luminati

import "sync"

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

// Get ...
func (p *Pool) Get() *URL {
	p.mx.Lock()
	url := p.urls[p.seek%len(p.urls)]
	p.seek++
	p.mx.Unlock()

	return url
}
