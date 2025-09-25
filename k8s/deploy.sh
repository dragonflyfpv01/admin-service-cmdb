#!/bin/bash

# Deploy script for Admin Service CMDB
echo "🚀 Deploying Admin Service CMDB to Kubernetes..."

# Check if kubectl is available
if ! command -v kubectl &> /dev/null; then
    echo "❌ kubectl is not installed. Please install kubectl first."
    exit 1
fi

# Build Docker image (optional - uncomment if needed)
echo "📦 Building Docker image..."
# docker build -t admin-service:latest .

# Apply Kubernetes manifests in order
echo "🔧 Applying Kubernetes manifests..."

kubectl apply -f k8s/00-namespace.yaml
echo "✅ Namespace created"

kubectl apply -f k8s/01-secrets.yaml
echo "✅ Secrets created"

kubectl apply -f k8s/02-configmap.yaml
echo "✅ ConfigMap created"

kubectl apply -f k8s/03-postgres.yaml
echo "✅ PostgreSQL deployed"

# Wait for PostgreSQL to be ready
echo "⏳ Waiting for PostgreSQL to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/postgres-deployment -n admin-service

kubectl apply -f k8s/04-admin-service.yaml
echo "✅ Admin Service deployed"

kubectl apply -f k8s/05-service.yaml
echo "✅ Services created"

kubectl apply -f k8s/06-ingress.yaml
echo "✅ Ingress created"

# Wait for admin service to be ready
echo "⏳ Waiting for Admin Service to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/admin-service-deployment -n admin-service

echo ""
echo "🎉 Deployment completed successfully!"
echo ""
echo "📋 Service Information:"
echo "Namespace: admin-service"
echo "NodePort: http://localhost:30080"
echo "Ingress: http://admin-service.local (add to /etc/hosts if using locally)"
echo ""
echo "📊 Check status:"
echo "kubectl get pods -n admin-service"
echo "kubectl get services -n admin-service"
echo "kubectl logs -f deployment/admin-service-deployment -n admin-service"