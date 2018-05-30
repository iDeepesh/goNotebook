Building container image and running it locally
- eval $(docker-machine env default)
- go build -o main .
- docker build -t scratch-app -f Dockerfile .
- docker run --publish 7080:7080 -it scratch-app
- curl localhost:7080/Arya

Tagging the image and uploading to hub.docker.com
- docker images
- docker tag 0cf5c5dbd51f ideepesh/k8s-demo
- docker push ideepesh/k8s-demo

Running the image in kubernetes
- minikube start
- kubectl create -f k8s-demo.yml
- kubectl get pods
- kubectl get pod scratch-app
- kubectl describe pod scratch-app
- kubectl port-forward scratch-app 7080:7080
- kubectl expose pod scratch-app --type=NodePort --name scratch-app-service
- kubectl get servic
- minikube service scratch-app-service --url
- <url>/Arya