# Arfcom Microservices App

This repo provides the source code for a sample microservice I created for my friends at Arfcom. It is the basis for my presentation, the *Hitchhiker's Guide to Microservices, Containers, and Kubernetes*. We take this source and compile/run it locally. Next, after a discussion of the move from microservices to containers, we create a Docker image and run that locally. Finally, we push the image to Docker Hub and deploy it to a live Kubernetes cluster in the Cloud and discuss the benefits of using a container orchestration platform.

The service itself is trivial. The goal isn't to provide an in-depth discussion of writing microservices but rather to provide something we can work with to see real-world examples of the features, functionalities, and capabilities in each of these three areas.

I have also provided Kubernetes manifests to demonstrate three key things:

- deploying a single container in a Pod
- deploying an application container with an Init container
- deploying an application as part of a Deployment for resiliency, better upgrades, etc.