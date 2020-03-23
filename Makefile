VERSION = `cat VERSION`
TAGS=-tags netgo
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
	LDFLAGS = '-extldflags "-static"'
endif

.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'; \
	echo "version: ${VERSION}"

.PHONY: build
build: generate			## Build binary
	@echo "+ $@"; \
	echo "Building"; \
	go build $(TAGS) -x -v -a -o dist/sampleapi; \
	echo "Finished building"

.PHONY: generate
generate:
	@echo "+ $@"
	cd cmd; go generate; cd -

.PHONY: clean
clean:				## Clean all artifacts
	@echo "+ $@"
	rm -fr dist

.PHONY: image
image: 	                        ## Build docker image
	@echo "+ $@"
	docker build -f Dockerfile -t sampleapi:latest ./

.PHONY: dockerrun
dockerrun:			## Run docker image
	@echo "+ $@"
	docker run --rm --name sampleapi -ti -p 8000:8000 sampleapi:latest 
