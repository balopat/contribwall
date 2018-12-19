#!/usr/bin/env bash

minikube start --container-runtime=containerd \
               --disk-size 30g \
               --vm-driver=hyperkit \
               --extra-config=kubelet.cadvisor-port=4194 \
               --alsologtostderr  \
               --v 10
