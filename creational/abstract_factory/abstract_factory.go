package abstractfactory

import "errors"

type UIFactory interface {
	CreateHeader(content string) (Header, error)
	CreateBody(content string) (Body, error)
}

type Header interface {
	Render()
}

type Body interface {
	Render()
	Clear()
}

// GUI Factory

type GUIFactory struct{}

func NewGUIFactory() UIFactory {
	return &GUIFactory{}
}

func (f *GUIFactory) CreateHeader(content string) (Header, error) {
	if content == "" {
		return nil, errors.New("header content cannot be empty")
	}
	return &GUIHeader{content: content}, nil
}

func (f *GUIFactory) CreateBody(content string) (Body, error) {
	if content == "" {
		return nil, errors.New("body content cannot be empty")
	}
	return &GUIBody{content: content}, nil
}

// GUI Header

type GUIHeader struct {
	content string
}

func (h GUIHeader) Render() {}

// GUI Body

type GUIBody struct {
	content string
}

func (b GUIBody) Render() {}

func (b *GUIBody) Clear() {
	b.content = ""
}

// TUI Factory

type TUIFactory struct{}

func NewTUIFactory() UIFactory {
	return &TUIFactory{}
}

func (f *TUIFactory) CreateHeader(content string) (Header, error) {
	if content == "" {
		return nil, errors.New("body content cannot be empty")
	}
	return &TUIHeader{content: content}, nil
}

func (f *TUIFactory) CreateBody(content string) (Body, error) {
	if content == "" {
		return nil, errors.New("body content cannot be empty")
	}
	return &TUIBody{content: content}, nil
}

// TUI Header

type TUIHeader struct {
	content string
}

func (h TUIHeader) Render() {}

// TUI Body

type TUIBody struct {
	content string
}

func (b TUIBody) Render() {}

func (b *TUIBody) Clear() {
	b.content = ""
}
