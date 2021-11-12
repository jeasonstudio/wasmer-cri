package oci

import (
	"context"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

// PushFromFile push to registry from file in file system
func (c *Client) PushFromFile(ref, modulePath string, middlewares ...OptPullPush) error {
	filename := filepath.Base(modulePath)
	ctx := context.Background()

	log.WithFields(log.Fields{
		"ref":      ref,
		"path":     modulePath,
		"filename": filename,
	}).Debug("PushFromFile")

	fileContent, err := ioutil.ReadFile(modulePath)
	if err != nil {
		log.WithError(err).Fatal("Failed to read Webassembly file from file-system")
		return err
	}

	_, err = c.Push(ctx, ref, filename, fileContent, middlewares...)
	if err != nil {
		log.WithError(err).Fatal("Failed to push Webassembly")
		return err
	}
	return nil
}
