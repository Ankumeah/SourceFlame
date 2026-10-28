default: run

run *args:
  go build -o build/out ./cmd/sourceflame/
  ./build/out {{ args }}

ci ci = "":
  #! /bin/sh

  if test -n '{{ ci }}'; then
    echo 'CI = true'
    export CI=true
  fi

  ./.ci.sh
