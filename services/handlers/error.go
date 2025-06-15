package handlers

type APIError struct {
	status int
	msg    string
}

func (a APIError) Error() string {
	return a.msg
}

func (a APIError) APIError() (int, string) {
	return a.status, a.msg
}
