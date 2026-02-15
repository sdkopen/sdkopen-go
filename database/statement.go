package database

import (
	"context"
	"database/sql"
	"errors"
	"io"

	"github.com/sdkopen/sdkopen-go/logging"
)

const (
	sqlTxContext         string = "SqlTxContext"
	closerErrorMsg       string = "Could not close statement: %v"
	queryIsEmptyErrorMsg string = "query is empty"
)

type Statement struct {
	ctx   context.Context
	query string
	args  []any
}

func NewStatement(ctx context.Context, query string, params ...any) *Statement {
	return &Statement{ctx, query, params}
}

func (s *Statement) Execute() error {
	return s.ExecuteInInstance(dbInstance)
}

func (s *Statement) ExecuteInInstance(instance *sql.DB) error {
	if err := s.validate(instance); err != nil {
		return err
	}

	stmt, err := s.createStatement(instance)
	if err != nil {
		return err
	}
	defer closer(stmt)

	if _, err = stmt.Exec(s.args...); err != nil {
		return err
	}

	return nil
}

func (s *Statement) createStatement(instance *sql.DB) (*sql.Stmt, error) {
	if tx := s.ctx.Value(sqlTxContext); tx != nil {
		return tx.(*sql.Tx).PrepareContext(s.ctx, s.query)
	}

	return instance.PrepareContext(s.ctx, s.query)
}

func (s *Statement) validate(instance *sql.DB) error {
	if instance == nil {
		return errors.New(dbNotInitializedErrorMsg)
	}

	if s.query == "" {
		return errors.New(queryIsEmptyErrorMsg)
	}

	return nil
}

func closer(o io.Closer) {
	if err := o.Close(); err != nil {
		logging.Error(closerErrorMsg, err)
	}
}
