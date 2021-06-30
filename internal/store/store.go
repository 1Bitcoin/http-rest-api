package store

import (
    "database/sql"

    _ "github.com/lib/pq" // анонимный импорт, чтобы сами методы не импортировались в код
)

type Store struct {
    config *Config
    db *sql.DB
}

func New(config *Config) *Store {
    return &Store{
        config: config,
    }
}

func (s *Store) Open() error {
    db, err := sql.Open("postgres", s.config.DatabaseURL)
    if err != nil {
        return err
    }

    // ping db
    if err := db.Ping(); err != nil {
        return  err
    }

    s.db = db
    return nil
}

func (s *Store) Close() {
    s.db.Close()
}
