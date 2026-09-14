package main
import("database/sql";"embed";"fmt";"net/http";"os";_ "modernc.org/sqlite")
//go:embed migrations/001_init.sql
var migration embed.FS
func main(){p:=os.Getenv("DATABASE_PATH");if p==""{p="lesson.db"};d,e:=sql.Open("sqlite",p);if e!=nil{panic(e)};defer d.Close();s,_:=migration.ReadFile("migrations/001_init.sql");if _,e=d.Exec(string(s));e!=nil{panic(e)};http.HandleFunc("/health",func(w http.ResponseWriter,_ *http.Request){fmt.Fprint(w,`{"status":"ok"}`)});http.ListenAndServe(":8080",nil)}
