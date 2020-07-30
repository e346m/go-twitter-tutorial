package handler

type ItemGetter interface {
    GetItem(id string) error
}

type Handler struct {
    ItemGetter
}
