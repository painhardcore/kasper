package file

import (
	"os"
	"testing"
	"time"
)

const testDataFolder = "./"

func TestNew(t *testing.T) {
	f := testDataFolder + "creating_NEW.file"
	_, err := New(f, time.Minute)
	if err != nil {
		t.Error(err)
	}

	_, err = os.Stat(f)
	if err != nil {
		t.Errorf("Didn't create file in testdata folder %s", err)
	}

	_ = os.Remove(f)
}

func Test_store_String(t *testing.T) {
	tests := []struct {
		name string
		s    store
		want string
	}{
		{"Zero count",
			store{
				dur:      time.Minute,
				filename: "test.base",
				data: &Timecount{
					cl:    mockClock{},
					Count: 0,
					Due:   time.Date(2018, time.January, 2, 15, 5, 0, 0, time.UTC),
				},
			},
			"0",
		},
		{"99 count",
			store{
				dur:      time.Minute,
				filename: "test.base",
				data: &Timecount{
					cl:    mockClock{},
					Count: 99,
					Due:   time.Date(2018, time.January, 2, 15, 5, 0, 0, time.UTC),
				},
			},
			"99",
		},
		{"99 count expired",
			store{
				dur:      time.Minute,
				filename: "test.base",
				data: &Timecount{
					cl:    mockClock{},
					Count: 99,
					Due:   time.Date(2018, time.January, 2, 15, 5, 0, 0, time.UTC),
				},
			},
			"99",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.String(); got != tt.want {
				t.Errorf("store.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_write_load_inc(t *testing.T) {
	f := testDataFolder + "notexpired99.file"
	tc := Timecount{
		Due:   time.Date(2018, time.January, 2, 15, 5, 0, 0, time.UTC),
		Count: 99,
		cl:    mockClock{},
	}
	newstore := store{filename: f, dur: time.Minute, data: &tc}
	err := newstore.write()
	if err != nil {
		t.Error(err)
	}

	data := Timecount{cl: mockClock{}}
	c := store{filename: f, dur: time.Minute, data: &data}
	err = c.load()
	if err != nil {
		t.Errorf("Didn't load file in testdata folder %s", err)
	}
	if c.String() != "99" {
		t.Errorf("Wrong number in data: got  %s", c.String())
	}
	err = c.Inc()
	if err != nil {
		t.Errorf("Error while increasing number: got  %s", err)
	}
	if c.String() != "100" {
		t.Errorf("Wrong number in data %s", c.String())
	}
	//change date to expired
	c.data.Due = time.Date(2018, time.January, 2, 15, 4, 0, 0, time.UTC)

	err = c.Inc()
	if err != nil {
		t.Errorf("Error while increasing number %s", err)
	}
	if c.String() != "1" {
		t.Errorf("Wrong number in data: got %s", c.String())
	}
	_ = os.Remove(f)

}
