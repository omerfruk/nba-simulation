# NBA - Simulation

## Kurulumlar

### DB Kurulum

DataBase'mizde nba-simulation adinda kullanici olusturalim ardindan

localhost'umuzda nba-simulation adli db olusturup kullanici olarak nba-simulation'i atayalim

db configure degiskenleri altta tarafta verilmistir

> ``
const (
HOST = "localhost"
DATABASE = "nba-simulation"
USER = "nba-simulation"
PASSWORD = "guzelsifre"
)
``


### Go

Golang bilgisayarimizda yuklu degilse ilk olarak [Golang](https://golang.org/)
adresinden indirip bilgisayarimiza kuruyoruz.

Sonrasinda  <code>`go mod init`,`go mod download` ve `go mod tidy`</code> komutlarini sirasi ile calistiriyoruz

### vue

Dizin olarak ./client'e gidiyoruz.

<code>yarn install or npm install</code>  dedikten sonra

<code>yarn serve or npm run serve</code> diyerek projemizin ayagı kalkmasini sagliyoruz


similasyonun gerceklik payı olması icin methodlarda time.sleep bulunmaktadır zaman kaybini engellemk icin yorum satirina alinmistir
