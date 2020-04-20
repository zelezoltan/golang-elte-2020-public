# HTTP Server examples

- [Simple http](ex1/README.md) - Simplest http server with simple GET handler
- [Multiple handlers](ex2/README.md) - Simple http server with form and I/O handling
- [HTTPS](ex3/README.md) - Add HTTPS support
- [Render HTML](ex4/README.md) - HTML rendering and static serving
- [Custom mux](ex5/README.md) - Add custom mux
- [Middleware magic](ex6/README.md) - Some middleware magic
- [Complex example](ex7/README.md) - Complex example with global context


# (Hazi) feladat:

Irjatok egy olyan webalkalmazast, ahol regisztralt felhasznalok tudnak uzeneteket irogatni es latni, mint egy faliujsagra.

- A felhasznaloi adatok johetnek barmilyen filebol (txt, csv, adatbazis), de eleg ha 3 user be van egetve a kodba. (Lehetoleg egy tombbol vagy hasonlo adatszerkezetbol jojjenek es ne egy nagy if-else legyen.) A lenyeg, hogy nem kell regisztracio es felhasznalo kezeles.
- Felhasznalo authentikalasa mehet siman `Basic auth`-al (szerver oldalon olvasasra: https://golang.org/pkg/net/http/#Request.BasicAuth), de az is oke, ha barmilyen Headerben megy fel az adat.
- Az uzenet mellett legyen letarolva a korabban authenticalt felhasznalo neve es a datum, hogy mikor irta

2 endpoint:

- GET /message

  Az uzenetek legyen az alabbi json formaban kilistazva:
  ```{"messages":[
    {
      "user":"user neve",
      "msg": "uzenet",
      "date": "a bekuldes datuma"
    },
    ....
  ]}
  ```

- POST /message

  A bekuldott uzeneteket a szerver fogadja az alabbi formaban:
    
  ```
  {"msg":"uzenet a felhasznalotol"}
  ```

  Az uzenet maga nem lehet ures, ebben az esetben az alkalmazas dobjon hibat es ne tarolja el azt.



**Extra** kiegeszitesi otletek, ha van kedvetek jatszani es gyakorolni kicsit a Go-t:
- Keszithettek hozza valamilyen UI-t, akar egy Go-s konzol alkalmazas formaban (egy konzol app, ami hivja a HTTP endpointokat)
- Csinaljatak opcionalis "anonymous" kapcsolot, aminek hatasara a listazasnal a `user` mezo uresen jelenik meg az adotxt uzenetnel
- Adjatok lehetoseget az uzenetek torlesere, de minden felhasznalo csak a sajat uzenetet tudja torolni
- Egeszitsetek ki az uzeneteket valaszhato nick nevekkel a user neve melle, amit a bekuldesnel megvalaszthat a bekuldo
- Adjatok egy endpointot amin keresztul uj felhasznalot lehet regisztralni a rendszerbe
- Keszithettek webes UI-t HTML es Javascript hasznalataval (React, vagy barmi johet)
