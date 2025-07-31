---

## 🚀 Golang CI/CD with Jenkins, Docker, Helm, and OpenShift

This repository demonstrates a full CI/CD pipeline for a **Golang application**, showing how code changes are built, containerized, and deployed to an **OpenShift** cluster using modern DevOps tools.

---

## 📁 Project Structure

```
.
├── app/                     # Go application code & Dockerfile
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
│
├── k8s/helm-chart/          # Helm chart for OpenShift deployment
│   ├── Chart.yaml
│   ├── values.yaml
│   └── templates/
│       ├── deployment.yaml
│       ├── service.yaml
│       └── route.yaml
│
└── Jenkinsfile              # Jenkins CI/CD pipeline definition
```

---

## 🔁 CI/CD Pipeline Overview

The pipeline is orchestrated using **Jenkins**, triggered via **GitHub Webhooks** on commits to the `main` branch.

### ✅ Continuous Integration (CI)

1. **Trigger**: GitHub push event (via webhook).
2. **Build**: Compiles the Go application as a static binary.
3. **Dockerize**: Creates a lightweight image using a multi-stage Dockerfile.
4. **Push**: Publishes the image to Docker Hub (`tusk03/golang-app`).

### 🚀 Continuous Delivery (CD)

1. **Login**: Jenkins authenticates with OpenShift via CLI (`oc login`) using a secret token.
2. **Deploy**: Runs `helm upgrade --install` using the chart in `k8s/helm-chart/`.
3. **Expose**: OpenShift `Route` exposes the service (similar to Ingress in Kubernetes).

---

## 🛠️ Technologies Used

| Tool/Platform   | Purpose                          |
| --------------- | -------------------------------- |
| **Go (Golang)** | Application logic                |
| **GitHub**      | Source control                   |
| **Jenkins**     | CI/CD pipeline automation        |
| **Docker**      | Containerization of app          |
| **Docker Hub**  | Container registry               |
| **Helm**        | Kubernetes/OpenShift deployment  |
| **OpenShift**   | Container orchestration platform |

---

## 🧩 Troubleshooting & Key Learnings

Building this pipeline involved solving several real-world DevOps issues:

### 🔹 1. Docker Build Context Error

* **Error**: `go.mod not found`, or `unable to prepare context`.
* **Fix**: Reorganized project so the Dockerfile and Go files live inside `/app`. Updated `docker build` path to `./app` in `Jenkinsfile`.

### 🔹 2. DockerHub Credential Issue

* **Error**: Jenkins failed to push image to Docker Hub.
* **Fix**: Corrected `withCredentials` block in `Jenkinsfile` using `usernamePassword` for DockerHub credentials.

### 🔹 3. OpenShift Token Expired

* **Error**: `oc login: token invalid or expired`.
* **Fix**: Generated a fresh token from OpenShift Console and updated Jenkins credentials.

### 🔹 4. Helm Chart Not Found

* **Error**: `Chart.yaml file is missing`.
* **Fix**: Corrected the path to `./k8s/helm-chart` and ensured file is named `Chart.yaml` (case-sensitive).

---

## ✅ Outcome

By the end of this CI/CD pipeline:

* A new Go binary is compiled and containerized automatically.
* Docker image is pushed to your registry.
* Helm deploys the updated image to your OpenShift cluster using rolling upgrades.
* The app is available via a public OpenShift `Route`.

---
