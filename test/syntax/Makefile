export DEBUG=TRUE

SHELL := /bin/bash
ARGS = $(filter-out $@,$(MAKECMDGOALS))
SRC_DIR = ../../
GO_ARG = run .

ENV_FILE = .env
YAML_FILE = docker-compose.yml

default:
	@cd $(SRC_DIR) && go $(GO_ARG)

stop:
	@cd $(SRC_DIR) && go $(GO_ARG) stop

stop%timeout:
	@cd $(SRC_DIR) && go $(GO_ARG) stop -t 10

stop%help:
	@cd $(SRC_DIR) && go $(GO_ARG) stop --help

stop%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) stop nginx

stop%error1:
	@cd $(SRC_DIR) && go $(GO_ARG) stop error1

stop%error2:
	@cd $(SRC_DIR) && go $(GO_ARG) stop error2

stop%mariadb:
	@cd $(SRC_DIR) && go $(GO_ARG) stop mariadb

start:
	@cd $(SRC_DIR) && go $(GO_ARG) start

start%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) start nginx

start%mariadb:
	@cd $(SRC_DIR) && go $(GO_ARG) start mariadb

start%all:
	@cd $(SRC_DIR) && go $(GO_ARG) start nginx
	@cd $(SRC_DIR) && go $(GO_ARG) start mariadb

restart:
	@cd $(SRC_DIR) && go $(GO_ARG) restart

restart%timeout:
	@cd $(SRC_DIR) && go $(GO_ARG) restart -t 20

restart%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) restart nginx

restart%nginxtime:
	@cd $(SRC_DIR) && go $(GO_ARG) restart nginx -t 20

down:
	@cd $(SRC_DIR) && go $(GO_ARG) down

kill:
	@cd $(SRC_DIR) && go $(GO_ARG) kill

kill%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) kill nginx

logs:
	@cd $(SRC_DIR) && go $(GO_ARG) logs

logs%mariadb:
	@cd $(SRC_DIR) && go $(GO_ARG) logs mariadb

logs%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) logs nginx

down%help:
	@cd $(SRC_DIR) && go $(GO_ARG) down --help

up:
	@cd $(SRC_DIR) && go $(GO_ARG) up

up%help:
	@cd $(SRC_DIR) && go $(GO_ARG) up --help

up%detach:
	@cd $(SRC_DIR) && go $(GO_ARG) up -d

version:
	@cd $(SRC_DIR) && go $(GO_ARG) version

version%help:
	@cd $(SRC_DIR) && go $(GO_ARG) version --help

ps:
	@cd $(SRC_DIR) && go $(GO_ARG) ps

ps%help:
	@cd $(SRC_DIR) && go $(GO_ARG) ps --help

ps%quiet:
	@cd $(SRC_DIR) && go $(GO_ARG) ps -q

pull:
	@cd $(SRC_DIR) && go $(GO_ARG) pull

pull%nginx:
	@cd $(SRC_DIR) && go $(GO_ARG) pull nginx

pull%mariadb:
	@cd $(SRC_DIR) && go $(GO_ARG) pull mariadb

help:
	@cd $(SRC_DIR) && go $(GO_ARG) --help

nocommand:
	@cd $(SRC_DIR) && go $(GO_ARG) nocommand

%:
	@: