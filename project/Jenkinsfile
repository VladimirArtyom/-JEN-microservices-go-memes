pipeline{
    agent any

    stages{
        stage('Checkout Code') {
            steps {
                echo 'Fetching lates code from Github...'
                checkout scm
            }
        }

        stage("Build & Start containers"){
            steps {
                echo 'Builiding go Binaries and starting docker containers ...'
                sh 'make up_build'
            }
        }

        stage("Verify Services are running") {
            steps {
                echo 'Checking if containers are healthy...'

                sh 'docker compose ps'

                sh 'sleep 5 && curl -f http://localhost:8080/broker || exit 1'
            }
        }

    }

    post{
        always {
            echo 'Tearing down containers and cleaning up...'
            sh 'make down'
            echo 'clean-up complete'
        }

        success {
            echo 'Verification successfull'
        }

        failure {
            echo 'Verification Failed'
        }
    }
}