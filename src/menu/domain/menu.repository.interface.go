package domain

type IGetMenuCollection = func(id string) (*MenuCollection, *MenuError)
type IStoreMenuCollection = func(collection *MenuCollection) *MenuError
type IEditMenuCollection = func(collection *MenuCollection) *MenuError
