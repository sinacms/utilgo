package utilgo

import "sync"

type RWLocker struct {
	mtx sync.RWMutex
}
func NewRWLock()*RWLocker{
	return &RWLocker{}
}

func (l *RWLocker)RLockDo(callback func()){
	l.mtx.RLock()
	defer l.mtx.RUnlock()
	callback()
}
func (l *RWLocker)RWLockDo(callback func()){
	l.mtx.Lock()
	defer l.mtx.Unlock()
	callback()
}

type Locker struct {
	mtx sync.Mutex
}
func NewLock()*Locker{
	return &Locker{}
}
func (l *Locker)LockDo(callback func()){
	l.mtx.Lock()
	defer l.mtx.Unlock()
	callback()
}

type MutexMap struct{
	m    map[interface{}]interface{}
	lock *sync.RWMutex

}
func NewMutexMap() *MutexMap {
	return &MutexMap{
		lock: new(sync.RWMutex),
		m:    make(map[interface{}]interface{}),
	}
}
func (m *MutexMap) Size() int{
	return len(m.m)
}
func (m *MutexMap) Raw() map[interface{}]interface{} {
	return m.m
}
//Get from maps return the k's value
func (m *MutexMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.m[k]; ok {
		return val
	}
	return nil
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *MutexMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.m[k]; !ok {
		m.m[k] = v
	} else if val != v {
		m.m[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *MutexMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.m[k]; !ok {
		return false
	}
	return true
}

func (m *MutexMap) Keys(ignoreNil  bool, keys ...interface{}) []interface{}{
	m.lock.RLock()
	defer m.lock.RUnlock()
	vals := []interface{}{}
	for _,k := range keys {
		if v,ok := m.m[k]; ok {
			vals = append(vals, v)
		}else{
			if !ignoreNil {
				vals = append(vals, nil)
			}
		}
	}
	return vals
}
func (m *MutexMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.m, k)
}