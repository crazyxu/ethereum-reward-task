package ethereum

type txReceiver struct {
	events map[string]chan <- interface{}
}

func newTxReceiver() (*txReceiver, error){
	txReceiver := txReceiver{
		events: make(map[string]chan <- interface{}, 0),
	}
	go txReceiver.startListen()
	return &txReceiver, nil
}

func (tx *txReceiver) startListen(){
}

func (tx *txReceiver) listen(serial string, r chan <- interface{} ){
	tx.events[serial] = r
}
