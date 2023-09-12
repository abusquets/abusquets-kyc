package presenter

type IPresenter[InT any, OutT any] interface {
	Present(data *InT)
	Result() OutT
}

type IPresenterIn[InT any] interface {
	Present(data *InT)
}

type IPresenterOut[OutT any] interface {
	Result() *OutT
}
