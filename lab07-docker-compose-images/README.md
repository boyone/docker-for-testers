# Multitier Architecture

1. Study the web application architecture

   - Requirements + `Diagram`: Services
   - Protocol + URI + PORT
   - Analysis and design how to use docker

2. Lab on local machine (Running with Docker compose)

   - Run Docker compose
   - Docker compose commands
   - Docker compose file(details)

   - [ ] image
   - [ ] container_name
   - [ ] build: context:
   - [ ] ports
   - [ ] environment:
   - [ ] volumes
   - [ ] restart

   ### Images hub.docker.com

   - [ ] nginx(frontend)
   - [ ] store-service
   - [ ] mysql
   - [ ] bank-gateway
   - [ ] shipping-gateway

3. Exercise

- Bank Gateway: `<source>`:`<destination>`
- Shipping Gateway: `<source>`:`<destination>`
- DBCONNECTION: `<username>`:`<password>`@(`<uri>`:`<port>`)/`<database>`
  > eg. root:password@(localhost:3306)/shopping-cart
- Database volumes : `<source>`:`<destination>` # for initial database

```sh
curl -i -H "Accept: application/json" http://localhost:8000/api/v1/product
```
