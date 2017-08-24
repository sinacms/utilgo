package utilgo


func If(exp bool, yes , no interface{})interface{}{
	if exp {
		return yes
	}else{
		return no
	}
}