package tokenstats

type Output interface {
	Output(stats TokenStats)
}
