package file

import (
	"log"
	"testing"
	"time"
)

type mockClock struct{}

const testtime = "2018-01-02T15:04:05Z"

//always return the same time
func (mockClock) Now() time.Time {
	tm, err := time.Parse(time.RFC3339, testtime)
	if err != nil {
		log.Fatalf("wrong time format in tests %s", err)
	}
	return tm
}

//Simple check reset function is working
func TestTimecount_reset(t *testing.T) {
	tc := Timecount{
		Due:   time.Date(1989, time.April, 10, 6, 0, 0, 0, time.UTC),
		Count: 99,
		cl:    mockClock{},
	}
	tc.reset(60 * time.Second)
	expected60sCount := int64(0)
	if tc.Count != expected60sCount {
		t.Errorf("Count not reseted. Value: %d, Expect: %d", tc.Count, expected60sCount)
	}
	expected60sDue := "2018-01-02T15:05:05Z"
	if tc.Due.Format(time.RFC3339) != expected60sDue {
		t.Errorf("Due not resetted. Value %s, Expect: %s", tc.Due.Format(time.RFC3339), expected60sDue)
	}
}

func TestTimecount_expired(t *testing.T) {
	type fields struct {
		cl    Clocker
		Count int64
		Due   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"Should be expired",
			fields{
				cl:    mockClock{},
				Count: 10,
				Due:   time.Date(2018, time.January, 2, 15, 4, 0, 0, time.UTC),
			},
			true,
		},
		{
			"Shouldn't be expired",
			fields{
				cl:    mockClock{},
				Count: 10,
				Due:   time.Date(2018, time.January, 2, 15, 5, 0, 0, time.UTC),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tc := &Timecount{
				cl:    tt.fields.cl,
				Count: tt.fields.Count,
				Due:   tt.fields.Due,
			}
			if got := tc.expired(); got != tt.want {
				t.Errorf("Timecount.expired() = %v, want %v", got, tt.want)
			}
		})
	}
}
