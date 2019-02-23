package midd

type Session struct {
	token string
}

// MiddelSession #middleware:"session"#
func MiddelSession(token string /* #in:"header" name:"x-token"# */) (session *Session, err error) {
	return &Session{token}, nil
}

// MiddelToken #middleware:"token"#
func MiddelToken(
	session *Session, /* #in:"middleware"# */
) (token string, err error) {
	return session.token, nil
}
