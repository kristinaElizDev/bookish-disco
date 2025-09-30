#!/usr/bin/env bash
set -e

if [ -z ${GITHUB_TOKEN} ]; then
    echo "GITHUB_TOKEN not set. Unable to set up goproxy." 1>&2
    exit 1
fi 

if [ -f ${HOME}/.netrc ] && grep -q "goproxy.githubapp.com" "${HOME}/.netrc"; then
    echo ".netrc already contains goproxy config"
    exit 0
fi

# Attempt to fallback to GOPROXY_TOKEN if GITHUB_TOKEN is a server-to-server token that isn't usable by goproxy.
# This fallback is useful for Codespace prebuilds.
if [[ "${GITHUB_TOKEN}" == "ghs_"* ]]; then
    if [ -z ${GOPROXY_TOKEN} ]; then
        echo "Server-to-server GITHUB_TOKEN detected. Could not find a GOPROXY_TOKEN value to fallback to. Unable to set up goproxy."
        exit 1
    else
        echo "Server-to-server GITHUB_TOKEN detected. Using GOPROXY_TOKEN instead."
        GITHUB_TOKEN=${GOPROXY_TOKEN}
    fi
fi

echo "machine goproxy.githubapp.com login nobody password ${GITHUB_TOKEN}" >> ${HOME}/.netrc
echo "setup ${HOME}/.netrc"