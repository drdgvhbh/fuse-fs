package main

import (
	"log"
	"os"

	"fuse-filesystem/pkg"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	_ "bazil.org/fuse/fs/fstestutil"
	"github.com/urfave/cli"
)

func main() {
	var mountDirectory string

	app := cli.NewApp()
	app.Name = "Fuse Filesystem"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "mnt-dir, m",
			Usage:       "Mount Directory",
			Required:    true,
			Destination: &mountDirectory,
		},
	}

	app.Action = func(context *cli.Context) error {
		c, err := fuse.Mount(
			mountDirectory,
			fuse.FSName("helloworld"),
			fuse.Subtype("hellofs"),
			fuse.LocalVolume(),
			fuse.VolumeName("Hello world!"),
		)
		if err != nil {
			return err
		}
		defer c.Close()

		err = fs.Serve(c, pkg.FS{})
		if err != nil {
			return err
		}

		<-c.Ready
		if err := c.MountError; err != nil {
			return err
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
