package domain

type Presenter interface {
	Render() string
	RenderChannel(chan<- string)
}
