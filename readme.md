![gekko](https://github.com/tonymtz/gekko/blob/master/static/gekko.png)

# Gekko

Sample project using [Echo](https://github.com/labstack/echo).

## Features
- `fasthttp` enabled by default
- OAuth2 with Google & Dropbox

## Quick Start
This lib requires Golang v1.6.

```sh
$ export GEKKO_ENV=dev && go run main.go
```

### Installation
```sh
$ go get
```

#### Database

```sql
CREATE TABLE "user" (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `id_provider` TEXT NOT NULL,
    `display_name` TEXT NOT NULL,
    `email` TEXT NOT NULL,
    `profile_picture` TEXT NOT NULL,
    `role` INTEGER NOT NULL DEFAULT 1,
    `token` TEXT
)
```

### Configuration
(pending...)

### Test
```sh
$ export GEKKO_ENV=test && go test
```
