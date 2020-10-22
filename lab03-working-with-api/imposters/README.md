# Store Service Stub

## Prerequisite

- Node.js

  [Download](https://nodejs.org/en/download/) and Install

- Mountebank

```sh
npm install -g mountebank
```

## Commands

- Start Stub

```sh
mb start --configfile store-service.json --allowInjection &
```

- Stop Stub

```sh
mb stop
```

- Restart Stub

```sh
mb restart --configfile store-service.json --allowInjection &
```

## Test Stub

- Product List

```sh
curl -i -X GET localhost:8000/api/v1/product --header 'Accept: application/json'
```

- Product Detail #1

```sh
curl -i -X GET localhost:8000/api/v1/product/1  --header 'Accept: application/json'
```

- Product Detail #2

```sh
curl -i -x GET localhost:8000/api/v1/product/2  --header 'Accept: application/json'
```

- Order Product

```sh
curl -i -X POST localhost:8000/api/v1/order --header 'Accept: application/json' --header 'Content-Type: application/json' -d '{
        "cart":[
                {
                        "product_id": 2,
                        "quantity": 1
                }
        ],
        "shipping_method": "Kerry",
        "shipping_address": "405/37 ถ.มหิดล",
        "shipping_sub_district": "ท่าศาลา",
        "shipping_district": "เมือง",
        "shipping_province": "เชียงใหม่",
        "shipping_zip_code": "50000",
        "recipient_name": "ณัฐญา ชุติบุตร",
        "recipient_phone_number": "0970809292"
}'
```

- Confirm Payment

```sh
curl -i -X POST localhost:8000/api/v1/confirmPayment --header 'Accept: application/json' --header 'Content-Type: application/json' -d '{
        "order_id": 8004359104,
        "payment_type": "credit",
        "type": "visa",
        "card_number": "4719700591590995",
        "cvv": "752",
        "expired_month": 7,
        "expired_year": 20,
        "card_name": "Karnwat Wongudom",
        "total_price": 14.95
}'
```
