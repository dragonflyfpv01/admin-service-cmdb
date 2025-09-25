#!/bin/bash
# delete.sh

set -e

for f in k8s-*.yaml; do
  echo "Deleting $f..."
  kubectl delete -f "$f" --ignore-not-found
done
