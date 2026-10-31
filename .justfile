default: run
set dotenv-filename := "env.justfile"

run *args:
  CGO_ENABLED=0 go build -o build/out ./cmd/sourceflame/
  ./build/out {{ args }}

ci ci = "":
  #! /bin/sh

  if test -n '{{ ci }}'; then
    echo 'CI = true'
    export CI=true
  fi

  ./.ci.sh
