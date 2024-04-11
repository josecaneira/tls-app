

***Modified version of Google Example TLS Web Application in Go***

Key features/changes from original:  
* Docker file changes:  
  Replaced Google base image "gcr.io/distroless/base-debian11" by Microsoft base image "mcr.microsoft.com/devcontainers/base:bullseye".  
  Replaced runtime user/group from "nonroot:nonroot" to "nobody:nobody".  
* Code changes:  
  Added domain display on page.  
  Included log and display of direct calls made to application as well, original Google code only had "X-Forwarded-For".  


CHANGELOG:  
* 11/04/2023:
  - Public release
  - Switch to "mcr.microsoft.com/cbl-mariner/distroless/minimal:2.0-nonroot" base image for size optimization.
		
TODO: Sugestions welcomed.  


Original Google source: https://github.com/googlecloudplatform/kubernetes-engine-samples/tree/394477ca2a99/quickstarts/hello-app-tls  

Built image available at: https://hub.docker.com/r/jcaneira/tls-app  

Docker pull command: docker pull jcaneira/tls-app  

AKS ready deployment: kubectl apply -f https://raw.githubusercontent.com/josecaneira/tls-app/main/tls-app.yaml  
Expects you have a TLS secret named "tls-app-ssl-cert" residing on same namespace.  
Exposing ingress not included.  

__________________________________________________________________________________________________________________________________________________________________________________________________________
