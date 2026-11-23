#! /bin/sh

curl -X POST \
  --data '{"title":"Add example feature","body":"This is a sample pull request with dummy data.\n\nstuff\tsome does studd\n","url":"test-url.com"}' \
  -H "Authorization: Basic $(printf 'user:pass' | base64)" \
  "http://localhost:8000/r/user/test/pr"
