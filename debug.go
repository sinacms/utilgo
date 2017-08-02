package lib


type Debugable interface {
	isDebug()(bool)
	setDebug(bool)
}
