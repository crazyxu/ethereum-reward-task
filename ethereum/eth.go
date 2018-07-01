package ethereum

import (
	"github.com/ethereum/go-ethereum/rpc"
	"fmt"
	"time"
	"errors"
)

const TransactionTimeout = 10 * time.Second

type EthManager struct {
	txSender
	txReceiver
}

func NewEthManager(ipcPath string) (*EthManager, error){
	client, err := rpc.DialIPC(nil, ipcPath)
	client.Subscribe()
	if err != nil{
		return nil, fmt.Errorf("NewEthManager %s", err)
	}

}


func (eth *EthManager) CreateEthAccount(key, passPhrase, serial string) error{
	return nil
}

func (eth *EthManager) PublishTask(taskType uint, name, desc string, reward uint64, ownerAddr, passPhrase string) (uint, error){
	serial := newSerial()
	r := make(chan interface{})
	//ctx := context.WithTimeout(context.Background(), TransactionTimeout)
	eth.listen(serial, r)
	eth.newTask(taskType, name, desc, reward, ownerAddr, passPhrase, serial)
	select {
		case <- time.After(TransactionTimeout):
		case <- r:{
			//.....
		}
	}
	return 0, errors.New("timeout")
}

func newSerial() string{
	return ""
}