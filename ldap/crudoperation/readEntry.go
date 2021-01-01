package crudoperation

import (
	"log"
	"github.com/go-ldap/ldap"
	"strings"
)

func errorHandle(err error, identifier string){
	if err !=nil {
		log.Fatalf("Error in %v %v", err, identifier)
	}
}

//ScopeBaseObject 
const (
	ScopeBaseObject		= 0
	ScopeSingleLevel	= 1
	ScopeWholeSubtree	= 2
)

// NeverDerefAliases
const (
	NeverDerefAliases	= 0
	DerefInSearching	= 1
	DerefFindingBaseObj	= 2
	DerefAlways			= 3
)

// SearchRequest is struct
type SearchRequest struct {
    BaseDN       string
	Scope        int
	Filter       string
	Attributes   []string
}

// LdapRead values for the attributes of item from LDAP
func LdapRead(lc *ldap.Conn, baseDN string, uid string, objClass string, filter string) (map[string]string, error) {
	filterSlice := strings.Split(filter, ",")
	searchRequest := ldap.NewSearchRequest(
		baseDN, 											// The base dn to search
		ScopeWholeSubtree, NeverDerefAliases, 0, 0, false,
		objClass,											// The filter to apply
		filterSlice,										// A list attributes to retrieve
		nil,
	)

	searchResults, err := lc.Search(searchRequest)
	userLoc := make(map[string]string)
	uidLoc := make(map[string]string)
	errorHandle(err, "Search Ldap")

	if len(searchResults.Entries) != 0 {
		for _, entry := range searchResults.Entries {
			cn := entry.GetAttributeValue("cn")
			userid := entry.GetAttributeValue("uid")
			if cn != "" {
				if cn == uid {
					uidLoc[cn] = entry.DN
				}
				userLoc[cn] = entry.DN
				// log.Printf("%v: %v", entry.DN, entry.GetAttributeValue("cn"))
			}
			if userid != "" {
				if userid == uid {
					uidLoc[userid] = entry.DN
				}
				userLoc[userid] = entry.DN
				// log.Printf("%v: %v", entry.DN, entry.GetAttributeValues("uid"))
			}
		}
	} else {
		log.Println("No Entry found in LDAP.")
	}

	// log.Println(userLoc)
	if uid != "" {
		return uidLoc, err
	}
	return userLoc, err
}