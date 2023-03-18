build_cli:
	go build -o image2css ./cmd/cli
	chmod +x image2css
build_wasm:
	GOOS=js GOARCH=wasm go build -o docs/main.wasm ./cmd/wasm
