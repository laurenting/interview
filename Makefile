# `make build` to build the binary
.PHONY: build
build:
	cd frontend/ && npm run build
	wails build

# `make dev` to build the binary
.PHONY: dev
dev:
	cd frontend/ && npm run build
	wails dev -debounce 500