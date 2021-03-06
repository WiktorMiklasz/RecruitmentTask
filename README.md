# RecruitmentTask
This repository contains Angular Hello World application hosted on nginx as well as Terraform module. Running Angular app isn't necessary since the image is already built and hosted on Dockerhub. <br> To run Terraform module Docker and Terraform must be installed.
# Setup for Linux
To run the module open terminal in the terraform_files folder. After making sure Terraform works correctly (for example by typing __terraform__ in console)next step should be initializing the project:
```
$ terraform init
```
Assuming the operation resulted in success, what is left is to run the module. If you want to use default options all that is needed is to type
```
$ sudo terraform apply
```
App should be accesible and visible in browser. For default port number 8300 localhost:8300 should work. It is possible to change docker image and port that is being used by the app. To do that you can use variables. Here is an example of using image from my DockerHub and using port 8081 
```
$ sudo terraform apply -var="image_name=promoqwerty/virtuslabtask:version1.0" -var="port=8081"
```

# Short documentation of the Terraform module

Terraform ,as a declarative tool, requires its user to only defy end result. The main.tf file is the root module. Other files in terraform_modules are automatically generated by Terraform. The module is divided by separate blocks. The first one is the terraform block. Here is where providers must be declared so Terraform can install them from its registry. 

```
terraform {
  required_providers {
  docker = {
    source = "kreuzwerker/docker"
    version = "~> 2.13.0"
    }
  }
}
```
Variable blocks can have several different arguments, the ones used here are type keywords string and number (type can also be a collection such as list or map)
and default. Option default is being used when module is run without flag -v.
```
variable "image_name" {
  description = "name of the image"
  type = string
  default = "nginx:latest"
}
  
variable "port" {
  type= number
  default = 8300

}
```
The last two blocks are resource blocks. Those are the objects that Terraform is supposed to manage. The arguments vary depending on the type of the object
In this case resource docker image takes name as a variable and is not stored locally, container on the other hand takes the variable name and uses it as its image to build itself, has predefined name and internal port. External port is also dependent on variable.
```
resource "docker_image" "nginx"{
  name = var.image_name
  keep_locally = false
}


resource "docker_container" "nginx"{
  image = var.image_name
  name = "Container"
  ports {
  internal = 8081
  external = var.port
  }
}
\\\


