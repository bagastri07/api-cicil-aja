migrate:
	go run ./migrator/gorm_migrator.go

mock:
	mockery --dir=./api --case=underscore --all --recursive --keeptree --disable-version-string