package bigint

import (
	"testing"
)

func TestNewInt(t *testing.T) {
	tests := []struct {
		num  string
		want string
		err  error
	}{
		{num: "12", want: "12"},
		{num: "-44", want: "-44"},
		{num: "0", want: "0"},
		{num: "00023", want: "23"},
		{num: "-00034", want: "-34"},
		{num: "+23", want: "23"},
	}
	for _, tt := range tests {
		got, _ := NewInt(tt.num)
		if got.Value != tt.want {
			t.Errorf("got %v but want %v, input value: %v", got.Value, tt.want, tt.num)
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		num  Bigint
		want error
	}{
		{Bigint{Value: "21"}, nil},
		{Bigint{Value: "-21"}, nil},
		{Bigint{Value: "++21"}, ErrorNotNumber},
	}
	for _, tt := range tests {
		got := tt.num.Set(tt.num.Value)
		if got != tt.want {
			t.Errorf("got %v but want %v, input value: %v", got, tt.want, tt.num.Value)
		}
	}
}

func TestAdd(t *testing.T) {
	got := Add(Bigint{Value: "12"}, Bigint{Value: "22"})
	want := Bigint{Value: "34"}

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}

}

func TestSub(t *testing.T) {
	got := Sub(Bigint{Value: "34"}, Bigint{Value: "22"})
	want := Bigint{Value: "12"}

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}

}
func TestMultiply(t *testing.T) {
	got := Multiply(Bigint{Value: "4"}, Bigint{Value: "3"})
	want := Bigint{Value: "12"}

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}

}
func TestMod(t *testing.T) {
	got := Mod(Bigint{Value: "36"}, Bigint{Value: "9"})
	want := Bigint{Value: "4"}

	if got != want {
		t.Errorf("got %v but want %v", got, want)
	}

}
func TestAbs(t *testing.T) {
	addTests := []struct {
		num, want Bigint
	}{
		{Bigint{Value: "123"},
			Bigint{Value: "123"}},
		{Bigint{Value: "-123"},
			Bigint{Value: "123"}},
	}
	for _, test := range addTests {
		got := test.num.Abs()
		if got != test.want {
			t.Errorf("got %v but want %v, input value: %v", got.Value, test.want.Value, test.num.Value)
		}
	}
}
