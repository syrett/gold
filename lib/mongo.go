package lib

/*-------------------------------------------------------------------
* @copyright 2015 有量(上海)信息技术有限公司
* @author liyouyou <youyou.li78@gmail.com>
#Time-stamp: <liyouyou 2015-08-14 17:36:36>
* @doc
* mongo.go
* @end
* Created : 2015/08/14 09:19:07 liyouyou

-------------------------------------------------------------------*/

import (
	"errors"
	"log"

	mgo "gopkg.in/mgo.v2"
)

type MongoDB struct {
	Dial_url string
	sess     *mgo.Session
}

const (
	RETRY_NUM = 3
)

func NewMongoDB(dial_url string) *MongoDB {
	return &MongoDB{Dial_url: dial_url}
}

func (m *MongoDB) Connection() (err error) {
	log.Printf("MongoDB:%s, connecting...", m.Dial_url)
	if m.Dial_url == "" {
		err = errors.New("not found dial_url")
	}
	var new_sess *mgo.Session
	new_sess, err = mgo.Dial(m.Dial_url)
	if err != nil {
		return
	}

	if new_sess == nil {
		err = errors.New("session is nil")
		return
	}
	m.sess = new_sess
	return
}

func (m *MongoDB) GetSession() (s *mgo.Session, err error) {
	s = m.sess.Clone()
	if s != nil {
		return
	} else {
		err = m.Reconn()
		if err != nil {
			return
		}
		s = m.sess.Clone()
		if s == nil {
			err = errors.New("Failed to fetch session")
		}
		return
	}
}

func (m *MongoDB) Reconn() (err error) {
	for i := 0; i < RETRY_NUM; i++ {
		err = m.Connection()
		if err == nil {
			break
		}
	}
	return
}
