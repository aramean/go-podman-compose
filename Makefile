SHELL := /bin/bash

default:
	@cd src && go run . --help

stop:
	@cd src && go run . stop

start:
	@cd src && go run . start

down:
	@cd src && go run . down

down%help:
	@cd src && go run . down -help

up:
	@cd src && go run . up

up%help:
	@cd src && go run . up -help

up%detach:
	@cd src && go run . up -d

version:
	@cd src && go run . version

version%help:
	@cd src && go run . version --help

ps:
	@cd src && go run . ps

help:
	@cd src && go run . --help