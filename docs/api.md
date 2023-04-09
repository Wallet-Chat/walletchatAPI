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

# Encryption
Ethereum based wallets (EVM compatible) have encryption for all DMs, using LIT Protocol

# Software Stack
![WC_SW_Stack drawio](https://user-images.githubusercontent.com/19207330/227810284-83324964-58b3-4335-bb7b-0d2128a3d62c.svg)

# API Key Holders
ADMIN API KEY functions allow vetted customer integrations to streamline the user experience.
To authenticate with an API key, the ADMIN_API_KEY must be used in place of the end user JWTs.

## Important Note for Security 
*We ask API key holders to make these requests in a protected manner, mainly from their own API.*

If the ADMIN API key is used in the client browser, it may be misused by malicious actors. 

Below are examples of each API which has an ADMIN API Key overrride. `curl` is just used as example,
replace with equivalent functionality as needed.

## Example for <API>/v1/name 
Update the wallet address to name mapping:

```
curl --location 'https://api-segmint.walletchat.fun/v1/name'
--header 'Authorization: Bearer AdminTestKey123'
--header 'Content-Type: text/plain'
--data '{
"name": "Nftz4Life",
"address": "0x14ffE94d2B5Bf47a8d55D713b3d6b35039167cfb"
}'
```

## Example for <API>/v1/update_settings
Update the email address to wallet address mapping:

```
curl --location 'https://api-segmint.walletchat.fun/v1/update_settings'
--header 'Authorization: Bearer MyTestApiKey123'
--header 'Content-Type: text/plain'
--data-raw '{
"email": "savemynft@gmail.com",
"walletaddr": "0x14fcE94d2B5Bf47a8d54D713b3d6b35039167cfb"
}'
```