package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

//ErrNoAvatar is the error that is return when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: Unable to get an avatar URL.")

// Avatar represents types capable of representing
// user profile pictures
type Avatar interface {
	//GetAvatarURL get the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is return if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}
type GravatarAvatar struct{}

var UseAuthAvatar AuthAvatar
var UseGravatarAvatar GravatarAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	url, ok := c.userData["avatar_url"]
	if !ok {
		return "", ErrNoAvatarURL
	}
	urlStr, ok := url.(string)
	if !ok {
		return "", ErrNoAvatarURL
	}
	return urlStr, nil
}

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	userID, ok := c.userData["userid"]
	if !ok {
		return "", ErrNoAvatarURL
	}
	userIDStr, ok := userID.(string)
	if !ok {
		return "", ErrNoAvatarURL
	}
	m := md5.New()
	io.WriteString(m, strings.ToLower(userIDStr))
	return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
}
