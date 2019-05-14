package cmd

import (
	"fmt"

	"github.com/moshloop/imageadm/pkg"
	"github.com/spf13/cobra"
)

var (
	//Build command
	Build = cobra.Command{
		Use:   "build",
		Short: "Build a new image",
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {

			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			image, _ := cmd.Flags().GetString("image")

			fmt.Println("Building " + image)
			packer := pkg.NewPacker(username, password, pkg.ISO[image])

			packer.Build()
		},
	}
)

func init() {
	Build.Flags().StringP("username", "u", "imageadm", "")
	Build.Flags().StringP("password", "p", "imageadm", "")
	Build.Flags().StringP("image", "i", "ubuntu1804", "One of: debian8, debian9, centos7, ubuntu1604, ubuntu1804")
}
