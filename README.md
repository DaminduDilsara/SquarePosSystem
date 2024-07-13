# Square POS API Integration

This project uses Square POS APIs to provide the following endpoints:

- `/location/create` - Create a new location
- `/order/create` - Create a new order
- `/orders/search` - Search for orders
- `/orders/find` - Find specific orders
- `/payment/create` - Create a payment

## Endpoints

### 1. Create Location
#### Sample cURL
```sh
curl --location 'localhost:8080/v1.0/location/create' \
--header 'Authorization: Bearer <auth_token> ' \
--header 'Content-Type: application/json' \
--data-raw '{
  "business_email": "business1@gmail.com",
  "description": "Business1"
}'
```
you can create locations using this endpoints and with the location ID, you can handle the orders separately

### 2. Create Order
#### Sample cURL
```sh
curl --location 'localhost:8080/v1.0/order/create' \
--header 'Authorization: Bearer  <auth_token> ' \
--header 'Content-Type: application/json' \
--data '{
    "location_id": "L3Q8ZJYHWS11N",
    "customer_id": "001",
    "reference_id": "001uiqydefiuvqe",
    "table_id": "002",
    "line_items": [
        {
            "catalog_object_id": "UXHZKPGKOKXRINRN2NNWYMYR",
            "name": "fish bun",
            "note": "toasted",
            "quantity": "5"
        }
    ]
}'
```
please note that the line_items catalog_object_id should be a pre created item from the square pos

### 3. Search Orders
#### Sample cURL
```sh
curl --location 'localhost:8080/v1.0/orders/search' \
--header 'Authorization: Bearer <auth_token> ' \
--header 'Content-Type: application/json' \
--data '{
    "location_id": "L3Q8ZJYHWS11N",
    "table_no": "002"
}'
```

in here, if you provide the table id, it will return all the orders relate for that table. if you leave table_id empty, it will return all the orders relate to the location

### 4. Find Orders with Order Id
#### Sample cURL
```sh
curl --location 'localhost:8080/v1.0/orders/find' \
--header 'Authorization: Bearer <auth_token> ' \
--header 'Content-Type: application/json' \
--data '{
    "order_ids": [
        "Z64rklvu3f4CRBHeBOggqXIdj4OZY"
    ],
    "location_id": "L3Q8ZJYHWS11N"
}'
```

you can provide multiple order ids

### 5. Make Payment For Order
#### Sample cURL
```sh
curl --location 'localhost:8080/v1.0/payment/create' \
--header 'Authorization: Bearer <auth_token> ' \
--header 'Content-Type: application/json' \
--data '{
    "order_id": "TrgENvxgUwAECYxGBjnIM3tRMyOZY",
    "bill_amount": 5000,
    "tip_amount": 10
}'
```