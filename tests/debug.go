package utilgo


type Debugable interface {
	isDebug()(bool)
	setDebug(bool)
}
