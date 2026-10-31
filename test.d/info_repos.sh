#! /bin/sh

curl \
  -o - \
  'http://"$HOST"/api/"$API_VERSION"/r/user/test/info/refs?service=git-upload-pack'
