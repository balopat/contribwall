#!/usr/bin/env bash

minikube start --container-runtime=containerd \
               --docker-opt containerd=/var/run/containerd/containerd.sock \
               --network-plugin=cni \
               --disk-size 30g \
               --vm-driver=hyperkit \
               --extra-config=kubelet.cadvisor-port=4194 \
               --alsologtostderr  \
               --v 10

minikube addons enable gvisor
