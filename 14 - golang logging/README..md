# Logger

## Level
| Level | Function |             Description              |
| :---: | :------: |:------------------------------------:|
| Trace | logger.Trace() |                  -                   |
| Debug | logger.Debug() |                  -                   |
| Info | logger.Info() |                  -                   |
| Warn | logger.Warn() |                  -                   |
| Error | logger.Error() |                  -                   |
| Fatal | logger.Fatal() | Memanggil os.Exit(1) setelah logging |
| Panic | logger.Panic() |   Memanggil panic setelah logging    |

### Logging Level
- Kenapa **Trace** dan **Debug** tidak keluar di console?
- Karena secara default, Logging Level yang digunakan adalah **Info**, artinya habya prioritas **Info keatas yang di log**
- Untuk mengubah Logging Level, kita bisa gunakan `logger.SetLevel()`

---

## Output
- Secara default, output dari logger dikirim ke Console oleh Logrus.
- Kadang kita butuh mengubah output tujuan log, misal ke File / Database.
- Output Logger adalah `io.Writer` jadi kita bisa menggunakan io.Writer dari Golang atau bisa membuat sendiri mengikuti kontrak `io.Writer`.

---

## Formatter
- Saat Logger mengirimkan data ke output, log yang kita kirim akan diformat menggunakan object Formatter.
- Logrus secara default memiliki **2 formatter**.
- TextFormatter, yang **secara default digunakan**
- JSONFormatter, yang bisa digunakan untuk memformat pesar log menjadi data JSON.
- Untuk mengubah formatter, kita bisa gunakan function `logger.SetFormatter()`.

---

## Field
- Saat kita mengirim informasi log, kadang kita ingin **menyisipkan sesuatu pada log tsb**.
- Misal saja, seperti informasi **siapa yang login di log-nya**
- Cara manual kita bisa menambahkan informasi di messagenya, tapi Logrus menyediakan cara yang lebih baik, yaitu menggunakan fitur **Field**.
- Dengan fitru **Field**, kita bisa menambahkan Field tambahan di informasi Log yang kita kirim.
- Sebelumnya melakukan logging, kita bisa menggunakan function `logger.WithField()` untuk menambahkan FIeld yang kita inginkan.

### Beberapa Fields
- Kita juga bisa langsung memasukkan beberapa Field dengan menggunakan **Fields**
- **Fields** adalah alias untuk map[string]interface{}
- Caranya kita bisa menggunakan function `logger.WithFields()`.

---

## Entry
- **Entry** adalah sebuah Struct representasi dari log yang kita kirim menggunakan Logrus.
- **Setiap log yang kita kirim**, maka akan dibuatkan object **Entry**
- Contohnya ketika kita **membuat Formatter sendiri**, maka parameter yang ktia dapat untuk melakukan fornmatting bukanlah string message, melainkan object **Entry**
- Untuk membuat Entry, kita bisa menggunakan function `logrus.NewEntry()`.

---

## Hook
- **Hook** adalah sebuah Struct yang bisa kita **tambahkan** ke Logger sebagai **callback** yang **akan dieksekusi** ketika **terdapat kejadian** log untuk **level tertentu**.
- Contohnya, **ketika ada log error**, kita **ingin mengirimkan notifikasi** via chat ke programmer, dan lain2.
- Kita bisa menambahkan Hook ke Logger dengan menggunakan function `logger.AddHook()`.
- Dan kita juga **bisa menambahkan lebih dari 1 Hook** ke Logger.

---

## Singleton
- Logrus sendiri memiliki **singleton object** untuk Logger, sehingga kita **tidak perlu membuat** object **Logger sendiri** sebenarnya.
- Namun artinya, **jika kita ubah data Logger singleton tersebut**, maka secara **otomatis** yang menggunakan Logger tersebut **akan berubah**.
- **Secara default**, Logger singleton yang ada di logrus menggunakan **TextFormatter** dan **Info Level**
- Cara menggunakan Logger singleton ini, kita **bisa langsung menggunakan** package logrus-nya saja.


