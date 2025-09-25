#!/bin/bash

# Cleanup script for Admin Service CMDB
echo "🧹 Cleaning up Admin Service CMDB from Kubernetes..."

# Delete all resources
kubectl delete -f k8s/06-ingress.yaml
kubectl delete -f k8s/05-service.yaml
kubectl delete -f k8s/04-admin-service.yaml
kubectl delete -f k8s/03-postgres.yaml
kubectl delete -f k8s/02-configmap.yaml
kubectl delete -f k8s/01-secrets.yaml
kubectl delete -f k8s/00-namespace.yaml

echo "✅ Cleanup completed!"