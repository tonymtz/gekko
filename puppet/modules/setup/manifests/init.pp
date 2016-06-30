class setup($user, $psql_password)
{
  include git

  # Default path values
  Exec {
    path => [
      '/usr/local/bin',
      '/usr/local/sbin',
      '/usr/bin/',
      '/usr/sbin',
      '/bin',
      '/sbin'
    ]
  }

  # Make sure our code directory has proper permissions
  file { '/vagrant':
    ensure => 'directory',
    mode   => 755
  }

  #####################
  # OH MY ZSH SECTION #
  #####################

  class { 'ohmyzsh': }
  ohmyzsh::install { 'root': }
  ohmyzsh::theme { ['root']: theme => 'dpoggi' }
  ohmyzsh::plugins { 'root': plugins => 'git' }

  ##################
  # GOLANG SECTION #
  ##################

  class { 'golang':
    version   => '1.6',
    workspace => '/usr/local/src/go',
  }

  file_line { 'adding_gopath':
    path => '/home/vagrant/.bashrc',
    line => 'GOPATH=/vagrant/'
  }

  file_line { 'adding_path':
    path => '/home/vagrant/.bashrc',
    line => 'PATH=$PATH:$GOPATH/bin:/usr/local/go/bin'
  }

  ####################
  # DATABASE SECTION #
  ####################

  # global PostgreSQL settings
  class { 'postgresql::globals':
    encoding => 'UTF8',
    version => '9.3',
    pg_hba_conf_defaults => false
  }

  class { 'postgresql::server':
    listen_addresses => '*',
    postgres_password => $psql_password,
    require => Class['postgresql::globals']
  }

  # create db + user
  postgresql::server::db { 'gekko_db':
    user     => 'gekko',
    password => postgresql_password('gekko', $psql_password),
  }

  # pg_hba rules (access lists)
  postgresql::server::pg_hba_rule { 'all local connections identified by user':
    type        => 'local',
    database    => 'all',
    user        => 'all',
    address     => ' ',
    auth_method => 'ident',
  }

  postgresql::server::pg_hba_rule { 'allow remote connections with password':
    type        => 'host',
    database    => 'all',
    user        => 'all',
    address     => '0.0.0.0/0',
    auth_method => 'md5',
  }

  # PostgreSQL password
  file {'.pgpass-vagrant':
    path    => '/home/vagrant/.pgpass',
    ensure  => present,
    mode    => 0600,
    content => 'localhost:5432:gekko_db:gekko:${psql_password}',
    owner  => 'vagrant',
    group  => 'vagrant',
  }

  # initialize the content of your new database
  # exec { 'populate_postgresql':
  #   command => '/usr/bin/psql -d testdb -U testuser -h localhost -p 5432 --no-password < /vagrant/psql-db/psql-dump.sql',
  #   path    => '/usr/vagrant/', # there is .pgpass  file
  #   user    => 'vagrant',
  #   logoutput => true,
  #   require => [
  #     File['.pgpass-vagrant'],
  #     Postgresql::Server::Db['gekko_db'],
  #     Postgresql::Server::Pg_hba_rule['all local connections identified by user'],
  #     Postgresql::Server::Pg_hba_rule['allow remote connections with password'],
  #   ]
  # }
}