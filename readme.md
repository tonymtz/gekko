# Gekko :dragon:

Sample project using [Echo](https://github.com/labstack/echo).

## Features
- `fasthttp` enabled :rocket:
- ready to work with postgresql :floppy_disk:
- OAuth2 with Google & Dropbox :cop:
- JWT for user authentication :godmode:

## Installation

The overall installation will take around 5 minutes downloading dependencies, depending of you internet bandwidth.

Clone repo :octocat::
```sh
$ git clone git@github.com:tonymtz/gekko.git --recursive
```

Start vagrant machine and log in (you might need enter your administrator password) :milky_way::
```sh
$ vagrant up
$ vagrant ssh
```

Install golang dependencies :minidisc::
```sh
$ cd src/github.com/tonymtz/gekko
$ go get bitbucket.org/liamstask/goose/cmd/goose
$ go get -u github.com/kardianos/govendor
$ govendor sync
```

## Running :red_car:
```sh
$ go run main.go
```

:bell: By default, your server will be available through this url: [http://10.0.0.10:3000](http://10.0.0.10:3000)

## Working from your vagrant machine

- [Install atom](https://codeforgeek.com/2014/09/install-atom-editor-ubuntu-14-04/)
- [Install goodies to atom](http://marcio.io/2015/07/supercharging-atom-editor-for-go-development/)
- [Speed up x11 forwarding](http://xmodulo.com/how-to-speed-up-x11-forwarding-in-ssh.html)
