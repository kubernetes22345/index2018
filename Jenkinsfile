podTemplate(
    cloud: 'kubernetes',
    label: 'dev', 
    inheritFrom: 'default',
    containers: [
        containerTemplate(
            name: 'golang', 
            image: 'golang:1.10-alpine',
            ttyEnabled: true,
            command: 'cat'
        ),
        containerTemplate(
            name: 'docker', 
            image: 'docker:18.02',
            ttyEnabled: true,
            command: 'cat'
        ), containerTemplate(
            name: 'kubectl', 
            image: 'lachlanevenson/k8s-kubectl:v1.6.6', 
            ttyEnabled: true, 
            command: 'cat'
        ),
        containerTemplate(
            name: 'helm', 
            image: 'alpine/helm',
            ttyEnabled: true,
            command: 'cat'
        )
    ],
    volumes: [
        hostPathVolume(
            hostPath: '/var/run/docker.sock',
            mountPath: '/var/run/docker.sock'
        )
    ]
) {
    node('dev') {
        def commitId
        stage ('Extract') {
            checkout scm
            commitId = sh(script: 'git rev-parse --short HEAD', returnStdout: true).trim()
        }
        stage ('Build') {
            container ('golang') {
                sh 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .'
            }
        }
        def app
        stage ('Docker') {
            container ('docker') {
                app = docker.build("rohan4494/hello:${commitId}")
                docker.withRegistry('https://registry.hub.docker.com', 'docker_hub_login') {
                    app.push("${env.BUILD_NUMBER}")
                    app.push("latest")
                }
            }
        }
        stage ('Deploy') {
           input "Deploy?"
           milestone(1)
           echo "Deploying"
            container ('helm') {
                //sh "kubectl get ns"
                // sh "helm init --client-only --skip-refresh"
                sh "helm ls"
                sh "helm install grafana stable/grafana"
                sh "helm upgrade --install --force --set image.repository=rohan4494/hello,image.tag=latest hello hello "
            }
        }
    }
}
