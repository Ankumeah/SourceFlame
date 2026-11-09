#! /bin/sh

curl "http://$HOST/api/$API_VERSION/r/user/test/blob/README.md?ref=$1"
