pipeline {
  agent {
    node {
      label 'eclipse'
    }

  }
  stages {
    stage('Install dependencies') {
      steps {
        echo 'Installing dependencies...'
        sh 'cd eclipse && yarn install'
      }
    }

    stage('Test') {
      steps {
        echo 'Running tests...'
        sh 'cd eclipse && yarn test'
      }
    }

  }
}
