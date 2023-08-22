package main

import (
    "database/sql"
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Open a connection to the MySQL database
    db, err := sql.Open("mysql", os.Getenv("MYSQL_USERNAME")+":"+os.Getenv("MYSQL_PASSWORD")+"@tcp("+os.Getenv("MYSQL_HOST")+":"+os.Getenv("MYSQL_PORT")+")/"+os.Getenv("MYSQL_DATABASE"))
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create a new HTTP server
    http.HandleFunc("/api/insert", func(w http.ResponseWriter, r *http.Request) {
        // Get the title, artist, and year from the request body
        title := r.FormValue("title")
        artist := r.FormValue("artist")
        year := r.FormValue("year")

        // Insert the record into the database
        stmt, err := db.Prepare("INSERT INTO music (title, artist, year) VALUES (?, ?, ?)")
        if err != nil {
            panic(err)
        }
        defer stmt.Close()

        _, err = stmt.Exec(title, artist, year)
        if err != nil {
            panic(err)
        }

        // Notify the client that the record was inserted successfully
        fmt.Fprintf(w, "Record inserted successfully")
    })

    http.HandleFunc("/api/select", func(w http.ResponseWriter, r *http.Request) {
        // Select all the records from the database
        rows, err := db.Query("SELECT * FROM music")
        if err != nil {
            panic(err)
        }
        defer rows.Close()

        // Iterate over the rows and print the results
        for rows.Next() {
            var title, artist, year string
            err := rows.Scan(&title, &artist, &year)
            if err != nil {
                panic(err)
            }

            fmt.Fprintf(w, "Title: %s\nArtist: %s\nYear: %s\n", title, artist, year)
        }
    })

    http.ListenAndServe(":8080", nil)
}
