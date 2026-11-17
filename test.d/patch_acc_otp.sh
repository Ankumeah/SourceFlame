#! /bin/sh

curl \
  -X PATCH \
  -H "Authorization: Basic $(printf user:pass | base64)" \
  -d "{\"email\":\"new@email.com\", \"otp\": $1}" \
  http://$HOST/api/acc
