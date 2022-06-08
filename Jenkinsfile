pipeline {
    agent any
	
    tools{
	    go 'go-config'
	}
	stages{
	    stage("Code Checkout"){
		    steps {
			    echo "*****Started git checkout*****"
				git changelog: false, poll: false, url: 'https://github.com/Purushothalam/go-cobra-dockercli.git'
				echo "*****Completed Git Checkout*****"
			}
		}
		stage("Build") {
		    steps {
				echo "*****Started Building the project*****"
			    sh 'go version'
				sh 'go build .'
				
			}
		}
		stage("Artifactory") {
		    steps {
			    echo "*****Started pushing binaries to the artifactory*****"
			}
		}
	}
}
