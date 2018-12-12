# Contributor wall 

A sample project for displaying the contributors of a github project. 
To try it out you'll need:


1. [Docker](http://docker.io)
1. [skaffold](http://skaffold.dev)
1. [minikube](http://github.com/kubernetes/minikube)   
1. [get an API token from github](https://github.com/settings/tokens) 
 
Steps: 

1. `minikube start && minikube tunnel`
1. `git clone github.com/balopat/contribwall`
1. replace TOKEN in `backend/svc-contributors/svc/contributors/contributors.go` with your token
1. run `skaffold dev`

