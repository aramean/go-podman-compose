default:
	@go run . --help

stop:
	@go run . stop

start:
	@go run . start

down:
	@go run . down

down%help:
	@go run . down -help

up:
	@go run . up

up%help:
	@go run . up -help

up%detach:
	@go run . up -d

version:
	@go run . version

version%help:
	@go run . version --help

ps:
	@go run . ps

help:
	@go run . --help