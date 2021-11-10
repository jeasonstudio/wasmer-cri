serve:
	go run debug/main.go

crictl_config = crictl.yml

test_create_container:
	crictl --config $(crictl_config) create POD fixtures/container-config.yaml fixtures/pod-config.yaml
