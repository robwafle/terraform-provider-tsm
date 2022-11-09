#!/bin/sh -e
tmp="$(mktemp)"
apt-get update 2>&1 | sed -En 's/.*NO_PUBKEY ([[:xdigit:]]+).*/\1/p' | sort -u > "${tmp}"
cat "${tmp}" | xargs  gpg --keyserver "hkps://keyserver.ubuntu.com:443" --recv-keys  # to /usr/share/keyrings/*
cat "${tmp}" | xargs -L 1 sh -c ' gpg --yes --output "/etc/apt/trusted.gpg.d/$1.gpg" --export "$1"' sh  # to /etc/apt/trusted.gpg.d/*
rm "${tmp}"
