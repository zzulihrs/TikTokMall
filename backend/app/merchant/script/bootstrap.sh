#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/merchant"
exec "$CURDIR/bin/merchant"
