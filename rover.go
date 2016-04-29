package rover

import(
	"fmt"
)

type Field struct {
	ip string
	role string
}

type Field222222232 struct {
	ip string
	role string
}
type Field333777898987 struct {
	ip string
	role string
}

func main(){
	fmt.Println("Hello World");
	c1:= Field {ip:"xxx.xxx.xxx.xxx", role:"eisx"}
	c2:= Field {ip:"xxx.xxx.yyy.xxx", role:"vcenter"}
    //var myarr [] Field
    myarr := [] Field {c1,c2}
	fmt.Println("myarr: ", myarr[0].ip)
    fmt.Println("myarr: ", myarr)

    var m map[string]string
    m = make(map[string]string)
	for _, ipcopy := range myarr{
	  fmt.Println("ip: ", ipcopy )
	  m[ipcopy.ip] = ipcopy.role
	}

	fmt.Println("printing map: ",  m["xxx.xxx.xxx.xxx"])
    fmt.Println("printing map: ",  m["xxx.xxx.yyy.xxx"])

}


