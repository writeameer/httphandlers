package handlers

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-ldap/ldap"
)

var (
	n            = 1
	server       = os.Getenv("ldap_server")
	username     = os.Getenv("ldap_username")
	ldapPassword = os.Getenv("ldap_password")
	domainName   = os.Getenv("ldap_domainname")
	tlsConfig    = &tls.Config{InsecureSkipVerify: true} // to allow for internal CAs
	conn         ldap.Conn
)

// AuthenticationHandler proxies requests to service-now.com
func AuthenticationHandler(next ...http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Run next handler only if auth succeeds
		username, password, _ := r.BasicAuth()
		log.Println("Username:" + username)
		log.Printf("password length is: %d \n", len(password))

		if username == "" {
			http.Error(w, "No username was provided.", http.StatusForbidden)
			return
		}

		if password == "" {
			http.Error(w, "No username was provided.", http.StatusForbidden)
			return
		}

		status, err := ldapLogin(username, password)

		if status == http.StatusOK {
			next[0].ServeHTTP(w, r)
		} else if next != nil {
			http.Error(w, " unauthorized:"+err.Error(), http.StatusForbidden)
		}
	})
}

// ldapLogin authenticates the user's request against Active Directory
func ldapLogin(username string, password string) (status int, err error) {

	// Connect to LDAP server
	conn, err := ldap.DialTLS("tcp", server, tlsConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Bind user to directory
	err = conn.Bind(username+"@"+domainName, password)
	if err != nil {
		status = http.StatusForbidden
	} else {
		status = http.StatusOK
	}

	return
}
