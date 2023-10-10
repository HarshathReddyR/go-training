package db

type Conn struct {
	db string
}

func New(db string) (Conn, bool) {
	if db != "" {
		return Conn{}, false
	}
	return Conn{db: db}, true
}
