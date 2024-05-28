package blockchain

func Init(c Connect,addr string)error{
	err:=c.New(addr)
	if err!=nil{
		return err
	}
	return nil
}