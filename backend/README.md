# LMS Backend

## Auto reload after save

Install **air** util.

`go install github.com/cosmtrek/air@latest`

Add your `go/bin` directory to **PATH** env.

Go bin path:

- Linux: `/home/<username>/go/bin`
- MacOs: `/Users/<username>/go/bin`

Type `air` command at project root directory.

## Migrating

Migrations will start automatically if env variable `MIGRATE` is set to `true`.

## API

**USERS GROUP**

`POST` */api/users/register*

`POST` */api/users/login*

`POST` */api/users/logout*

`GET` */api/users/get-user*

**BOOKS GROUP**

`GET` */api/books/get-book*

`POST` */api/users/add-book*
