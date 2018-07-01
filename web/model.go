package web

type User struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	EthAddress string `json:"eth_address"`
}

type Task struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	Name string `json:"name"`
	Description string `json:"description"`
	Reward uint64 `json:"reward"`
}

type TaskTemplate struct {
	Id uint
	Name string `json:"name"`
	Description string `json:"description"`
	Constructor TaskConstructor `json:"constructor"`
	Actions TaskActions `json:"actions"`
	Abi string `json:"abi"`
}

type TaskConstructor struct {

}

type TaskActions struct {

}

type TaskAction struct {
	MethodId string
	Params []MethodParam
}

type MethodParam struct {

}

func (t *TaskTemplate) match(obj interface{}) bool{
	return false
}