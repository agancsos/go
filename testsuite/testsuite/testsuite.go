package testsuite
import (
	"./tests"
)

var UnitTests = map[string]tests.UnitTest {
	"Test1" : &tests.Test1{},
};

type TestSuite struct {
}


func (a TestSuite) Invoke(b tests.UnitTest) bool {
	println("Running: " + b.GetName());
	b.OnInvoke();
	return true;
}

func (a TestSuite) Initialize() {
	/*files, err := ioutil.ReadDir("./testsuite/tests");
	if (err == nil) {
		for _, f := range files {
			var name = strings.Replace(f.Name(), ".go", "", -1);
			UnitTests[name] = &(name){}
		}
	}*/
}
func (a TestSuite) ListTests() {
	for key := range UnitTests {
		println(" * " + key);
	}
}

