.SILENT:

REST=cmd/app/main.go
BUILDFOLDER=./cmd/build

#
run:
	go run $(REST)

# build program and move it in build directory
build:
	go build cmd/app/main.go && move main.exe $(BUILDFOLDER)

buildAndRun: build
	cd $(BUILDFOLDER) && start main.exe