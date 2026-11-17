#! /bin/sh

curl \
  -X PATCH \
  -H "Authorization: Basic $(printf user:pass | base64)" \
  -d '{"email":"new@email.com"}' \
  http://$HOST/api/acc

echo "Check mail and run ./post_acc_otp.sh <otp> where otp is the otp you got"
