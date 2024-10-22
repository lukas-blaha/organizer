# BACKEND_BINARY=backendApp

up:
	@echo "Starting containers..."
	docker compose up -d
	@echo "Containers started!"

up_build:
	@echo "Stopping containers"
	docker compose down
	@echo "Building and starting containers..."
	docker compose up --build -d
	@echo "Docker images built and containers started!"

down:
	@echo "Stopping docker containers"
	docker compose down
	@echo "Done!"

build:
	@echo "Building organizer images"
	cd ../backend && docker build --platform linux/arm64 -t lukasblaha/organizer-backend:latest --push .
	cd ../reminder && docker build --platform linux/arm64 -t lukasblaha/organizer-reminder:latest --push .
	@echo "All images built and published"

deploy: build
	@echo "Deploying services into home-k3s"
	@kubectl delete \
		-n organizer \
		--kubeconfig=/Users/lukas/.kube/home-env.config \
		$(shell kubectl get pods \
			-n organizer \
			-o name \
			-l app=backend \
			--kubeconfig=/Users/lukas/.kube/home-env.config) \
		$(shell kubectl get pods \
			-n organizer \
			-o name \
			-l app=reminder \
			--kubeconfig=/Users/lukas/.kube/home-env.config) 1>/dev/null
	@echo "Containers started"
