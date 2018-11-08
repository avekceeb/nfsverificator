package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"fmt"
	"time"
	"errors"
	"strings"
)

//var (
//    c *V40Test
//    export string
//    rootFH NfsFh4
//)

func TimedClient() (error) {
	name := RandString(8)
	if strings.HasPrefix(name, "d") {
		return errors.New("Bad name: " + name)
	}
	for i:=0;i<3;i++ {
		fmt.Println(name, " - ", i)
		time.Sleep(time.Second)
	}
	return nil
}

func Task(j int) {
	for i:=0;i<3;i++ {
		fmt.Println("Box", j, ":", i)
		time.Sleep(time.Second)
	}
}

func BackgroundClient() (error) {
	for i:=0;i<1000;i++ {
		c := NewNFSv40Client(Config.GetHost(),
			Config.GetPort(),
			RandString(8) + ".flash.mob", 0, 0, RandString(8))
		r, _ := c.Compound(
			Setclientid(c.GetClientID(), c.GetCallBack(), 1))
		c.ClientId = r.Resarray[0].Opsetclientid.Resok4.Clientid
		c.Verifier = r.Resarray[0].Opsetclientid.Resok4.SetclientidConfirm
		c.Compound(
			SetclientidConfirm(c.ClientId, c.Verifier))
		var e error
		for x := 0; x < 3; x++ {
			e = c.Null()
			if nil != e {
				return e
			}
			time.Sleep(time.Millisecond * time.Duration(RandInt(1, 1000)))
			c.GetRootFH()
		}
		time.Sleep(time.Millisecond * time.Duration(RandInt(1, 1000)))
	}
	return nil
}

var _ = Describe("Stress", func() {


	Context("Helper", func() {

		It("Pulti", func() {
			c := make(chan string)
			for i := 1; i <= 5; i++ {
    			go func(i int, co chan<- string) {
        			for j := 1; j <= 5; j++ {
            			co <- fmt.Sprintf("hi from %d.%d", i, j)
						time.Sleep(time.Second)
        			}
    			}(i, c)
			}
			for i := 1; i <= 25; i++ {
			    fmt.Println(<-c)
			}
		})

		It("Zulti", func() {
			var N int = 1000
			ack := make(chan bool, N)
			for i := 0; i < N; i++ {
				go func(arg int) {
					Task(arg)
					ack <- true
				}(i)
			}
			for i := 0; i < N; i++ {
				<-ack
			}
		})

		It("Multi", func() {
			var N int = 100
			var err error
			errCh := make(chan error, N)
			for i := 0; i < N; i++ {
				go func() {
					errCh <- BackgroundClient()
				}()
			}
			for i := 0; i < N; i++ {
				err = <-errCh
				if nil != err {
					fmt.Println ("### Client error:", err.Error())
				}
			}
		})

	})

})
