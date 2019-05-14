package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type BeforeBootHandler func(vars Variables) func()

type iso struct {
	Name        string
	URL         string
	Checksum    string
	BootCommand []string
	Preseed     string
	BeforeBoot  BeforeBootHandler
}

var ISO = make(map[string]iso)

type Provisioner struct {
	EnvironmentVars []string `json:"environment_vars"`
	ExecuteCommand  string   `json:"execute_command"`
	Scripts         []string `json:"scripts"`
	Type            string   `json:"type"`
}

type Variables struct {
	BootCommandPrefix string `json:"boot_command_prefix"`
	CleanupPause      string `json:"cleanup_pause"`
	Cpus              string `json:"cpus"`
	CustomScript      string `json:"custom_script"`
	Desktop           string `json:"desktop"`
	DiskSize          string `json:"disk_size"`
	FtpProxy          string `json:"ftp_proxy"`
	Headless          string `json:"headless"`
	HTTPProxy         string `json:"http_proxy"`
	HTTPSProxy        string `json:"https_proxy"`
	Locale            string `json:"locale"`
	Memory            string `json:"memory"`
	NoProxy           string `json:"no_proxy"`
	Preseed           string `json:"preseed"`
	RsyncProxy        string `json:"rsync_proxy"`
	Hostname          string `json:"hostname"`
	Password          string `json:"ssh_password"`
	Username          string `json:"ssh_username"`
	Update            string `json:"update"`
	Comment           string `json:"_comment"`
	VMName            string `json:"vm_name"`
	IsoChecksum       string `json:"iso_checksum"`
	IsoChecksumType   string `json:"iso_checksum_type"`
	IsoURL            string `json:"iso_url"`
	Packages          string
	ISO               iso `json:"-"`
}

type Builder struct {
	Type              string     `json:"type"`
	BootCommand       []string   `json:"boot_command"`
	DiskSize          string     `json:"disk_size"`
	FloppyFiles       []string   `json:"floppy_files"`
	Headless          string     `json:"headless"`
	Format            string     `json:"format"`
	SkipCompaction    bool       `json:"skip_compaction"`
	HTTPDirectory     string     `json:"http_directory"`
	IsoChecksum       string     `json:"iso_checksum"`
	IsoChecksumType   string     `json:"iso_checksum_type"`
	IsoUrls           []string   `json:"iso_urls"`
	OutputDirectory   string     `json:"output_directory"`
	ShutdownCommand   string     `json:"shutdown_command"`
	SSHPassword       string     `json:"ssh_password"`
	SSHUsername       string     `json:"ssh_username"`
	SSHWaitTimeout    string     `json:"ssh_wait_timeout"`
	BootKeyInterval   string     `json:"boot_key_interval"`
	UseDefaultDisplay string     `json:"use_default_display"`
	VMName            string     `json:"vm_name"`
	Qemuargs          [][]string `json:"qemuargs"`
}
type Packer struct {
	Comment      string        `json:"_comment"`
	Builders     []Builder     `json:"builders"`
	Provisioners []Provisioner `json:"provisioners"`
	Variables    Variables     `json:"variables"`
}

func NewPacker(username string, password string, iso iso) Packer {

	vars := Variables{
		Username: username,
		Password: password,
		DiskSize: fmt.Sprintf("%d", 16*1024),
		Cpus:     fmt.Sprintf("%d", 1),
		Locale:   "en_US",
		Packages: "openssh-server net-tools curl nano sudo bzip2 acpid cryptsetup zlib1g-dev wget dkms make nfs-common",
		Memory:   fmt.Sprintf("%d", 1024),
		Preseed:  "preseed.cfg",
		Hostname: username,
		ISO:      iso,
		Update:   "false",
	}
	return Packer{
		Variables: vars,
		Builders: []Builder{{
			Type:              "qemu",
			Qemuargs:          getQemuArgs(vars),
			BootCommand:       InterpolateStrings(iso.BootCommand, vars),
			DiskSize:          vars.DiskSize,
			UseDefaultDisplay: "true",
			Format:            "raw",
			Headless:          "true",
			SkipCompaction:    true,
			BootKeyInterval:   "100ms",
			SSHWaitTimeout:    "10000s",
			ShutdownCommand:   fmt.Sprintf("echo '%s' | sudo -S shutdown -P now", vars.Password),
			SSHUsername:       vars.Username,
			SSHPassword:       vars.Password,
			HTTPDirectory:     "./",
			VMName:            vars.Hostname,
			OutputDirectory:   "output-" + iso.Name,
			FloppyFiles:       []string{vars.Preseed},
			IsoChecksum:       iso.Checksum,
			IsoChecksumType:   "sha256",
			IsoUrls:           []string{path.Base(iso.URL), iso.URL},
		}},
	}
}

func (p Packer) Build() {
	cleanup := p.Variables.ISO.BeforeBoot(p.Variables)
	// cleanup := SavePreseed(p.Variables)
	defer cleanup()
	data, _ := json.Marshal(p)
	_ = os.Remove("packer.json")
	err := ioutil.WriteFile("packer.json", data, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer os.Remove("packer.json")
	Exec("PACKER_LOG=1 packer build packer.json")

}
