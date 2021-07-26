PUBLIC_IMAGE_NAME=pthomison/message-board
PRIVATE_IMAGE_NAME=localhost:9267/message-board
CLUSTER_IMAGE_NAME=k3d-message-board:5000/message-board
IMAGE_TAG=$(shell git rev-parse --short HEAD)
# IMAGE_TAG=$(shell find . -type f -print0 | grep -v dist | sort -z | xargs -0 sha1sum | sha1sum | cut -f 1 -d ' ')

gen-tag:
	awk -v min=1 -v max=9999999 'BEGIN{srand(); print int(min+rand()*(max-min+1))}' > .tag

create-cluster:
	k3d registry create message-board -p 9267
	k3d cluster create message-board --registry-use k3d-message-board:9267 -v /dev/mapper:/dev/mapper 

delete-cluster:
	k3d cluster delete message-board
	k3d registry delete message-board

build: gen-tag
	goreleaser release --skip-publish --snapshot --rm-dist
	skopeo copy --dest-tls-verify=false docker-daemon:${PUBLIC_IMAGE_NAME}:latest docker://${PRIVATE_IMAGE_NAME}:$(shell cat ./.tag)

deploy: build
	kubectl get ns message-board || kubectl create ns message-board
	cd k8s && kustomize edit set image ${CLUSTER_IMAGE_NAME}:$(shell cat ./.tag)
	kubectl --cluster="k3d-message-board" apply -k ./k8s --wait=true

port-forward: deploy
	kubectl -n message-board rollout status deploy/message-board
	sleep 1
	kubectl port-forward -n message-board service/message-board 8080

clean-deploy:
	kubectl --cluster="k3d-message-board" delete -k ./k8s
