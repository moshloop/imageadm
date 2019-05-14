package pkg

func init() {
	ISO["debian9"] = iso{
		Name:     "debian-9",
		URL:      "https://saimei.ftp.acc.umu.se/debian-cd/current/amd64/iso-cd/debian-9.9.0-amd64-netinst.iso",
		Checksum: "d4a22c81c76a66558fb92e690ef70a5d67c685a08216701b15746586520f6e8e",
		BootCommand: []string{
			"<esc><wait>",
			"install <wait>",
			// " preseed/file=/floppy/preseed.cfg<wait>",
			" preseed/url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg <wait>",
			" debian-installer=en_US.UTF-8 <wait>",
			" auto <wait>",
			" locale=en_US.UTF-8 <wait>",
			" kbd-chooser/method=us <wait>",
			" keyboard-configuration/xkb-keymap=us <wait>",
			" netcfg/get_hostname={{ .Hostname }} <wait>",
			" netcfg/get_domain=local <wait>",
			" fb=false <wait>",
			" debconf/frontend=noninteractive <wait>",
			" console-setup/ask_detect=false <wait>",
			" console-keymaps-at/keymap=us <wait>",
			"<enter><wait>",
		},
	}

	ISO["debian8"] = iso{
		Name:     "debian-8",
		URL:      "https://cdimage.debian.org/cdimage/archive/8.10.0/amd64/iso-cd/debian-8.10.0-amd64-netinst.iso",
		Checksum: "896cc42998edf65f1db4eba83581941fb2a584f2214976432b841af96b17ccda",
		BootCommand: []string{
			"<esc><wait>",
			"install <wait>",
			// " preseed/file=/floppy/preseed.cfg<wait>",
			" preseed/url=http://{{ .HTTPIP }}:{{ .HTTPPort }}/preseed.cfg <wait>",
			" debian-installer=en_US.UTF-8 <wait>",
			" auto <wait>",
			" locale=en_US.UTF-8 <wait>",
			" kbd-chooser/method=us <wait>",
			" keyboard-configuration/xkb-keymap=us <wait>",
			" netcfg/get_hostname={{ .Hostname }} <wait>",
			" netcfg/get_domain=local <wait>",
			" fb=false <wait>",
			" debconf/frontend=noninteractive <wait>",
			" console-setup/ask_detect=false <wait>",
			" console-keymaps-at/keymap=us <wait>",
			"<enter><wait>",
		},
	}

	ISO["centos7"] = iso{
		Name:     "centos7",
		URL:      "http://mirrors.kernel.org/centos/7.6.1810/isos/x86_64/CentOS-7-x86_64-DVD-1810.iso",
		Checksum: "6d44331cc4f6c506c7bbe9feb8468fad6c51a88ca1393ca6b8b486ea04bec3c1",
		BootCommand: []string{
			"<up><wait><tab> text ks=http://{{ .HTTPIP }}:{{ .HTTPPort }}/ks.cfg<enter><wait>",
		},
	}
	ISO["ubuntu1804"] = iso{
		Name:     "ubuntu-bionic",
		Checksum: "a2cb36dc010d98ad9253ea5ad5a07fd6b409e3412c48f1860536970b073c98f5",
		URL:      "http://cdimage.ubuntu.com/ubuntu/releases/18.04/release/ubuntu-18.04.2-server-amd64.iso",
		BootCommand: []string{
			"<esc><wait>",
			"<esc><wait>",
			"<enter><wait>",
			"/install/vmlinuz<wait>",
			" auto<wait>",
			" console-setup/ask_detect=false<wait>",
			" console-setup/layoutcode=us<wait>",
			" console-setup/modelcode=SKIP<wait>",
			" debconf/frontend=noninteractive<wait>",
			" debian-installer=en_US<wait>",
			" fb=false<wait>",
			" initrd=/install/initrd.gz<wait>",
			" kbd-chooser/method=us<wait>",
			" keyboard-configuration/layout=USA<wait>",
			" keyboard-configuration/variant=USA<wait>",
			" locale=en_US<wait>",
			" netcfg/get_domain=vm<wait>",
			" netcfg/get_hostname={{.Hostname}}<wait>",
			" noapic<wait>",
			" preseed/file=/floppy/preseed.cfg<wait>",
			" -- <wait>",
			"<enter><wait>",
		},
	}

	ISO["ubuntu1604"] = iso{
		Name:     "ubuntu-xenial",
		Checksum: "16afb1375372c57471ea5e29803a89a5a6bd1f6aabea2e5e34ac1ab7eb9786ac",
		URL:      "http://releases.ubuntu.com/16.04/ubuntu-16.04.6-server-amd64.iso",
		BootCommand: []string{
			"<enter><wait><f6><esc><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
			"<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
			"<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
			"<bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs><bs>",
			"/install/vmlinuz<wait>",
			" auto<wait>",
			" console-setup/ask_detect=false<wait>",
			" console-setup/layoutcode=us<wait>",
			" console-setup/modelcode=SKIP<wait>",
			" debconf/frontend=noninteractive<wait>",
			" debian-installer=en_US<wait>",
			" fb=false<wait>",
			" initrd=/install/initrd.gz<wait>",
			" kbd-chooser/method=us<wait>",
			" keyboard-configuration/layout=USA<wait>",
			" keyboard-configuration/variant=USA<wait>",
			" locale=en_US<wait>",
			" netcfg/get_domain=vm<wait>",
			" netcfg/get_hostname={{.Hostname}}<wait>",
			" noapic<wait>",
			" preseed/file=/floppy/preseed.cfg<wait>",
			" -- <wait>",
			"<enter><wait>",
		},
	}
}
