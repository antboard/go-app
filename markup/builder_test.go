package markup

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewCompoBuilder(t *testing.T) {
	NewCompoBuilder()
}

func TestCompoBuilderRegister(t *testing.T) {
	c := &ValidCompo{}
	ct := reflect.TypeOf(*c)
	cname := strings.ToLower(ct.String())

	b := make(compoBuilder)
	err := b.Register(c)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := b[cname]; !ok {
		t.Fatalf("%s should have been registered", cname)
	}

	empc := &EmptyCompo{}
	if err = b.Register(empc); err == nil {
		t.Fatal("register empty component should returns an error")
	}
}

func TestCompoBuilderNew(t *testing.T) {
	c := &ValidCompo{}
	cname := "markup.validcompo"
	b := make(compoBuilder)
	b.Register(c)

	n, err := b.New(cname)
	if err != nil {
		t.Fatal(err)
	}
	if n == nil {
		t.Fatalf("%s should have been created: %v", cname, n)
	}

	if _, err = b.New("unknown"); err == nil {
		t.Fatal("unknown should not have been created")
	}
}

func TestNormalizeCompoName(t *testing.T) {
	if name := "lib.FooBar"; normalizeCompoName(name) != "lib.foobar" {
		t.Errorf(`name should be "lib.foobar": "%s"`, name)
	}

	if name := "main.FooBar"; normalizeCompoName(name) != "foobar" {
		t.Errorf(`name should be "foobar": "%s"`, name)
	}
}
