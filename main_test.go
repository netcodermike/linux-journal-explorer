package main
import "testing"
func TestRecords(t *testing.T){rows,e:=parse("{\"__REALTIME_TIMESTAMP\":\"123\",\"MESSAGE\":\"hello\",\"_SYSTEMD_UNIT\":\"demo.service\",\"PRIORITY\":\"6\"}\n");if e!=nil||len(rows)!=1||rows[0].Message!="hello"||rows[0].Unit!="demo.service"{t.Fatal(rows,e)};if _,e=parse("bad");e==nil{t.Fatal("malformed JSON accepted")}}
