package crudoperation

import (
	"log"
	"strings"
	"github.com/go-ldap/ldap"
)

func modify(modRequest *ldap.ModifyRequest , lc *ldap.Conn) error {
	err := lc.Modify(modRequest)
	if err != nil {
		log.Println("Entry NOT modify", err)
	} else {
		log.Println("Entry MODIFIED", err)
	}
	return err
}

// LdapModify to delete the request
func LdapModify(lc *ldap.Conn, uid string, orgunit string, object string, userDC string) error {
	objectSlice := strings.Split(object, ",")
	userReq := uid + "," + orgunit + "," + userDC
	log.Println(userReq)
	modReq := ldap.NewModifyRequest(userReq, []ldap.Control{})
	modReq.Replace("objectClass", objectSlice)
	return modify(modReq, lc)
}