ifeq ($(OS),Windows_NT)
    SHELL := powershell.exe #change shell for windows
    .SHELLFLAGS := -Command
    ending := exe
else
    ending := out
endif


cmdpath = ./cmd
files = $(cmdpath)/main.go


run:
	go run $(files)

build:
	go build -o ./builds/main.$(ending) $(files)

fmt: 
	go fmt ./cmd ./internal/db ./internal/shortlinkgen ./internal/api