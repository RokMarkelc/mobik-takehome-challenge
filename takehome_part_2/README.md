# Mobik Take-Home Challenge:

This project demonstrates an automated deployment of a three-tier application using Vagrant, Ansible, Minikube, Kubernetes, and Docker. The goal is to provision a local environment and deploy the application with minimal manual intervention.

## Architecture

The solution implements a three-tier architecture on a single virtual machine (VM) running Debian 12:

- **Tier 1: Frontend** - An Angular web application that communicates with backend.
- **Tier 2: Backend** - A Go REST API that communicates with the database.
- **Tier 3: Database** - A PostgreSQL database.

**Deployment Flow:**
1. Vagrant creates and manages the Debian 12 virtual machine.
2. Ansible provisions the VM by installing Docker, Minikube, and kubectl.
3. Ansible start minikube kubernetes cluster, makes shell use minikube's docker daemon and deploys applications (postgres database, go backend and angular frontend) to the cluster with kubectl

Key components:
- **Vagrant**: Manages the VM lifecycle.
- **Ansible**: Handles VM provisioning (installing Docker, Minikube, kubectl) and application deployment (building images, applying manifests).
- **Minikube**: Provides the local Kubernetes cluster.
- **Docker**: Containerizes the frontend, backend, and database.
- **Kubernetes**: Orchestrates the deployments and services.

## Prerequisites
To provision and deploy the whole architecture automatically you will need:
   - Vagrant
   - VirtualBox
   - Make

## Instructions

1. **Provision VM and deploy applications**:
    To start the VM, provision it with Vagrant, deploy all applications and expose frontend on the host, run:
    ```
    make full-deploy
    ```

    To run these 4 steps separately run:
    ```
    make start-vm         # Starts the vm without provisioning
    make provision-vm     # Provisions the vm using ansible/setup.yml playbook
    make deploy-apps      # Starts minikube and deploys all apps using ansible/deploy-app.yml playbook
    make expose-frontend  # Makes frontend accessible on host machine
    ```

2. **Verify the Deployment**:
    - Make sure all pods are running:
        ```
        make check-pods
        ```
    - Check services:
        ```
        make check-services
        ```
    - The frontend should be accessible on http://localhost:30080 and new todos can be added.

3. **Cleanup**:
    To delete the VM and everything inside run:
    ```
    make destroy-vm
    ```

## Accessing the Application

After deployment, the frontend is exposed via a Kubernetes NodePort Service. At this point it is only accessible from inside the VM. To access it on host machine and open browser: 

- Run:
  ```
  make expose-frontend
  ```
  This will open an ssh terminal. As long as it is opened frontend will be accessible on host machine on http://localhost:30080
- Open a browser and navigate to `http://localhost:30080` to view the frontend page.
- The backend API and database can be accessed internally via `http://backend:8080` (e.g., for API calls from frontend) and via `http://postgres:5432` (e.g., for SQL calls from backend).
- All services can be exec-ed into by running `make ssh-pod service=<service-name>` (postgres/frontend/backend)

## Troubleshooting

- **Pods not starting**: Check pod status with `make check-pods` and check logs with `make logs service=<service-name>` (postgres/frontend/backend). Common issues: image build failures, resource constraints or shaky network connection (backend sometimes can't access GH to download `migrate` tool).
- **Services not accessible**: SSH into the vm with `make ssh` and verify with `kubectl get services` and `kubectl describe svc <service-name>` (postgres/frontend/backend).
- **Services not working correctly**: Exec into container's shell with `make ssh-pod service=<service-name>` (postgres/frontend/backend) and check if everything is configured correctly. Check logs with `make logs service=<service-name>`
- **Minikube issues**: SSH into the vm with `make ssh` and restart with `minikube stop && minikube start`. Check status with `minikube status`.
- **Ansible failures**: Re-run playbooks (they are idempotent). Check Ansible output for errors, e.g., permission issues during Docker setup.
