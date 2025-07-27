package core

const (
	DatabaseTypeSqlite     DatabaseTypeT = "sqlite"
	DatabaseTypeMysql      DatabaseTypeT = "mysql"
	DatabaseTypePostgresql DatabaseTypeT = "postgresql"
	DatabaseTypeMongoDB    DatabaseTypeT = "mongodb"
)

type (
	DatabaseTypeT string

	Database struct {
		ResourceObject

		DatabaseType DatabaseTypeT `json:"database_type"` // e.g., "mysql", "postgresql", "mongodb"
		Username     string        `json:"username"`
		Password     string        `json:"password,omitempty"` // Optional, can be empty if not required
	}
)

func NewDatabase() (*Database, error) {
	return &Database{
		ResourceObject: ResourceObject{
			Name: "database",
		},
	}, nil
}

func (d *Database) Deploy() error {
	// Implement deployment logic here
	return nil
}
