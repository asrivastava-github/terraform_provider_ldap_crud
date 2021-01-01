uid         = "uid=terraform-user10"
host        = "localhost"
port        = "389"
# User name to connect to ldap
cn          = "cn=admin"
# Where is user being created/Deleted/Modified exists 
userdc      = "dc=localldap,dc=com"
# Where is Admin user resides
basedn      = "dc=localldap,dc=com"
objectclass = "(&(objectClass=*))"
# What all to search for
searchfilter= "dn,cn,ou,uid"
# What will be object class of user being created. this will be part of modify request with .Replace
object      = "account,simpleSecurityObject"
# Organisation unit of user being created
orgunit     = "ou=admins,ou=uk,ou=secservice"
# Need to be sent via -var option
password    = "admin"