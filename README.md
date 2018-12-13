# Contributor wall 

A sample project for displaying the contributors of a github project. 
To try it out you'll need:


1. [Docker](http://docker.io)
1. [skaffold](http://skaffold.dev)
1. [minikube](http://github.com/kubernetes/minikube)   
1. kubectl
1. [get an API token from github](https://github.com/settings/tokens) 
 
Steps for running with minikube with docker: 

1. `minikube start && minikube tunnel`
1. `git clone github.com/balopat/contribwall`
1. `printf TOKEN > backend/svc-contributors/token` with your TOKEN
1. run `skaffold dev`
1. get the ExternalIP field from `kubectl get svc frontend` and navigate there with your browser! 


Steps for running with minikube with contained & gvisor and external registry (`myrepo`):

1. `git clone github.com/balopat/contribwall`
1. `printf TOKEN > backend/svc-contributors/token` with your TOKEN
1. `bash minikube-gvisor.sh`
1. start `minikube tunnel` in a separate terminal session and keep it running 
1. change `skaffold.yaml` to have `push: true`
1. run `skaffold dev --default-repo myrepo` with your `myrepo`
1. get the ExternalIP field from `kubectl get svc frontend` and navigate there with your browser! 
