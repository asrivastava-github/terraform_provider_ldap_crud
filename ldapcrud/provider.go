package ldapcrud

import (
	"log"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"../ldap/client"
)

func errorHandle(err error, identifier string){
	if err !=nil {
		log.Fatalf("%v %v", identifier, err)
	}
}

// Provider function get the required values from terraform file
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_HOST", ""),
			},
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_PORT", ""),
			},
			"cn": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_CN", ""),
			},
			"basedn": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_DN", ""),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("LDAP_PASSWD", ""),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ldapcrud_operation": crudOperation(),
		},
		// responsible for making the connection
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	host := d.Get("host").(string)
	port := d.Get("port").(string)
	cn := d.Get("cn").(string)
	baseDN := d.Get("basedn").(string)
	passwd := d.Get("password").(string)
	// ldapClient, bindErr := lc.LDAPClient()
	// defer ldapClient.Close()
	return client.NewConn(host, port, cn, baseDN, passwd, false), nil
}
