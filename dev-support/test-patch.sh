#!/bin/bash

ROOT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )"/.. && pwd )"
GOCACHE=/go-cache
GOPATH="/go"
export GOCACHE GOPATH

function check_go_fmt {
    result=`find ./ -name "*.go" -exec goimports -l {} \;`
    if [ -n "${result}" ]; then
        echo "Go format [formatter: goimports] error within the following files:" >> $ROOT_DIR/.phabricator-comment
        printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
}

function check_go_lint {
    result=`golint .`
    if [ -n "${result}" ]; then
        echo "Go lint error:" >> $ROOT_DIR/.phabricator-comment
        printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
}


function check_go_vet {
    result=`go vet **/*.go 2>&1`
    if [ -n "${result}" ]; then
        echo "Go vet error:" >> $ROOT_DIR/.phabricator-comment
        printf "%s\n" "${result}" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
}

function check_build {
    rm -rf $ROOT_DIR/build && mkdir -p $ROOT_DIR/build
    pushd $ROOT_DIR/build >/dev/null
    cmake -DGOPATH=${GOPATH} -GNinja .. && ninja && ctest -V
    if [ $? -ne 0 ]; then
        echo "build failed" >> $ROOT_DIR/.phabricator-comment
        exit 1
    fi
    popd > /dev/null
}

cat /dev/null > $ROOT_DIR/.phabricator-comment
mkdir -p $ROOT_DIR/build
cd ${ROOT_DIR}

echo "Checking Go repos"

HAS_GO=$(/usr/bin/git diff HEAD --name-only --diff-filter=ACMRT | grep -E ".*\.go$")
if [ -n "${HAS_GO}" ];then
    cd ${ROOT_DIR}/golang
    check_go_fmt
    check_go_lint
    check_build
fi

echo "Build succeed." > $ROOT_DIR/.phabricator-comment
