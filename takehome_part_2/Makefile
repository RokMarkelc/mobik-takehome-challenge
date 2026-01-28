
MINIKUBE_IP = $(shell vagrant ssh -c "minikube ip")

.PHONY: start-vm destroy-vm provision-vm deploy-apps full-deploy ssh-vm ssh-pod logs check-pods check-services expose-frontend

start-vm: 
	vagrant up --no-provision

destroy-vm: 
	vagrant destroy -f

provision-vm: 
	vagrant provision

deploy-apps:
	vagrant ssh -c "cd /vagrant && ansible-playbook -i ansible/hosts.ini ansible/deploy-app.yml"

full-deploy: start-vm provision-vm deploy-apps expose-frontend
	@echo "✅ Deployment complete"

ssh-vm:
	vagrant ssh

ssh-pod:
	@echo "Usage: make ssh-pod service=service_name (postgres/backend/frontend)"
	vagrant ssh -c "kubectl exec -it svc/$(service) -- bash"

logs:
	@echo "Usage: make logs service=service_name (postgres/backend/frontend)"
	vagrant ssh -c "kubectl logs -l app=$(service) --tail=100 --follow"

check-pods:
	vagrant ssh -c "kubectl get pods"

check-services:
	vagrant ssh -c "kubectl get services"

expose-frontend:
	@echo "Connecting to Minikube at $(MINIKUBE_IP)..."
	@echo "As long as the following terminal is open, you can access the frontend at http://localhost:30080"
	vagrant ssh -- -L 30080:$(MINIKUBE_IP):30080
