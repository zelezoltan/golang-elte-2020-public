# Feladat

Irjatok egy olyan webalkalmazast, ahol regisztralt felhasznalok tudnak uzeneteket irogatni es latni, mint egy faliujsagra.

- A felhasznaloi adatok az sqlite adatbazisbol jonnek. (Van egy user tabla, benne egy `name` es `code` oszloppal. Elerheto user/code parok: `test1|skey1`, `test2|key2`)
- Felhasznalo authentikalasahoz az `Authorization` headerben menjen fel a kod a szerverre es egy contextbe keruljon a felhasznalo neve.
- Az uzenet mellett legyen letarolva a korabban authenticalt felhasznalo neve es a datum, hogy mikor irta

2 endpoint:

- GET /api/message

  Az uzenetek legyen az alabbi json formaban kilistazva:
  ```{"messages":[
    {
      "id": uzenet azonositoja,
      "user":"user neve",
      "msg": "uzenet",
      "date": "a bekuldes datuma"
    },
    ....
  ]}
  ```

- POST /api/message

  A bekuldott uzeneteket a szerver fogadja az alabbi formaban:
    
  ```
  {"msg":"uzenet a felhasznalotol"}
  ```

  Az uzenet maga nem lehet ures, ebben az esetben az alkalmazas dobjon hibat es ne tarolja el azt.



**Extra** kiegeszitesi otletek, ha van kedvetek jatszani es gyakorolni kicsit a Go-t:
- Keszithettek hozza valamilyen UI-t, akar egy Go-s konzol alkalmazas formaban (egy konzol app, ami hivja a HTTP endpointokat)
- Csinaljatak opcionalis "anonymous" kapcsolot, aminek hatasara a listazasnal a `user` mezo uresen jelenik meg az adott uzenetnel
- Adjatok lehetoseget az uzenetek torlesere, de minden felhasznalo csak a sajat uzenetet tudja torolni
- Egeszitsetek ki az uzeneteket valaszhato nick nevekkel a user neve melle, amit a bekuldesnel megvalaszthat a bekuldo
- Adjatok egy endpointot amin keresztul uj felhasznalot lehet regisztralni a rendszerbe
- Keszithettek webes UI-t HTML es Javascript hasznalataval (React, vagy barmi johet)



A feladat elkesziteshez ebben a mappaban talahato egy felkesz alkalmazas.
A kiegeszitendo reszeket az alabbi komment sor formaval jeloltem:
```
// TODO(feladat): <feladat szovege>
```