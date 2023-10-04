minikube start
kubectl apply -f ./kubernetes-deploy/db-persistent-volume.yaml
kubectl apply -f ./kubernetes-deploy/db-volume-claim.yaml
kubectl apply -f ./kubernetes-deploy/db-configmap.yaml
kubectl apply -f ./kubernetes-deploy/db-deployment.yaml
kubectl apply -f ./kubernetes-deploy/db-service.yaml