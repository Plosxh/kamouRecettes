package main

import (
  "database/sql"
	"log"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
//  "math/rand"
)


func main()  {
for i := 0; i < 10; i++ {
  genRecette()
}



  return
}

type ingredient struct{
  id int
  name string
  leType string

}

type recette struct{
  nom string
  ingr1 ingredient
  ingr2 ingredient
  ingr3 ingredient
}

func getIngredient(ingr string) ingredient{
var monIngr ingredient
//var nbIngr int
//var nbRand int

db, err := sql.Open("sqlite3","./recettes.db")
if err != nil{
  log.Fatal(err)
}
defer db.Close()

/*nb, err := db.Query("select count(*) from ingredient")
defer nb.Close()
for nb.Next(){
err = nb.Scan(&nbIngr)
if err != nil {
  log.Fatal(err)
}
}
nbRand = rand.Intn(nbIngr)
fmt.Println(nbRand)*/
var id int
var name string
var monType string

rows, err := db.Query("select * from ingredient where type=\""+ingr+"\""+" order by random() limit 1")
defer rows.Close()
for rows.Next(){
err = rows.Scan(&id, &name,&monType)
if err != nil {
  log.Fatal(err)
}
}

err = rows.Err()
if err != nil{
  log.Fatal(err)
}

monIngr.id = id
monIngr.name = name
monIngr.leType = monType
//fmt.Println(monIngr.name)
return monIngr
}


func genRecette(){
var maRecette recette

maRecette.ingr1 = getIngredient("viande")
maRecette.ingr2 = getIngredient("legume")
maRecette.ingr3 = getIngredient("sauce")
maRecette.nom = maRecette.ingr1.name + " accompagné de " + maRecette.ingr2.name + " et sa sauce " + maRecette.ingr3.name

fmt.Println(maRecette.nom)



}
