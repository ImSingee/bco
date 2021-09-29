package bco

type Collector struct {
	TrueCount  int
	FalseCount int
}

func New() Collector {
	return Collector{}
}

// True returns `true` if any `true` is collected
func (e Collector) True() bool {
	return e.TrueCount != 0
}

// False returns `true` if any `false` is collected
func (e Collector) False() bool {
	return e.FalseCount != 0
}

// Add the specific bool
func (e *Collector) Add(b bool) {
	if b {
		e.TrueCount++
	} else {
		e.FalseCount++
	}
}

func (e *Collector) Reset() {
	e.TrueCount = 0
	e.FalseCount = 0
}

func (e *Collector) AddTrue() {
	e.TrueCount++
}

func (e *Collector) AddFalse() {
	e.FalseCount++
}

// C is same as Add, but also return received bool
func (e *Collector) C(b bool) bool {
	if b {
		e.TrueCount++
	} else {
		e.FalseCount++
	}

	return b
}

// Do a function and collect the returned bool
func (e *Collector) Do(f func() bool) {
	e.Add(f())
}

// D is same as Do, but also return received bool
func (e *Collector) D(f func() bool) bool {
	return e.C(f())
}

// CollectAll execute some functions and collect all
func (e *Collector) CollectAll(processes ...func() bool) {
	for _, p := range processes {
		e.Do(p)
	}
}

// Do some things and collect all
func Do(processes ...func() bool) Collector {
	e := New()
	e.CollectAll(processes...)
	return e
}
