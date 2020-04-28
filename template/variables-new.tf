# Common variables
variable "environment" {}

variable "image_tag" {}

variable "service_name" {}

variable "current_version" {
  default = ""
}

variable "area" {
  default = "edgeservices"
}

variable "podnumber" {
  default = "2"
}

variable "cpu_req" {
  default = "0.5"
}

variable "cpu_limit" {
  default = "1"
}

variable "mem_req" {
  default = "2Gi"
}

variable "mem_limit" {
  default = "2Gi"
}

///////// Service related variables /////////

variable "service_port" { default = "8292" }
variable "edgeapprepository_url" { default = "http://edgeapprepository-svc:8120" }
variable "multilanguagedocumentstore_url" { default = "http://multilanguagedocumentstore-svc:8289" } 

//variable "nas_server" {  default = "" }
//variable "SW_AGENT_ENABLE" {  default = "" }
variable "nas_server" {}

variable "SW_AGENT_ENABLE" { default = ""}
