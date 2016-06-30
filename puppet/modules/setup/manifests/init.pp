class setup($user)
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
      "/sbin"
    ]
  }

# Make sure our code directory has proper permissions
  # file { "/vagrant":
  #   ensure => "directory",
  #   mode   => 755
  # }

# oh my zsh
  class { "ohmyzsh": }
  ohmyzsh::install { "root": }
  ohmyzsh::theme { ["root"]: theme => "dpoggi" }
  ohmyzsh::plugins { "root": plugins => "git" }

  class { 'golang':
    version   => '1.6',
    workspace => '/usr/local/src/go',
  }

  # file { "/etc/environment":
  #   content => inline_template("GOPATH=/vagrant/")
  # }

  # file { "/etc/environment":
  #   content => inline_template("export PATH=$PATH:$GOPATH/bin:/usr/local/go/bin")
  # }

  file_line { 'adding_gopath':
    path => '/home/vagrant/.bashrc',
    line => 'GOPATH=/vagrant/'
  }

  file_line { 'adding_path':
    path => '/home/vagrant/.bashrc',
    line => 'PATH=$PATH:$GOPATH/bin:/usr/local/go/bin'
  }

# cachefilesd
#   exec { 'cachefilesd install':
#     cwd     => "/vagrant/setup",
#     command => "bash ./cachefilesd.sh"
#   }
}