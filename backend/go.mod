module lms

go 1.19

replace lms/db => ./db

replace lms/utils => ./utils

replace lms/controllers => ./controllers

require (
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo v3.3.10+incompatible
	lms/controllers v0.0.0-00010101000000-000000000000
	lms/utils v0.0.0-00010101000000-000000000000
)

require (
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
)
