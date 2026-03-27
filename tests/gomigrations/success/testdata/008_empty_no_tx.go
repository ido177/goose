package gomigrations

import (
	"github.com/ido177/goose/v3"
)

func init() {
	goose.AddMigrationNoTx(nil, nil)
}
