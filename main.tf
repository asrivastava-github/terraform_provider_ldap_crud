variable host           {}
variable port           {}
variable cn             {}
variable userdc         {}
variable password       {}
variable basedn         {}
variable objectclass    {}
variable uid            {}
variable searchfilter   {}
variable object         {}
variable orgunit        {}

provider "ldapcrud" {
    host        = var.host
    port        = var.port
    cn          = var.cn
    basedn      = var.basedn
    password    = var.password
}

resource "ldapcrud_operation" "user_operation" {
    userdc          = var.userdc
    objectclass     = var.objectclass
    searchfilter    = var.searchfilter
    uid             = var.uid
    object          = var.object
    orgunit         = var.orgunit
}