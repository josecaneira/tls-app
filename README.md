

***Modified version of Google Example TLS Web Application in Go***

Key features/changes:  
* Docker file changes:  
  Replaced Google base image "gcr.io/distroless/base-debian11" by Microsoft base image "mcr.microsoft.com/devcontainers/base:bullseye".  
  Replaced runtime user/group from "nonroot:nonroot" to "nobody:nobody".  
* Code changes:  
  Added domain display on page.  
  Included log and display of direct calls made to application as well, original Google code only had "X-Forwarded-For".  


CHANGELOG:  
* 11/04/2023:
  - Public release.
		
TODO: Sugestions welcomed.  


Original Google source: https://github.com/googlecloudplatform/kubernetes-engine-samples/tree/394477ca2a99/quickstarts/hello-app-tls  

Built image available at: https://hub.docker.com/r/jcaneira/tls-app  

Docker pull command: docker pull jcaneira/tls-app  

__________________________________________________________________________________________________________________________________________________________________________________________________________
