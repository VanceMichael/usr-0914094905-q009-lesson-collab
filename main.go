package main
import("database/sql";"embed";"net/http";"os";"github.com/labstack/echo/v4";_ "modernc.org/sqlite")
//go:embed migrations/001_init.sql
var migration embed.FS
func main(){p:=os.Getenv("DATABASE_PATH");if p==""{p="lesson.db"};d,e:=sql.Open("sqlite",p);if e!=nil{panic(e)};defer d.Close();s,_:=migration.ReadFile("migrations/001_init.sql");if _,e=d.Exec(string(s));e!=nil{panic(e)};app:=echo.New();app.GET("/health",func(c echo.Context)error{return c.JSON(http.StatusOK,map[string]string{"status":"ok"})});app.Start(":8080")}
