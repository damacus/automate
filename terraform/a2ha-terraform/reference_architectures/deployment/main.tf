resource "random_id" "cluster_id" {
  byte_length = 4
}

module "system-tuning-automate" {
  source                             = "./modules/system"
  automate_archive_disk_fs_path      = var.automate_archive_disk_fs_path
  elasticsearch_archive_disk_fs_path = var.elasticsearch_archive_disk_fs_path
  instance_count                     = var.automate_instance_count
  postgresql_archive_disk_fs_path    = var.postgresql_archive_disk_fs_path
  public_ips                         = var.automate_public_ips
  ssh_key_file                       = var.ssh_key_file
  ssh_user                           = var.ssh_user
  ssh_user_sudo_password             = local.fe_sudo_password
  sudo_cmd                           = var.sudo_cmd
}

module "system-tuning-chef_server" {
  source                             = "./modules/system"
  automate_archive_disk_fs_path      = var.automate_archive_disk_fs_path
  elasticsearch_archive_disk_fs_path = var.elasticsearch_archive_disk_fs_path
  instance_count                     = var.chef_server_instance_count
  postgresql_archive_disk_fs_path    = var.postgresql_archive_disk_fs_path
  public_ips                         = var.chef_server_public_ips
  ssh_key_file                       = var.ssh_key_file
  ssh_user                           = var.ssh_user
  ssh_user_sudo_password             = local.fe_sudo_password
  sudo_cmd                           = var.sudo_cmd
}

module "system-tuning-elasticsearch" {
  source                             = "./modules/system"
  automate_archive_disk_fs_path      = var.automate_archive_disk_fs_path
  elasticsearch_archive_disk_fs_path = var.elasticsearch_archive_disk_fs_path
  instance_count                     = var.elasticsearch_instance_count
  postgresql_archive_disk_fs_path    = var.postgresql_archive_disk_fs_path
  public_ips                         = var.elasticsearch_public_ips
  ssh_key_file                       = var.ssh_key_file
  ssh_user                           = var.ssh_user
  ssh_user_sudo_password             = local.be_sudo_password
  sudo_cmd                           = var.sudo_cmd
}

module "system-tuning-postgresql" {
  source                             = "./modules/system"
  automate_archive_disk_fs_path      = var.automate_archive_disk_fs_path
  elasticsearch_archive_disk_fs_path = var.elasticsearch_archive_disk_fs_path
  instance_count                     = var.postgresql_instance_count
  postgresql_archive_disk_fs_path    = var.postgresql_archive_disk_fs_path
  public_ips                         = var.postgresql_public_ips
  ssh_key_file                       = var.ssh_key_file
  ssh_user                           = var.ssh_user
  ssh_user_sudo_password             = local.be_sudo_password
  sudo_cmd                           = var.sudo_cmd
}

module "airgap_bundle-elasticsearch" {
  source            = "./modules/airgap_bundle"
  archive_disk_info = module.system-tuning-elasticsearch.archive_disk_info
  instance_count    = var.elasticsearch_instance_count
  public_ips        = var.elasticsearch_public_ips
  bundle_files = [{
    source      = var.backend_aib_local_file
    destination = var.backend_aib_dest_file
  }]
  ssh_key_file = var.ssh_key_file
  ssh_user     = var.ssh_user
  tmp_path     = var.tmp_path
}

module "airgap_bundle-postgresql" {
  source            = "./modules/airgap_bundle"
  archive_disk_info = module.system-tuning-postgresql.archive_disk_info
  instance_count    = var.postgresql_instance_count
  public_ips        = var.postgresql_public_ips
  bundle_files = [{
    source      = var.backend_aib_local_file
    destination = var.backend_aib_dest_file
  }]
  ssh_key_file = var.ssh_key_file
  ssh_user     = var.ssh_user
  tmp_path     = var.tmp_path
}

module "airgap_bundle-automate" {
  source            = "./modules/airgap_bundle"
  archive_disk_info = module.system-tuning-automate.archive_disk_info
  instance_count    = var.automate_instance_count
  public_ips        = var.automate_public_ips
  bundle_files = [{
    source      = var.backend_aib_local_file
    destination = var.backend_aib_dest_file
    }, {
    source      = var.frontend_aib_local_file
    destination = var.frontend_aib_dest_file
  }]
  ssh_key_file = var.ssh_key_file
  ssh_user     = var.ssh_user
  tmp_path     = var.tmp_path
}

module "airgap_bundle-chef_server" {
  source            = "./modules/airgap_bundle"
  archive_disk_info = module.system-tuning-chef_server.archive_disk_info
  instance_count    = var.chef_server_instance_count
  public_ips        = var.chef_server_public_ips
  bundle_files = [{
    source      = var.backend_aib_local_file
    destination = var.backend_aib_dest_file
    }, {
    source      = var.frontend_aib_local_file
    destination = var.frontend_aib_dest_file
  }]
  ssh_key_file = var.ssh_key_file
  ssh_user     = var.ssh_user
  tmp_path     = var.tmp_path
}

module "habitat-elasticsearch" {
  source                          = "./modules/habitat"
  airgap_info                     = module.airgap_bundle-elasticsearch.airgap_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  hab_sup_http_gateway_ca_cert    = var.hab_sup_http_gateway_ca_cert
  hab_sup_http_gateway_priv_key   = var.hab_sup_http_gateway_priv_key
  hab_sup_http_gateway_pub_cert   = var.hab_sup_http_gateway_pub_cert
  hab_sup_ring_key                = var.hab_sup_ring_key
  hab_sup_run_args                = var.hab_sup_run_args
  install_hab_sh_args             = ""
  instance_count                  = var.elasticsearch_instance_count
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  private_ips                     = var.elasticsearch_private_ips
  public_ips                      = var.elasticsearch_public_ips
  peer_ips = concat(
    var.elasticsearch_private_ips,
    var.postgresql_private_ips
  )
  ssh_key_file           = var.ssh_key_file
  ssh_user               = var.ssh_user
  ssh_user_sudo_password = local.be_sudo_password
  sudo_cmd               = var.sudo_cmd
  habitat_uid_gid        = var.habitat_uid_gid
}

module "habitat-postgresql" {
  source                          = "./modules/habitat"
  airgap_info                     = module.airgap_bundle-postgresql.airgap_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  hab_sup_http_gateway_ca_cert    = var.hab_sup_http_gateway_ca_cert
  hab_sup_http_gateway_priv_key   = var.hab_sup_http_gateway_priv_key
  hab_sup_http_gateway_pub_cert   = var.hab_sup_http_gateway_pub_cert
  hab_sup_ring_key                = var.hab_sup_ring_key
  hab_sup_run_args                = var.hab_sup_run_args
  install_hab_sh_args             = ""
  instance_count                  = var.postgresql_instance_count
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  private_ips                     = var.postgresql_private_ips
  public_ips                      = var.postgresql_public_ips
  peer_ips = concat(
    var.elasticsearch_private_ips,
    var.postgresql_private_ips
  )
  ssh_key_file           = var.ssh_key_file
  ssh_user               = var.ssh_user
  ssh_user_sudo_password = local.be_sudo_password
  sudo_cmd               = var.sudo_cmd
  habitat_uid_gid        = var.habitat_uid_gid
}

module "habitat-automate" {
  source                          = "./modules/habitat"
  airgap_info                     = module.airgap_bundle-automate.airgap_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  hab_sup_http_gateway_ca_cert    = var.hab_sup_http_gateway_ca_cert
  hab_sup_http_gateway_priv_key   = var.hab_sup_http_gateway_priv_key
  hab_sup_http_gateway_pub_cert   = var.hab_sup_http_gateway_pub_cert
  hab_sup_ring_key                = var.hab_sup_ring_key
  hab_sup_run_args                = var.hab_sup_run_args
  install_hab_sh_args             = "--no-service"
  instance_count                  = var.automate_instance_count
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  private_ips                     = var.automate_private_ips
  public_ips                      = var.automate_public_ips
  peer_ips                        = var.automate_private_ips
  ssh_key_file                    = var.ssh_key_file
  ssh_user                        = var.ssh_user
  ssh_user_sudo_password          = local.fe_sudo_password
  sudo_cmd                        = var.sudo_cmd
  habitat_uid_gid                 = var.habitat_uid_gid
}

module "habitat-chef_server" {
  source                          = "./modules/habitat"
  airgap_info                     = module.airgap_bundle-chef_server.airgap_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  hab_sup_http_gateway_ca_cert    = var.hab_sup_http_gateway_ca_cert
  hab_sup_http_gateway_priv_key   = var.hab_sup_http_gateway_priv_key
  hab_sup_http_gateway_pub_cert   = var.hab_sup_http_gateway_pub_cert
  hab_sup_ring_key                = var.hab_sup_ring_key
  hab_sup_run_args                = var.hab_sup_run_args
  install_hab_sh_args             = "--no-service"
  instance_count                  = var.chef_server_instance_count
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  private_ips                     = var.chef_server_private_ips
  public_ips                      = var.chef_server_public_ips
  peer_ips                        = var.chef_server_private_ips
  ssh_key_file                    = var.ssh_key_file
  ssh_user                        = var.ssh_user
  ssh_user_sudo_password          = local.fe_sudo_password
  sudo_cmd                        = var.sudo_cmd
  habitat_uid_gid                 = var.habitat_uid_gid
}

module "elasticsearch" {
  source                       = "./modules/elasticsearch"
  airgap_info                  = module.airgap_bundle-elasticsearch.airgap_info
  backend_aib_dest_file        = var.backend_aib_dest_file
  backend_aib_local_file       = var.backend_aib_local_file
  curator_pkg_ident            = var.curator_pkg_ident
  elasticsearch_instance_count = var.elasticsearch_instance_count
  elasticsearch_listen_port    = var.elasticsearch_listen_port
  elasticsearch_pkg_ident      = var.elasticsearch_pkg_ident
  elasticsearch_svc_load_args  = var.elasticsearch_svc_load_args
  elasticsidecar_pkg_ident     = var.elasticsidecar_pkg_ident
  elasticsidecar_svc_load_args = var.elasticsidecar_svc_load_args
  habitat_info                 = module.habitat-elasticsearch.habitat_info
  journalbeat_pkg_ident        = var.journalbeat_pkg_ident
  kibana_pkg_ident             = var.kibana_pkg_ident
  metricbeat_pkg_ident         = var.metricbeat_pkg_ident
  private_ips                  = var.elasticsearch_private_ips
  public_ips                   = var.elasticsearch_public_ips
  ssh_key_file                 = var.ssh_key_file
  ssh_user                     = var.ssh_user
  ssh_user_sudo_password       = local.be_sudo_password
  sudo_cmd                     = var.sudo_cmd
}

module "postgresql" {
  source                          = "./modules/postgresql"
  airgap_info                     = module.airgap_bundle-postgresql.airgap_info
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  elasticsearch_listen_port       = var.elasticsearch_listen_port
  elasticsearch_private_ips       = var.elasticsearch_private_ips
  habitat_info                    = module.habitat-postgresql.habitat_info
  journalbeat_pkg_ident           = var.journalbeat_pkg_ident
  metricbeat_pkg_ident            = var.metricbeat_pkg_ident
  pgleaderchk_listen_port         = var.pgleaderchk_listen_port
  pgleaderchk_pkg_ident           = var.pgleaderchk_pkg_ident
  pgleaderchk_svc_load_args       = var.pgleaderchk_svc_load_args
  postgresql_archive_disk_fs_path = var.postgresql_archive_disk_fs_path
  postgresql_instance_count       = var.postgresql_instance_count
  postgresql_listen_port          = var.postgresql_listen_port
  postgresql_pkg_ident            = var.postgresql_pkg_ident
  postgresql_pg_dump_enabled      = var.postgresql_pg_dump_enabled
  postgresql_ssl_enable           = var.postgresql_ssl_enable
  postgresql_svc_load_args        = var.postgresql_svc_load_args
  postgresql_wal_archive_enabled  = var.postgresql_wal_archive_enabled
  proxy_listen_port               = var.proxy_listen_port
  proxy_pkg_ident                 = var.proxy_pkg_ident
  proxy_svc_load_args             = var.proxy_svc_load_args
  private_ips                     = var.postgresql_private_ips
  public_ips                      = var.postgresql_public_ips
  ssh_key_file                    = var.ssh_key_file
  ssh_user                        = var.ssh_user
  ssh_user_sudo_password          = local.be_sudo_password
  sudo_cmd                        = var.sudo_cmd
}

module "bootstrap_automate" {
  source                          = "./modules/automate"
  airgap_info                     = module.airgap_bundle-automate.airgap_info
  automate_admin_email            = var.automate_admin_email
  automate_admin_username         = var.automate_admin_username
  automate_admin_password         = var.automate_admin_password
  automate_config                 = file(var.automate_config_file)
  automate_dc_token               = var.automate_dc_token
  automate_fqdn                   = var.automate_fqdn
  automate_instance_count         = 1
  automate_role                   = "bootstrap_automate"
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  cluster_id                      = random_id.cluster_id.hex
  frontend_aib_dest_file          = var.frontend_aib_dest_file
  frontend_aib_local_file         = var.frontend_aib_local_file
  habitat_info                    = module.habitat-automate.habitat_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  elasticsearch_listen_port       = var.elasticsearch_listen_port
  elasticsearch_private_ips       = var.elasticsearch_private_ips
  postgresql_private_ips          = var.postgresql_private_ips
  postgresql_ssl_enable           = var.postgresql_ssl_enable
  private_ips                     = slice(var.automate_private_ips, 0, 1)
  proxy_listen_port               = var.proxy_listen_port
  public_ips                      = slice(var.automate_public_ips, 0, 1)
  ssh_key_file                    = var.ssh_key_file
  ssh_user                        = var.ssh_user
  ssh_user_sudo_password          = local.fe_sudo_password
  sudo_cmd                        = var.sudo_cmd
  teams_port                      = var.teams_port
}

module "automate" {
  source                          = "./modules/automate"
  airgap_info                     = module.airgap_bundle-automate.airgap_info
  automate_admin_email            = var.automate_admin_email
  automate_admin_username         = var.automate_admin_username
  automate_admin_password         = var.automate_admin_password
  automate_config                 = file(var.automate_config_file)
  automate_dc_token               = var.automate_dc_token
  automate_fqdn                   = var.automate_fqdn
  automate_instance_count         = var.automate_instance_count - 1
  automate_role                   = "automate"
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  cluster_id                      = random_id.cluster_id.hex
  frontend_aib_dest_file          = var.frontend_aib_dest_file
  frontend_aib_local_file         = var.frontend_aib_local_file
  habitat_info                    = module.habitat-automate.habitat_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  elasticsearch_listen_port       = var.elasticsearch_listen_port
  elasticsearch_private_ips       = var.elasticsearch_private_ips
  postgresql_private_ips          = var.postgresql_private_ips
  postgresql_ssl_enable           = var.postgresql_ssl_enable
  private_ips = slice(
    var.automate_private_ips,
    1,
    length(var.automate_private_ips),
  )
  proxy_listen_port = var.proxy_listen_port
  public_ips = slice(
    var.automate_public_ips,
    1,
    length(var.automate_public_ips),
  )
  ssh_key_file           = var.ssh_key_file
  ssh_user               = var.ssh_user
  ssh_user_sudo_password = local.fe_sudo_password
  sudo_cmd               = var.sudo_cmd
  teams_port             = var.teams_port
}

module "chef_server" {
  source                          = "./modules/automate"
  airgap_info                     = module.airgap_bundle-chef_server.airgap_info
  automate_admin_email            = var.automate_admin_email
  automate_admin_username         = var.automate_admin_username
  automate_admin_password         = var.automate_admin_password
  automate_config                 = file(var.automate_config_file)
  automate_dc_token               = var.automate_dc_token
  automate_fqdn                   = var.automate_fqdn
  automate_instance_count         = var.chef_server_instance_count
  automate_role                   = "chef_api"
  backend_aib_dest_file           = var.backend_aib_dest_file
  backend_aib_local_file          = var.backend_aib_local_file
  cluster_id                      = random_id.cluster_id.hex
  frontend_aib_dest_file          = var.frontend_aib_dest_file
  frontend_aib_local_file         = var.frontend_aib_local_file
  habitat_info                    = module.habitat-chef_server.habitat_info
  hab_sup_http_gateway_auth_token = var.hab_sup_http_gateway_auth_token
  elasticsearch_listen_port       = var.elasticsearch_listen_port
  elasticsearch_private_ips       = var.elasticsearch_private_ips
  postgresql_private_ips          = var.postgresql_private_ips
  postgresql_ssl_enable           = var.postgresql_ssl_enable
  private_ips                     = var.automate_private_ips
  proxy_listen_port               = var.proxy_listen_port
  public_ips                      = var.chef_server_public_ips
  ssh_key_file                    = var.ssh_key_file
  ssh_user                        = var.ssh_user
  ssh_user_sudo_password          = local.fe_sudo_password
  sudo_cmd                        = var.sudo_cmd
  teams_port                      = var.teams_port
}
