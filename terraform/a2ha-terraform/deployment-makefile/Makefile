# Global defaults
TERRAFORM_PATH = $(shell pwd)/terraform
SORT_PY = $(shell pwd)/scripts/tf_var_sort.py
AIRGAP_SH = $(shell pwd)/scripts/bundle_creation.sh
SUP_KEY_SH = $(shell pwd)/scripts/sup-keys.sh
TMPOUT = '/tmp/make.out'
REPO_ROOT = $(shell pwd)
DATESTAMP := $(shell date +"%Y%m%d%H%M%S")
BACKENDAIB = $(TERRAFORM_PATH)/transfer_files/backend-$(DATESTAMP).aib
BACKENDAIB_TFVARS = $(shell pwd)/terraform/a2ha_aib_be.auto.tfvars
FRONTENDAIB = $(TERRAFORM_PATH)/transfer_files/frontend-$(DATESTAMP).aib
FRONTENDAIB_TFVARS = $(shell pwd)/terraform/a2ha_aib_fe.auto.tfvars
TMPFILE := $(shell mktemp)
VARFILES := $(shell find $(TERRAFORM_PATH) -name variables.tf -o -name inputs.tf)
HABITAT_TFVARS = $(shell pwd)/terraform/a2ha_habitat.auto.tfvars

# Help menu colors
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

# Habitat variables - can be passed in via VARNAME=XXX make <target>
BECHANNEL?=stable
BEVERSION?=latest
FECHANNEL?=current
FEVERSION?=latest

# Targets

# https://gist.github.com/prwhite/8168133#gistcomment-2278355
# any targets that have double '#' above them will be picked up and added to the menu
## Display this help text
help:
	@echo ''
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-20s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
	@echo ''
	@echo 'Overrides:'
	@printf "  ${YELLOW}%-20s${RESET} ${GREEN}%s${RESET}\n" "BECHANNEL|FECHANNEL" "Bldr channel to pull from for Backend|Frontend Components. ex. BECHANNEL=acceptance make airgap"
	@printf "  ${YELLOW}%-20s${RESET} ${GREEN}%s${RESET}\n" "BEVERSION|FEVERSION" "Semver of Backend|Frontend Components to pull ex. BEVERSION=0.1.63 make airgap"
	@echo ''

$(FRONTENDAIB):
	@echo "# Updating FE component package versions: $(FRONTENDAIB_TFVARS)" && \
		$(AIRGAP_SH) -t frontend -o $(FRONTENDAIB) -b $(BACKENDAIB) > $(FRONTENDAIB_TFVARS) && \
	  $$(hab pkg path core/coreutils)/bin/md5sum $(FRONTENDAIB) > $(FRONTENDAIB).md5 && \
	  $$(hab pkg path core/coreutils)/bin/md5sum $(BACKENDAIB) > $(BACKENDAIB).md5

$(HABITAT_TFVARS):
	@echo -e "# Generating habitat encryption keys in $(HABITAT_TFVARS)" && \
	  $(SUP_KEY_SH) > $(HABITAT_TFVARS)

## Create a frontend .aib
fe-bundle: $(FRONTENDAIB)

## Create Airgap bundles
airgap: fe-bundle

.PHONY: clean help

## Clean up any old manifest and airgap bundles
clean:
	@rm -f $(MANIFESTJSON) $(MANIFEST_TFVARS) $(BACKENDAIB_TFVARS) $(FRONTENDAIB_TFVARS) $(HABITAT_TFVARS) \
		$(TERRAFORM_PATH)/transfer_files/*.aib $(TERRAFORM_PATH)/transfer_files/*.md5

## Generate random Habitat Supervisor ring key and HTTP gateway bearer token
sup-keys: $(HABITAT_TFVARS)

## Generate habitat certificate for cluster internal secure communication 
habitat-sup-cert:
	bash scripts/cert.sh

## Force a re-push of the .aib files by changing their destination filename.
bump-aib:
	bash scripts/bumpaib.sh $(TERRAFORM_PATH)/terraform.tfvars

## Setup habitat keys and airgap bundles
setup: sup-keys fe-bundle habitat-sup-cert
