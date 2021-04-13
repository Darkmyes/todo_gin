package main

import "todo_gin/pkg/migrations"

func main() {
	migrations.Migrate()
}
