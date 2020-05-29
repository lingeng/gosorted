package gosorted

type EleInterface interface {
	Less(ele EleInterface) bool
	Equal(ele EleInterface) bool
}
