include ../../common.mk

build-in-docker: clean
	docker run --rm \
		-v $$PWD/../..:$(GO_REPO_ROOT) \
		-w $(GO_REPO_ROOT)/plugins/repo \
		$(BUILD_IMAGE) \
		bash -c "make build" || exit $$?

build: commands subcommands
subcommands: subcommands/create subcommands/default subcommands/destroy

commands: **/**/commands.go
	go build -a -o commands src/commands/commands.go

subcommands/create: **/**/**/create.go
	go build -a -o subcommands/create src/subcommands/create/create.go

subcommands/default: **/**/**/default.go
	go build -a -o subcommands/default src/subcommands/default/default.go

subcommands/destroy: **/**/**/destroy.go
	go build -a -o subcommands/destroy src/subcommands/destroy/destroy.go

clean:
	rm -rf commands subcommands/create subcommands/default subcommands/destroy

src-clean:
	rm -rf .gitignore src vendor
