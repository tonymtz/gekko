class setup($user, $source_folder, $golang_version, $npm_packages)
{
  include git

# Default path values
  Exec {
    path => [
      "/usr/local/bin",
      "/usr/local/sbin",
      "/usr/bin/",
      "/usr/sbin",
      "/bin",
      "/sbin",
      "/usr/local/node/node-default/bin"
    ]
  }

# Make sure our code directory has proper permissions
  file { "/vagrant/${source_folder}":
    ensure => "directory",
    mode   => 755
  }

# oh my zsh
  class { "ohmyzsh": }
  ohmyzsh::install { "root": }
  ohmyzsh::theme { ["root"]: theme => "dpoggi" }
  ohmyzsh::plugins { "root": plugins => "git" }

  class { 'golang':
    version   => $golang_version,
    workspace => '/usr/local/src/go',
  }

# cachefilesd
#   exec { 'cachefilesd install':
#     cwd     => "/vagrant/setup",
#     command => "bash ./cachefilesd.sh"
#   }
}