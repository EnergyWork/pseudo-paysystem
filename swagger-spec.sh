#!/bin/bash

file=swagger.yaml

if [ -e "$file" ]; then
  rm swagger.yaml
fi
exec D:\\docker\\Swagger\\swagger generate spec -o ./swagger.yaml --scan-models