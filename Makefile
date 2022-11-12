# Based on that awesome makefile https://github.com/dunglas/symfony-docker/blob/main/docs/makefile.md#the-template

.DEFAULT_GOAL = default

.PHONY: default
default: build

##—— 📚 Help ——————————————————————————————————————————————————————————————
.PHONY: help
help: ## ❓ Dislay this help
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' \
		| sed -e 's/\[32m##——————————/[33m           /'  \
		| sed -e 's/\[32m##——/[33m ——/' \
		| sed -e 's/\[32m####/[34m                                 /' \
		| sed -e 's/\[32m###/[36m                                 /' \
		| sed -e 's/\[32m##\?/[35m /'  \
		| sed -e 's/\[32m##/[33m/'

##—— ️⚙️  Environments ——————————————————————————————————————————————————————
.PHONY: configure-dev-env
configure-dev-env: ## 🤖 Install required libraries for dev environment
configure-dev-env:
	go install github.com/posener/goreadme/cmd/goreadme@latest
.PHONY: configure-test-env
configure-test-env: ## 🤖 Install required libraries for test environment (golint, staticcheck, etc)
configure-test-env: configure-dev-env
configure-test-env:
	go install golang.org/x/lint/golint@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest

##—— 📝 Documentation —————————————————————————————————————————————————
.PHONY: build-doc
build-doc: ## 🗜️  Generate and enhance packages doc
build-doc: generate-doc enhance-doc

.PHONY: generate-doc
generate-doc: ## 🗜️  Generate packages doc
generate-doc:
	goreadme -constants -variabless -types -methods -functions -factories > DOC.md
	@find * -prune -type d -name "gh*" | while IFS= read -r d; do \
  		cd $$d; \
  		goreadme -constants -variabless -types -methods -functions -factories > README.md; \
  		cd ..; \
	done

.PHONY: enhance-doc
enhance-doc: ## 🗜️  Enhance packages doc
enhance-doc:
	# Add terraform style for raw blocks and fix links for sub-packages
	@sed -i "~" -E -e ':a' -e 'N' -e '$$!ba' -e 's/```(\n)(resource ")/```terraform\1\2/g' DOC.md
	@find * -prune -type d -name "gh*" | while IFS= read -r d; do \
		cd $$d; \
		sed -i "~" -E -e ':a' -e 'N' -e '$$!ba' -e 's/```(\n)(resource ")/```terraform\1\2/g' README.md; \
		sed -i "~" -E "s/]\((\/.+)\.go/](.\1.go/g" README.md; \
		cd ..; \
	done

##—— 🐹 Golang —————————————————————————————————————————————————
.PHONY: build
build: ## 🗜️  Build package
#### Use build_o="..." to specify build options
$(eval build_o ?=)
build:
	go build -v $(build_o)

.PHONY: verify
verify: ## 🗜️  Verify dependencies
verify:
	go mod verify

.PHONY: format
format: ## 🗜️  Format code with go fmt command
#### Use format_o="..." to specify format options
$(eval format_o ?=)
format:
	gofmt -w -s $(format_o) .


##—— 🧪️ Tests —————————————————————————————————————————————————————————————
.PHONY: test
test: ## 🏃 Launch all tests
test: test-vet test-lint test-staticcheck test-go

test-go: ## 🏃 Launch go test
#### Use gotest_o="..." to specify options
$(eval gotest_o ?=)
test-go:
	go test -v  $(gotest_o) ./...

test-vet: ## 🏃 Launch go vet
#### Use vet_o="..." to specify options
$(eval vet_o ?=)
test-vet:
	go vet $(vet_o) ./...

test-lint: ## 🏃 Launch go lint
#### Use lint_o="..." to specify options (-set_exit_status for instance)
$(eval lint_o ?=-set_exit_status)
test-lint:
	golint $(lint_o) ./...

test-staticcheck: ## 🏃 Launch staticcheck
#### Use staticcheck_o="..." to specify options
$(eval staticcheck_o ?=)
test-staticcheck:
	staticcheck $(staticcheck_o) ./...
