pipeline {
    agent any
    
    tools{
        go 'go-config'
    }
    
    stages{
        stage("Build"){
            steps {
                echo "Started Building the Jenkins File"   
                sh 'go version'
            }
        }   
    }
}
