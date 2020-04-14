package luminati

import "net"

// URL ...
type URL struct {
	user    string
	pass    string
	ip      string
	Session *Session
}

// CreateURL ...
func CreateURL(user string, pass string) *URL {
	return &URL{
		user:    user,
		pass:    pass,
		ip:      getHostByName("zproxy.lum-superproxy.io"),
		Session: CreateSession(),
	}
}

// ToString ...
func (u *URL) ToString() string {
	return u.user + "-session-" + u.Session.Get() + ":" + u.pass + "@" + u.ip + ":22225"
}

func getHostByName(name string) string {
	ips, _ := net.LookupHost(name)
	return ips[0]
}
