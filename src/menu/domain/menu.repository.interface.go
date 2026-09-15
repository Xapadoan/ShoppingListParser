package domain

type IGetMenuCollection = func(id string) (*MenuCollection, *MenuError)
