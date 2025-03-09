package handler

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}
