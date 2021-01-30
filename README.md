# Introduction

This is code for a very basic LDAP crud operation on LDAP using terraform and developing in-house terraform provider.

# TODO

Need to finish assigining the values to resource schema for comparision purpose. Refer link https://github.com/Pryz/terraform-provider-ldap/blob/master/resource_ldap_object.go

# How it works
- Build the provider in tmp directory but with prefix of dir "registry.terraform.io/hashicorp" followed by <provider_name>/<version>/<os>_<arch>/<terraform-provider>-<provider_name>. E.g below
    ``` cli
    go build -o /tmp/registry.terraform.io/hashicorp/ldapcrud/0.0.1/darwin_amd64/terraform-provider-ldapcrud
    ```

    Once the terraform provider is build. Run Terraform init from code base dir where the tf code exists.

    ```cli
    terraform init -plugin-dir=/tmp
    ```

# Steps

- go build -o terraform-provider-demo
- terraform provider snippet
    ``` terraform
    provider "example" {
        address = "http://localhost"
        port    = "3001"
        token   = "superSecretToken"
    }
    ```
- terraform resourse snippet
    ``` terraform
    resource "example_item" "test" {
        name = "this_is_an_item1"
        description = "this is an item1"
        tags = [
            "hello",
            "Me"
        ]
    }
    ```

# Go function bullets
1. NewConn will be needed to pass the values to struct UserDetails. while passing or calling from another function.
2. LDAPClient, any function need ()
3. If you are using a receiver for a function then function get tagged to it    like class of that function. Consider it as self of python. Calling a function  will be self.function_name
    ``` go
    func (lc *LDAPConn) enableDebugStr (val string){
        if lc.Debug {
            log.Printf("%v", val)
        }
    }
    ```
    ``` go
    lc.enableDebugStr('Error Message')
    ```

# Docker ldap:
- Reference to write docker-compose.yml: https://github.com/Ramhm/openldap 

# Misc

- Testing main.go

        ```go
        package main

        import (
            // "github.com/hashicorp/terraform/plugin"
            "./ldap/client"
            "github.com/go-ldap/ldap"
            "log"
            "./ldap/crudoperation"
        )

        // func main() {
        //     plugin.Serve(&plugin.ServeOpts{
        //         ProviderFunc: ldapcrud.Provider,
        //     })
        // }

        func errorHandle(err error, identifier string){
            if err !=nil {
                log.Fatalf("%v %v", identifier, err)
            }
        }


        func main() {
            lc := *client.NewConn("localhost", "389", "cn=admin", "dc=localldap,dc=com", "admin", false)
            ldapServer := lc.Hostname + ":" + lc.Port
            log.Printf("Connecting to Server: %v", ldapServer)
            // Docker image of openLDAP does not have TLS cert
            // conn, err := ldap.DialTLS("tcp", ldapServer, &tls.Config{InsecureSkipVerify: true})
            ldapClient, err := ldap.Dial("tcp", ldapServer)
            errorHandle(err, "connecting to LDAP")
            ldapUser := lc.BindUser + "," + lc.BaseDN
            ldapPass := lc.BindPassword
            ConnErr := ldapClient.Bind(ldapUser, ldapPass)
            if ConnErr != nil {
                log.Fatalln(ConnErr)
            }
            readEntry, err := crudoperation.LdapRead(ldapClient, "dc=localldap,dc=com", "user-wrt", "(&(objectClass=*))", "dn,cn,uid")
            addErr := crudoperation.LdapCreate(ldapClient, "uid=goldap-user4", "ou=admins,ou=uk,ou=secservice", "account,simpleSecurityObject", "dc=localldap,dc=com")
            delErr := crudoperation.LdapDelete(ldapClient, "uid=goldap-user4", "ou=admins,ou=uk,ou=secservice", "dc=localldap,dc=com")
            ModErr := crudoperation.LdapModify(ldapClient, "uid=goldap-user4", "ou=admins,ou=uk,ou=secservice", "account,simpleSecurityObject,uidObject", "dc=localldap,dc=com")
            //lc *ldap.Conn, uid string, ou []string, object []string, dc []string
            defer ldapClient.Close()
            for k, v := range readEntry {
                if k != "" {
                    log.Printf("%v: %v", k, v)
                }
            }
            log.Println(err)
            log.Println(delErr)
            log.Println(addErr)
            log.Println(ModErr)
        }
        ```

-----------------------------------

# MISC

CI: Azure DevOps Build Pipeline Build CI is not new to you
CD: Azure DevOps release pipeline —> Jira raised, All Run team needs to do is to create a release Branch with that Jira Number

			    Jira
		 		 |
		    Release branch with Jira
				 |
		    ADO pipeline triggered
		  		|
		ADO with Connect to Jira to fetch the details
		  		|
		        Create the config
		  		|
		        Validate the config
		 		|
        If validation passed, Create a PR (git request-pull release/Jira https://git.ko.xz/project master) No need for that as well but just a gatekeeper before the changes
		  		| 
		        PR Approved
		  		|
        Trigger Build ADO pipeline which will create Assyst record based on Jira template
		  		|
        And post Assyst is in implementation stage, triggers a Jenkins job to deploy the changes (Again Jenkins will have review gate keeping)
	  			|
			    /      \
	            Update/del	    New
  		        /		\
        Jenkins will 		First Jenkins will create the certificate and Trigger the Step Function as well
        handle		            \
					Step function is a set of lambda functions arranged in a logical formation to address the specific need.
					        \ 
					First lambda will keep monitoring the Pending certificate and keep sending notification to DNS admin for DCV
			                            \
					Upon Certificate issued stack implementation plan will be email to DevOps Team for verification and approval
			   			        \
					Upon Approval Lambda will update the listener certificates, create rules/Stack, attach respective certificates, Email the IPs to DNS Admin and requester
						            \
					Lambda will trigger ADO pipeline ? Possible ? For a testing.
						                \
					Merge PR and Close the Assyst record


