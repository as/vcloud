package session

type session struct {
	Host  string
	Port  string
	Org   string
	User  string
	Pass  string
	Token string

	client *http.Client
	rx     int64
	tx     int64
}

func New(host, port, org, user, pass string) *session{
	return &Session{ 
		Host: host,
		Port: port,
		Org: org,
		User: user,
		Pass: pass,
	}
}

func Old(host, port, token string) *session{
	return &Session{ Host: host,
		Host: host,
		Port: port,
		Token: token,
	}
}

func (s *Session) Rx() int64 {
	return s.rx
}

func (s *Session) Tx() int64 {
	return s.tx
}
