pipeline {
  agent any
  
  tools {
    go '1.21'
  }

  stages {
    stage('Github Checkout Repo') {
      steps {
        git branch: 'main', credentialsId: 'JENKINS_GITHUB_TOKEN', url: 'https://github.com/siddhantprateek/qdrant.git'
      }
    }

    stage('Setup Go Environment') {
      steps {
        
      }
    } 
  }
}
