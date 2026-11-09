#! /bin/sh

curl "http://$HOST/api/$API_VERSION/r/user/test/git-log?ref=$1"
