package main

import(
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "net/http"
    "html/template"
    "log"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
 
        err := r.ParseForm()
        if err != nil {
            log.Println(err)
        }
        model := r.FormValue("model")
        company := r.FormValue("company")
        price := r.FormValue("price")
 
        _, err = database.Exec("insert into productdb.Products (model, company, price) values (?, ?, ?)", 
          model, company, price)
 
        if err != nil {
            log.Println(err)
          }
        http.Redirect(w, r, "/", 301)
    }else{
        http.ServeFile(w,r, "templates/create.html")
    }
}
 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
 
    rows, err := database.Query("select * from productdb.Products")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()
    products := []Product{}
     
    for rows.Next(){
        p := Product{}
        err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
        if err != nil{
            fmt.Println(err)
            continue
        }
        products = append(products, p)
    }
 
    tmpl, _ := template.ParseFiles("templates/index.html")
    tmpl.Execute(w, products)
}

func main() {
    db, err := sql.Open("mysql", "root:password@/productdb")
     
    if err != nil {
        log.Println(err)
    }
    database = db
    defer db.Close()
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/create", CreateHandler)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)



	// validSolution([
	// 	[5, 3, 4, 6, 7, 8, 9, 1, 2],
	// 	[6, 7, 2, 1, 9, 5, 3, 4, 8],
	// 	[1, 9, 8, 3, 4, 2, 5, 6, 7],
	// 	[8, 5, 9, 7, 6, 1, 4, 2, 3],
	// 	[4, 2, 6, 8, 5, 3, 7, 9, 1],
	// 	[7, 1, 3, 9, 2, 4, 8, 5, 6],
	// 	[9, 6, 1, 5, 3, 7, 2, 8, 4],
	// 	[2, 8, 7, 4, 1, 9, 6, 3, 5],
	// 	[3, 4, 5, 2, 8, 6, 1, 7, 9]
	//  ]);

	var hor1, hor2, hor3, hor4, hor5, hor6, hor7, hor8, hor9  [9]int
	hor1[0] = 1
	hor2[0] = 2
	hor3[0] = 3
	hor4[0] = 4
	hor5[0] = 5
	hor6[0] = 6
	hor7[0] = 7
	hor8[0] = 8
	hor9[0] = 9

	var ver1 = [9]int{hor1[0], hor2[0], hor3[0], hor4[0], hor5[0], hor6[0], hor7[0], hor8[0], hor9[0] }
	var ver2 = [9]int{hor1[1], hor2[1], hor3[1], hor4[1], hor5[1], hor6[1], hor7[1], hor8[1], hor9[1] }
	var ver3 = [9]int{hor1[2], hor2[2], hor3[2], hor4[2], hor5[2], hor6[2], hor7[2], hor8[2], hor9[2] }
	var ver4 = [9]int{hor1[3], hor2[3], hor3[3], hor4[3], hor5[3], hor6[3], hor7[3], hor8[3], hor9[3] }
	var ver5 = [9]int{hor1[4], hor2[4], hor3[4], hor4[4], hor5[4], hor6[4], hor7[4], hor8[4], hor9[4] }
	var ver6 = [9]int{hor1[5], hor2[5], hor3[5], hor4[5], hor5[5], hor6[5], hor7[5], hor8[5], hor9[5] }
	var ver7 = [9]int{hor1[6], hor2[6], hor3[6], hor4[6], hor5[6], hor6[6], hor7[6], hor8[6], hor9[6] }
	var ver8 = [9]int{hor1[7], hor2[7], hor3[7], hor4[7], hor5[7], hor6[7], hor7[7], hor8[7], hor9[7] }
	var ver9 = [9]int{hor1[8], hor2[8], hor3[8], hor4[8], hor5[8], hor6[8], hor7[8], hor8[8], hor9[8] }
	

	a, _ := fmt.Println(hor1, hor2, hor3, hor4, hor5, hor6, hor7, hor8, hor9, ver1, ver2, ver3, ver4, ver5, ver6, ver7, ver8, ver9 )
	if false{
		fmt.Println(a )
	}
	ver1[0] = 10

	fmt.Println()

}
