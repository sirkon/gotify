package gotify

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type quadro struct {
	orig    string
	public  string
	private string
	pkg     string
}

func TestGotify(t *testing.T) {
	data := []quadro{
		{
			orig:    "data_tape_loader",
			public:  "DataTapeLoader",
			private: "dataTapeLoader",
			pkg:     "datatapeloader",
		},
		{
			orig:    "Donald_duck",
			public:  "DonaldDuck",
			private: "donaldDuck",
			pkg:     "donaldduck",
		},
		{
			orig:    "userId",
			public:  "UserID",
			private: "userID",
			pkg:     "userid",
		},
		{
			orig:    "user_id",
			public:  "UserID",
			private: "userID",
			pkg:     "userid",
		},
		{
			orig:    "St.Loop",
			public:  "StLoop",
			private: "stLoop",
			pkg:     "stloop",
		},
		{
			orig:    "ownerUid",
			public:  "OwnerUID",
			private: "ownerUID",
			pkg:     "owneruid",
		},
		{
			orig:    "ï»¿",
			public:  "UnrecognizedSequence",
			private: "unrecognizedSequence",
			pkg:     "ï»¿",
		},
	}

	goish := New(map[string]string{
		"uid": "UID",
	})

	for _, x := range data {
		require.Equal(t, x.public, goish.Public(x.orig))
		require.Equal(t, x.private, goish.Private(x.orig))
		require.Equal(t, x.pkg, goish.Package(x.orig))
	}

}
