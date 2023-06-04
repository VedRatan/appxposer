deploy:
	kubectl apply -f manifests/namespace.yaml
	kubectl apply -f manifests/sa.yaml
	kubectl apply -f manifests/clusterrole.yaml
	kubectl apply -f manifests/clusterrolebinding.yaml
	kubectl apply -f manifests/deployment.yaml

delete:
	kubectl delete -f manifests/namespace.yaml
	kubectl delete -f manifests/clusterrole.yaml
	kubectl delete -f manifests/clusterrolebinding.yaml

