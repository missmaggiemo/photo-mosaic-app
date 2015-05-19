#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

export GOPATH="${DIR}"

export PATH="$PATH:${DIR}/bin/"

cd ${DIR}/src/photomosaic

# go get github.com/astaxie/beego
# go get github.com/nfnt/resize

