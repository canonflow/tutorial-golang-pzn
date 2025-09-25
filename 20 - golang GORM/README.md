# Golang Object Relational Mapping

## Pengenalan GORM

- ORM atau Object Relational Mapping adalah teknik untuk memetakan data dari database relational ke dalam object dalam pemrograman.
- Ketika kita belajar Golang Database, kita belajar pattern bernama **Repository**, yang digunakan sbg jembatan komunikasi ke Database.
- Saat meembuat Repository, kita membuat struct `Entity` sebagai **representasi** dari **table** di database.
- Hal itu sebenarnya sudah bisa dibilang sebuah `ORM`, namun masih dilakukan secara **manual**.

### Diagram Repository Pattern

![Diagram Repository Pattern](./assets/1.png)

### GORM

- GORM adalah salah satu library untuk implementasi ORM secara otomatis di Golang.
- Dengan menggunakan GORM, kita bisa fokus membuat pemetaan struct `Entity`, **tanpa harus memikirkan** detail dari implementasi `SQL` yang harus kita buat **untuk memanipulasi datanya**.
- GORM juga **mendukung relasi** antar `Entity` atau `Table`, baik **One to One**, **One to Many**, sampai **Many to Many**
- [https://gorm.io](https://gorm.io)

---

## Database Connection

- Untuk membuat koneksi ke database, kita bisa menggunakan `gorm.Open()`.
- **Tiap database** memiliki `config` masing-masing, kita bisa lihat semua config di database pada halaman: [https://gorm.io/docs/connecting_to_the_database.html](https://gorm.io/docs/connecting_to_the_database.html)

### Kode: Database Connection

```go

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
    dialect := mysql.Open("root:@tcp(127.0.0.1:3306)/belajar_golang_gorm?charset=utfmb4&parseTime=True&loc=Local")

    db, err := gorm.Open(dialect, &gorm.Config{})
    if err != nil {
        panic(err)
    }

    return db
}
```

### Kode: Test Database Connection

```go
var db = OpenConnection()

func TestConnection(t *testing.T) {
    assert.NotNil(t, db)
}
```

---

## Raw SQL

- Raw SQL artinya membuat query SQL secara manual.
- Terdapat **2 jenis** SQL, untuk melakukan Query (Select) atau untuk mengubah data (Insert, Update, Delete).
- Untuk melakukan query, kita bisa menggunakan method `Raw(sql)` di `gorm.DB`.
- Dan untuk melakukan manipulasi data, kita bisa gunakan method `Exec(sql)` di `gorm.DB`.

### Kode: Table Sample

```sql
create table sample (
    id varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) engine = InnoDB
```

### Kode: Execute SQL

```go
func TestExecuteSQL(t *testing.T) {
    err := db.Exec("insert into sample(id, name) values (?, ?)", "1", "nathan").Error

    assert.Nil(t, err)

    err = db.Exec("insert into sample(id, name) values (?, ?)", "2", "garzya").Error

    assert.Nil(t, err)

    err = db.Exec("insert into sample(id, name) values (?, ?)", "3", "santoso").Error

    assert.Nil(t, err)

    err = db.Exec("insert into sample(id, name) values (?, ?)", "4", "canonflow").Error

    assert.Nil(t, err)
}
```

### Kode: Query SQL

```go

type Sample struct {
    ID string
    Name string
}

func TestRawSQL(t *testing.T) {
    var sample Sampel

    err := db.Raw("select id, name from sample where id = ?", "1").Scan(&sample).Error

    assert.Nil(t, err)
    assert.Equal(t, "1", sample.ID)

    var samples []Sample
    err = db.Raw("select id, name from sample").Scan(&samples).Error

    assert.Nil(t, err)
    assert.Equal(t, 4, len(samples))
}
```

### `sql.Row` & `sql.Rows`

- GORM sendiri sebenarnya didalamnya tetap menggunakan package `sql` bawaan dari Golang.
- Jika kita ingin mendapatkan hasil query dalam bentuk `sql.Rows`, kita bisa menggunakan method `Rows()` setelah melakukan query.

### Kode: `sql.Row`

```go
func TestSQLRow(t *testing.T) {
    var samples []Sample

    rows, err := db.Raw("select id, name from sample").Rows()
    assert.Nil(t, err)
    defer rows.Close()

    for rows.Next() {
        var id string
        var name string

        err := rows.Scan(&id, &name)
        assert.Nil(t, err)

        samples = append(samples, Sample{
            ID: id,
            Name: name
        })
    }

    assert.Equal(t, 4, len(samples))
}
```

### Kode: `gorm.DB.ScanRows()`

```go
func TestScanRows(t *testing.T) {
    var samples []Sample

    rows, err := db.Raw("select id, name from sample").Rows()
    assert.Nil(t, err)
    defer rows.Close()

    for rows.Next() {
        err := db.ScanRows(rows, &samples)
        assert.Nil(t, err)
    }

    assert.Equal(t, 4, len(sample))
}
```

---

## Model

- Model atau Entity adalah `Struct` **representasi** dari table di database.
- Saat kita membuat tabel di database, direkomendasikan dibuatkan struct representasinya.
- Hal ini agar kita **tidak perlu** melakukan pembuatan perintah SQL **secara manual lagi**.

### Kode: Tabel User

```sql
create table users (
    id varchar(100) not null,
    password varchar(100) not null,
    name varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB
```

### Kode: User Entity

```go
import "time"

type User struct {
    ID string
    Password string
    Name string
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Convention

- Saat kita membuat struct, **secara default** GORM akan melakukan **mapping secara otomatis**, dimana:
  - **nama tabel** akan dipilih dari **nama struc**t menggunakan **lower_case jamak**.
  - **nama kolom** akan dipilih menggunakan **lower_case**.
- Selain itu, **secara otomatis** GORM akan memilih **field ID** sebagai **Primary Key**.
- Namun, sebenarnya **disarankan** kita **deklarasi secara manual** menggunakan `tag` dibanding secara otomatis.
- [https://gorm.io/docs/models.html#Fields-Tags](https://gorm.io/docs/models.html#Fields-Tags)

### Kode: User Entity dengan Tag

```go
type User struct {
    ID string `gorm:"primaryKey;column:id"`
    Password string `gorm:"column:password"`
    Name string `gorm:"name"`
    CreatedAt time.Time `gorm:"created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
}
```

### Table Name

- Secara default, nama table akan menggunakan **lower_case** dan **jamak**.
- Misal struct `User` akan menggunakan table `users`.
- Misal struct `OrderDetail` akan menggunakan table `order_details`.
- Namun, jika kita ingin menggunakan manual nama tabel-nya, kita bisa menggunakan **interface Tabler**, yang **mewajibkan** membuat **method** dengan nama `TableName()`.

### Kode: User Entity

```go
type User struct {
    ID string `gorm:"primaryKey;column:id"`
    Password string `gorm:"column:password"`
    Name string `gorm:"name"`
    CreatedAt time.Time `gorm:"created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *User) TableName() string {
    return "users"
}
```

---

## Model Conventions

- Saat kita membuat struct untuk Model, jika kita **ingin mengikuti aturan** dari GORM, sebenarnya kita tidak perlu menggunakan `tag`.
- Sekarang kita akan bahas beberapa convention yang digunakan oleh GORM.

### ID Sebagai Primary Key

- GORM secara default menggunakan **field** `ID` sbg **primary key**.
- Jika kita membuat field `ID` di struct Model, secara default akan digunakan sbg primary key di tabel.
- Jika kita ingin memilih field lain selain ID, maka kita harus menggunakan `tag` GORM `gorm:"primaryKey"`.

### Table Name

- GORM secara default akan menggunakan nama tabel dengan format `snake_cases` (**lowercase**, menggunakan `_` sebagai pemisah, dan dijadikan **jamak**).
- Misal untuk struct `User`, secara default akan menggunakan nama tabel `users`.
- Untuk struct `OrderDetail` secara default akan menggunakan nama tabel `order_details`.
- Jika kita ingin **mengubah nama default tabel nya**, maka kita harus implementasi **interface Tabler** dengan menambahkan method `TableName()`.

### Column Name

- GORM secara default akan menggunakan **snake_case** untuk nama kolom dari struct Model yang kita buat.
- Misal jika field `ID` artinya kolomnya id.
- Jika field `FirstName` artinya kolomnya `first_name`.
- Jika kita ingin mengubah nama kolom, kita bisa menggunakan tag `gorm:"column:nama_kolom"`.

### Timestamp Tracking

- GORM memiliki fitur `timestamp`, yaitu melakukan **perubahan otomatis** untuk **waktu dibuat** dan **diubah**-nya Model menggunakan field `CreatedAt` dan `UpdatedAt`.
- `CreatedAt` dan `UpdatedAt` **secara default** akan menggunakan `time.Now()` ketika dibuat
- `UpdatedAt` akan selalu diubah menjadi `time.Now()` ketika diupdate datanya.
- jika kita ingin mengubah nilai field-nya, kita bisa tambahkan `gorm:"autoCreateTime:true"` untuk `CreatedAt`.
- Dan menggunakan `gorm:"autoUpdateTime:True"` untuk `UpdatedAt`.

---

## Field Permission

- Secara **default**, semua field di struct Model **akan dianggap** kolom di tabel.
- Dan **semua perubahan di field**, akan **diupdate** ke tabel di database.
- Namun, kadang - kadang mungkin **kita ingin membuat field** yang **tidak merepresentasikan kolom di tabel** sehingga **tidak perlu** di `create` / `update`.
- Atau **mungkin terdapat kolom** yang **tidak perlu diupdate** lagi.
- Untuk mendukung kasus seperti itu, GORM menyediakan **Field Permission** menggunakan tag `gorm`.

### GORM Field Permission

| Tag  | Keterangan                                                                                                        |
| :--: | :---------------------------------------------------------------------------------------------------------------- |
| `<-` | **WRITE PERMISSION**, `<-:create` untuk create only, `<-:update` untuk update only, `<-` untuk create dan update. |
| `->` | **READ PERMISSION**, `->:false` untuk no read permission.                                                         |
| `-`  | **Ignore field ini**, tidak ada `write` / `read` permission.                                                      |

### Kode: User Entity

```go
type User struct {
    ID string `gorm:"primaryKey;column:id;<-:create"`
    Password string `gorm:"column:password"`
    Name string `gorm:"name"`
    CreatedAt time.Time `gorm:"created_at;autoCreateTime;<-:create"`
    UpdatedAt time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
    Information string `gorm:"-"`
}

func (u *User) TableName() string {
    return "users"
}
```

---

## Embedded Struct

- **Secara default**, field di Model akan **menjadi kolom di tabel**, kecuali yang menggunakan field `-` permission.
- Namun bagaimana jika ternyata field yang kita buat sangat banyak?
- Kadang-kadang, ada baiknya kita **simpan field yang sejenis** dalam `struct` **yang berbeda**
- Untungnya, GORM memiliki fitur bernama **embedded struct**, dimana kita bisa melakukan **embed struct** dalam field di Model, sehingga **seluruh isi field** di embedded struct akan **dianggap field di Model utamanya**.

### Kode: Alter Table `users`

```sql
alter table users
    rename column name to first_name;
alter table users
    add column middle_name varchar(100) null after first_name;
alter table users
    add column last_name varchar(100) null after middle_name;
```

### Kode: Struct Name

```go
type Name struct {
    FirstName string `gorm:"column:first_name"`
    MiddleName string `gorm:"column:middle_name"`
    LastName string `gorm:"column:last_name"`
}
```

### Kode: User Model

```go
type User struct {
    ID string `gorm:"primaryKey;column:id;<-:create"`
    Password string `gorm:"column:password"`
    Name Name `gorm:"embedded"`
    CreatedAt time.Time `gorm:"created_at;autoCreateTime;<-:create"`
    UpdatedAt time.Time `gorm:"updated_at;autoCreateTime;autoUpdateTime"`
    Information string `gorm:"-"`
}

func (u *User) TableName() string {
    return "users"
}
```

---

## Create

- Untuk **memasukkan** data ke database, kita **tidak perlu membuat** SQL Insert **secara manual**.
- GORM bisa membaca data dari Model yang sudah kita buat, lalu **secara otomatis** akan membuatkan perintah SQL **sesuai dengan data Model** yang kita buat.
- Kita bisa menggunakan method `Create()` pada `gorm.DB`.

### Kode: Create

```go
func TestCreateUser(t *testing.T) {
    user := User{
        ID: "1",
        Password: "rahasia",
        Name: Name {
            FirstName: "Nathan",
            MiddleName: "Garzya",
            LastName: "Santoso"
        },
        Information: "Ini akan di-ignore oleh GORM"
    }

    response := db.Create(&user)
    assert.Nil(t, response.Error)
    assert.Equal(t, 1, int(response.RowsAffected))
}
```

### Batch Insert

- Saat kita menggunakan `Create()` Method, maka tiap data akan dibuatkan SQL Insert.
- Kadang ketika kita ingin **memasukkan banyak data** sekaligus, ada baiknya kita menggunakan sekali SQL Insert agar lebih efektif.
- GORM mendukung hal ini menggunakan method `Create(slices)` atau jika ingin **menentukan jumlah data** per SQL Insert, kita bisa gunakan `CreateInBatches(slices, size)`.

### Kode: Batch Insert

```go
func TestBatchInsert(t *testing.T) {
    var users []User

    for i := 2; i < 10; i++ {
        users := append(users, User{
            ID: strconv.Itoa(i),
            Name: Name{
                FirstName: "User " + strconv.Itoa(i),
            },
            Password: "rahasia"
        })
    }

    result := db.Create(&users)
    asset.Nil(t, result.Error)
    assert.Equal(t, 8, int(result.RowsAffected))
}
```

---

## Logger

- **Secara Default**, **informasi** perintah SQL yang dieksekusi oleh GORM **tidak akan di-log**.
- Kita bisa mengubah level log dari GORM menggunakan `gorm.Config` ketika membuat koneksi database.

---

## Transaction

- Transaction hanya bisa terjadi jika kita menggunakan Database Connection **yang sama**.
- Saat kita **menggunakan Method** di `gorm.DB`, **bisa saja** tiap Method **akan menggunakan** Database Connection **yang berbeda**, karena `Connection Pool`-nya **diatur oleh GORM**.
- Jika kita ingin melakukan `transaction`, kita bisa menggunakan method `Transaction(callback)`, dan di dalam function `callback` kita bisa buat semua kode transactionnya.

### Kode: Transaction Success

```go
func TestTransaction(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&User{ID: "11", Password: "rahasia", Name: Name{FirstName: "User 11"} }).Error

        if err != nil {
            return err
        }

        err = tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"} }).Error

        if err != nil {
            return err
        }

        err = tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"} }).Error

        if err != nil {
            return err
        }

        return nil
    })

    assert.Nil(t, err)
}
```

### Kode: Transaction Error

```go
func TestTransactionError(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&User{ID: "12", Password: "rahasia", Name: Name{FirstName: "User 12"} }).Error

        if err != nil {
            return err
        }

        err = tx.Create(&User{ID: "13", Password: "rahasia", Name: Name{FirstName: "User 13"} }).Error

        if err != nil {
            return err
        }

        return nil
    })

    assert.NotNil(t, err)
}
```

### Manual Transaction

- Selain menggunakan Method `Transaction(callback)`, kita juga bisa melakukan manajemen transaksi secara manual.
- Kita bisa membuat object gorm.DB baru ketika menjalankan transaksi menggunakan method `Begin()`.
- Ketika kita ingin melakukan `commit`, gunakan `Commit()`.
- Ketika kita ingin melakukan `rollback`, gunakan `Rollback()`.

### Kode: Manual Transaction Success

```go

func TestManualTransactionSuccess(t *testing.T) {
    tx := db.Begin()
    defer tx.Rollback()

    err := tx.Create(&User{ID: "14", Password: "rahasia", Name: Name{FirstName: "User 14"} }).Error

    assert.Nil(t, err)

    err = tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"} }).Error

    assert.Nil(t, err)

    if err == nil {
        tx.Commit()
    }
}
```

### Kode: Manual Transaction Error

```go

func TestManualTransactionError(t *testing.T) {
    tx := db.Begin()
    defer tx.Rollback()

    err := tx.Create(&User{ID: "16", Password: "rahasia", Name: Name{FirstName: "User 14"} }).Error

    assert.Nil(t, err)

    err = tx.Create(&User{ID: "15", Password: "rahasia", Name: Name{FirstName: "User 15"} }).Error

    assert.NotNil(t, err)

    if err == nil {
        tx.Commit()
    }
}
```

---

## Query

- Sama dengan `Create()`, GORM juga bisa secara otomatis membuat SQL Select dari Model yang kita buat.
- Sehingga kita tidak perlu lagi membuat SQL Select secara manual.

### Single Object

- GORM menyediakan **beberapa method** untuk mengambil **single object** dari database.
- `First()`: untuk mengambil **data pertama** yang **diurutkan** berdasarkan `id`.
- `Take()`: untuk mengambil **satu data**, **tanpa diurutkan**.
- `Last()`: untuk mengambil **data terakhir** yang **diurutkan** berdasarkan `id`.
- Jika **data tidak ditemukan**, maka akan `error: gorm.ErrRecordNotFound`.

### Kode: Single Object

```go
func TestQuerySingleObject(t *testing.T) {
    user := User{}
    result := db.First(&user)

    assert.Nil(t, result.Error)
    assert.Equal(t, "1", user.ID)

    user = User{}
    result = db.Last(&user)
    assert.Nil(t, result.Error)
    assert.Equal(t, "9", user.ID)
}
```

### Inline Condition

- Saat kita menggunakan method `First()`, `Take()`, atau `Last()`, terdapat **parameter selanjutnya** bernama `Inline Condition`.
- **Inline Condition** tersebut secara otomatis akan menjadi kondisi **WHERE** di SQL SELECT-nya.

### Kode: Inline Condition

```go
func TestQueryInlineCondition(t *testing.T) {
    user := User{}

    result := db.First(&user, "id = ?", "5")
    assert.Nil(t, result.Error)
    assert.Equal(t, "5", user.ID)
}
```

### Query All Objects

- GORM juga bisa digunakan untuk melakukan Query untuk semua data di tabel menggunakan method `Find()`.
- Sama seperti method sebelumnya, method `Find()` juga mendukung **Inline Condition** jika kita mau tambahkan kondisi `Where`-nya.
- Untuk detail query `Where` akan dibahas di materi [Query Condition](README.md#query-condition).

### Kode: Query All Objects

```go
func TestQueryAllObjects(t *testing.T) {
    var users []User
    result := db.Find(&users, "id in ?", []string{"1", "2", "3", "4"})

    assert.Nil(t, result.Error)
    assert.Equal(t, 4, len(users))
}
```

---

## Advance Query

### Query Condition

- Sebelumnya kita sudah bisa menggunakan `Inline Condition` ketika melakukan query.
- Selain menggunakan `Inline Condition`, kita juga bisa menggunakan method `Where()` untuk mengubah kondisi query yang akan kita buat.

### Kode: Query Condition

```go
func TestQueryCondition(t *testing.T) {
    var users []User

    result := db.Where("first_name LIKE ?", "%User%").
        Where("password = ?", "rahasia").
        Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 15, len(users))
}
```

### Or Condition

- Secara default saat kita menggunakan `Where()`, kondisi akan digabungkan menggunakan **AND** operator.
- Jika kita ingin menggunakan **OR** operator, kita bisa menggunakan method `Or()`.

### Kode: Or Condition

```go
func TestOrCondition(t *testing.T) {
    var users []User

    result := db.Where("first_name LIKE ?", "%User%").
        Or("password = ?", "rahasia").
        Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 15, len(users))
}
```

### Not Condition

- Dan jika kita ingin menggunakan `NOT` operator, kita bisa gunakan method `Not()`.

### Kode: Not Condition

```go
func TestNotCondition(t *testing.T) {
    var users []User

    result := db.Not("first_name LIKE ?", "%User%").
        Where("password = ?", "rahasia").
        Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 0, len(users))
}
```

### Select Fields

- Secara default, semua kolom akan di-select dan dimasukkan ke field Model.
- Jika misal kita ingin menentukan kolom apa saja yang mau di-select, kita bisa menggunakan method `Select(columns)`.

### Kode: Select Fields

```go
func TestSelectFields(t *testing.T) {
    var users []User

    result := db.Select("id", "first_name").Find(&users)
    assert.Nil(t, result.Error)

    for _, user := range users {
        assert.NotNil(t, user.ID)
        assert.Equal(t, "", user.Name.FirstName)
    }

    assert.Equal(t, 15, len(users))
}
```

### Struct and Map Condition

- Saat kita menggunakan `Where()`, `Not()`, atau `Or()`, kita juga bisa menggunakan parameter `Struct` atau `Map`.
- Secara otomatis, `field` atau `key` akan **dijadikan kolom query**, dan `value` akan **dijadikan value query**.
- Ini cocok ketika pada kasus kita butuh query yang dinamis, sehingga kolom yang dicari bisa berbeda - beda sesuai kondisi pencarian.

### Kode: Struct Condition

```go
func TestStructCondition(t *testing.T) {
    userCondition := User{
        Name: Name{
            FirstName: "User 5",
        },
    }

    var users []User

    result := db.Where(userCondition).Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 1, len(users))
}
```

### Kode: Map Condition

```go
func TestMapCondition(t *testing.T) {
    mapCondition  := map[string]interface{}{
        "middle_name": "",
    }

    var users []User
    result := db.Where(mapCondition).Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 15, len(users))
}
```

### Order, Limit, dan Offset

- Untuk melakukan **sorting**, kita juga bisa menggunakan method `Order()`.
- Dan untuk melakukan **paging**, kita bisa menggunakan method `Limit()` dan `Offset()`.

### Kode: Order, Limit, dan Offset

```go
func TestOrderLimitOffset(t *testing.T) {
    var users []User

    result := db.Order("id asc, first_name asc").
        Limit(5).
        Offset(5).
        Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 5, len(users))
    assert.Equal(t, "14", users[0].ID)
}
```

---

## Query Non Model

- Saat kita menggunakan `First()`, `Take()`, `Last()`, dan `Find()`, GORM **melihat struktur** tabel **ke Model** yang kita gunakan.
- Namun GORM memiliki fitur untuk **menyimpan data ke data yang bukan model**.
  - Contoh, kita hanya ingin melakukan query `first_name` dan `last_name` saja di tabel `users` misalnya. Kita **bisa membuat struct berbeda** dibanding menggunakan model User.
- Jika kita ingin melakukan hal ini, **GORM tetap harus tahu**, Model mana yang digunakan.
  - Caranya adalah kita bisa menggunakan method `Model()` untuk **menentukan Model yang digunakan**.

### Kode: Query Non Model

```go
type UserResponse struct {
    ID string
    FirstName string
    LastName string
}

func TestQueryNonModel(t *testing.T) {
    var users []UserResponse

    result := db.Model(&User{}).Select("id", "first_name", "last_name").
        Find(&users)

    assert.Nil(t, result.Error)
    assert.Equal(t, 15, len(users))
}
```

---

## Update

- Untuk melakukan update data Model yang sudah kita modifikasi, kita bisa menggunakan method `Save()` di `gorm.DB`.
- Secara otomatis semua kolom yang memang **memiliki permission untuk di-update**, akan di-update ke database.

### Kode: Update

```go
func TestUpdate(t *testing.T) {
    user := User{}
    result := db.First(&user, "id = ?", "1")
    assert.Nil(t, result.Error)

    user.Name.FirstName := "Nathan"
    user.Name.MiddleName := "Garzya"
    user.Name.LastName := "Santoso"
    user.Password := "password123"

    result = db.Save(&user)
    assert.Nil(result.Error)
}
```

### Update Selected Column

- **Secara default**, melakukan `Save()` untuk data Model, akan melakukan **update semua kolom walaupun tidak berubah**.
- Jika kita ingin menentukan hanya **beberapa kolom yang ingin di-update**, kita bisa menggunakan method `Updates()`.
- Atau menggunakan `Update(kolom, value)` jika hanya ingin melakukan update **satu kolom** saja.

### Kode: Update Selected Column

```go
func TestSelectedColumns(t *testing.T) {
    result := db.Model(&User{}).Where("id = ?", "1").
        Updates(map[string]interface{}{
            "middle_name": "",
            "last_name": "Morro",
        })

    assert.Nil(t, result.Error)

    result = db.Model(&User{}).Where("id = ?", "1").
        Update("password", "ubahlagi")

    assert.Nil(t, result.Error)

    result = db.Where("id = ?", "1").
        Updates(User{
			Name: Name{
				FirstName: "Nathan",
				LastName:  "Santoso",
			},
		})

    assert.Nil(t, result.Error)
}
```

---

## Auto Increment

- Salah satu fitur yang biasa ada di database adalah **Auto Increment** untuk Primary Key.
- Contoh di MySQL ada `AUTO_INCREMENT` atau di PostgreSQL ada `SERIAL`.
- GORM mendukung pembuatan `ID` yang **Auto Increment** dan **secara otomatis akan melakukan query data ID** setelah `Create()` datanya, sehingga tidak perlu melakukan query manual lagi.
- Untuk memberitahu bahwa field adalah auto increment, kita harus gunakan tag `gorm:"autoIncremet"`

### Kode: Tabel user_logs

```sql
create table user_logs
(
    id int auto_increment,
    user_id varchar(100) not null,
    action varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;
```

### Kode: UserLog Model

```go
import "time"

type UserLog struct {
    ID int `gorm:"primaryKey;column:id;autoIncrement"`
    UserID string `gorm:"column:user_id"`
    Action string `gorm:"column:action"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *UserLog) TableName() string {
    return "user_logs"
}
```

### Kode: Insert User Log

```go
func TestAutoIncrement(t *testing.T) {
    for i := 0; i < 10; i++ {
        userLog := UserLog{
            UserID: "1",
            Action: "Test Action",
        }

        result := db.Create(&userLog)
        assert.Nil(t, result.Error)
        assert.NotEqual(t, 0, userLog.ID)
        fmt.Println(userLog.ID)
    }
}
```

---

## Timestamp Tracking

- Seperti yang sudah dibahaw di materi [Convention](README.md#convention), bahwa GORM menggunakan field `CreatedAt` dan `UpdatedAt` sebagai **Timestamp Tracking**.
- Atau jika **ingin menggunakan field yang berbeda**, kita bisa tambahkan tag `gorm:"autoCreateTime"` atau `gorm:"autoUpdateTime"`.

### Tipe Data Timestamp Tracking

- Sebenarnya, kita juga bisa ubah tipe datanya jika kita mau.
- GORM mendukung tipe data dalam bentuk `number`, dimana **satuannya bisa kita ganti** menjadi `milli` untuk **millisecond**, atau `nano` untuk **nanosecond**, semuanya disimpan dalam **waktu epoch unix time**.
- Misal `gorm:"autoCreateTime:milli"` atau `gorm:"autoUpdateTime:milli"`.
- [https://currentmillis.com](https://currentmillis.com)

### Kode: Alter Table User Log

```sql
DELETE FROM user_logs;

ALTER TABLE user_logs
    modify created_at bigint not null;

ALTER TABLE user_logs
    modify updated_at bigint not null;
```

### Kode: UserLog Model

```go
type UserLog struct {
    ID int `gorm:"primaryKey;column:id;autoIncrement"`
    UserID string `gorm:"column:user_id"`
    Action string `gorm:"column:action"`
    CreatedAt int64 `gorm:"column:created_at;autoCreateTime:milli"`
    UpdatedAt int64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (u *UserLog) TableName() string {
    return "user_logs"
}
```

### Kode: Insert User Log

```go
func TestAutoIncrement(t *testing.T) {
    for i := 0; i < 10; i++ {
        userLog := UserLog{
            UserID: "1",
            Action: "Test Action",
        }

        result := db.Create(&userLog)
        assert.Nil(t, result.Error)
        assert.NotEqual(t, 0, userLog.ID)
        fmt.Println(userLog.ID)
    }
}
```

---

## Upsert

### Save

- Sebelumnya, kita telah menggunakan method `Save()` untuk melakukan **UPDATE**.
- Method `Save()` sebenarnya memiliki kemampuan **untuk mendeteksi apakah harus melakukan UPDATE atau CREATE**.
- Jika data yang kita kirim **tidak memiliki value ID**, maka secara default akan melakukan `CREATE`.
- Jika data yang kita kirim **memiliki value ID**, maka akan melakukan `UPDATE`.
- Hal ini mungkin cocok untuk jenis data yang ID nya adalah **Auto Increment**, karena kita **tidak butuh ID** ketika melakukan `CREATE`.

### Kode: Save

```go
func TestSaveOrUpdate(t *testing.T) {
    userLog := UserLog{
        UserID: "1",
        Action: "Test Action",
    }

    result := db.Save(&userLog)  // create
    assert.Nil(t, result.Error)

    userLog.UserId = "2"
    result = db.Save(&userLog)  // update
    assert.Nil(result.Error)
}
```

### Data Non Auto Increment

- Bagaimana dengan jenis data yang memiliki **ID tidak auto increment**, misalnya data user sebelumnya.
- Untungnya, `Save()` juga bisa digunakna untuk proses otomatis `CREATE`.
- Jadi `Save()` akan mencoba melakukan **UPDATE terlebih dahulu**, ketika mendeteksi jumlah `EffectedRow`-nya adalah 0, secara otomatis `Save()` akan melakukan proses `CREATE`.

### Kode: Save Non Auto Increment Data

```go
func TestSaveOrUpdateNonAutoIncrement(t *testing.T) {
    user := User{
        ID: "99",
        Name: Name{
            FirstName: "User 99",
        },
    }

    result := db.Save(&user) // create
    assert.Nil(t, result.Error)

    user.Name.FirstName = "User 99 Updated"
    result = db.Save(&user) // create
    assert.Nil(t, result.Error)
}
```

### On Conflict

- GORM juga menawarkan **pengaturan Conflict** di Method `Create()`.
- Dengan pengaturan ini, kita bisa menentukan ketika kita coba `Create()` data, lalu **terjadi conflict** (`data sudah ada`), apa yang mau kita lakukan?
- Kita bisa **mengubah pengaturan conflict** ini menggunakan method `Clauses()`.

### Kode: On Conflict

```go
func TestConflict(t *testing.T) {
    user := User{
        ID: "88",
        Name: Name{
            FirstName: "User 88",
        },
    }

    result := db.Clauses(clause.OnConflict{
        UpdateAll: true,
    }).Save(&user) // create

    assert.Nil(t, result.Error)
}
```

---

## Delete

- Untuk melakukan penghapusan data, kita bisa menggunakan method `Delete()`.
- Method `Delete()` juga mendukung **Inline Condition** atau menggunakan method `Where()`.

### Kode: Delete

```go
func TestDelete(t *testing.T) {
    var user User
    // Pakai model yang sudah ada ID-nya
    result := db.First(&user, "id = ?", "88")
    assert.Nil(t, result.Error)

    result = db.Delete(&user)
    assert.Nil(t, result.Error)

    // Pakai model, ambil id -> langsung delete (pakai inline where)
    result = db.Delete(&User{}, "id = ?", "99")
    assert.Nil(t, result.Error)

    // Pakai where method -> lalu pass strut User (untuk memberi tahu table apa yang dipakai)
    result = db.Where("id = ?", "15").Delete(&User{})
    assert.Nil(t, result.Error)
}
```

---

## Soft Delete

- Soft Delete adalah salah satu prakter yang sering dilakukan ketika membuat aplikasi.
- Soft Delete merupakan praktek **menghapus data**, **tanpa menghapus** data dari database.
- Praktek ini **membuat satu kolom** biasanya berupa tipe data `timestamp` yang **berisi waktu dihapus**.
- Jika kolom tersebut **berisi data**, otomatis data tersebut **dianggap sudah di-delete**.

### GORM Soft Delete

- GORM mendukung fitur **Soft Delete secara otomatis**, caranya **kita cukup membuat field** `DeletedAt` dengan tipe `gorm.DeletedAt` (**alias untuk time.Time**).
- Jika GORM **mendeteksi terdapat field** dengan nama `DeletedAt`, **secara otomatis** GORM akan melakukan **Soft Delete**.
- Selain itu, **ketika melakukan query**, secara **otomatis** juga GORM akan **menambahkan filter** Soft Delete yang artinya hasil query hanya **data yang belum di-delete**.

### Kode: Table Todo

```sql
create table todos
(
    id bigint not null auto_increment,
    user_id varchar(100) not null,
    title varchar(100) not null,
    description text null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    deleted_at timestamp null,
    primary key (id)
) engine = InnoDB;
```

### Kode: Todo Model

```go
type Todo struct {
    ID int64 `gorm:"primaryKey;column:id;autoIncrement"`
    UserID string `gorm:"column:user_id"`
    Title string `gorm:"column:title"`
    Description string `gorm:"column:description"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
    DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (t *Todo) TableName() string {
    return "todos"
}
```

### Kode: Delete Todo

```go
func TestSoftDelete(t *testing.T) {
    todo := Todo{
        UserID: "1",
        Title: "Todo 1",
        Description: "Isi Todo 1",
    }

    // create
    result := db.Create(&todo)
    assert.Nil(t, result.Error)

    // delete
    result = db.Delete(&todo)
    assert.Nil(t, result.Error)
    assert.NotNil(t, todo.DeletedAt)

    var todos []Todo
    result = db.Find(&todos)
    assert.Nil(t, result.Error)
    assert.Equal(t, 0, len(todos))
}
```

### Unscoped

- Ketika kita ingin mengambil data, **termasuk data yang sudah di-soft delete**, ktia bisa menggunakan method `Unscoped()`.
- Method `Unscoped()` **juga bisa digunakan** jika kita benar-benar mau melakukan **hard delete** permanen di database.

### Kode: Unscoped

```go
func TestUnscoped(t *testing.T) {
    var todo Todo

    result := db.Unscoped().First(&todo, "id = ?", "1")
    assert.Nil(t, result.Error)

    result = db.Unscoped().Delete(&todo)
    assert.Nil(t, result.Error)

    var todos []Todo
    result = db.Unscoped().Find(&todos)
    assert.Nil(t, result.Error)
    assert.Equal(t, 0, len(todos))
}
```

### Peringatan

- Saat menggunakan **Soft Delete**, perhatikan penggunaan **Primary Key** atay **Unique Index**.
- ketika **data sudah dihapus secara** **Soft Delete**, sebenarnya **data masih ada di table**. Oleh karena itu, pastikan data primary key atau unique index **tidak duplicate** dengan **data yang sudah dihapus** secara `soft delete`.

---

## Model Struct

- GORM menyediakan sebuah struct bernama `Model` yang berisi field `ID`, `CreatedAt`, `UpadtedAt`, dan `DeletedAt`.
- Ini cocok digunakan ketika kita menggunakan field yang sesuai dengan **convention-nya GORM**
- Contoh misal kita bisa gunakan struct Model ini ketika membuat model `Todo`.

### Kode: Todo Model

```go
type Todo struct {
    gorm.Model
    UserID      string         `gorm:"column:user_id"`
	Title       string         `gorm:"column:title"`
	Description string         `gorm:"column:description"`
}
```

---

## Lock

- Hal yang biasa kita lakukan saat menggunakan database adalah melakukan **Lock data**.
- Biasanya, ini dilakukan **agar tidak terjadi RACE CONDITION** ketika memanipulasi **data yang sama** oleh **beberapa request**.
- Untuk melakukan `Lock` menggunakan GORM, kita bisa tambahkan `Clauses()` Locking.
- Kita **bisa menentukan jenis Lock-nya**, apakah itu **UPDATE**, **SHARE**, atau yang lainnya. **Sesuai dengan dukungan database** yang kita gunakan.

### Kode: Lock

```go
func TestLock(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        var user User
        err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
            First(&user, "id = ?", "1").Error

        if err != nil {
            return err
        }

        user.Name.FirstName = "Joko"
        user.Name.LastName = "Morro"
        return tx.Save(&user).Error
    })

    assert.Nil(t, err)
}
```

---

## One to One

- Relationship dalam database yang paling sederhana adalan **One to One**, dimana data di tabel berelasi dengan satu data di tabel lain.
- Di GORM, One to One disebut juga dengan **Has One**.
- Untuk membuat relasi One to One, kita **cukup buat field** dengan **tipe Model yang berelasi**.
- Lalu kita bisa tambahkan informasi di `tag`.
  - `gorm:"foreignKey:nama_kolom"` untuk kolom yang dijadikan **foreign key**.
  - `gorm:"references:nama_kolom"` untuk kolom yang dijadikan **reference**.

### kode: Tabel Wallets

```sql
create table wallets
(
    id varchar(100) not null,
    user_id varchar(100) not null,
    balance bigint not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
) engine = InnoDB;
```

### Wallet Model

```go
type Wallet struct {
    ID string `gorm:"primaryKey;column:id"`
    UserId string `gorm:"column:user_id"`
    Balance int64 `gorm:"column:balance"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (w *Wallet) TableName() string {
    return "wallets"
}
```

### Kode: User Model

```go
type User struct {
	ID          string    `gorm:"primaryKey;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
    Wallet Wallet `gorm:"foreignKey:user_id;references:id"`
}
```

### Kode: Create Wallet

```go
func TestCreateWallet(t *testing.T) {
    wallet := Wallet{
        ID: "1",
        UserId: "1",
        Balance: 1000000,
    }

    err := db.Create(&wallet).Error
    assert.Nil(t, err)
}
```

### Preload

- Secara **default**, **relasi tidak akan di-query oleh GORM**, artinya sifatnya adalah `LAZY`.
- Jika kita ingin melakukan query relation (`EAGER`) **secara langsung ketika melakukan query Model**, kita **bisa sebutkan relasi** yang ingin kita load menggunakan method `Preload()`.

### Kode: Retrieve Relation (1)

```go
func TestRetrieveRelation(t *testing.T) {
    var user User
    err := db.Model(&User{}).Preload("Wallet").First(&user).Error

    assert.Nil(t, err)

    assert.Equal(t, "1", user.ID)
    assert.Equal(t, "1", user.Wallet.ID)
}
```

### Join

- Menggunakan `Preload()`, GORM akan melakukan pengambilan data relasi **menggunakan query yang terpisah**.
- Hal ini **cocok ketika** menggunakan relasi **One to Many** atau **Many to Many**.
- Namun, pada kasus **One to One**, kadang ada baiknya kita **lakukan sekali query** saja menggunakan `JOIN`, karena hasilnya hanya satu data, **jadi lebih cepat**.
- Jika kita ingin menggunakan `JOIN`, kita bisa menggunakan method `Joins()`, lalu **menyebutkan field mana** yang akan kita lakukan `JOIN`.
- Method `Joins()` akan menggunakan `LEFT JOIN` untuk defaultnya

### Kode: Join

```go
func TestRetrieveRelationJoin(t *testing.T) {
    var users []User
    err := db.Model(&User{}).Joins("Wallet").Find(&users).Error

    assert.Nil(t, err)

    assert.Equal(t, 14, len(users))
}
```

---

## Auto Upsert Relation

- Saat kita menggunakan relasi, lalu kita ingin melakukan **create/update** data Model, secara **default** GORM akan **mengecek relasi** yang terdapat **di data tersebut**.
- **Jika terdapat data relasi**, GORM akan melaukan proses `Upsert` **terhadap data relasinya**, sehingga kita **tidak perlu melakukan** create/update data relasi **secara manual**.

### Kode: Auto Create / Update

```go
func TestAutoCreateUpdate(t *testing.T) {
    user := User{
        ID: "20",
        Password: "rahasia",
        Name: Name{
            FirstName: "User 20",
        },
        Wallet: Wallet{
            ID: "20",
            UserId: "20",
            Balance: 1000000,
        },
    }

    err := db.Create(&user).Error
    assert.Nil(t, err)
}
```

### Skip Auto Create / Update

- Jika kita **tidak mau** melakuakn **auto create/update** `data relasi`, kita bisa gunakan method `Omit()` yang berisi `clause.Associations`.
- Ini **memberitahu** GORM **bahwa kita tidak mau** melakukan **auto create/update** untuk data relasinya.

### Kode: Skip Auto Create / Update

```go
func TestSkipAutoCreateUpdate(t *testing.T) {
    user := User{
        ID: "21",
        Password: "rahasia",
        Name: Name{
            FirstName: "User 21",
        },
        Wallet: Wallet{
            ID: "21",
            UserId: "21",
            Balance: 1000000,
        },
    }

    err := db.Omit(clause.Associations).Create(&user).Error
    assert.Nil(t, err)
}
```

---

## One to Many

- Relasi One to Many adalah relasi dimana data di tabel bisa memiliki relasi **ke banyak data** di tabel lain.
- Di GORM, One to Many juga disebut relasi **Has Many**.
- Untuk **membuat relasi One to Many**, kita bisa gunakan **field** dengan tipe **Slice Model** yang berelasi.
- Kita juga bisa menentukan informasi seperti `foreignKey` dan `references`-nya, sama seperti ketika menggunakan relasi One to One.

### Kode: Table Addresses

```sql
create table addresses
(
    id bigint not null auto_increment,
    user_id varchar(100) not null,
    address varchar(100) not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    foreign key (user_id) references users (id)
) engine = InnoDB;
```

### Kode: Address Model

```go
type Address struct {
    ID int64 `gorm:"primaryKey;column:id;autoIncrement"`
    UserId string `gorm:"column:user_id"`
    Address string `gorm:"column:address"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
```

### Kode: User Model

```go
type User struct {
	ID          string    `gorm:"primaryKey;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        Name      `gorm:"embedded"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
    Wallet Wallet `gorm:"foreignKey:user_id;references:id"`
    Addresses []Address `gorm:"foreignKey:user_id;references:id"`
}
```

### Kode: Auto Create / Update

```go
func TestUserAndAddresses(t *testing.T) {
    user := User{
        ID: "50",
        Password: "rahasia",
        Name: Name{
            FirstName: "User 50",
        },
        Wallet: Wallet{
            ID: "50",
            UserId: "50",
            Balance: 1000000,
        },
        Addresses: []Address{
            {
                UserId: "50",
                Address: "jalan A",
            },
            {
                UserId: "50",
                Address: "Jalan B",
            },
        },
    }

    err := db.Create(&user).Error
    assert.Nil(t, err)
}
```

### Kode: Preload & Join

```go
func TestPreloadJoinOneToMany(t *testing.T) {
    var userPreload []User

    err := db.Model(&User{}).
        Preload("Addresses").
        Joins("Wallet").
        Find(&userPreload).Error

    assert.Nil(t, err)
}
```

---

## Belongs To

### Belongs To di One to Many

- Saat kita membuat relasi One to Many, **ada sudut pandang lain** dari Model sebelahnya, yaitu relasi **Many to One**.
- Pada Kasus ini, kita bisa menggunakan relasi **Belongs To (milik)** di GORM.
- Contoh sebelumnya kita tahu bahwa User punya banyak Address, artinya `Address` milik (belongs to) `User`.
- Kita bisa **tambahkan relasi ini** di model `Address`, agar ketika kita melakukan **Query ke Model** `Address`, kita **juga bisa mendapatkan informasi** relasi `User`-nya
- Cara membuatnya mirip seperti ketika membuat relasi One to One.

### Kode: Address Model

```go
type Address struct {
	ID        int64     `gorm:"primaryKey;column:id;autoIncrement"`
	UserId    string    `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
    User User           `gorm:"foreignKey:user_id;references:id"`
}
```

### Kode: Preload atau Join

```go
func TestBelongsTo(t *testing.T) {
    fmt.Println("Preload")
    var addresses []Address
    err := db.Preload("User").Find(&addresses).Error
    assert.Nil(t, err)


    fmt.Println("Joins")
    addresses = []Address{}
    err = db.Joins("User").Find(&addresses).Error
    assert.Nil(t, err)
}
```

### Belongs To di One to One

- Selain di One to Many, Belongs To bisa diimplementasikan di relasi One to One.
- Sebelumnya kita tahu bahwa `User` punya satu (Has One) `Wallet`, artinya `Wallet` itu milik (Belongs To) `User`
- Kita bisa tambahkan field `User` di `Wallet` sebagai relasi **Belongs To**.
- Namun, karena di Golang, **cyclic itu tidak boleh**, maka untuk menambahkan relasi Belongs To di One to One, kita perlu menggunakan `pointer`.

### Kode: Wallet Model

```go
type Wallet struct {
	ID        string    `gorm:"primaryKey;column:id"`
	UserId    string    `gorm:"column:user_id"`
	Balance   int64     `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
    User      *User     `gorm:"foreignKey:user_id;references:id"`
    // Gunakan pointer untuk menghindari cyclic dependency
}
```

### Kode: Belongs To di One to One

```go
func TestBelongsToOneToOne(t *testing.T) {
    fmt.Println("Preload")
    var wallets []Wallet

    err := db.Preload("User").Find(&wallets).Error
    assert.Nil(t, err)

    fmt.Println("Joins")
    wallets = []Wallet{}
    err = db.Joins("User").Find(&wallets).Error
    assert.Nil(t, err)
}
```

---

## Many to Many

- Relasi yang paling kompleks adalah relasi Many to Many.
- Seperti yg kita tahu, bahwa relasi Many to Many, kita harus membuat tabel jembatan penghubung antara 2 tabel.
- GORM juga mendukung relasi Many to Many, caranya mudah kita **hanya perlu membuat field** berupa `Slice` di **kedua Model** yang berelasi.
- Untuk **memberitahu** **tabel penghubung** dan juga **kolom untuk join**, kita bisa menggunakan tag.

### Many to Many Tag

- Untuk memberitahu **nama tabel penghubung**, kita bisa menggunakan tag `gorm:"many2many:nama_tabel"`.
- Saat melakukan query Many to Many, terdapat banyak sekali kolom yang perlu diketahui, sepert:
  - Kolom `ID` di Model **pertama**.
  - Kolom `Foreign Key` Model **pertama** di **tabel penghubung**.
  - Kolom `ID` di Model **kedua**.
  - Kolokm `Foregin Key` Model **kedua** di **tabel penghubung**.
- Semua bisa kita lakukan dengan menggunakan Tag:
  - `gorm:"foreignKey:kolom_id"`: untuk `ID` di Model pertama.
  - `gorm:"joinForeignKey:kolom_id"`: untuk `Foreign Key` Model pertama di tabel penghubung.
  - `gorm:"references:kolom_id"`: untuk `ID` di Model kedua.
  - `gorm:"referencesForeignKey:kolom_id"`: untuk `Foreign Key` Model kedua di tabel penghubung.

### Contoh Kasus

- Misal kita akan membuat fitur dimana `User` bisa like banyak `Product`, dan satu `Product` bisa dilike banyak `User`.
- Artinya kita akan membuat model `Product`, lalu akan membuat Relasi Many to Many antara `User` dan `Product`.

### Kode: Tabel Products

```sql
create table products
(
    id varchar(100) not null,
    name varchar(100) not null,
    price bigint not null,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    primary key (id)
) engine = InnoDB;
```

### Kode: Product Model

```go
type Product struct {
    ID string `gorm:"primaryKey;column:id"`
    Name string `gorm:"column:name"`
    Price int64 `gorm:"column:price"`
    CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
    UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
```

### Kode: Tabel Penghubung user_like_product

```sql
create table user_like_product
(
    user_id varchar(100) not null,
    product_id varchar(100) not null,
    primary key (user_id, product_id),
    foreign key (user_id) references users (id),
    foreign key (product_id) references products (id)
) engine = InnoDB;
```

### Kode: User Model

```go
type User struct {
	ID           string    `gorm:"primaryKey;column:id;<-:create"`
	Password     string    `gorm:"column:password"`
	Name         Name      `gorm:"embedded"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information  string    `gorm:"-"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses    []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}
```

### Kode: Product Model with Users

```go
type Product struct {
	ID           string    `gorm:"primaryKey;column:id"`
	Name         string    `gorm:"column:name"`
	Price        int64     `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User    `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"`
}
```

### Create / Update / Delete Relasi Many to Many

- Salah satu tantangan relasi Many to Many adalah: Bagaimana cara melakukan **Create/Update/Delete** relasi Many to Many?
- Hal ini karena **tabel penghubung** **tidak memiliki representasi** Model di GORM.
- Untungnya, GORM bisa digunakan untuk memanipulasi data, tanpa harus menggunakan Model
- Kita bisa gunakan method `Table()` pada `gorm.DB` untuk **menyebutkan tabel** mana yang mau **kita pilih**.
- Walaupun cara ini bisa dilakukan, tapi sebenarnya cara yang lebih baik itu menggunakan fitur bernama [Association Mode](README.md#association-mode)

### Kode: Create Many to Many

```go
func TestCreateManyToMany(t *testing.T) {
    product := Product{
        ID: "P001",
        Name: "Contoh Product",
        Price: 10000,
    }

    err := db.Create(&product).Error
    assert.Nil(t, err)

    // Create data into user_like_product table
    err = db.Table("user_like_product").
        Create(map[string]interface{}{
            "user_id": "1",
            "product_id": "P001",
        }).Error
    assert.Nil(t, err)


    err = db.Table("user_like_product").
        Create(map[string]interface{}{
            "user_id": "2",
            "product_id": "P001",
        }).Error
    assert.Nil(t, err)
}
```

### Kode: Preload

```go
func TestPreloadManyToMany(t *testing.T) {
    var product Product
    err := db.Preload("LikedByUsers").First(&product, "id = ?", "P001").Error

    assert.Nil(t, err)
    assert.Equal(t, 2, len(product.LikedByUsers))
}
```

---

## Association Mode

- GORM memiliki fitur bernama **Association Mode**, yang digunakan **untuk memanipulasi data relasi** menggunakan object `gorm.Association`.
- Fitur ini sangat berguna, misal ketika **kita ingin menambah relasi ke Model**, **terutama** untuk relasi **Many to Many** yang tidak memiliki Model untuk tabel penghubungnya.
- Untuk membuat object `Association`, kita cukup menggunakan method `Association(relasi)`.

### Kode: Mencari Relasi

```go
func TestAssociationFind(t *testing.T) {
    var product Product
    err := db.First(&product, "id = ?", "P001").Error
    assert.Nil(t, err)

    var users []User
    err = db.Model(&product).
        Where("first_name LIKE ?", "User%").
        Association("LikedByUsers").
        Find(&users)
    assert.Nil(t, err)
    assert.Equal(t, 1, len(users))
}
```

### Kode: Add Relation

```go
func TestAssociationAdd(t *testing.T) {
    var user User
    err := db.First(&user, "id = ?", "3").Error
    assert.Nil(t, err)

    var product Product
    err = db.First(&product, "id = ?", "P001").Error
    assert.Nil(t, err)

    err = db.Model(&product).Association("LikedByUsers").Append(&user)
    assert.Nil(t, err)
}
```

### Kode: Replace Relation

```go
func TestAssociationReplace(t *testing.T) {
    err := db.Transaction(func (tx *gorm.DB) error {
        var user User
        err := tx.First(&user, "id = ?", "1").Error
        assert.Nil(t, err)

        wallet := Wallet{
            ID: "01",
            UserId: "1",
            Balance: 2000000,
        }

        err = tx.Model(&user).Association("Wallet").Replace(&wallet)

        return err
    })

    assert.Nil(t, err)
}
```

### Kenapa Error?

- Hal ini karena di Tabel `wallets`, kita menambahkan aturan **constraint** `NOT NULL`.
- Jika ketika GORM membuat Wallet dengan `ID` **01**, dan dia coba **menghapus** `user_id` di wallet **sebelumnya** yaitu **wallet 1**, maka terjadi `error`.
- Namun dari sini kita bisa tahu bahwa, GORM akan **menghapus relasi ke data sebelumnya** ketika kita melakukan `Replace()`.
- `Replace()` ini **hanya cocok** untuk relasi **One to One** atau **Belongs To**, dan ketika kita coba menggunakan `Append()` di relasi tersebut, secara otomatis GORM akan mengubah menjadi operasi `Replace()`, **bukan** `Append()` lagi.

### Kode: Delete Relation

```go
func TestAssociationDelete(t *testing.T) {
    var user User
    err := db.First(&user, "id = ?", "3").Error
    assert.Nil(t, err)

    var product Product
    err = db.First(&product, "id = ?", "P001").Error
    assert.Nil(t, err)

    err = db.Model(&product).Association("LikedByUsers").Delete(&user)
    assert.Nil(t, err)
}
```

### Clear Relation

```go
func TestAssocationClear(t *testing.T) {
    var product Product
    err := db.First(&product, "id = ?", "P001").Error
    assert.Nil(t, err)

    err = db.Model(&product).Association("LikedByUsers").Clear()
    assert.Nil(t, err)
    fmt.Println(product)
}
```

---

## Preloading

- Sebelumnya kita sudah tahu bahwa untuk melakukan loading relasi, kita bisa menggunakan `Preloading()`.
- Bagaimana jika kita ingin menambahkan kondisi ketika melakukan `Preloading`?
- Kita bisa tambahkan `Inline Condition` ketika melakukan **Preloading**.

### Kode: Preloading with Condition

```go
func TestPreloadingWithCondition(t *testing.T) {
    var user User

    err := db.Preload("Wallet", "balance > ?", 1000000).First(&user, "id = ?", "1").Error
    assert.Nil(t, err)
    fmt.Println(user)
}
```

### Nested Preloading

- Preloading juga bisa dilakukan untuk relasi yang **Nested**.
- Misal kita akan melakukan query untuk Model `Wallet`, kita ingin melakukan **Preload** ke `User` dan juga ke `Address`-nya.
- Kita bisa gunakan `Preloading` dengan menggunakan `.` (titik), misal `User.Addresses`.

### Kode: Nested Preloading

```go
func TestNestedPreloading(t *testing.T) {
    var wallet Wallet
    err := db.Preload("User.Addresses").Find(&wallet, "id = ? ", "1").Error
    assert.Nil(t, err)
}
```

### Preload All

- Jika kita ingin melakukan `Preload` **semua relasi** di Model, kita bisa menggunakan `clauses.Associations` ketika melakukan `Preloading()`.
- Namun perli diingat, bahwa **Preload All** **tidak akan melakukan load** `Nested Relation`.

### Kode: Preload All

```go
func TestPreloadAll(t *testing.T) {
    var user User
    err := db.Preload(clause.Associations).First(&user, "id = ?", "1").Error
    assert.Nil(t, err)
}
```

---

## Joins

- Sebelumnya kita sudah melakukan `Joins()` dengan menyebutkan nama field Relation-nya.
- Selaing menggunakan nama Relation, kita juga **bisa menggunakan Query manual** ketika melakukan `Joins`.
- Bedanya adalah, **kolom Relation-nya tidak akan di Query**.

### Kode: Join Query

```go
func TestJoinQuery(t *testing.T) {
    var user []User
    err := db.Joins("join wallets on wallets.user_id = users.id").Find(&users).Error
    assert.Nil(t, err)
    assert.Equal(t, 3, len(users))

    users = []User{}
    err = db.Joins("Wallet").Find(&users).Error // Using Left Join
    assert.Nil(t, err)
    assert.Equal(t, 17, len(users))
}
```

### Join Condition

- Saat menggunakan `Joins`, ketika kita ingin menambahkan kondisi di Join Table-nya, **jangan lupa untuk menyebutkan nama tabel-nya**.
- Namun saat menggunakan `Joins()` **menggunakan nama relasi**, **secara otomatis** GORM akan **membuat alias** untuk nama relasi, jadi **kita harus menggunakan nama relasi ketika menambahkan kondisi**.

### Kode: Join Condition

```go
func TestJoinQueryCondition(t *testing.T) {
    var users []User
    err := db.Joins("join wallets on wallets.user_id = users.id AND wallets.balance > ?", 500000).Find(&users).Error

    assert.Nil(t, err)
    assert.Equal(t, 3, len(users))

    users = []User{}
    err = db.Joins("Wallet").Where("Wallet.balance > ?", 500000).Find(&users).Error // Alias menggunakan nama field
    assert.Nil(t, err)
    assert.Equal(t, 3, len(users))
}
```

---

## Query Aggregation

- Ada banyak **Query Aggregation** yang biasanya tersedia di Database.
- Di GORM, hanya menyediakan method `Aggregation` untuk **Count** saja.
- Jika kita ingin melakukan Aggregation misal `Avg`, `Sum`, `Max`, dan lainnnya, maka kita **harus melakukan secara manual** menggunakan `Select()`.

### Kode: Count Aggregation

```go
func TestCount(t *testing.T) {
    var count int64

    err := db.Model(&User{}).
        Joins("Wallet").
        Where("Wallet.balance > ?", 500000).
        Count(&count).Error

    assert.Nil(t, err)
    assert.Equal(t, int64(3), count)
}
```

### Kode: Other Aggregation

```go
type AggregationResult struct {
    TotalBalance int64
    MinBalance int64
    MaxBalance int64
    AvgBalance float64
}

func TestOtherAggregation(t *testing.T) {
    var result AggregationResult

    err := db.Model(&Wallet{}).
        Select("SUM(balance) as total_balance", "MIN(balance) as min_balance", "MAX(balance) as max_balance", "AVG(balance) as avg_balance").
        Take(&result).Error

    assert.Nil(t, err)
    assert.Equal(t, int64(4000000), result.TotalBalance)
    assert.Equal(t, int64(1000000), result.MinBalance)
    assert.Equal(t, int64(3000000), result.MaxBalance)
    assert.Equal(t, float64(1333333.3333), result.AvgBalance)
}
```

### Group By and Having

- GORM juga menyediakan method untuk melakukan **Group By** menggunakan method `Group()` dan untuk melakukan **Having** menggunakan method `Having()`.

### Kode: Group By and Having

```go
func TestGroupByHaving(t *testing.T) {
    var result []AggregationResult

    err := db.Model(&Wallet{}).
        Select("SUM(balance) as total_balance", "MIN(balance) as min_balance", "MAX(balance) as max_balance", "AVG(balance) as avg_balance").
        Joins("User").
        Group("User.id").
        Having("SUM(balance) > ?", 1000000).
        Find(&result).Error
    assert.Nil(t, err)
    assert.Equal(t, 1, len(result))
}
```

---

## Context

- Saat kita menggunakan Golang, biasanya kita akan menggunakan `context.Context`.
- Bagaimana dengan GORM? Dari awal kita belum pernah menggunakan `Context`.
- GORM juga mendukung penggunaan `Context`, kita bisa menggunakan method `WithContext()` ketika mau melakukan operasi menggunakan GORM.

### Kode: Context

```go
func TestContext(t *testing.T) {
    ctx := context.Background()

    var users []User
    err := db.WithContext(ctx).
        Find(&users).Error

    assert.Nil(t, err)
    assert.Equal(t, 17, len(users))
}
```

---

## Scopes

- Sebelumnya kita sudah pernah menggunakan method `Scopes()`.
- Method `Scopes()` juga **bisa digunakan untuk menambahkan custom logic** yang mungkin sering kita gunakan.
- Kita cukup menambahkan `function(*gorm.DB) *gorm.DB`

### Kode: Sample Function

```go
func BrokeWalletBalance(db *gorm.DB) *gorm.DB {
    return db.Where("balance = ?", 0)
}

func SultanWalletBalance(db *gorm.DB) *gorm.DB {
    return db.Where("balance > ?", 1000000)
}
```

### Kode: Scopes

```go
func TestScopes(t *testing.T) {
    var wallets []Wallet

    err := db.Scopes(BrokeWalletBalance).Find(&wallets).Error
    assert.Nil(t, err)

    wallets = []Wallet{}
    err = db.Scopes(SultanWalletBalance).Find(&wallets).Error
    assert.Nil(t, err)
}
```

---

## Connection Pool

- Saat kita belajar di Kelas [Golang Database](https://github.com/canonflow/tutorial-golang-pzn/tree/main/07%20-%20golang%20mysql), kita belajar tentang Pool, dimana Golang mengatur mantain koneksi yang terbuka dan tertutup secara otomatis.
- Kita hanya cukup menggunakan saja, tanpa harus pusing mengaturnya.
- Bagaimana dengan GORM?
- GORM sendiri sebenarnya di dalamnya tetap menggunakan `sql.DB`.
- Jadi jika kita ingin mengubah pengaturan Pool-nya, kita bisa menggunakan `sql.DB`.

### Kode: Connection Pool

```go
sqlDb, err := db.DB()

if err != nil {
    panic(err)
}

sqlDB.SetMaxOpenConns(100)
sqlDB.SetMaxIdleConns(10)
sqlDB.SetConnMaxLifetime(30 * time.Minute)
sqlDB.SetConnMaxIdleTime(5 * time.Minute)
```

---

## Migrator

- Sebelumnya, kita pernah belajar tentang Database migration di kelas [Golang Database Migration](https://github.com/canonflow/tutorial-golang-pzn/tree/main/16%20-%20golang%20database%20migration).
- Sebenarnya GORM sendiri memiliki fitur untuk melakukan **Migration secara otomatis**.
- Namun **lebih direkomendasikan** menggunakan database migration **yang mendukung version** agar tidak terjadi kesalahan ketika melakukan perubahan schema database.
- [https://pkg.go.dev/gorm.io/gorm#Migrator](https://pkg.go.dev/gorm.io/gorm#Migrator).
- Kita tidak akan membahas terlalu panjang fitur Migrator ini, karena jarang sekali digunakan di aplikasi nyata, biasanya hanya digunakan untuk melakukan pengetesan di komputer programmer saja

### Kode: GuestBook Model

```go
type GuestBook struct {
	ID        int64  `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Message   string `gorm:"column:message"`
	CreatedAt string `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt string `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (g *GuestBook) TableName() string {
	return "guest_books"
}
```

### Kode: Migrator

```go
func TestMigrator(t *testing.T) {
    err := db.Migrator().AutoMigrate(&GuestBook{})
    assert.Nil(t, err)
}
```
