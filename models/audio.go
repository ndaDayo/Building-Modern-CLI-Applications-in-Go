package models

type Audio struct {
	Id       string
	Path     string
	Metadata Metadata
	Status   string
	Error    []error
}
