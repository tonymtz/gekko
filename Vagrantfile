# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.hostname = "gekko.dev"

  config.vm.network "private_network", ip: "10.0.0.10"
  config.ssh.forward_agent = true

  config.vm.synced_folder ".", "/vagrant", nfs: true, mount_options: ['rw', 'vers=3', 'tcp', 'fsc']

  config.vm.provider "virtualbox" do |vb|
    host = RbConfig::CONFIG['host_os']
    if host =~ /darwin/
      cpus = `sysctl -n hw.ncpu`.to_i
      mem = `sysctl -n hw.memsize`.to_i / 1024 / 1024 / 4
    elsif host =~ /linux/
      cpus = `nproc`.to_i
      mem = `grep 'MemTotal' /proc/meminfo | sed -e 's/MemTotal://' -e 's/ kB//'`.to_i / 1024 / 4
    else
      cpus = 2
      mem = 1024
    end

    vb.customize ["modifyvm", :id, "--memory", mem]
    vb.customize ["modifyvm", :id, "--cpus", cpus]
  end

  config.vm.provision "puppet" do |puppet|
    puppet.hiera_config_path = "puppet/hiera.yaml"
    puppet.manifests_path = "puppet/manifests"
    puppet.manifest_file  = "default.pp"
    puppet.module_path = "puppet/modules"
    # puppet.options = "--verbose --debug"
  end
end
