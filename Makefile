serve:
	LOGLEVEL=debug go run debug/main.go

crictl_config = crictl.yml

test_create_container:
	crictl --config $(crictl_config) create POD fixtures/container-config.yaml fixtures/pod-config.yaml
test_list_images:
	crictl --config $(crictl_config) images
