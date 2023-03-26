To interact with the live API via this page, you will need a valid JWT Token, see next section for details.

Some JSON data structures are shared for both input/return values.
Required input parameters will have a red * in the data type outline at
the bottom of the page, along with a comment. 
This means when executing API functionality from this API documentation page, some fields 
can be left blank or may be removed from the JSON struct before submitting.
Please email the developers with any issues.
Some JSON data structures are output only, and will be marked as such.

# JWT Authentication
Except for AUTH functions, all /v1 endpoints must include `Bearer <JWT>` token (requests showing the Lock Icon)

For this API Doc, use the `Authorize` button on the right hand side to enter `Bearer <JWT>` where the JWT will
come from the return value of the `/signin` endpoint. Please read the `/users/<>/nonce` and `/signin`
descriptions to understand the login workflow via JWT Auth. 

Another easy way to obtain a JWT for use here is to grab it from the web app:
1) Log into https://app.walletchat.fun, with the wallet address desired to act on behalf of
2) Right click the web page and select `inspect`
3) Go to the `Application` tab
4) Find local storage, and select https://app.walletchat.fun 
5) Select the value of `jwt` or `jwt_*` variable, this is your JWT for the signed in wallet

#Software Stack

![WC_SW_Stack drawio](https://user-images.githubusercontent.com/19207330/227810284-83324964-58b3-4335-bb7b-0d2128a3d62c.svg)
