package client
// package main

import (
	"log"
	// "crypto/tls"
	"github.com/go-ldap/ldap"
)

func errorHandle(err error, identifier string){
	if err !=nil {
		log.Fatalf("Error in %v %v", err, identifier)
	}
}

// LDAPConn struct contains the info about client
type LDAPConn struct {
	Hostname     	string `json:"hostname"`
	Port         	string `json:"port"`
	BindUser    	string `json:"bindUser"`
	BaseDN       	string `json:"baseDN"`
	BindPassword 	string
	Debug 			bool
}

// UserDetails struct contains the info about User details
type UserDetails struct {
	UID          string `json:"Uid"`
	CN           string
	SN           string
	Mail         string
	UserPassword string
}

// NewConn will be needed to pass the values to struct UserDetails. while passing or calling from another function
func NewConn(hostname string, port string, binduser string, baseDN string, bindPassword string, debug bool) *LDAPConn {
	return &LDAPConn{
		Hostname:   	hostname,
		Port:       	port,
		BindUser:		binduser,
		BaseDN:			baseDN,
		BindPassword:	bindPassword,
		Debug:			debug,
	}
}

func (lc *LDAPConn) enableDebugStr (val string){
	if lc.Debug {
		log.Printf("%v", val)
	}
}

func (lc *LDAPConn) enableDebugldapConn (val *ldap.Conn){
	if lc.Debug {
		log.Printf("%v", val)
	}
}

// func main() {
// 	lc := *NewConn("localhost", "389", "cn=admin", "dc=localldap,dc=com", "None", "None", "admin", true)
// 	ldapClient := lc.LDAPClient()
// 	log.Println(ldapClient)
// 	// ldapClient.Seach()

// 	// log.Println(client)
// 	defer ldapClient.Close()
// }
