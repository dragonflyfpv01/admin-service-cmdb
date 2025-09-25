#!/bin/bash
# apply.sh

set -e

for f in k8s-*.yaml; do
  echo "Applying $f..."
  kubectl apply -f "$f"
done
