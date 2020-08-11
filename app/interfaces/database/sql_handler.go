package database

// SQLHandler interface
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Result interface
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row interface
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
