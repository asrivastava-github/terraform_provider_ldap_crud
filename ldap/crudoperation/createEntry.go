package crudoperation

import (
	"log"
	"strings"
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

// func removeComma(dcStr string) string {
// 	s := dcStr
// 	sz := len(s)

// 	if sz > 0 && s[sz-1] == ',' {
// 		s = s[:sz-1]
// 	}
// 	return s
// }

// func nestedEntry(attrType string, dc []string) string {
// 	var fullDC string
// 	for _, d := range dc {
// 		fullDC += attrType + "=" + d + ","
// 	}
// 	return removeComma(fullDC)
// }

func add(addRequest *ldap.AddRequest, lc *ldap.Conn) error {
	err := lc.Add(addRequest)
	if err != nil {
		log.Println("Entry NOT added", err)
	} else {
		log.Println("Entry ADDED", err)
	}
	return err
}

// LdapCreate function is to create the user
func LdapCreate(lc *ldap.Conn, uid string, orgunit string, object string, userDC string) error {
	objectSlice := strings.Split(object, ",")
	userReq := uid + "," + orgunit + "," + userDC
	log.Println(userReq)
	addReq := ldap.NewAddRequest(userReq, []ldap.Control{})
	addReq.Attribute("objectClass", objectSlice)
	addReq.Attribute("userPassword", []string{"readUser1"})
	addReq.Attribute("uid" ,[]string{uid})
	log.Println(addReq)
	// a.Attribute("uid" ,[]string{"readUser1"})
	// a.Attribute("email" ,[]string{"readUser1@gmail.com"})
	// a.Attribute("member" ,[]string{"cn=readonly,ou=groups,dc=mview,dc=example,dc=com"})
	// a.Attribute("name" ,[]string{"Guser"})
	// a.Attribute("sn" ,[]string{"XXX"})
	return add(addReq, lc)
}
