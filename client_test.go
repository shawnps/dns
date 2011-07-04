package dns

import (
	"testing"
)

func TestClientSync(t *testing.T) {
	m := new(Msg)
        m.SetQuestion("miek.nl", TypeSOA)

        c := NewClient()
        r := c.Exchange(m, "85.223.71.124:53")

	if r != nil && r.Rcode != RcodeSuccess {
		t.Log("Failed to get an valid answer")
		t.Fail()
		t.Logf("%v\n", r)
	}
}


func helloMiek(w RequestWriter, r *Msg) {
        w.Send(r)
        reply, _ := w.Receive()
        w.Write(reply)
}

func TestClientASync(t *testing.T) {
        HandleQueryFunc("miek.nl", helloMiek) // All queries for miek.nl will be handled by HelloMiek
        ListenAndQuery(nil, nil)

	m := new(Msg)
        m.SetQuestion("miek.nl", TypeSOA)

                println("SENDING")
        c := NewClient()
                println("SENDING II")
        c.Do(m, "85.223.71.124:53")

                println("REPLY")
forever:
        for {
                select {
                case n := <-DefaultReplyChan:
                println("REPLY")
                        if n[1] != nil && n[1].Rcode != RcodeSuccess {
                                t.Log("Failed to get an valid answer")
                                t.Fail()
                                t.Logf("%v\n", n[1])
                        }
                        break forever
                }
        }
}

// TestClientEDNS
/*
func TestResolverEdns(t *testing.T) {


	// Add EDNS rr
	edns := new(RR_OPT)
	edns.Hdr.Name = "." // must . be for edns
	edns.Hdr.Rrtype = TypeOPT
	// You can handle an OTP RR as any other, but there
	// are some convience functions
	edns.SetUDPSize(2048)
	edns.SetDo()
	edns.Option = make([]Option, 1)
	edns.SetNsid("") // Empty to request it

	// ask something
	m.Question[0] = Question{"powerdns.nl", TypeDNSKEY, ClassINET}
	m.Extra = make([]RR, 1)
	m.Extra[0] = edns

	in, _ := res.Query(m)
	if in != nil {
		if in.Rcode != RcodeSuccess {
			t.Logf("%v\n", in)
			t.Log("Failed to get an valid answer")
			t.Fail()
		}
	}
}
*/

/*
func TestResolverTsig(t *testing.T) {
	res := new(Resolver)
	res.Servers = []string{"127.0.0.1"}
	res.Timeout = 2
	res.Attempts = 1

	m := new(Msg)
	m.MsgHdr.RecursionDesired = true //only set this bit
	m.Question = make([]Question, 1)

	// ask something
	m.Question[0] = Question{"powerdns.nl", TypeDNSKEY, ClassINET}
	m.Extra = make([]RR, 1)
	m.Id = Id()


        tsig := new(Tsig)
        tsig.Name = "miek.nl."
        tsig.Algorithm = HmacMD5
        tsig.Fudge = 300
        tsig.TimeSigned = uint64(time.Seconds())
        tsig.Secret = "ZGZqc2tmZAo="

	in, _ := res.QueryTsig(m,tsig)
	if in != nil {
		if in.Rcode != RcodeSuccess {
			t.Logf("%v\n", in)
			t.Log("Failed to get an valid answer")
		//	t.Fail()
		}
	}
}

func TestAXFR(t *testing.T) {
	res := new(Resolver)
	res.Servers = []string{"127.0.0.1"}
	m := new(Msg)
	m.Question = make([]Question, 1)
	m.Question[0] = Question{"miek.nl", TypeAXFR, ClassINET}

        ch := make(chan Xfr)
        go res.Xfr(m, ch)
	for x := range ch {
		var _ = x
		// fmt.Printf("%v\n",dm.Dns)
	}
	// channel is closed by Axfr() 
}
*/
