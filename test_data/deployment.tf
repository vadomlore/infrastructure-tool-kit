## AgentIAM ##

data "template_file" "deployment" {
  template = "${file("${path.module}/deployment.yaml")}"

  vars {
    # Below are common setting.
    environment     = "${var.environment}"
    area            = "${var.area}"
    namespace       = "${var.area}-${var.environment}"
    service_name    = "${var.service_name}"
    service_port    = "${var.service_port}"
    podnumber       = "${var.podnumber}"
    image_tag       = "${var.image_tag}"
    deployment      = "${var.environment == "gadev" || var.environment == "gainteg" ? var.service_name : "${var.service_name}-${var.image_tag}"}"
    current_version = "${var.environment == "gadev" || var.environment == "gainteg" ? var.image_tag : var.current_version}"
    mem_req         = "${var.mem_req}"
    mem_limit       = "${var.mem_limit}"
    cpu_req         = "${var.cpu_req}"
    cpu_limit       = "${var.cpu_limit}"

    # Below are service related setting.
    spring_profiles_active = "secured"

    # URL
    apigateway_url = "${data.terraform_remote_state.infrastructure.apigateway_url}"
    oauth2_jwks    = "${data.terraform_remote_state.infrastructure.coreiam_jwks}"
    jwks_url    = "${data.terraform_remote_state.infrastructure.coreiam_jwks}"
    issuer_url = "${data.terraform_remote_state.infrastructure.issuer_url}"

    edgeapprepository_url = "${var.environment == "gadev" ? var.edgeapprepository_url : "https://edgeapprepository-svc.${var.area}-${var.environment}.svc.cluster.local"}"
    multilanguagedocumentstore_url = "${var.environment == "gadev" ? var.multilanguagedocumentstore_url : "https://multilanguagedocumentstore-svc.${var.area}-${var.environment}.svc.cluster.local"}"
    nas_server = "${var.nas_server}"
    sw_agent_enable = "${var.SW_AGENT_ENABLE}"
  }
}
