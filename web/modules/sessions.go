package modules

import (
	"time"
)

type Session struct {
	username string
	expiry   time.Time
}


func (s *Session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type ServiceSession struct {
	// db Storage
	cache *Cache
}


func NewServiceSession() *ServiceSession{
	ss := &ServiceSession{
		cache: NewCache(),
	}
	return ss 
}


func (ss *ServiceSession) Login(){




}