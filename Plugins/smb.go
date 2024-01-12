package Plugins

import (
	"errors"
	"fmt"
	"github.com/stacktitan/smb/smb"
	"github.com/wjlin0/fscan/common"
	"strings"
	"time"
)

func SmbScan(info *common.HostInfo) (tmperr error) {
	if common.IsBrute {
		return nil
	}
	starttime := time.Now().Unix()
	for _, user := range common.Userdict["smb"] {
		for _, pass := range common.Passwords {
			pass = strings.Replace(pass, "{user}", user, -1)
			flag, err := doWithTimeOut(info, user, pass)
			if flag == true && err == nil {
				var result string
				if common.Domain != "" {
					result = fmt.Sprintf("[+] SMB %v:%v:%v\\%v %v", info.Host, info.Ports, common.Domain, user, pass)
				} else {
					result = fmt.Sprintf("[+] SMB %v:%v:%v %v", info.Host, info.Ports, user, pass)
				}
				common.LogSuccess(result)
				return err
			} else {
				errlog := fmt.Sprintf("[-] smb %v:%v %v %v %v", info.Host, 445, user, pass, err)
				errlog = strings.Replace(errlog, "\n", "", -1)
				common.LogError(errlog)
				tmperr = err
				if common.CheckErrs(err) {
					return err
				}
				if time.Now().Unix()-starttime > (int64(len(common.Userdict["smb"])*len(common.Passwords)) * common.Timeout) {
					return err
				}
			}
		}
	}
	return tmperr
}

func SmblConn(info *common.HostInfo, user string, pass string, signal chan struct{}) (flag bool, err error) {
	flag = false
	Host, Username, Password := info.Host, user, pass
	options := smb.Options{
		Host:        Host,
		Port:        445,
		User:        Username,
		Password:    Password,
		Domain:      common.Domain,
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			flag = true
		}
	}
	signal <- struct{}{}
	return flag, err
}

func doWithTimeOut(info *common.HostInfo, user string, pass string) (flag bool, err error) {
	signal := make(chan struct{})
	go func() {
		flag, err = SmblConn(info, user, pass, signal)
	}()
	select {
	case <-signal:
		return flag, err
	case <-time.After(time.Duration(common.Timeout) * time.Second):
		return false, errors.New("time out")
	}
}
