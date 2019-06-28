package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
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
	GetAvatarURL(ChatUser) (string, error)
}
type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	var avatarURL string
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			avatarURL = url
			break
		}
	}
	if avatarURL == "" {
		return "", ErrNoAvatarURL
	}
	return avatarURL, nil
}

type AuthAvatar struct{}
type GravatarAvatar struct{}
type FileSystemAvatar struct{}

var UseAuthAvatar AuthAvatar
var UseGravatarAvatar GravatarAvatar
var UseFileSystemAvatar FileSystemAvatar

func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return url, nil
}

func (GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return fmt.Sprintf("//www.gravatar.com/avatar/%s", u.UniqueID()), nil
}

func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	var returnPath string
	if _, err := ioutil.ReadDir("avatars"); err != nil {
		return "", ErrNoAvatarURL
	}
	files, _ := ioutil.ReadDir("avatars")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if match, _ := path.Match(u.UniqueID()+"*", file.Name()); match {
			returnPath = "/avatars/" + file.Name()
			break
		}
	}
	if returnPath == "" {
		return "", ErrNoAvatarURL
	}
	return returnPath, nil
}
