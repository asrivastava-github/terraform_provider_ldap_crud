package crudoperation

import (
	"log"
	"github.com/go-ldap/ldap"
)

// func NewAddRequest(dn string, controls []Control) *AddRequest
// NewAddRequest returns an AddRequest for the given DN, with no attributes
// func (*AddRequest) Attribute
// func (req *AddRequest) Attribute(attrType string, attrVals []string)
// Attribute adds an attribute with the given type and values
// type Attribute
// type Attribute struct {
//     // Type is the name of the LDAP attribute
//     Type string
//     // Vals are the LDAP attribute values
//     Vals []string
// }

func del(delRequest *ldap.DelRequest , lc *ldap.Conn) error {
	err := lc.Del(delRequest)
	if err != nil {
		log.Println("Entry NOT deleted",err)
	} else {
		log.Println("Entry DELETED",err)
	}
	return err
}

// LdapDelete to delete the request DN	ou=users,dc=localldap,dc=com
func LdapDelete(lc *ldap.Conn, uid string, orgunit string, userDC string) error {
	userReq := uid + "," + orgunit + "," + userDC
	// log.Println(userReq)
	delReq := ldap.NewDelRequest(userReq, []ldap.Control{})
	return del(delReq, lc)
}