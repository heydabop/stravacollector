package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Emyrk/strava/database/migrations"
	"github.com/Emyrk/strava/database/postgres"
)

const minimumPostgreSQLVersion = 13

func main() {
	connection, closeFn, err := postgres.Open()
	if err != nil {
		panic(err)
	}
	defer closeFn()

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	err = migrations.Up(db)
	if err != nil {
		panic(err)
	}

	hasPGDump := false
	if _, err = exec.LookPath("pg_dump"); err == nil {
		out, err := exec.Command("pg_dump", "--version").Output()
		if err == nil {
			// Parse output:
			// pg_dump (PostgreSQL) 14.5 (Ubuntu 14.5-0ubuntu0.22.04.1)
			parts := strings.Split(string(out), " ")
			if len(parts) > 2 {
				version, err := strconv.Atoi(strings.Split(parts[2], ".")[0])
				if err == nil && version >= minimumPostgreSQLVersion {
					hasPGDump = true
				}
			}
		}
	}

	cmdArgs := []string{
		"pg_dump",
		"--schema-only",
		connection,
		"--no-privileges",
		"--no-owner",

		// We never want to manually generate
		// queries executing against this table.
		"--exclude-table=schema_migrations",
	}

	if !hasPGDump {
		cmdArgs = append([]string{
			"docker",
			"run",
			"--rm",
			"--network=host",
			fmt.Sprintf("postgres:%d", minimumPostgreSQLVersion),
		}, cmdArgs...)
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...) //#nosec
	cmd.Env = append(os.Environ(), []string{
		"PGTZ=UTC",
		"PGCLIENTENCODING=UTF8",
	}...)
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		panic(err)
	}

	for _, sed := range []string{
		// Remove all comments.
		"/^--/d",
		// Public is implicit in the schema.
		"s/ public\\./ /g",
		"s/::public\\./::/g",
		"s/'public\\./'/g",
		// Remove database settings.
		"s/SET .* = .*;//g",
		// Remove select statements. These aren't useful
		// to a reader of the dump.
		"s/SELECT.*;//g",
		// Removes multiple newlines.
		"/^$/N;/^\\n$/D",
	} {
		cmd := exec.Command("sed", "-e", sed)
		cmd.Stdin = bytes.NewReader(output.Bytes())
		output = bytes.Buffer{}
		cmd.Stdout = &output
		cmd.Stderr = os.Stderr
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

	dump := fmt.Sprintf("-- Code generated by 'make coderd/database/generate'. DO NOT EDIT.\n%s", output.Bytes())
	_, mainPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("couldn't get caller path")
	}
	err = os.WriteFile(filepath.Join(mainPath, "..", "..", "..", "dump.sql"), []byte(dump), 0o600)
	if err != nil {
		panic(err)
	}
}
