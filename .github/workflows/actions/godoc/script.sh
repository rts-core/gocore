#!/bin/bash

set -eo pipefail

# Make sure background processes are killed on script leave
trap "trap - SIGTERM && kill -- -$$" SIGINT SIGTERM EXIT

html_dir = $1

# find go mod and make sure we're using it's parent as our root.
cd "$(dirname "$(find . -name 'go.mod' | head -n 1)")" || exit 1

# clean and make our docs dir
rm -rf $html_dir
mkdir -p $html_dir

# host godoc
godoc -http=:8080 &
for (( ; ; )); do
  sleep 0.5
  if [[ $(curl -so /dev/null -w '%{http_code}' "http://localhost:8080/pkg/") -eq 200 ]]; then
    break
  fi
done

# get our html
cd $html_dir
wget -m -k -q -erobots=off --no-host-directories --no-use-server-timestamps http://localhost:8080

exit 0