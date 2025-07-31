pipeline {
  agent any

  environment {
    IMAGE = "docker.io/tusk03/golang-app:latest"
    DOCKERHUB_USER = "tusk03"
    OPENSHIFT_API = "https://api.rm3.7wse.p1.openshiftapps.com:6443"
    NAMESPACE = "pritam-03-dev"
    RELEASE_NAME = "golang-app"
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
            docker build -f ./app/Dockerfile -t $IMAGE ./app
        '''
      }
    }

    stage('Push to Dockerhub') {
      steps {
        withCredentials([usernamePassword(credentialsId: 'dockerhub-creds', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
            sh '''
              echo "Logging into Docker Hub..."
              echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
              echo "Pushing Docker Image..."
              docker push $IMAGE
         '''
        }
      }
    }
    stage('Deploy to OpenShift via Helm') {
      steps {
        withCredentials([string(credentialsId: 'openshift-token', variable: 'OC_TOKEN')]) {
          sh '''
            echo "Logging into OpenShift..."
            oc login $OPENSHIFT_API --token=$OC_TOKEN --insecure-skip-tls-verify

            echo "Helm Deploying..."
            helm upgrade --install $RELEASE_NAME ./k8s/helm-chart \
              --set image.repository=docker.io/tusk03/golang-app \
              --set image.tag=latest \
              --namespace=$NAMESPACE
          '''
        }
      }
    }
  }
}
  
