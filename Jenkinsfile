pipeline {
  agent none
  stages {
    stage('Deploy') {
      agent any
      steps {
        echo 'Deploying'
        sh 'chmod 777 deploy.sh'
        sh './deploy.sh'
      }
    }
  }
}