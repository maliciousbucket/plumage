#!/usr/bin/env bash


curl --silent --retry 3 "https://update.k3s.io/v1-release/channels/stable" | egrep -o '/v[^ ]+"' | sed -E 's/\/|\"//g' | sed -E 's/\+/\-/'

