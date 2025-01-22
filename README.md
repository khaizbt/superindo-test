# Superindo Pre Test Backend Engineer
Here are the results of the Superindo Backend Engineer pre test project

## Technology Used
![Project Structure](https://raw.githubusercontent.com/khaizbt/superindo-test/refs/heads/develop/helper/Screenshot%202025-01-22%20at%2019.20.32.png)
1. This Repo using [Gin-Gonic](https://github.com/gin-gonic/gin)
2. Gorm for query
3. Database using MySQL
4. Redis for save temporary data
5. JWT for Authentication

## Installation

To Run this Repo, you need to install [Go](https://golang.org/dl/)
## Usage

1. Clone this repo using https ```git clone https://github.com/khaizbt/superindo-test.git```
2. in your bash, run ```cd superindo-test```
3. run ```cp .env.example .env```
4. You can use MySql, for configuration the database is contained in the file [.env](https://github.com/khaizbt/superindo-test/-/blob/master/.env) and Docs of database in [here](https://gorm.io/docs/connecting_to_the_database.html)
5. You can use Redis for save temporary data. for configuration is contained in the file  [.env](https://github.com/khaizbt/superindo-test/-/blob/master/.env)
6. If you have configured the Database and Secret Key on [.env](https://github.com/khaizbt/superindo-test/-/blob/master/.env) you can immediately run  ```go run main.go ``` use bash in Project Directory
7. If it's run successfully, you can import postman.json on your postman in the file [postman_collection](https://github.com/khaizbt/superindo-test/blob/main/Superindo.postman_collection.json) and  [postman_environtment](https://github.com/khaizbt/superindo-test/blob/main/superindo.postman_environment.json ) or via the link [here](https://documenter.getpostman.com/view/12945074/2sAYQdjVbW)



## Code Explanation:
- ~/controller/* is used for code related to context, middleware, validation, etc
- ~/config/* is file used to contains config file that are used in the application
- ~/entity/* is file used to handle struct for user input and response to user
- ~/helper/* is file used to contains helper files that are used in the application
- ~/middleware/* is file used to handle middleware ex : auth, validating user agent
- ~/model/* is file used to handle struct for database
- ~/repository/* is used for code related to the database, such as create, read, update and delete.
- ~/route/* is route file used to handle URLs requested in the web application
- ~/service/* is used for code related to workflow, logic, etc


## Features

- MySql for save all data User, Product, Category
- Redis for save temporary data like UserID that is currently logged in, This is of course so that the system does not always retrieve repeated data in the MySql database which impacts the speed of a system. 

## Thanks

Thank you Superindo team, sorry for the shortcomings in this project, I am waiting for good news from the Superindo team ;)
