// Copyright © 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause
//
package rpc

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"net"
	"sync/atomic"
	"time"
	. "github.com/avekceeb/nfsverificator/common"
	"github.com/avekceeb/nfsverificator/xdr"
	"os"
	"syscall"
	"os/user"
)

const (
	MsgAccepted = iota
	MsgDenied
)

const (
	Success = iota
	ProgUnavail
	ProgMismatch
	ProcUnavail
	GarbageArgs
	SystemErr
)

const (
	RpcMismatch = iota
)

var xid uint32

func init() {
	// seed the XID (which is set by the client)
	xid = rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
}

type Client struct {
	*tcpTransport
}

func DialTCP(network string, ldr *net.TCPAddr, addr string) (*Client, error) {
	a, err := net.ResolveTCPAddr(network, addr)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP(a.Network(), ldr, a)
	if err != nil {
		return nil, err
	}

	t := &tcpTransport{
		r:  bufio.NewReader(conn),
		wc: conn,
	}

	return &Client{t}, nil
}

type message struct {
	Xid     uint32
	Msgtype uint32
	Body    interface{}
}

func maybe41CallBack(r *io.ReadSeeker) (bool) {
	mtype, _ := xdr.ReadUint32(*r)
	if 0 != mtype {
		return false
	}
	// rpcVersion
	xdr.ReadUint32(*r)
	// program
	xdr.ReadUint32(*r)
	return true
}

func (c *Client) Call(call interface{}) (io.ReadSeeker, error) {
	retries := 1

	msg := &message{
		Xid:  atomic.AddUint32(&xid, 1),
		Body: call,
	}

retry:
	w := new(bytes.Buffer)
	if err := xdr.Write(w, msg); err != nil {
		return nil, err
	}

	if _, err := c.Write(w.Bytes()); err != nil {
		return nil, err
	}

listen:
	res, err := c.recv()
	if err != nil {
		return nil, err
	}

	xid, err := xdr.ReadUint32(res)
	if err != nil {
		return nil, err
	}

	if xid != msg.Xid {
		// TODO: this is temporary workaround
		if maybe41CallBack(&res) {
			goto listen
		} else {
			return nil, fmt.Errorf(
				"xid did not match, expected: %x, received: %x",
				msg.Xid, xid)
		}
	}

	mtype, err := xdr.ReadUint32(res)
	if err != nil {
		return nil, err
	}

	if mtype != 1 {
		return nil, fmt.Errorf("message as not a reply: %d", mtype)
	}

	status, err := xdr.ReadUint32(res)
	if err != nil {
		return nil, err
	}

	switch status {
	case MsgAccepted:

		// padding
		_, err = xdr.ReadUint32(res)
		if err != nil {
			panic(err.Error())
		}

		opaque_len, err := xdr.ReadUint32(res)
		if err != nil {
			panic(err.Error())
		}

		_, err = res.Seek(int64(opaque_len), io.SeekCurrent)
		if err != nil {
			panic(err.Error())
		}

		acceptStatus, _ := xdr.ReadUint32(res)

		switch acceptStatus {
		case Success:
			return res, nil
		case ProgUnavail:
			return nil, fmt.Errorf("%s", "rpc: PROG_UNAVAIL - server does not recognize the program number")
		case ProgMismatch:
			return nil, fmt.Errorf("rpc: PROG_MISMATCH - program version does not exist on the server")
		case ProcUnavail:
			return nil, fmt.Errorf("rpc: PROC_UNAVAIL - unrecognized procedure number")
		case GarbageArgs:
			// emulate Linux behaviour for GARBAGE_ARGS
			if retries > 0 {
				Debugf("Retrying on GARBAGE_ARGS per linux semantics")
				retries--
				goto retry
			}

			return nil, fmt.Errorf("rpc: GARBAGE_ARGS - rpc arguments cannot be XDR decoded")
		case SystemErr:
			return nil, fmt.Errorf("rpc: SYSTEM_ERR - unknown error on server")
		default:
			return nil, fmt.Errorf("rpc: unknown accepted status error: %d", acceptStatus)
		}

	case MsgDenied:
		rejectStatus, _ := xdr.ReadUint32(res)
		switch rejectStatus {
		case RpcMismatch:

		default:
			return nil, fmt.Errorf("rejectedStatus was not valid: %d", rejectStatus)
		}

	default:
		return nil, fmt.Errorf("rejectedStatus was not valid: %d", status)
	}

	panic("unreachable")
}

func isAddrInUse(err error) bool {
	if er := (err.(*net.OpError)); er != nil {
		if syser, ok := er.Err.(*os.SyscallError); ok {
			return syser.Err == syscall.EADDRINUSE
		}
	}

	return false
}

func DialService(addr string, port int) (*Client, error) {
	var (
		ldr    *net.TCPAddr
		client *Client
	)

	usr, err := user.Current()

	// Unless explicitly configured, the target will likely reject connections
	// from non-privileged ports.
	if err == nil && usr.Uid == "0" {
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

		var p int
		for {
			p = r1.Intn(1024)
			if p < 0 {
				continue
			}

			ldr = &net.TCPAddr{
				Port: p,
			}

			raddr := fmt.Sprintf("%s:%d", addr, port)
			Debugf("Connecting to %s", raddr)

			client, err = DialTCP("tcp", ldr, raddr)
			if err == nil {
				break
			}
			// bind error, try again
			if isAddrInUse(err) {
				continue
			}

			return nil, err
		}

		Debugf("using random port %d -> %d", p, port)
	} else {
		raddr := fmt.Sprintf("%s:%d", addr, port)
		Debugf("Connecting to %s from unprivileged port", raddr)

		client, err = DialTCP("tcp", ldr, raddr)
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}

