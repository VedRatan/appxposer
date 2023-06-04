# appxposer
A simple k8s controller to expose your application to the extenal world


## Steps to deploy this controller in your k8s cluster

- Clone the repo in your local machine
- Make sure you have your k8s cluster running
- run `go build`
- build your docker image by running `docker build -t <your docker user name>/appxposer:1.0 .`
- push the docker image to your remote docker hub repo `docker push <your docker user name>/appxposer:1.0`
- now just change your image inside the deployment.yaml file at line 22 by `<your docker user name>/appxposer:1.0`
- run `make deploy` to deploy this controller to the k8s cluster
- to delete this controller from your cluster run `make delete`


### Drop a star, if found interesting :)
