SHELL := /bin/bash
ARGS = $(filter-out $@,$(MAKECMDGOALS))

default:
	@cd src && go run . --help

stop:
	@cd src && go run . stop

stop%help:
	@cd src && go run . stop --help

stop%nginx:
	@cd src && go run . stop nginx

stop%mariadb:
	@cd src && go run . stop mariadb

start:
	@cd src && go run . start

start%nginx:
	@cd src && go run . start nginx

start%mariadb:
	@cd src && go run . start mariadb

start%all:
	@cd src && go run . start nginx
	@cd src && go run . start mariadb

down:
	@cd src && go run . down

down%help:
	@cd src && go run . down --help

up:
	@cd src && go run . up

up%help:
	@cd src && go run . up --help

up%detach:
	@cd src && go run . up -d

version:
	@cd src && go run . version

version%help:
	@cd src && go run . version --help

ps:
	@cd src && go run . ps

ps%help:
	@cd src && go run . ps --help

help:
	@cd src && go run . --help

%:
	@: