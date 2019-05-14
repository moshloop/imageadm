package pkg

import (
	"io/ioutil"
	"os"
	"strings"
)

func common(vars Variables) string {
	return Interpolate(`
choose-mirror-bin mirror/http/proxy string
d-i keymap select us
d-i debian-installer/framebuffer boolean false
d-i debconf/frontend select noninteractive
d-i base-installer/kernel/override-image string linux-server
d-i clock-setup/utc boolean true
d-i clock-setup/utc-auto boolean true
d-i clock-setup/ntp boolean false
d-i finish-install/reboot_in_progress note
d-i pkgsel/include string {{.Packages}}
d-i pkgsel/install-language-support boolean false
d-i pkgsel/update-policy select none
d-i pkgsel/upgrade select full-upgrade
d-i time/zone string UTC

popularity-contest popularity-contest/participate boolean false
`, vars)

}

func lvm(vars Variables) string {
	return Interpolate(`
d-i partman-auto/method string lvm
d-i partman-auto-lvm/guided_size string max
d-i partman-lvm/confirm boolean true
d-i partman-lvm/confirm_nooverwrite boolean true
d-i partman-auto/choose_recipe select atomic
d-i partman/choose_partition select finish
d-i partman/confirm boolean true
d-i partman/confirm_nooverwrite boolean true
d-i partman/confirm_write_new_label string true
d-i grub-installer/only_debian boolean true
d-i grub-installer/with_other_os boolean true
d-i grub-installer/bootdev  string /dev/vda
`, vars)
}

func user(vars Variables) string {
	// d-i passwd/root-login boolean false
	// d-i passwd/root-password-again password vagrant
	// d-i passwd/root-password password vagrant
	return Interpolate(`
d-i passwd/root-login boolean false
d-i passwd/user-fullname string {{ .Username }}
d-i passwd/username string {{ .Username }}
d-i passwd/user-password password {{ .Password }}
d-i passwd/user-password-again password {{ .Password }}
d-i passwd/user-default-groups {{ .Username }} sudo
d-i user-setup/allow-password-weak boolean true
d-i user-setup/encrypt-home boolean false
`, vars)
}

func getPreseed(vars Variables, iso iso) string {

	if iso.Name == "debian-8" {
		return common(vars) + lvm(vars) + user(vars) + `
d-i apt-setup/use_mirror boolean true
d-i mirror/country string manual
d-i mirror/http/directory string /debian
d-i mirror/http/hostname string httpredir.debian.org
d-i mirror/http/proxy string
d-i preseed/late_command string sed -i '/^deb cdrom:/s/^/#/' /target/etc/apt/sources.list
apt-cdrom-setup apt-setup/cdrom/set-first boolean false
apt-mirror-setup apt-setup/use_mirror boolean true
tasksel tasksel/first multiselect standard, ssh-server`
	}

	if iso.Name == "debian-9" {
		return common(vars) + lvm(vars) + user(vars) + `
tasksel tasksel/first multiselect standard, ssh-server
d-i apt-setup/use_mirror boolean false
apt-cdrom-setup apt-setup/cdrom/set-first boolean false
d-i cdrom-checker/nextcd boolean false`
	}

	if strings.HasPrefix(iso.Name, "ubuntu") {
		return common(vars) + lvm(vars) + user(vars) + `
tasksel tasksel/first multiselect standard, ubuntu-server
`
	}
	return common(vars) + lvm(vars) + user(vars)
}

func SavePreseed(vars Variables) func() {
	err := ioutil.WriteFile("preseed.cfg", []byte(getPreseed(vars, vars.ISO)), os.ModePerm)
	if err != nil {
		panic(err)
	}

	return func() {
		os.Remove("preseed.cfg")
	}

}
