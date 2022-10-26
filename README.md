<div align="center">
  <h1>Kutsuya-E-Commerce-App</h1>
</div>
</br>

## About
Kutsuya is an online shop that provides products in the form of shoes. In this application, the user can act as a buyer or seller.

## Tech Stack
- Go
- Echo Framework
- MySQL
- GORM
- Docker
- AWS


## Installation
1. Clone this project

```
git clone https://github.com/muhdwiar/Kutsuya-E-Commerce-App.git
```
2. Open the Kutsuya-E-Commerce-App folder and run this code:

```
cd Kutsuya-E-Commerce-App
rm -rf .git
git init
```

## Features
In some features, using JWT Token. This token is obtained when logging in. To find out the complete documentation of each feature, please check the provided <a href="https://app.swaggerhub.com/apis-docs/raorafarhan/open-api_kutsuya/1.0.0#/" target="_blank">Swagger</a>. The available features are as follows:

### :adult: User
| Endpoint | Feature | Description | Token |
| --- | --- | --- | --- |
| /users | POST | User register to make account | :x: |
| /login | POST | Login for user, as buyer or seller | :x: |
| /users | GET | Open user profile | :heavy_check_mark: |
| /users | PUT | Edit user profile or change pasword | :heavy_check_mark: |
| /users | DELETE | Delete user account | :heavy_check_mark: |

### :athletic_shoe: Products
| Endpoint | Feature | Description | Token |
| --- | --- | --- | --- |
| /products | POST | Add new product | :heavy_check_mark: |
| /products | GET | Show all products | :x: |
| /products/:id | GET | Get detail of a product | :x: |
| /products/:id | PUT | Edit detail of a product | :heavy_check_mark: |
| /products/:id | DELETE | Delete product data | :heavy_check_mark: |

### :shopping_cart: Shopping Cart
| Endpoint | Feature | Description | Token |
| --- | --- | --- | --- |
| /carts | POST | Add a product to user cart | :heavy_check_mark: |
| /carts | GET | Show all products in user cart | :heavy_check_mark: |

### :shopping: Order History
| Endpoint | Feature | Description | Token |
| --- | --- | --- | --- |
| /orders | POST | Orders product in the cart | :heavy_check_mark: |
| /cancel | POST | Cancel order product in the cart | :heavy_check_mark: |
| /orders | GET | Show all order history | :heavy_check_mark: |

## Collaborator
* Raora Farhan Al Abrar ~ [Github](https://github.com/raorafarhan) ~ [LinkedIn](https://www.linkedin.com/in/raorafarhanalabrar/)
* Muhamad Dwi Arifianto ~ [Github](https://github.com/muhdwiar) ~ [LinkedIn](https://www.linkedin.com/in/muhamad-dwi-arifianto-b76147238/)
