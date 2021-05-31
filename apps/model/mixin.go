package model

type MixinThing struct {
	Name        string `cast:"name" validate:"required"`
	Url         string `cast:"url"`
	Image       string `cast:"image"`
	Description string `cast:"description"`
}
