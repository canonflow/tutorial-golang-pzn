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