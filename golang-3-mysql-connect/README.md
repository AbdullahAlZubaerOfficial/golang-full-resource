# 🚀 MySQL Setup & Go (Golang) Connection Guide

এই README ফাইলটি Windows এ **MySQL install করা**, **database ও user তৈরি করা**, এবং **Go দিয়ে MySQL connect করা**—এই পুরো প্রক্রিয়াটা step-by-step বুঝিয়ে দেবে 🧩💻  

---

## 🟢 Step 1: MySQL Download & Install

###  Download MySQL
- Official website থেকে MySQL download করো  
  👉 https://dev.mysql.com/downloads/mysql/
- **Windows এর জন্য MySQL Installer** ডাউনলোড করো

### ⚙️ Install MySQL
1. Installer ওপেন করো  
2. **Developer Default** select করো  
3. `Next` দিয়ে default settings এ install করো  
4. MySQL Server automatically run হবে  
5. **Root password সেট করো** (এই password পরে ব্যবহার হবে)  

**Example Password:**  
John1234


---

## 🟢 Step 2: MySQL Server Run & Command Line Open

###  MySQL Command Line Client
- Windows search bar এ গিয়ে **MySQL Command Line Client** ওপেন করো  
- Installation এর সময় দেওয়া **root password** ইনপুট করো  

### ✅ Test Connection
```sql
SHOW DATABASES;
➡️ এখানে existing database গুলো দেখতে পাবে 

🟢 Step 3: Database Create করা
আমরা simplerest নামে database ব্যবহার করবো।

🗄️ Database Create
CREATE DATABASE simplerest;
🔍 Check Database
SHOW DATABASES;
🟢 Step 4: User Create & Privileges Set
এখন john নামে একটি user তৈরি করে তাকে simplerest database এর full access দেবো 🔐

CREATE USER 'john'@'localhost' IDENTIFIED BY 'John1234';
GRANT ALL PRIVILEGES ON simplerest.* TO 'john'@'localhost';
FLUSH PRIVILEGES;
ℹ️ Explanation
localhost মানে এই user শুধু এই machine থেকেই connect করতে পারবে

FLUSH PRIVILEGES permission update করে

🟢 Step 5: Go (Golang) দিয়ে MySQL Connect করা
🔗 DSN (Data Source Name)
john:John1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local
 Go MySQL Driver Install
go get -u github.com/go-sql-driver/mysql
 Example Go Code
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
Explanation
sql.Open() → MySQL driver দিয়ে database connect করে

db.Ping() → connection ঠিক আছে কিনা চেক করে

💡 Tip (Important)
যদি password এ special character থাকে (#, +, @ ইত্যাদি), তাহলে backtick (`) ব্যবহার করো 

dsn := `john:John1234@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local`

➡️ Backtick ব্যবহার করলে পুরো string safe থাকে 
