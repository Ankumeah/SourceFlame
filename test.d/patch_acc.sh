#! /bin/sh

curl \
  -X PATCH \
  -H "Authorization: Basic $(printf user:pass | base64)" \
  -d '{"email":"email"}' \
  http://$HOST/api/$API_VERSION/acc
