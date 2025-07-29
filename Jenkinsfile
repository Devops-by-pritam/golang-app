pipeline {
  agent any

  environment {
    IMAGE = "docker.io/tusk03/golang-app:latest"
    DOCKERHUB_USER = "tusk03"
  }

  stages {
    stage('Checkout') {
      steps {
        git url: 'https://github.com/Devops-by-pritam/golang-app.git', branch: 'main'
      }
    }

    stage('Build Docker Image') {
      steps {
        sh '''
            export DOCKER_BUILDKIT=0
            docker build -t $IMAGE .
        '''
      }
    }

    stage('Push to Dockerhub') {
      steps {
        withCredentials([string(credentialsId: 'dockerhub-creds', variable: 'DOCKER_PASS')]) {
          sh 'echo $DOCKER_PASS | docker login -u $DOCKERHUB_USER --password-stdin'
          sh 'docker push $IMAGE'
        }
      }
    }

    stage('Trigger Argo CD Sync') {
      steps {
        // Optional: This stage triggers Argo CD sync via API or webhook
        // Replace with actual ArgoCD API URL/token if needed
        echo "Trigger Argo CD here if needed"
        // Example (unsecured):
        // sh 'curl -X POST http://argocd-server:8080/api/v1/applications/golang-app/sync'
      }
    }
  }
}
