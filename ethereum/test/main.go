package main

import (
	"github.com/ethereum/go-ethereum/rpc"
	"context"
	"github.com/prometheus/common/log"
	"fmt"
	"time"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/crazyxu/ethereum-reward-task/ethereum/test/contract"
	"sync"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

const TaskContractAddress = "0xc40024e25dc6d2ff9227528c98e604533996ad52"
const ownerAddress = "0x331a6b286e596ae0e1c8818537d50c3f92a74920"
const OnwerKeyStore = `{"address":"331a6b286e596ae0e1c8818537d50c3f92a74920","crypto":{"cipher":"aes-128-ctr","ciphertext":"78a2f38c5e96f16f1adc83d17e90995112d6e2529a89f53d687ee5bd03a872bd","cipherparams":{"iv":"4fed6bf456fe3adcbe289228f182e7b7"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"46d59a7d402557d93b5a465caf582bcb93069528ac5af8a78433098230904234"},"mac":"4d26f662910f311e358a92a6a4a2e6fd8cc9b39143f9e8f2968e2c6a9d58b519"},"id":"97b243cd-0b36-4237-b25c-46b0df7b1754","version":3}`
const workerAddress = "0x0f7284598a33be8cb2e29cc474b18557003e757f"
const IPC = "/Users/xucan/ethereum/chain/geth.ipc"
const KEYSTORE = "/Users/xucan/ethereum/chain/keystore"

func main() {
	//rpc client
	client, err := rpc.DialIPC(context.Background(), IPC)
	if err != nil{
		log.Fatalf("DialIPC %s", err)
	}
	//eth client
	ethClient := ethclient.NewClient(client)
	stdTask, err := contract.NewStdTasks(common.HexToAddress(TaskContractAddress), ethClient)
	if err != nil{
		log.Fatalf("NewStdTasks %s", err)
	}
	//context
	ctx, _ := context.WithTimeout(context.Background(), time.Minute * 5)
	var wg sync.WaitGroup
	//watch event
	wg.Add(1)
	go func() {
		defer wg.Done()
		addTaskEvents := make(chan *contract.StdTasksAddTask)
		opts := &bind.WatchOpts{}
		subscription, err := stdTask.WatchAddTask(opts, addTaskEvents)
		if err != nil{
			log.Fatalf("WatchAddTask %s", err)
		}
		for{
			select {
			case err := <- subscription.Err():
				log.Fatalf("subscription %s", err)
				return
			case event := <-addTaskEvents:
				fmt.Print(event)
			case <-ctx.Done():
				subscription.Unsubscribe()
				return
			}
		}
	}()
	//create a account
	wg.Add(1)
	go func() {
		defer wg.Done()
		ks := keystore.NewKeyStore(KEYSTORE, 8, 8)
		a, err := ks.NewAccount("123456")
		if err != nil{
			log.Fatalf("NewAccount %s", err)
		}
		log.Info(a.Address.Hex())
	}()

	//add a task
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	auth, err := bind.NewTransactor(strings.NewReader(OnwerKeyStore), "123456")
	//	if err != nil{
	//		log.Fatalf("NewTransactor %s", err)
	//	}
	//	opts := bind.TransactOpts{
	//		From: common.HexToAddress(ownerAddress),
	//		Value: big.NewInt(int64(100000)),
	//		Signer:auth.Signer,
	//	}
	//	tx, err := stdTask.AddTask(&opts,"标题", "描述", "111222")
	//	if err != nil{
	//		log.Fatalf("AddTask %s", err)
	//	}
	//	log.Info(tx)
	//}()
	wg.Wait()
	fmt.Print("bye")
}