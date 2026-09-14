package main
import("testing";"os")
func TestCompile(t *testing.T){if _,err:=os.Stat("migrations/001_init.sql");err!=nil{t.Fatal(err)}}
