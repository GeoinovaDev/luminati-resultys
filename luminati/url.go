package luminati

import "net"

// URL ...
type URL struct {
	user    string
	pass    string
	ip      string
	port    string
	Session *Session
}

// CreateURL ...
func CreateURL(user string, pass string, host string, port string) *URL {
	return &URL{
		user:    user,
		pass:    pass,
		ip:      host,
		port:    port,
		Session: CreateSession(),
		// ip:      getHostByName(host),
	}
}

// ToString ...
func (u *URL) ToString() string {
	return u.user + "-session-" + u.Session.Get() + ":" + u.pass + "@" + u.ip + ":" + u.port
}

// Equal ...
func (u *URL) Equal(url *URL) bool {
	return u.user == url.user && u.pass == url.pass && u.ip == url.ip && u.port == url.port
}

// Raw ...
func (u *URL) Raw() string {
	return u.user + "-session-" + u.Session.ID + ":" + u.pass + "@" + u.ip + ":" + u.port
}

func getHostByName(name string) string {
	ips, _ := net.LookupHost(name)
	return ips[0]
}
