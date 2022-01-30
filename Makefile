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
	@if [ -d $(SRC_DIR)/$(GO_OUTPUT_DIR) ] ; then \
		cd $(SRC_DIR) && $(GO_BIN) build -o $(GO_OUTPUT_DIR)$(GO_OUTPUT_APP_NAME) ; \
		if [ -d $(GO_OUTPUT_DIR) ] ; then \
			echo -e "$(GREEN)Build for $(ARCH) architecture successfuly created.$(NC)" ; \
		fi \
	else \
		echo -e "$(RED)\"$(GO_OUTPUT_DIR)\" folder not found!$(NC)" ; \
		exit 1 ; \
	fi \

install:
	@if [ -d $(SRC_DIR)/$(GO_OUTPUT_DIR) ] ; then \
		dir=$(SRC_DIR)/$(GO_OUTPUT_DIR) ; \
		list="`echo "$$dir"/*`" ; \
		test "$$list" = "$$dir/*" && echo -e "$(RED)No build found in "$$dir" folder!$(NC)" && exit 1 || true ; \
 	else \
		echo -e "$(RED)\"$(GO_OUTPUT_DIR)\" folder not found!$(NC)" ; \
		exit 1 ; \
	fi \

	@while [ -z "$$CONTINUE" ] ; do \
		echo -e "\n$(CYAN)INSTALLATION ($(ARCH)):$(NC)\nThis will install $(WHITE)$(GO_OUTPUT_APP_NAME)$(NC) to /usr/local/bin\n" ; \
		read -r -p "Do you want to continue? [y/n]: " CONTINUE ; \
	done ; \
	[ $$CONTINUE = "y" ] || [ $$CONTINUE = "Y" ] || (echo -e "$(RED)Installation cancelled.$(NC)" ; exit 1;)

	@sudo cp $(SRC_DIR)/$(GO_OUTPUT_DIR)/$(GO_OUTPUT_APP_NAME) $(INSTALL_DIR)/$(INSTALL_NAME) ; \
	if [ -d $(INSTALL_DIR)/$(INSTALL_NAME) ] ; then \
		echo -e "$(GREEN)The installation was successful.$(NC)\n\nStart command:\n\n$(WHITE)$(NC)" ; \
	fi \

%:
	@: