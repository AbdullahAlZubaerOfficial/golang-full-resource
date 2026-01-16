Step 1: MySQL Download & Install 

Download MySQL:

Official website theke download koro: https://dev.mysql.com/downloads/mysql/

Windows er jonno MySQL Installer download koro.

Install MySQL:

Installer open koro, Developer Default select koro.

Next kore default settings e install koro.

MySQL server run korbe automatically.

Root password set korte bolbe: ei password porer step e tumi use korbe.
(Example: John1234)

Step 2: MySQL Server Run & Command Line Open 

Command Prompt / MySQL Shell:

Windows e search bar e MySQL Command Line Client open koro.

Password input koro (je password tumi installation e diyecho).

Test Connection:

SHOW DATABASES;


Ekhane existing databases dekhte parbe.

Step 3: Database Create করা 

Tumi simplerest naam e database use korcho. Create korte:

CREATE DATABASE simplerest;


Check korte:

SHOW DATABASES;

Step 4: User Create & Privileges Set 

Tumi john naam user banaccho. Ei user ke full access dao simplerest database e:

CREATE USER 'john'@'localhost' IDENTIFIED BY 'John1234';
GRANT ALL PRIVILEGES ON simplerest.* TO 'john'@'localhost';
FLUSH PRIVILEGES;


Ekhane localhost mane ei user sudhu tumi je machine e install korcho sekhanei connect korte parbe.

FLUSH PRIVILEGES; user permission update kore.

Step 5: Go Te MySQL Connect করা 

Tumi dsn likhecho:

dsn := "john:John1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"


Go MySQL Driver Install:

go get -u github.com/go-sql-driver/mysql


Example Code:

package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    dsn := "john:John1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
    
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Connection failed:", err)
    }

    fmt.Println("Successfully connected to MySQL!")
}


sql.Open MySQL driver diye database connect kore.

db.Ping() check kore connection thik ache kina.

💡 Tip:

Password e special characters (#, +) thakle, tumi quotes use korte paro like:

dsn := `john:John1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local`


Backtick (`) diye puro string safe thakbe.
