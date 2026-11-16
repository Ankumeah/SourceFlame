#! /bin/sh

curl -o - "http://$HOST/r/user/test/info/refs?service=git-upload-pack"
