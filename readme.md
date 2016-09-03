![gekko](https://github.com/tonymtz/gekko/blob/master/static/gekko.png)

# Gekko :dragon:

Sample project using [Echo](https://github.com/labstack/echo).

## Features
- :rocket: `fasthttp` enabled
- :floppy_disk: ready to work with postgresql
- :cop: OAuth2 with Google & Dropbox
- :godmode: JWT for user authentication

## Installation

The overall installation will take around 5 minutes downloading dependencies, depending of you internet bandwidth.

:octocat: Clone repo:
```sh
$ git clone git@github.com:tonymtz/gekko.git --recursive
```

:milky_way: Start vagrant machine and log in (you might need enter your administrator password):
```sh
$ vagrant up
$ vagrant ssh
```

:minidisc: Install golang dependencies:
```sh
$ cd src/github.com/tonymtz/gekko
$ go get -u github.com/kardianos/govendor
$ govendor sync
```

### Running :red_car:
```sh
$ go run main.go
```

:bell: By default, your server will be available through this url: [http://10.0.0.10:3000](http://10.0.0.10:3000)
