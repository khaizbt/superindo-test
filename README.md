# Superindo Pre Test Backend Engineer
Here are the results of the Superindo Backend Engineer pre test project

## Overview

1. This Repo using [Gin-Gonic](https://github.com/gin-gonic/gin)
2. Gorm for query
3. Database using MySQL
4. JWT for Authentication

## Installation

To Run this Repo, you need to install [Go](https://golang.org/dl/)
## Usage

1. Clone this repo using https ```git clone https://github.com/khaizbt/superindo-test.git```
2. in your bash, run ```cd Superindo-test```
3. You can use any RDBMS Database, for configuration the database is contained in the file [.env](https://github.com/khaizbt/superindo-test/-/blob/master/.env) and Docs of database in [here](https://gorm.io/docs/connecting_to_the_database.html)
4. If you have configured the Database and Secret Key on [.env](https://github.com/khaizbt/superindo-test/-/blob/master/.env) you can immediately run  ```go run main.go ``` use bash in Project Directory
5. If it's run successfully, you can import postman.json on your postman in the file [postman_collection](https://github.com/khaizbt/superindo-test/blob/main/Superindo.postman_collection.json) and  [postman_environtment](https://github.com/khaizbt/superindo-test/blob/main/Superindo.postman_environment.json ) or via the link [here](https://documenter.getpostman.com/view/12945074/UVeCQoNQ)
6. ### **Make sure you have turned off seeding when running the program if the database has been successfully migrated.**

## Technologies Used

- **Go**: The primary programming language for application development.
- **Goroutines**: To handle concurrent operations.
- **Mutex**: To avoid race conditions when accessing shared data.

## Code Explanation:
- Cart structure: Uses a map to store items and their quantities, and a mutex (sync.Mutex) to protect access to this data.

- Checkout Method: Locks the mutex before performing operations on the cart. Once the process is complete, the cart is emptied.
By using mutex, we ensure that when one goroutine is processing a checkout, no other goroutine can change the contents of the cart, thus maintaining data integrity.

## Features

- Add items to the cart with handle race condition
- Buy Now with handle race condition
- View cart contents
- Perform checkout
- Prevent race conditions when accessing shared cart data

## Thanks

Thank you Superindo team, sorry for the shortcomings in this project, I am waiting for good news from the Superindo team ;)
