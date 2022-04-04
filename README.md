# RecruitmentTask
This repository contains Angular Hello World application hosted on nginx as well as Terraform module. Running Angular app isn't necessary since the image is already built and hosted on Dockerhub. <br> To run Terraform module Docker and Terraform must be installed.
# Setup for Linux
To run the module open terminal in the terraform_files folder. After making sure Terraform works correctly (for example by typing __terraform__ in console)next step should be initializing the project:
```
$ Terraform init
```
Assuming the operation resulted in success, what is left is to run the module. If you want to use default options all that is needed is to type
```
$ Terraform apply
```
Is it possible to change docker image and port that is being used by the app. To do that you can use variables. Here is an example of using image from my DockerHub and using port 8081 
```
$ Terraform apply -var="image_name=promoqwerty/virtuslabtask:version1.0" -var="port=8081"
```
Is is also worth mentioning that some of the commands require root permissions, so it is worth to use __sudo__ before all of them. 
