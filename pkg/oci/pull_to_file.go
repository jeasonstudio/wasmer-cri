package oci

import (
	"context"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// PullToFile Pull and write to file system
func (c *Client) PullToFile(ref, filename string, middlewares ...OptPullPush) (*Image, error) {
	ctx := context.Background()
	log.WithFields(log.Fields{
		"ref": ref,
	}).Debug("Pull image to file")

	image, err := c.Pull(ctx, ref, middlewares...)
	if err != nil {
		log.WithError(err).Fatal("Failed to pull image")
		return nil, err
	}

	err = ioutil.WriteFile(filename, image.Content, 0755)
	if err != nil {
		log.WithError(err).Fatal("Failed to write image into file-system")
		return nil, err
	}
	return image, nil
}
