provider "google" {
  project = var.provider_id
  credentials = var.credentials
}

module "storage" {
  source = "../../modules/storage"

  storage_instance_name = var.infrastructure_id
}

module "compute" {
  source = "../../modules/compute"

  # Pool configuration
  compute_name      = var.infrastructure_id
  storage_host      = module.storage.storage_host
  storage_port      = module.storage.storage_port
  storage_name      = module.storage.storage_name
  storage_user      = module.storage.storage_user
  storage_password  = module.storage.storage_password
}
