/*
//* Database query logic।
//* Database access logic write here.
//* In future if you want to add row sql or chage to the framework and add new framework here
*/

package users

import (
	"errors"

	"gorm.io/gorm"
)

var ErrAlreadyExist = errors.New("user email already exists")

type Repository interface{
	CreateUser(u *User) error
}

// repository take db 
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(user *User) error {
	result := r.db.Create(user)

	if result.Error != nil {

		if errors.Is(result.Error, gorm.ErrDuplicatedKey){
			return ErrAlreadyExist
		}

		return result.Error
	}
	
	return nil
}

/*


পুরো কোডটা কী করছে?

এটা একটা **User Repository layer**
👉 মানে database-এর সাথে সব communication এখান থেকে হবে

---

# 🧠 Package

```go id="pkg1"
package users
```

👉 এই code `users` নামে একটি module/package এর অংশ

📌 মানে:

* user-related সব logic এখানে থাকবে
* অন্য file থেকে `users` হিসেবে import করা যাবে

---

# ⚠️ Error define করা

```go id="err1"
var ErrorAleradyExist = errors.New("User email already exist")
```

## 👉 মানে কী?

* একটা custom error বানানো হয়েছে
* যদি same email আগে থেকে থাকে তাহলে এই error return করবে

📌 সহজ ভাষায়:
👉 "এই user আগে থেকেই আছে"

---

# 🧱 Struct (Repository implementation)

```go id="st1"
type repository struct {
	db *gorm.DB
}
```

## 👉 মানে কী?

* এটা একটা struct (object)
* এর ভিতরে আছে:

  * `db` → database connection

📌 সহজভাবে:
👉 এই struct database handle করবে

---

# 🔌 Interface (Contract)

```go id="if1"
type Repository interface{
	CreateUser(u *User) error
}
```

## 👉 মানে কী?

এটা একটা rule set:

👉 যেই struct এই interface implement করবে
👉 তার অবশ্যই `CreateUser()` function থাকতে হবে

📌 মানে:

* এটা বলে দিচ্ছে “কি করতে হবে”
* কিন্তু “কিভাবে করবে” না

---

# 🏗️ Constructor function

```go id="fn1"
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
```

## 👉 মানে কী?

এটা object create করার function

### কাজ:

* বাইরে থেকে database নেয়
* repository struct তৈরি করে
* return করে interface হিসেবে

📌 সহজভাবে:
👉 database inject করে repository বানানো হচ্ছে

---

# 🚀 Main function (CreateUser)

```go id="fn2"
func (r repository) CreateUser(user User) error {
```

## 👉 মানে কী?

* এটা repository struct-এর method
* নাম: CreateUser
* input: user data
* output: error (থাকলে)

---

## 💾 Database insert

```go id="db1"
result := r.db.Create(user)
```

## 👉 মানে কী?

* user কে database-এ insert করা হচ্ছে

📌 সহজভাবে:
👉 নতুন user save করা হচ্ছে DB-তে

---

# ❌ Error check

```go id="err2"
if result.Error != nil {
```

👉 যদি database operation fail করে

---

## ❗ Duplicate email check

```go id="err3"
if errors.Is(result.Error, gorm.ErrDuplicatedKey){
	return ErrorAleradyExist
}
```

## 👉 মানে কী?

* যদি email already exist করে
* তাহলে custom error return করবে

📌 সহজভাবে:
👉 "এই email আগে থেকেই আছে"

---

## ❗ Other error return

```go id="err4"
return result.Error
```

👉 অন্য যেকোনো DB error return করবে

---

## ✅ Success case

```go id="ok1"
return nil
```

👉 কোন error না থাকলে:
✔ success

---

# 🧠 পুরো flow সহজভাবে

### যখন user create হবে:

1. user data আসবে
2. DB-তে insert হবে
3. যদি email duplicate হয় → error দিবে
4. অন্য error হলে সেটাই দিবে
5. success হলে nil return

---

# 🔥 Real life example

ধরো তুমি signup করছো:

### Case 1:

✔ নতুন email
👉 user save হবে

### Case 2:

❌ same email আবার দিলা
👉 "User email already exist"

### Case 3:

❌ database problem
👉 DB error return

---

# ⚡ Important concept (Interview point)

এই code শেখায়:

### ✔ Repository pattern

👉 database logic আলাদা layer-এ রাখা

### ✔ Dependency injection

👉 DB বাইরে থেকে pass করা

### ✔ Error handling

👉 custom + system error handle করা

### ✔ Interface usage

👉 future-এ DB change করলেও code break হবে না

---

# 🚀 One line summary

👉 এই code database-এ user create করে এবং duplicate email handle করে clean architecture style-এ

---

যদি চাও, আমি next step এ তোমাকে দেখাতে পারি:

👉 এই code-এর **clean architecture diagram**
👉 service layer কিভাবে add করবে
👉 real production Go backend structure

বললেই আমি সেটাও সহজভাবে বুঝিয়ে দেব 👍


*/