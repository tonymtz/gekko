$user = hiera("user")
$source_folder = hiera("source_folder")
$golang_version = hiera("golang_version")
$npm_packages = hiera_array("npm_packages")

file { "/etc/motd":
  content => "
*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*
=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=
     Node/MongoDB Dev Server
- OS:      Ubuntu 14.04 (trusty 64)
- Node:    ${golang_version}
- IP:      10.0.0.10
*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*
=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=*=
\n"
}

class { "setup":
  user            => $user
}