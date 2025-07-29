```markdown
# DevOpsified Golang Web App 🚀
```
This is a simple, production-ready **Golang web application** built to showcase DevOps practices using tools like Docker, Kubernetes, Jenkins, and ArgoCD. The app displays a personal profile page with static content served through a Go web server.


## 📁 Project Structure

```

golang-app/
│
├── app /
|   ├── main.go                      # Main Go web server
|   ├── Dockerfile                   # Containerization
|   └── README.md
│
├── k8s/
│   ├── deployment.yaml          # Kubernetes Deployment
│   ├── service.yaml             # Kubernetes Service
│   └── route.yaml               # OpenShift Route
│
├── Jenkinsfile                  # Jenkins CI/CD pipeline
└── .argocd/
     └── application.yaml         # ArgoCD GitOps config

````

---

## 🎯 What You'll Learn

✅ How to build and containerize a Golang app  
✅ Write Kubernetes manifests for deployment  
✅ Create CI/CD pipeline with Jenkins  
✅ Automate GitOps deployment with ArgoCD  
✅ Use OpenShift or any K8s cluster for deployment  
✅ End-to-end DevOps project setup (Infrastructure + Code + Automation)

---

## 🛠️ Prerequisites

- Docker
- Kubernetes or OpenShift
- Jenkins server
- ArgoCD setup (optional for GitOps)
- GitHub repo with access tokens if using GitOps

---

## 🚀 Deployment Steps

### 1. Build & Push Docker Image
```bash
docker build -t yourdockerhub/golang-app .
docker push yourdockerhub/golang-app
````

### 2. Apply Kubernetes Manifests

```bash
kubectl apply -f k8s/
```

### 3. ArgoCD GitOps Setup

Make sure the ArgoCD app watches your repo and namespace.

---

## 📬 Contact

Feel free to reach out on [LinkedIn](https://www.linkedin.com/in/pritam-mane03/) or [Email](mailto:pritammane7666@gmail.com) if you have any questions.

---

> Built with ❤️ by [Pritam Mane](https://github.com/prritam)

```

