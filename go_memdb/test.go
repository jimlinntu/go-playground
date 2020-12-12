package main

import (
    "github.com/hashicorp/go-memdb"
    "fmt"
)

type Person struct{
    Email string
    Name string
    Age int
}

func main(){
    schema := &memdb.DBSchema{
        Tables: map[string] *memdb.TableSchema{
            "person": &memdb.TableSchema{
                Name: "person",
                Indexes: map[string] *memdb.IndexSchema{
                    "id": &memdb.IndexSchema{
                        Name: "id",
                        Unique: true,
                        Indexer: &memdb.StringFieldIndex{Field: "Email"},
                    },
                    "age": &memdb.IndexSchema{
                        Name: "age",
                        Unique: false,
                        Indexer: &memdb.IntFieldIndex{Field: "Age"},
                    },
                },
            },
        },
    }
    db, err := memdb.NewMemDB(schema)
    if err != nil {
        panic(err)
    }
    txn := db.Txn(true)
    people := []*Person{
        &Person{"kk@gmail.com", "Kin", 20},
    }
    for _, p := range people {
        if err := txn.Insert("person", p); err != nil {
            panic(err)
        }
    }

    txn.Commit()

    txn = db.Txn(false)
    defer txn.Abort()
    raw, err := txn.First("person", "id", "kk@gmail.com")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%v\n", raw.(*Person))
}
