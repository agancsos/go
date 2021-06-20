package tests
type UnitTest interface {
    OnInvoke()
    GetName()   string
}

type EmptyUnitTest struct {
}

func (a *EmptyUnitTest) OnInvoke() { }
func (a *EmptyUnitTest) GetName() string { return "Empty"; }
