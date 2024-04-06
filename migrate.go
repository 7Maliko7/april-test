package main

import (
	"embed"
	"fmt"
	"log"
	"path"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/7Maliko7/april-test/internal/config"
	"github.com/7Maliko7/april-test/migration"
)

func doMigration(cfg *config.Config) error {
	pgMigrationsDir := "database/postgres"

	fmt.Println("Found migration files")
	fmt.Println("Postgres:")

	fileList, err := getAllFilenames(&migration.Database, pgMigrationsDir)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	for i := range fileList {
		fmt.Println(fileList[i])
	}

	m, err := migrate.New("file://migration/"+pgMigrationsDir, cfg.RWDB.ConnectionString)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	ver, dirt, _ := m.Version()
	fmt.Printf("DB Version: %v Dirty: %v\n", ver, dirt)

	err = migrateUp(m)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	ver, dirt, _ = m.Version()
	fmt.Printf("New DB Version: %v Dirty: %v\n", ver, dirt)

	return nil
}

func getAllFilenames(fs *embed.FS, dir string) (out []string, err error) {
	if len(dir) == 0 {
		dir = "."
	}
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, entry := range entries {
		fp := path.Join(dir, entry.Name())
		if entry.IsDir() {
			res, err := getAllFilenames(fs, fp)
			if err != nil {
				return nil, err
			}
			out = append(out, res...)
			continue
		}
		out = append(out, fp)
	}
	return
}

func migrateUp(m *migrate.Migrate) error {
	log.Println("Migrate up started")

	if err := m.Up(); err != nil {
		log.Println(err)
		return err
	}

	log.Println("Migrate up finished")

	return nil
}
