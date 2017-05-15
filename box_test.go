package ricebox

import (
	"testing"

	rice "github.com/GeertJohan/go.rice"
	st "github.com/mattes/migrate/source/testing"
)

func TestFS(t *testing.T) {
	box := rice.MustFindBox("testbox")

	d, err := WithInstance(box)
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}

func TestEmbedded(t *testing.T) {
	box := rice.MustFindBox("testbox-embedd")

	d, err := WithInstance(box)
	if err != nil {
		t.Fatal(err)
	}

	st.Test(t, d)
}
