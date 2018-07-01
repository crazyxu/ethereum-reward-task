package model

type BackendApi interface{

	//CreateEthAccount create a ethereum account by  key and passPhrase
	CreateEthAccount(key, passPhrase string) error

	//PublishTask public a Task, backend will commit a transaction waiting a miner to execute.
	//A `AddTask` event will be emit after that, the serial will be the context,
	//different task of type will be stored in different contract.
	PublishTask(taskType uint, name, desc string, reward uint64, ownerAddr, passPhrase string) (uint, error)
}