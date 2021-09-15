package cri

// ensureSandboxImageExists pulls the image when it's not present.
func (c *CRIManager) ensureSandboxImageExists(imageRef string) error {
	// TODO: needs image service
	// _, _, _, err := c.ImageMgr.CheckReference(ctx, imageRef)
	// if err == nil {
	// 	return nil
	// }
	// if errtypes.IsNotfound(err) {
	// 	err = c.ImageMgr.PullImage(ctx, imageRef, nil, bytes.NewBuffer([]byte{}))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to pull sandbox image %q: %v", imageRef, err)
	// 	}
	// 	return nil
	// }
	// return fmt.Errorf("failed to check sandbox image %q: %v", imageRef, err)
	return nil
}
