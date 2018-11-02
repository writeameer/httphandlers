# Overview

Some sample http middleware


## Authentication Handler

[authenticationHandler](./handlers/authenticationHandler.go) authenticates against LDAP


### Setup credential file

Create a file called `~/.ldap`
```
export ldap_username="<your lan id>" 
export ldap_password="<your password>"
export ldap_server="blah.com:636"
export ldap_domainname="blah.com"
```

Dot source the file to load the creds in memory:

```
source ~/.ldap
```


### Start the server

Clone this repo then type:

```
go run main.go
```

### Test Login

- Test with your username (successful auth):

```
curl --user $ldap_username:$ldap_password http://localhost:8080
```

Yo should get output similar ot the following:

```
<HTML><HEAD><meta http-equiv="content-type" content="text/html;charset=utf-8">
<TITLE>301 Moved</TITLE></HEAD><BODY>
<H1>301 Moved</H1>
The document has moved
<A HREF="https://www.google.com/">here</A>.
</BODY></HTML>
```


Test with non-existant username (successful auth):

```
curl --user blah:blah http://localhost:8080
```

You should get output similar to the following:

```
unauthorized:LDAP Result Code 49 "Invalid Credentials"...
```


## Reverse Proxy  Handler

[reverseProxyHandler](./handlers/reverseProxyHandler.go) just forwards requests to `originHost`. In the following example, `google.com` is set as the origin host so all traffic to localhost:8080 will be proxied to Google.

```
func main() {

	// Create new router
	mux := http.NewServeMux()

	originHost := "google.com"
	mux.Handle("/", handlers.AuthenticationHandler(handlers.ReverseProxyHandler(originHost)))

	// Listen and Server
	port := ":8080"
	log.Println("Server started on port" + port)
	log.Fatal(http.ListenAndServe(port, mux))
}

```
