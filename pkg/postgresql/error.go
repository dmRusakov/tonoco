package postgresql

import (
	"github.com/dmRusakov/tonoco/pkg/common/errors"
)

var (
	ErrNothingInserted = errors.New("nothing inserted")
)

func ErrCommit(err error) error {
	return errors.Wrap(err, "failed to commit Tx")
}

func ErrRollback(err error) error {
	return errors.Wrap(err, "failed to rollback Tx")
}

func ErrCreateTx(err error) error {
	return errors.Wrap(err, "failed to create Tx")
}

func ErrCreateQuery(err error) error {
	return errors.Wrap(err, "failed to create SQL Query")
}

func ErrScan(err error) error {
	return errors.Wrap(err, "failed to scan")
}

func ErrExec(err error) error {
	return errors.Wrap(err, "failed to execute")
}

func ErrDoQuery(err error) error {
	return errors.Wrap(err, "failed to query")
}

func ErrNoRows() error {
	return errors.Wrap(errors.New("no rows found"), "failed to query")
}

func ErrNoRowForID(id string) error {
	return errors.New("no rows found for id: " + id)
}

func ErrNoRowForURL(url string) error {
	return errors.New("no rows found for url: " + url)
}
func ErrNoRowForUpdatedAt(id string) error {
	return errors.New("No rows found for updated_at with id: " + id)
}
func ErrNoRowForTableUpdated() error {
	return errors.New("No rows found for table updated_at")
}
