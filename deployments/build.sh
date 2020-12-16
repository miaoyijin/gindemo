#!/bin/bash

APP_PKG_NAME="$1"
BASE_DIR=$(cd "$(dirname "$0")";pwd)

APP_BINARY="startup"

#declare -A APP_PKGNAME_TO_MAIN
#APP_PKGNAME_TO_MAIN["api"]="cmd/api/main.go"

function build() {
    [ -d "${BASE_DIR}/bin" ] && rm -rf "${BASE_DIR}/bin"

    LDFLAGS="-X \"gindemo/helpers/version.GoVersion=$(go version)\""
		echo $LDFLAGS

    #export GOOS=linux

    export GO111MODULE=on
    export GOPROXY=https://goproxy.cn
    go mod vendor
    go mod verify
    go build -ldflags "${LDFLAGS}" -o ./bin/"${APP_BINARY}" ../main.go
}

function dev() {
    [ -d "${BASE_DIR}/bin" ] && rm -rf "${BASE_DIR}/bin"

    LDFLAGS="-X \"gindemo/helpers/version.GoVersion=$(go version)\""
    		echo $LDFLAGS


		#echo $LDFLAGS

    #export GOOS=linux

    go mod vendor
    go mod verify
    go build -ldflags "${LDFLAGS}" -o ./bin/"${APP_BINARY}" ./main.go
}

#check
function dispatcher(){
    if [ $# -lt 1 ] ;then
        usage
        exit -1
    fi

    local args=$1
    case "$args" in
    force)
        force_Instance
        ;;
    start)
        startup_Instance
        ;;
    status)
        status_Instance
        ;;
    *)
        usage
        ;;
    esac
}

#dispatcher "$@"
build
