module lms

go 1.19

replace lms/db => ./db

replace lms/db/models => ./db/models

replace lms/utils => ./utils

replace lms/controllers => ./controllers

require (
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	gorm.io/gorm v1.24.3
	lms/controllers v0.0.0-00010101000000-000000000000
	lms/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.2.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	gorm.io/driver/postgres v1.4.6 // indirect
	lms/db v0.0.0-00010101000000-000000000000 // indirect
	lms/db/models v0.0.0-00010101000000-000000000000 // indirect
)
