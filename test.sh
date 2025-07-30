#!/opt/homebrew/bin/bash
source host_init.sh

./terraformer import infoblox \
  --resources=record_a \
  --path-output=generated \
  --compact=false

