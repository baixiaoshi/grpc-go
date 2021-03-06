// +build go1.9,!appengine

/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package credentials

import (
	"crypto/tls"
	"errors"
	"net"
	"syscall"
)

type tlsConn struct {
	*tls.Conn
	rawConn net.Conn
}

// implements the syscall.Conn interface
func (c tlsConn) SyscallConn() (syscall.RawConn, error) {
	conn, ok := c.rawConn.(syscall.Conn)
	if !ok {
		return nil, errors.New("RawConn does not implement syscall.Conn")
	}
	return conn.SyscallConn()
}
