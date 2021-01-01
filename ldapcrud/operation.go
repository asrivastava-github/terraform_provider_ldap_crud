package ldapcrud

import (
	"fmt"
	"regexp"
	"log"
	// "time"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/go-ldap/ldap"
	"../ldap/crudoperation"
	"../ldap/client"
)

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("name cannot contain whitespace. Got %s", value))
		return warns, errs
	}
	return warns, errs
}

func crudOperation() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Required:     true,
				Description: "The name of the resource, also acts as it's unique ID",
				ForceNew:    true,
				ValidateFunc: validateName,
			},
			"userdc": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource, also acts as it's unique ID",
				ValidateFunc: validateName,
			},
			"objectclass": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object classes for the user",
				ValidateFunc: validateName,
			},
			"searchfilter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter to be applied on a search",
			},
			"object": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object in which user has to be part of",
			},
			"orgunit": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "All the OU User has to be part of",
			},
		},
		Create: ldapCreateItem,
		Read:   ldapReadItem,
		Update: ldapModifyItem,
		Delete: ldapDeleteItem,
		Exists: ldapExistsItem,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}


func ldClient(m interface{}) *ldap.Conn {
	ldConn 		:= m.(*client.LDAPConn)
	ldapServer	:= ldConn.Hostname + ":" + ldConn.Port
	log.Printf("Connecting to Server: %v", ldapServer)
	// Docker image of openLDAP does not have TLS cert
	// conn, err := ldap.DialTLS("tcp", ldapServer, &tls.Config{InsecureSkipVerify: true})
	ldapClient, err := ldap.Dial("tcp", ldapServer)
	errorHandle(err, "connecting to LDAP")
	ldapUser 	:= ldConn.BindUser + "," + ldConn.BaseDN
	ldapPass 	:= ldConn.BindPassword
	ConnErr		:= ldapClient.Bind(ldapUser, ldapPass)
	if ConnErr != nil {
		log.Fatalln(ConnErr)
	}
	return ldapClient
}

var objectClass 	string
var searchfilter 	string
var object 			string
var orgunit			string
var uid 			string
var userDC 			string

func getValues(d *schema.ResourceData) (string, string, string, string, string, string) {
	userDC 			:= d.Get("userdc").(string)
	objectClass 	:= d.Get("objectclass").(string)
	searchfilter	:= d.Get("searchfilter").(string)
	object 			:= d.Get("object").(string)
	orgunit 		:= d.Get("orgunit").(string)
	uid 			:= d.Get("uid").(string)
	return userDC, objectClass, searchfilter, object, orgunit, uid
}

// m is meta data of client created in provider.go
func ldapCreateItem(d *schema.ResourceData, meta interface{}) error {
	userDC, objectClass, searchfilter, object, orgunit, uid = getValues(d)
	ldapClient := ldClient(meta)
	createErr := crudoperation.LdapCreate(ldapClient, uid, orgunit, object, userDC)
	// addErr := crudoperation.LdapCreate(ldapClient, "goldap-user4", []string{"admins", "uk", "secservice"}, []string{"account", "simpleSecurityObject"}, "dc=localldap,dc=com")
	// defer ldapClient.Close()
	if createErr != nil {
		log.Printf("Create Error: %v", createErr)
		return createErr
	}
	d.SetId(uid)
	// We can set OAN as id
	// d.SetId(uid)
	return nil
}

func ldapReadItem(d *schema.ResourceData, meta interface{}) error {
	userDC, objectClass, searchfilter, object, orgunit, uid = getValues(d)
	ldapClient := ldClient(meta)
	readEntry, err := crudoperation.LdapRead(ldapClient, userDC, uid, objectClass, searchfilter)
	// readEntry, err := crudoperation.LdapRead(ldapClient, "dc=localldap,dc=com", "user-wrt", "(&(objectClass=*))", []string{"dn", "cn", "uid"})
	log.Println(readEntry)
	// defer ldapClient.Close()
	if err != nil{
		log.Printf("Error %v", err)
	}
	d.SetId(uid)
	d.Set("object", object)
	d.Set("objectclass", objectClass)
	d.Set("orgunit", orgunit)
	d.Set("searchfilter", searchfilter)
	d.Set("uid", uid)
	d.Set("userdc", userDC)
	return nil
}

func ldapModifyItem(d *schema.ResourceData, meta interface{}) error {
	userDC, objectClass, searchfilter, object, orgunit, uid = getValues(d)
	ldapClient := ldClient(meta)
	modErr := crudoperation.LdapModify(ldapClient, uid, orgunit, object, userDC)
	// ModErr := crudoperation.LdapModify(ldapClient, "goldap-user4", []string{"admins", "uk", "secservice"}, []string{"account", "simpleSecurityObject", "uidObject"}, "dc=localldap,dc=com")
	// defer ldapClient.Close()
	if modErr != nil {
		log.Printf("Modify Error: %v", modErr)
		return modErr
	}
	return nil
}

func ldapDeleteItem(d *schema.ResourceData, meta interface{}) error {
	userDC, objectClass, searchfilter, object, orgunit, uid = getValues(d)
	ldapClient := ldClient(meta)
	delError := crudoperation.LdapDelete(ldapClient, uid, orgunit, userDC)
	// delErr := crudoperation.LdapDelete(ldapClient, "goldap-user4", []string{"admins", "uk", "secservice"}, "dc=localldap,dc=com")
	// defer ldapClient.Close()
	if delError != nil {
		log.Printf("Delete Error: %v", delError)
		return delError
	}
	d.SetId("")
	return nil
}

func ldapExistsItem(d *schema.ResourceData, meta interface{}) (bool, error) {
	userDC, objectClass, searchfilter, object, orgunit, uid = getValues(d)
	var exists bool
	ldapClient := ldClient(meta)
	readEntry, err := crudoperation.LdapRead(ldapClient, userDC, uid, objectClass, object)
	// readEntry, err := crudoperation.LdapRead(ldapClient, "dc=localldap,dc=com", "user-wrt", "(&(objectClass=*))", []string{"dn", "cn", "uid"})
	for k, v := range readEntry {
		if k == uid {
			log.Printf("%v: %v", k, v)
			exists = true
		} else {
			exists = false
		}	
	}
	if err != nil {
		log.Printf("Exists Check Error: %v", err)
		return exists, err
	}
	// defer ldapClient.Close()
	return exists, nil
}