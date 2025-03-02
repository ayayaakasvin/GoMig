package migrationconfig

type MigrationType int

const (
	Up MigrationType = iota
	Down
	Unknown
)

type MigrationConfig struct {
	MigrationType
	SourcePath string
}

// New creates a new MigrationConfig with the specified migration type and source path.
// It returns a pointer to the MigrationConfig instance.
//
// Parameters:
//   - migrationType: The type of migration to be performed
//   - path: The source path for the migration
//
// Returns:
//   - *MigrationConfig: A pointer to the newly created MigrationConfig
func New(migrationType MigrationType, path string) *MigrationConfig {
	return &MigrationConfig{
		SourcePath: path,
		MigrationType: migrationType,
	}
}
