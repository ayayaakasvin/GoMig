package scripts

import (
	"io"
	"log"
	"os"
	joiner "path" // used alias, because i already used "path" as variable
	"strings"

	"migrationtool/internal/config/migrationconfig"
)

// ParseMigrationFiles reads the migration files and prepares them
func ParseMigrationFiles (mconf *migrationconfig.MigrationConfig) ([]string, error) {
	if mconf == nil {
		return nil, EmptyMigrationConfig	
	}

	sqlScripts , err:= parse(mconf.MigrationType, mconf.SourcePath)
	if len(sqlScripts) == 0 {
		return nil, err
	}

	return sqlScripts, nil
}

// parse processes migration files based on the specified migration type and path.
// It handles both single SQL file and directory containing multiple SQL files.
//
// For a single file:
// - Reads and returns the SQL content as a single-element slice
//
// For a directory:
// - Reads all files in the directory
// - Filters files based on migration type (up/down)
// - Collects SQL content from matching files
//
// Parameters:
//   - mtype: Type of migration (Up/Down)
//   - path: File path or directory path containing SQL migration files
//
// Returns:
//   - []string: Slice of SQL scripts content
//   - error: Returns nil on success, or one of:
//     * UnknownMigration: If migration type is unknown
//     * DirectionNoFile: If directory is empty
//     * EmptyDirectory: If no matching migration files found
//     * Other errors from file operations
func parse (mtype migrationconfig.MigrationType, path string) ([]string, error) {
	if mtype == migrationconfig.Unknown {
		return nil, UnknownMigration
	}

	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	} else if !fileInfo.IsDir() {
		script, err := readSqlFile(path)
		if err != nil {
            log.Printf("failed to read file %s: %v", path, err)
            return nil, err
        }
		return []string{script}, nil
	}
	
	var sqlScripts []string
	if fileInfo.IsDir() {
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Printf("failed to read directory: %v", err)
			return nil, err
		} else if len(entries) == 0 {
			log.Printf("directory %s is empty", path)
			return nil, DirectionNoFile
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			entryPath := joiner.Join(path, entry.Name())
			if (mtype == migrationconfig.Up && isUpFile(entryPath) || 
				mtype == migrationconfig.Down && isDownFile(entryPath)) {
					sqlScript, err := readSqlFile(entryPath)
					if err != nil {
						log.Printf("failed to read file %s: %v", entryPath, err)
						return nil, err
					}
					sqlScripts = append(sqlScripts, sqlScript)
			}
		}
	}

	if len(sqlScripts) == 0 {
		log.Printf("directory %s doesn't contain migration files", path)
		return nil, EmptyDirectory
	}
	
	return sqlScripts, nil
}

func readSqlFile(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("failed to open file %s: %v", filepath, err)
		return "", err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Printf("error during reading file %s: %v", file.Name(), err)
		return "", err
	} else if len(content) == 0 {
		log.Printf("file %s is empty", filepath)
        return "", EmptySqlFile
	}

	return string(content), nil
}

func isUpFile(filename string) bool {
	return strings.HasSuffix(filename, ".up.sql")
}

func isDownFile(filename string) bool {
	return strings.HasSuffix(filename, ".down.sql")
}