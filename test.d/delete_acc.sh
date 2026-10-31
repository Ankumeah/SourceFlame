#! /bin/sh

curl \
  -X DELETE \
  -H "Authorization: Basic $(printf user:pass | base64)" \
  http://$HOST/api/$API_VERSION/acc
