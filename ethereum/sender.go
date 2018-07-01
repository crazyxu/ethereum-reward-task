package ethereum

import "github.com/ethereum/go-ethereum/rpc"

type txSender struct {

}

func newEthSender(conn rpc.Client) (*txSender, error){

}

func (*txSender) newTask(taskType uint, name, desc string, reward uint64, owner, passPhrase, serial string){

}