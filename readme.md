![gekko](https://github.com/tonymtz/gekko/blob/master/static/gekko.png)

# Gekko

Sample project using [Echo](https://github.com/labstack/echo).

## Features
- `fasthttp` enabled by default
- OAuth2 with Google & Dropbox
- JWT for user authentication

## Quick Start
This lib requires:

- Golang v1.6
- sqlite3 v3.8.10.2

### Installation
```sh
$ go get
```

#### Database

Create database file using sqlite3 (you must have sqlite3 installed)
```sh
$ sqlite3 gekko.db
.databases
```

and then create the `user` table:

```sql
CREATE TABLE "user" (
    `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    `id_provider` TEXT NOT NULL,
    `display_name` TEXT NOT NULL,
    `email` TEXT NOT NULL,
    `profile_picture` TEXT NOT NULL,
    `role` INTEGER NOT NULL DEFAULT 1,
    `token` TEXT,
    `jwt` TEXT
);
```

### Configuration

Create your config file within `config` directory. Its name must match with the pattern `[env].conf`.
Recommended envs: `dev`, `test`.

Use `env.conf.sample` as base since it has the expected format. Just replace the values!

### Test
```sh
$ export GEKKO_ENV=test && go test
```

### Running
```sh
$ export GEKKO_ENV=[env] && go run main.go
```
