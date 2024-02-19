jenkins_password=$(docker exec -it jenkins-blueocean cat /var/jenkins_home/secrets/initialAdminPassword)

echo "Jenkins Unlock Password: ${jenkins_password}"
echo "Please visit http://localhost:8082 and use the above password to unlock Jenkins."