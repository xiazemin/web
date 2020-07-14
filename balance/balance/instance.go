package balance
import (
	"fmt"
	"hash/crc32"
	"strings"
	"strconv"
)

type Instance struct {
	Ip   string
	Port int
}

func NewInstance(ip string, port int) *Instance {
	return &Instance{
		Ip:   ip,
		Port: port,
	}
}

func (this *Instance) String() string {
	return this.Ip + ":" + fmt.Sprintf("%d", this.Port)
}

//这里我选择crc32，具体情况具体安排
func (this *Instance) hashKey(host Instance) uint32 {
	scratch := []byte(fmt.Sprint(host))
	return crc32.ChecksumIEEE(scratch)
}

func GetInstance(host string)(*Instance,error){
	aIns:=strings.Split(host,":")
	port,err:=strconv.ParseInt(aIns[1],10,10)
	if aIns[0]=="" || err!=nil{
		return nil,err
	}
	return &Instance{
		Ip:aIns[0],
		Port:int(port),
	},nil
}
