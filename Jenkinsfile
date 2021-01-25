pipeline {
  agent none
  stages {
    stage('Deploy') {
      agent any
      steps {
        echo 'Deploying'
        sh 'chmod +x deploy.sh'
        sh './deploy.sh'
      }
    }
  }
}