package dialects

import (
	"fmt"
	"strings"

	"github.com/ido177/goose/v3/database/dialect"
)

// DefaultSparkStorageFormat is the default storage format used for the goose_db_version table in Spark SQL.
var DefaultSparkStorageFormat = "ICEBERG"

// NewSpark returns a new [dialect.Querier] for the Spark SQL dialect.
// It initializes the querier using the default storage format (ICEBERG).
func NewSpark() dialect.Querier {
	return NewSparkWithFormat(DefaultSparkStorageFormat)
}

// NewSparkWithFormat returns a new [dialect.Querier] for the Spark SQL dialect
// using a specific storage format (e.g., "ICEBERG" or "PAIMON").
func NewSparkWithFormat(format string) dialect.Querier {
	f := strings.ToUpper(strings.TrimSpace(format))
	// Strictly enforce supported ACID formats to ensure DELETE operations work.
	if f != "ICEBERG" && f != "PAIMON" {
		f = "ICEBERG"
	}
	return &spark{storageFormat: f}
}

// spark implements the dialect.Querier interface for Spark SQL.
type spark struct {
	storageFormat string
}

var _ dialect.Querier = (*spark)(nil)

// CreateTable generates the SQL to create the goose_db_version table.
func (s *spark) CreateTable(tableName string) string {
	// Spark SQL uses the 'USING' clause for table formats instead of 'STORED BY'.
	// Since Spark does not support AUTO_INCREMENT or IDENTITY intuitively here,
	// the 'id' field remains a standard bigint.
	q := `CREATE TABLE IF NOT EXISTS %s (
		id bigint,
		version_id bigint,
		is_applied boolean,
		tstamp timestamp
	) USING %s`
	return fmt.Sprintf(q, tableName, s.storageFormat)
}

// InsertVersion generates the SQL to insert a new migration record.
func (s *spark) InsertVersion(tableName string) string {
	q := `INSERT INTO %s (version_id, is_applied, tstamp) VALUES (?, ?, CURRENT_TIMESTAMP)`
	return fmt.Sprintf(q, tableName)
}

// DeleteVersion generates the SQL to delete a migration record.
func (s *spark) DeleteVersion(tableName string) string {
	q := `DELETE FROM %s WHERE version_id=?`
	return fmt.Sprintf(q, tableName)
}

// GetMigrationByVersion generates the SQL to retrieve a specific migration.
func (s *spark) GetMigrationByVersion(tableName string) string {
	q := `SELECT tstamp, is_applied FROM %s WHERE version_id=? ORDER BY tstamp DESC LIMIT 1`
	return fmt.Sprintf(q, tableName)
}

// ListMigrations generates the SQL to list all migrations.
func (s *spark) ListMigrations(tableName string) string {
	q := `SELECT version_id, is_applied FROM %s ORDER BY version_id DESC`
	return fmt.Sprintf(q, tableName)
}

// GetLatestVersion generates the SQL to get the maximum version_id.
func (s *spark) GetLatestVersion(tableName string) string {
	q := `SELECT max(version_id) FROM %s`
	return fmt.Sprintf(q, tableName)
}
