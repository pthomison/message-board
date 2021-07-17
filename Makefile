PUBLIC_IMAGE_NAME=pthomison/message-board
PRIVATE_IMAGE_NAME=localhost:9267/message-board
CLUSTER_IMAGE_NAME=k3d-message-board:5000/message-board
IMAGE_TAG=$(shell git rev-parse --short HEAD)


create-cluster:
	k3d registry create message-board -p 9267
	k3d cluster create message-board --registry-use k3d-message-board:9267 

delete-cluster:
	k3d cluster delete message-board
	k3d registry delete message-board

build:
	goreleaser release --skip-publish --snapshot --rm-dist
	skopeo copy --dest-tls-verify=false docker-daemon:${PUBLIC_IMAGE_NAME}:latest docker://${PRIVATE_IMAGE_NAME}:${IMAGE_TAG}

deploy: build
	kubectl get ns message-board || kubectl create ns message-board
	cd k8s && kustomize edit set image ${CLUSTER_IMAGE_NAME}:${IMAGE_TAG}
	kubectl --cluster="k3d-message-board" apply -k ./k8s

clean-deploy:
	kubectl --cluster="k3d-message-board" delete -k ./k8s