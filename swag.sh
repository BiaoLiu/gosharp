#!/bin/bash

echo "start gen swagger doc"

if [ ! -d docs ]; then
  mkdir docs
fi

path="./docs/swagger.json"

swagger generate spec -o  $path
echo "finish gen swagger doc"

if [[ $1 = serve ]]; then
  echo "start swagger serve"
  swagger generate spec -o $path &&
  swagger serve -F=swagger $path
fi
