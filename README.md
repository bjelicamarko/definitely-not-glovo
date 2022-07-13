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
  
- Radnik

  - Pregled poruđbenica
  - Prihvatanje poruđbenice
  
- Administrator
  
  - CRUD nad entitetima sistema (korisnici, restorani, artikli)
  - Izvještaji o poslovanju (prihodi restorana, najprodavaniji artikli)
  
## Arhitektura sistema
Web aplikacija će biti zasnovana na mikroservisnoj arhitekturi.

- Gateway servis - Go
- Korisnički servis - Go
- Servis za restorane - Rust
- Servis za artikle - Go
- Servis za poruđbenice (kreiranje, pregled, izvještaji) - Go
- Klijentska web aplikacija - Angular (možda React)

Podaci će biti čuvani u SQL bazi (PostgreSQL) <br />
