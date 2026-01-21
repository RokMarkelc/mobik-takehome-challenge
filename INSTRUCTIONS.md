# Mobik take-home assignment

Welcome to the Mobik take-home assignment.
This assignment is structured in two parts.

First part will test your savviness in Linux server administration.
Second part will test your knowledge in software development and delivery.

Store your answers in a public Git repository (use GitHub or your platform of choice).

---

## Part one: Linux server administration

Log in to the Linux server (you will get login information via e-mail) and follow the instructions presented by the welcome banner.
Write your answers to a file named `takehome_part_1.md` in your Git project.

---

## Part two: Automated Three-Tier App Deployment

This test task is designed to evaluate a candidate's ability to use a full DevOps toolchain to automate the provisioning of a virtual machine and the deployment of a three-tier application. The goal is to see how you integrate different technologies and manage a complete end-to-end workflow.

---

### The Scenario

Your task is to automate the deployment of a three-tier application onto a local Kubernetes cluster. The entire environment must be provisioned and managed using **Vagrant** and **Ansible**.

The target system is a virtual machine running **Debian 12**.

---

### The Application Stack

* **Tier 1: Frontend:** A simple web application that serves a static page.
* **Tier 2: Backend:** A REST API that connects to the database.
* **Tier 3: Database:** A MySQL or PostgreSQL database.

### The Toolchain

* **Vagrant:** To create and manage the virtual machine.
* **Ansible:** To provision the VM and deploy the application.
* **Minikube & `kubectl`:** To run the local Kubernetes cluster.
* **Docker:** To containerize the application components.

---

### Your Deliverables

Please provide a Git repository containing the following files and directories. The final solution should be fully automated so that a reviewer can run a single command (or a few commands) to see the entire process from start to finish.

#### Part 1: Infrastructure as Code with Vagrant and Ansible

Create the files necessary to provision a Debian 12 virtual machine and install the required tools.

* `Vagrantfile`: Your Vagrant configuration file. It should use the `debian/bookworm64` box.
* `ansible/` directory:
    * `hosts.ini`: An inventory file targeting the Vagrant VM.
    * `setup.yml`: An Ansible playbook that installs Docker, Minikube, and `kubectl` on the Debian 12 guest.

#### Part 2: Containerization and Kubernetes Manifests

Containerize the provided three-tier application and write the Kubernetes manifests to deploy it.

* `app/` directory:
    * `frontend/Dockerfile`: The Dockerfile for the frontend application.
    * `backend/Dockerfile`: The Dockerfile for the backend API.
* `kubernetes/` directory:
    * `db-deployment.yml`: The Kubernetes `Deployment` and `Service` for the database.
    * `backend-deployment.yml`: The `Deployment` and `Service` for the backend, including environment variables to connect to the database.
    * `frontend-deployment.yml`: The `Deployment` and `Service` for the frontend.

#### Part 3: Automated Application Deployment

Create a second Ansible playbook to deploy the application onto the Minikube cluster running inside the VM.

* `deploy-app.yml`: An Ansible playbook that:
    * Starts Minikube if it isn't running.
    * Configures the shell to use Minikube's Docker daemon.
    * Builds the application images and applies the Kubernetes manifests in the correct order (database, then backend, then frontend).
    * This playbook should be **idempotent**.

---

### Final Submission and Evaluation

Your submission should be a link to your Git repository.

A comprehensive `README.md` file is crucial. It should contain:
* A clear explanation of the architecture.
* Instructions on how to run your entire solution, from `vagrant up` to verifying the final deployment.
* Instructions on how to access the application after it has been deployed.
* A brief section on how you would troubleshoot potential issues (e.g., checking pod logs or service status).

Your solution will be evaluated on the following criteria:
* **Automation:** Is the process fully automated and easy to run?
* **Code Quality:** Is the code clean, well-structured, and easy to read?
* **Idempotence:** Can the playbooks be run multiple times without issues?
* **Documentation:** Is the `README.md` helpful and clear?
* **Tool Integration:** Do you demonstrate a solid understanding of how all the tools work together?
