package utilgo


type errContainer struct {
	errs []error
}

func NewErrContainer()*errContainer{
	ec := &errContainer{
		errs: []error{},
	}
	return ec
}
func (this *errContainer)Get()[]error{
	return this.errs
}
func (this *errContainer)HasErr()bool{
	return len(this.errs)!=0
}
func (this *errContainer)GetStrings()*[]string{
	strs := []string{}
	for _,e := range this.errs {
		strs = append(strs, e.Error())
	}
	return &strs
}


func (this *errContainer)Reset(){
	this.errs = []error{}
}
func (this *errContainer)Add(err error){
	if err != nil {
		this.errs = append(this.errs, err)
		IgnoreErr(err)
	}
}