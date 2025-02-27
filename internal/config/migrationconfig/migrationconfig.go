package migrationconfig

type MigrationType int

const (
	Up = iota
	Down
)

type MigrationConfig struct {
	MigrationType
	SourcePath string
}

func New(migrationType MigrationType, path string) *MigrationConfig {
	return &MigrationConfig{
		SourcePath: path,
		MigrationType: migrationType,
	}
}
