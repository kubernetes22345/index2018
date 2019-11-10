pipeline {
    agent { label 'helm-deploy-slave' }
    environment {
        
        DOCKER_IMAGE_NAME = "rohan4494/hello"
    }
    stages {
        stage('Build Docker Image') {
            steps {
                script {
                    app = docker.build(DOCKER_IMAGE_NAME)
                }
            }
        }
        stage('Push Docker Image') {
            
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker_hub_login') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }
        stage('DeployToProduction') {
            
            steps {
                input 'Deploy to Production?'
                milestone(1)
                sh "/helm upgrade --install --wait --set image.repository=rohan4494/hello,image.tag=${env.BUILD_NUMBER} hello hello --namespace jenkins-master"
            }
        }
    }
}
