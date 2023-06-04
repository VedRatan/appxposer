# appxposer
A simple k8s controller to expose your application to the extenal world


## Steps to deploy this controller in your k8s cluster

- Clone the repo in your local machine
- Make sure you have your k8s cluster running
- run `go build`
- run `make deploy` to deploy this controller to the k8s cluster
- to delete this controller from your cluster run `make delete`
