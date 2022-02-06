SHELL := /bin/bash
ARGS = $(filter-out $@,$(MAKECMDGOALS))
ARCH = $(shell uname -p)
SRC_DIR = src
INSTALL_DIR = /usr/local/bin/
INSTALL_NAME = podman-compose
GO_OUTPUT_DIR = bin
GO_OUTPUT_APP_NAME = podman-compose
GO_BIN = go

CYAN=\033[1;36m
GREEN=\033[1;32m
RED=\033[1;31m
WHITE=\033[1;37m
NC=\033[0m

default: build

build:
	@mkdir -p $(SRC_DIR)/$(GO_OUTPUT_DIR)
	@cd $(SRC_DIR) && $(GO_BIN) build -o $(GO_OUTPUT_DIR)/$(GO_OUTPUT_APP_NAME) \
	&& echo -e "$(GREEN)Build for $(ARCH) architecture successfuly created.$(NC)" \
	|| echo -e "$(RED)Error: Build not successful.$(NC)"

install:
	@mkdir -p $(SRC_DIR)/$(GO_OUTPUT_DIR)
	@if [ ! -e $(SRC_DIR)/$(GO_OUTPUT_DIR)/$(GO_OUTPUT_APP_NAME) ]; then \
		echo -e "$(RED)No build found in \"$(GO_OUTPUT_DIR)\" folder!$(NC)" ; \
		exit 1 ; \
	fi \

	@while [ -z "$$CONTINUE" ] ; do \
		echo -e "\n$(CYAN)INSTALLATION ($(ARCH)):$(NC)\nThis will install $(WHITE)$(GO_OUTPUT_APP_NAME)$(NC) to /usr/local/bin\n" ; \
		read -r -p "Do you want to continue? [y/n]: " CONTINUE ; \
	done ; \
	[ $$CONTINUE = "y" ] || [ $$CONTINUE = "Y" ] || (echo -e "$(RED)Installation cancelled.$(NC)" ; exit 1;)

	@sudo cp $(SRC_DIR)/$(GO_OUTPUT_DIR)/$(GO_OUTPUT_APP_NAME) $(INSTALL_DIR)/$(INSTALL_NAME) ; \
	if [ -e $(INSTALL_DIR)/$(INSTALL_NAME) ] ; then \
		echo -e "\n\n$(GREEN)The installation was successful.$(NC)\n\nStart command:\n\n$(WHITE)$(INSTALL_NAME)$(NC)\n" ; \
	fi \

%:
	@: