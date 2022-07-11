# Glovo na aparatima
## Projekat iz predmeta Napredne tehnike programiranja
**Glovo na aparatima je web aplikacija za naručivanje hrane i pića zasnovana na mikroservisnoj arhitekturi.**

## Funkcionalnosti
- Neregistrovani korisnik

  - Prijava
  - Registracija
  
- Registrovani korisnik
  
  - Pregled i izmjena profila
  - Pregled i pretraga restorana
  - Pregled i pretraga hrane i pića unutar restorana
  - Naručivanje 
  - Pregled istorije poruđbenica
  
- Administrator
  
  - CRUD nad entitema sistema (korisnici, restorani, artikli)
  - Izvještaji o poslovanju (prihodi restorana, najprodavaniji artikli)
  
## Arhitektura sistema
Web aplikacija će biti zasnovana na mikroservisnoj arhitekturi.

- Gateway servis - Go
- Korisnički servis - Go
- Servis za restorane - Go
- Servis za poruđbenice - Go
- Servis za izvještaje - Rust
- Klijentska web aplikacija - Angular (možda React)

Podaci će biti čuvani u SQL bazi (PostgreSQL) <br />
Za kontejnerizaciju koristiće se Docker. <br />
