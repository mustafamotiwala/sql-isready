GREEN = \033[0;32m
BLUE = \033[0;34m
NOCOLOR = \033[0m

help: Makefile
	@echo "$(GREEN)"
	@echo "\t Choose a command to run:"
	@echo "$(NOCOLOR)"
	@sed -n "s/^##//p" $< | column -t -s ':' |  sed -e "s/^/ /"
	@echo

## build – Build the standard executable for our project.
build:
	@echo "$(BLUE)"
	@echo "\t>  Building binary..."
	@echo "$(NOCOLOR)"
	go build -o bin/sql-isready cmd/sql-isready/main.go

.PHONY: help
all: help