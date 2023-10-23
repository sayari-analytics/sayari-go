include .env
export

push-test-readme-guides:
	rdme docs docs/go_sdk.md --key=$K
