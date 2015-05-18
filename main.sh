#!/bin/bash

DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

export GOPATH="${DIR}"

export PATH="$PATH:${DIR}/bin/"

cd ${DIR}/src/photomosaic
