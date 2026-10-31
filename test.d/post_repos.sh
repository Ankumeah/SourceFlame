#! /bin/sh

curl \
  -X POST \
  -H "Authorization: Basic $(printf user:pass | base64)" \
  http://"$HOST"/api/"$API_VERSION"/repos/test
