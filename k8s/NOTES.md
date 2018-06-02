#Building container image and running it locally
- eval $(docker-machine env default)
- go build -o main .
- docker build -t scratch-app -f Dockerfile .
- docker run --publish 7080:7080 -it scratch-app
- curl localhost:7080/Arya

#Tagging the image and uploading to hub.docker.com
- docker images
- docker tag 0cf5c5dbd51f ideepesh/k8s-demo:v2
- docker push ideepesh/k8s-demo
- docker tag 0adflkadfjlf ideepesh/k8s-demo:latest
- docker push ideepesh/k8s-demo

#Running the image in kubernetes for single pod
- minikube start
- kubectl create -f config/k8sDemoOnePod.yml
- kubectl get pods
- kubectl get pod scratch-app
- kubectl describe pod scratch-app
- Expose service:
  - kubectl expose pod scratch-app --type=NodePort --name scratch-app-pod-svc
  - kubectl get service
  - kubectl describe service scratch-app-pod-svc
  - minikube service scratch-app-pod-svc --url
  - URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya
- kubectl exec scratch-app -- ANY_COMMAND
- kubectl delete service scratch-app-pod-svc
- kubectl delete pod scratch-app

#Running the image in kubernetes with replication controller
- minikube start
- kubectl create -f config/k8sDemoReplicas.yml
- kubectl get replicationcontrollers
- kubectl describe rc scratch-app-controller
- kubectl get pods
- kubectl scale --replicas=3 -f config/k8sDemoReplicas.yml
- kubectl scale --replicas=4 replicationcontroller scratch-app-controller
- Expose service
  - kubectl expose rc scratch-app-controller --type=NodePort --name scratch-app-rc-svc
  - kubectl get service
  - kubectl describe service scratct-app-rc-svc
  - minikube service scratch-app-rc-svc --url
  - URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya
- kubectl delete service scratch-app-rc-svc
- kubectl delete replicationcontroller scratch-app-controller

#Running the image in kubernetes with deployment
- minikube start
- kubectl create -f config/k8sDemoDeployment.yml
- kubectl get deployments
- kubectl describe deployment scratch-app-deployment
- kubectl get replicaset
- kubectl describe rs XXXXX
- kubectl get pods
- kubectl 
- Expose service
  - kubectl expose deployment scratch-app-deployment --type=NodePort --name scratch-app-dep-svc
  - kubectl get service
  - kubectl describe service scratct-app-dep-svc
  - minikube service scratch-app-dep-svc --url
  - URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya
- kubectl set image deployment/scratch-app-deployment k8s-demo=ideepesh/k8s-demo:v1
- kubectl rollout status deployment/scratch-app-deployment
- URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya
- kubectl rollout undo deployment/scratch-app-deployment
- kubectl rollout status deployment/scratch-app-deployment
- URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya
- kubectl delete service scratch-app-dep-svc
- kubectl delete deployment scratch-app-deployment

#Creating service from config file
- Create deployment as mentioned above
- kubectl create -f config/k8sDemoSvc.yml
- minikube service scratch-app-dep-svc --url
- URL_RETRIEVED_WITH_PREVIOUS_COMMAND/Arya

#Using Node selectors:
- kubectl get nodes
- kubectl get nodes --show-labels
- kubectl create -f config/k8sDemoDepNodeSelector.yml
- kubectl get po
- kubectl get deploy
- kubectl describe po
- kubectl label nodes minikube hardware=commodity
- kubectl get nodes --show-labels
- kubectl get po
- kubectl label nodes minikube hardware=commodity
- kubectl get po
- kubectl delete po ANY_ONE_POD
- kubectl get po

#Debugging tricks
- Simple port forwarding on localhost to pod:
  - kubectl port-forward scratch-app 6080:7080
  - localhost:6080/Arya
  - Checks if the pod is working