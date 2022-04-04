terraform {
  required_providers {
  docker = {
    source = "kreuzwerker/docker"
    version = "~> 2.13.0"
    }
  }
}

variable "image_name" {
  description = "name of the image"
  type = string
  default = "nginx:latest"
}
  
variable "port" {
  type= number
  default = 8300

}
provider "docker" {}

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


