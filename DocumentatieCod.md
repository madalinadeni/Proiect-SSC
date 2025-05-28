# Proiect-SSC
## Documentatie Cod.

Ce este certificatul unui site? 
  -> Certificatul unui site se numeste SSL ( Secure Sockets Layer), iar rolul sau este de atesta ca pagina web este sigura si ca poate fi accesata fara probleme. 
  ->  Certificatul SSL transforma adresa web a unui site din HTTP in HTTPS (s de la securizat) 
  -> Certificatul SSL cripteaza informatiile utilizatorului, astfel incat, chiar daca ar fi interceptate, nu pot sa fie citite. 

  -> Tipuri de certificate SSL 
        > Certificat SSL de validare domeniu – GoGetSSL DV SSL: validarea pentru un domeniu, ideal pentru majoritatea website-urilor de prezentare. 
        > Certificat SSL Wildcard – Wildcard SSL: acopera domeniul principal si toate subdomeniile sale, destinat pentru acele website-uri care isi extind       activitatea pe mai multe subdomenii. 
        > Certificat SSL de validare extinsa – BusinessTrust EV SSL: validare companie, utilizat pentru companii mari care gestioneaza date sensibile. 

 

Acest program Go verifică validitatea certificatului TLS al unui site specificat prin linia de comandă. Programul efectuează o cerere HTTPS către domeniul dat, extrage certificatul serverului și verifică două aspecte principale:
        > Validitatea temporală a certificatului (dacă este încă valabil).
        > Dacă numele domeniului coincide cu cel din certificat.
În cazul în care oricare din verificări eșuează, programul avertizează asupra unui posibil atac „Man in the middle”.

## Structura programului
  ## Funcția verifyCertificate
  Are urmatorii parametri:
        > cert — certificatul TLS extras de la server (tip *x509.Certificate).
        > hostname — numele domeniului pentru care se face verificarea.
  Funcționalitate:
        > Verifică dacă certificatul este valabil în intervalul său de valabilitate (între NotBefore și NotAfter).
        > Verifică dacă numele domeniului (hostname) este inclus în certificatul TLS (folosind cert.VerifyHostname).
  Aceasta functie returneaza:
        > nil dacă certificatul este valid și numele domeniului corespunde.
        > O eroare descriptivă în caz contrar.

  ## Funcția main
  Verifică dacă programul a fost apelat cu exact un argument (numele domeniului).
  Construiește URL-ul HTTPS corespunzător domeniului.
  Efectuează o cerere HTTP GET către URL.
  Obține starea conexiunii TLS (resp.TLS).
  Extrage certificatul serverului din lista PeerCertificates.
  Afișează data până la care certificatul este valabil (NotAfter).
  Apelează funcția verifyCertificate pentru a valida certificatul.
  Afișează un mesaj corespunzător dacă site-ul este sigur sau dacă există risc de atac „Man in the middle”.
  ## Modalitate de rulare
  go run program.go example.com
        > unde example.com este domeniul căruia îi verifici certificatul TLS.

  ## Exemple de output
  Dacă certificatul este valid și corespunde domeniului:
        > Certificatul este valabil pana la: 10 Dec 2025 23:59:59
          Site-ul este sigur!
  Dacă certificatul nu este valabil (expirat sau încă nu activ):
        > Certificatul este valabil pana la: 10 Dec 2020 23:59:59
          Site-ul este nesigur. Posibil un atac 'Man in the middle'!:
          Certificatul este invalid. Posibil un atac 'Man in the middle'!
  Dacă numele domeniului nu corespunde:
        > Certificatul este valabil pana la: 10 Dec 2025 23:59:59
          Site-ul este nesigur. Posibil un atac 'Man in the middle'!:
          Numele domeniului nu corespunde certificatului.
  Dacă argumentele sunt greșite:
        > Eroare la argumentele din linia de comanda!
  Dacă nu se poate obține certificatul TLS:
        > Certificatul TLS nu a putut fi obtinut!
  ## Surse de documentare:
  https://www.datahost.ro/blog/tot-ce-trebuie-sa-stii-despre-un-certificat-ssl/
  https://pkg.go.dev/crypto/x509
  https://pkg.go.dev/net/http
  https://pkg.go.dev/crypto/tls#ConnectionState
  https://pkg.go.dev/os#Args
  https://pkg.go.dev/time
  https://go.dev/tour/welcome/1
      
