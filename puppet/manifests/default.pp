$user = hiera('user')
$psql_password = hiera('psql_password')

file { '/etc/motd':
  content => "
*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*
=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=
     ~~~ RELEASE THE GEKKO ~~~
    Golang/Postgresql Dev Server
- Name:    gekko.dev
- OS:      Ubuntu 14.04 (trusty 64)
- IP:      10.0.0.10
*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*
=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=
\n"
}

class { "setup":
  user => $user,
  psql_password => $psql_password
}