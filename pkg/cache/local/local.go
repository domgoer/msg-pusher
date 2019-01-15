/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : local.go
#   Created       : 2019/1/15 16:48
#   Last Modified : 2019/1/15 16:48
#   Describe      :
#
# ====================================================*/
package local

import (
	"github.com/patrickmn/go-cache"
	"time"
	cache2 "uuabc.com/sendmsg/pkg/cache"
)

type Client struct {
	c *cache.Cache
}

func (c *Client) Get(s string) ([]byte, error) {
	v, b := c.c.Get(s)
	if !b {
		return nil, cache2.ErrCacheMiss
	}
	return v.([]byte), nil
}

func (c *Client) Put(k string, v []byte, ttl int64) error {
	c.c.Set(k, v, time.Second*time.Duration(ttl))
	return nil
}

func (c *Client) Del(k string) error {
	c.c.Delete(k)
	return nil
}

func (c *Client) Add(k string, v []byte, ttl int64) error {
	return c.c.Add(k, v, time.Second*time.Duration(ttl))
}

// 本地缓存不实现append
// Deprecated: don't use this method.
func (c *Client) Append(k string, v []byte) error {
	return nil
}

// 本地缓存不实现IsMember
// Deprecated: don't use this method.
func (c *Client) IsMember(k string, v []byte) (bool, error) {
	return false, nil
}

func (c *Client) Close() error {
	return nil
}

func NewClient() *Client {
	return &Client{
		c: cache.New(time.Minute*5, time.Minute*10),
	}
}