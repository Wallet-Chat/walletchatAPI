# walletchatAPI
API side of walltchatFE

To run both front-end and API on localhost and use Chrome, npm package local-cors-proxy is required:

lcp --proxyUrl http://localhost:8080 --origin http://localhost:3000 --credentials 

Point to new proxy API in walletchatFE .env file:
REACT_APP_REST_API=http://localhost:8010/proxy

Replace your port numbers as needed, the are the defaults.

DO NOT USE cors proxy in production, allow correct headers/origins from deployed infrastructure.
