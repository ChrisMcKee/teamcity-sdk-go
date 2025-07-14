#!/bin/sh

echo "Removing old data"
rm -rf /data/teamcity_server/datadir/*
echo "Copying fresh data"
cp -a /test-data/* /data/teamcity_server/datadir/

exec /run-services.sh