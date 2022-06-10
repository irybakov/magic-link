# Intro
This is the PoC of using Magic-link with Go-based API

The `auth-magic` folder cotains login form and ufter login provides  the `token`.

To login user must enter email on which the auth link will be sent.

Click on the link provides the token on the client side (for example, browser)

# Token

To create token [Decentralized ID Token Specification](https://w3c-ccg.github.io/did-primer/#the-format-of-a-did) is used.
Token is signed with `public key`

# Access API
To access API client uses token in the `Authorization header`
```
curl --location --request GET 'http://localhost:8090/hello' \
--header 'Authorization: Bearer WyI... token ...jFiXCJ9Il0='
```

API (resource endpoints) should verify the `token` with `private key`


# How to launch the demo

### Start Auth service
 ```
 cd auth-magic
 yarn start
 ```
 Follow instractions to get the link.

 ### Start Go Api service
 ```
 cd .. 
 go run main.go
 ```

### Test all together
Copy token, replace `WyI... token ...jFiXCJ9Il0=` and make a request ot API
 ```
curl --location --request GET 'http://localhost:8090/hello' \
--header 'Authorization: Bearer WyI... token ...jFiXCJ9Il0='
```