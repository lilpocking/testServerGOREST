# start makefile in silent mode
.SILENT:

MAINGOFILE=cmd/app/main.go
BUILD=./cmd/build/app

# parametrs for build
BUILD_CGO_PARAM=1

#BUILD_OS
#windows
#linux
#darwin (for macOS systems)
BUILD_OS=windows

#BUILD_ARCH
# 32-bit systems
#386 
# 64-bit systems
#amd64
#arm64
BUILD_ARCH=amd64

run:
	go run $(MAINGOFILE)

build:
	go mod download && go build -o $(BUILD).exe $(MAINGOFILE)

buildWithParam:
	go mod download && CGO_ENABLED=$(BUILD_CGO_PARAM) GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH) go build -o $(BUILD)-$(BUILD_OS)-$(BUILD_ARCH).exe $(MAINGOFILE)

# perform start comand
runApp:
	start $(BUILD)