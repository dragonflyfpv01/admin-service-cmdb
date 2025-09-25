# Admin Service CMDB - Kubernetes Deployment

## Prerequisites

- Kubernetes cluster (local or cloud)
- kubectl configured
- Docker (for building images)
- NGINX Ingress Controller (for Ingress support)

## Quick Start

### 1. Build Docker Image (if needed)
```bash
docker build -t admin-service:latest .
```

### 2. Deploy to Kubernetes
```bash
chmod +x k8s/deploy.sh
./k8s/deploy.sh
```

### 3. Access the Application

#### Via NodePort
```bash
curl http://localhost:30080/health
```

#### Via Ingress (add to /etc/hosts)
```bash
echo "127.0.0.1 admin-service.local" >> /etc/hosts
curl http://admin-service.local/health
```

## Manual Deployment

If you prefer to deploy manually:

```bash
# Apply manifests in order
kubectl apply -f k8s/00-namespace.yaml
kubectl apply -f k8s/01-secrets.yaml
kubectl apply -f k8s/02-configmap.yaml
kubectl apply -f k8s/03-postgres.yaml
kubectl apply -f k8s/04-admin-service.yaml
kubectl apply -f k8s/05-service.yaml
kubectl apply -f k8s/06-ingress.yaml
```

## Configuration

### Environment Variables (ConfigMap)
- `DB_HOST`: PostgreSQL host (default: postgres-service)
- `DB_PORT`: PostgreSQL port (default: 5432)
- `DB_NAME`: Database name (default: mydb)
- `APP_ENV`: Application environment (default: production)
- `APP_PORT`: Application port (default: 3000)

### Secrets
- `DB_USER`: Database username (base64 encoded)
- `DB_PASSWORD`: Database password (base64 encoded)
- `JWT_SECRET`: JWT secret key (base64 encoded)

### To update secrets:
```bash
echo -n "newpassword" | base64
kubectl patch secret admin-service-secret -n admin-service -p '{"data":{"DB_PASSWORD":"bmV3cGFzc3dvcmQ="}}'
```

## Monitoring

### Check Pod Status
```bash
kubectl get pods -n admin-service
```

### View Logs
```bash
# Admin Service logs
kubectl logs -f deployment/admin-service-deployment -n admin-service

# PostgreSQL logs
kubectl logs -f deployment/postgres-deployment -n admin-service
```

### Check Services
```bash
kubectl get services -n admin-service
```

### Check Ingress
```bash
kubectl get ingress -n admin-service
```

## Scaling

### Scale Admin Service
```bash
kubectl scale deployment admin-service-deployment --replicas=3 -n admin-service
```

### Update Image
```bash
kubectl set image deployment/admin-service-deployment admin-service=admin-service:v2 -n admin-service
```

## Troubleshooting

### Admin Service Not Starting
1. Check if PostgreSQL is ready:
```bash
kubectl get pods -n admin-service
kubectl logs deployment/postgres-deployment -n admin-service
```

2. Check admin service logs:
```bash
kubectl logs deployment/admin-service-deployment -n admin-service
```

3. Check environment variables:
```bash
kubectl describe deployment admin-service-deployment -n admin-service
```

### Database Connection Issues
1. Test connection from admin service pod:
```bash
kubectl exec -it deployment/admin-service-deployment -n admin-service -- nc -z postgres-service 5432
```

2. Check database credentials:
```bash
kubectl get secret admin-service-secret -n admin-service -o yaml
```

### Ingress Not Working
1. Check if NGINX Ingress Controller is installed:
```bash
kubectl get pods -n ingress-nginx
```

2. Check ingress status:
```bash
kubectl describe ingress admin-service-ingress -n admin-service
```

## Cleanup

To remove all resources:
```bash
chmod +x k8s/cleanup.sh
./k8s/cleanup.sh
```

## API Endpoints

Once deployed, the following endpoints are available:

### Public Endpoints
- `GET /health` - Health check
- `POST /admin/login` - Admin login
- `POST /admin/signup` - Admin signup

### Protected Endpoints (Require JWT)
- `GET /admin/profile` - Get admin profile
- `GET /admin/users` - Get all users (admin only)
- `GET /admin/infra-components` - Get infra components (paginated)
- `GET /admin/infra-components/all` - Get all infra components
- `GET /admin/infra-components/pending` - Get pending infra components
- `POST /admin/infra-components` - Create new infra component
- `PUT /admin/infra-components/status` - Update infra component status
- `PUT /admin/infra-components` - Update infra component

## Security Notes

- Database passwords are stored in Kubernetes Secrets
- JWT tokens are used for authentication
- All admin endpoints require valid JWT tokens
- CORS is configured for cross-origin requests
- Resource limits are set for all containers