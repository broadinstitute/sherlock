package db

import "embed"

// MigrationFiles stores the migration SQL files themselves inside Sherlock's binary.
// If we were to reference these files from the machine's filesystem, the relative path
// would change based on the entrypoint file, which is quite inconvenient to handle.
// Instead, we reference them with a relative path here with `embed` so we can import it
// like any other variable throughout Sherlock.
//
//go:embed migrations/*.sql
var MigrationFiles embed.FS
