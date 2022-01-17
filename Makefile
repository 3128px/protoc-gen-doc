name := protoc-gen-doc

# Currently we resolve it using which. But more sophisticated approach is to use infer GOROOT.
go     := $(shell which go)
goarch := $(shell $(go) env GOARCH)
goexe  := $(shell $(go) env GOEXE)
goos   := $(shell $(go) env GOOS)

all_go_sources := $(wildcard cmd/*/*.go *.go extensions/*/*.go)
main_go_sources := cmd/$(name)/main.go $(wildcard $(filter-out %_test.go,$(all_go_sources)))
binary_$(name)_sources := cmd/$(name)/main.go cmd/$(name)/flags.go # TODO(dio): Make sure we have automated way for getting these files.

current_binary_path := build/$(name)_$(goos)_$(goarch)
current_binary      := $(current_binary_path)/$(name)$(goexe)

linux_platforms       := linux_amd64 linux_arm64
non_windows_platforms := darwin_amd64 darwin_arm64 $(linux_platforms)
windows_platforms     := windows_amd64

build: $(current_binary) ## Build the protoc-gen-doc binary

clean:
	@rm -rf $(current_binary)

build/$(name)_%/$(name): $(main_go_sources)
	@$(call go-build,$@,$(binary_$(name)_sources))

build/$(name)_%/$(name).exe: $(main_go_sources)
	$(call go-build,$@,$(binary_$(name)_sources))

go-arch = $(if $(findstring amd64,$1),amd64,arm64)
go-os   = $(if $(findstring .exe,$1),windows,$(if $(findstring linux,$1),linux,darwin))
define go-build
	@CGO_ENABLED=0 GOOS=$(call go-os,$1) GOARCH=$(call go-arch,$1) $(go) build \
		-ldflags "-s -w" \
		-o $1 $2
endef
