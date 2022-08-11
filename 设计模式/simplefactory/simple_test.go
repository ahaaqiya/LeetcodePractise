package simplefactory

import "testing"

func TestType(t *testing.T)  {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi Tom" {
		t.Fatal("type1 fail")
	}
}

func TestType2(t *testing.T)  {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello Tom"{
		t.Fatal("type2 fail")
	}
}

/*func TestHellpAPI_Say(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hello := &HellpAPI{}
			if got := hello.Say(tt.args.name); got != tt.want {
				t.Errorf("Say() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewAPI(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want API
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAPI(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hiAPI_Say(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hi := &hiAPI{}
			if got := hi.Say(tt.args.name); got != tt.want {
				t.Errorf("Say() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/