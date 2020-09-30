#!/bin/sh

SCRIPT=$(readlink -f "$0")
# Absolute path this script is in
BASEDIR=$(dirname "$SCRIPT")
sh ${BASEDIR}/../../server/tools/installer/install.sh "$@" -c "${BASEDIR}"
