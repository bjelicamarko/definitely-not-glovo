# Nije Glovo
## Projekat iz predmeta Napredne tehnike programiranja
**Nije Glovo je web aplikacija za naručivanje hrane i pića zasnovana na mikroservisnoj arhitekturi.**

## Funkcionalnosti
- Neregistrovani korisnik

  - Prijava
  - Registracija
  
- Registrovani korisnik
  
  - Pregled i izmjena profila
  - Pregled i pretraga restorana
  - Pregled i pretraga hrane i pića unutar restorana
  - Naručivanje (kreiranje poruđbenice)
  - Pregled istorije poruđbenica
  - Ostavljanje recenzije u vidu komentara i ocjene nakon primljene porudžbine
  - Mogućnost ostavljanja napojnice dostavljaču (+)
  
- Radnik

  - Pregled poruđbenica
  - Prihvatanje poruđbenice
  - Prijava neprikladnih komentara
  
- Administrator
  
  - Pregled i pretraga nad entitetima sistema (korisnici, restorani, artikli) 
  - CRUD nad entitetima sistema (korisnici, restorani, artikli)
  - Izvještaji o poslovanju (prihodi restorana, najprodavaniji artikli)
  - Uvid u neprikladne komentare
  - Mogućnost blokiranja i odblokiranja korisnika
  
- Dostavljač (+)
  
  - Preuzimanje i dostavljanje porudžbine (+)
  
## Arhitektura sistema
Web aplikacija će biti zasnovana na mikroservisnoj arhitekturi.

- Gateway servis - Go
- Korisnički servis - Go
- Servis za restorane - Go
- Servis za artikle - Go
- Servis za poruđbenice - Go
- Servis za recenzije - Go
- Servis za izvještaje - Rust
- Klijentska web aplikacija - Angular

Podaci će biti čuvani u SQL bazi (PostgreSQL), svaki mikroservis će imati odvojenu instancu SQL baze. <br />
Za kontejnerizaciju biće korišćen Docker. (+)  
Kreiranje lokacije restorana uz pomoć mape i prikaz lokacije restorana na mapi. (+) <br />  <br />

**_Funkcionalnosti označene sa (+) su proširenja za diplomski rad._**
