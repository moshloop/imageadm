package pkg

import (
	"io/ioutil"
	"os"
)

var kickstart = `
install
cdrom
lang en_US.UTF-8
keyboard us
network --bootproto=dhcp --onboot=on --device=eth0
rootpw {{.Password}}
firewall --disabled
selinux --permissive
timezone UTC
unsupported_hardware
bootloader --location=mbr --append="net.ifnames=0 biosdevname=0"
text
skipx
zerombr
clearpart --all --initlabel
autopart
auth --enableshadow --passalgo=sha512 --kickstart
firstboot --disabled
reboot --eject
user --name={{.Username}} --plaintext --password {{.Password}}

%packages --nobase --ignoremissing --excludedocs --instLangs=en_US.utf8
# vagrant needs this to copy initial files via scp
openssh-clients
sudo
kernel-headers
kernel-devel
gcc
make
perl
selinux-policy-devel
wget
nfs-utils
net-tools
bzip2
deltarpm
-fprintd-pam
-intltool

# unnecessary firmware
-aic94xx-firmware
-alsa-firmware
-alsa-tools-firmware
-ivtv-firmware
-iwl100-firmware
-iwl105-firmware
-iwl135-firmware
-iwl1000-firmware
-iwl2000-firmware
-iwl2030-firmware
-iwl3160-firmware
-iwl3945-firmware
-iwl4965-firmware
-iwl5000-firmware
-iwl5150-firmware
-iwl6000-firmware
-iwl6000g2a-firmware
-iwl6000g2b-firmware
-iwl6050-firmware
-iwl7260-firmware
-iwl7265-firmware
%end

%post
# sudo
echo "%{{.Username}} ALL=(ALL) NOPASSWD: ALL" >> /etc/sudoers.d/{{.Username}}
chmod 0440 /etc/sudoers.d/{{.Username}}

#Enable hyper-v daemons only if using hyper-v virtualization
if [ $(virt-what) == "hyperv" ]; then
    yum -y install hyperv-daemons cifs-utils
    systemctl enable hypervvssd
    systemctl enable hypervkvpd
fi

# Since we disable consistent network naming, we need to make sure the eth0
# configuration file is in place so it will come up.
# Delete other network configuration first because RHEL/C7 networking will not
# restart successfully if there are configuration files for devices that do not
# exist.
rm -f /etc/sysconfig/network-scripts/ifcfg-e*
cat > /etc/sysconfig/network-scripts/ifcfg-eth0 << _EOF_
TYPE=Ethernet
PROXY_METHOD=none
BROWSER_ONLY=no
BOOTPROTO=dhcp
DEFROUTE=yes
IPV4_FAILURE_FATAL=no
IPV6INIT=yes
IPV6_AUTOCONF=yes
IPV6_DEFROUTE=yes
IPV6_FAILURE_FATAL=no
IPV6_ADDR_GEN_MODE=stable-privacy
NAME=eth0
DEVICE=eth0
ONBOOT=yes
_EOF_
%end
`

func SaveKickstart(vars Variables) func() {
	err := ioutil.WriteFile("ks.cfg", []byte(Interpolate(kickstart, vars)), os.ModePerm)
	if err != nil {
		panic(err)
	}

	return func() {
		os.Remove("ks.cfg")
	}
}
