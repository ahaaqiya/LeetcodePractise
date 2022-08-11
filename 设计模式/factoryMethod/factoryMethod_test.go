package factoryMethod


/*func TestMinusOperatorFactory_Create(t *testing.T) {
	tests := []struct {
		name string
		want Operator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := MinusOperatorFactory{}
			if got := o.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinusOperator_Result(t *testing.T) {
	type fields struct {
		OperatorBase *OperatorBase
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := MinusOperator{
				OperatorBase: tt.fields.OperatorBase,
			}
			if got := o.Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOperatorBase_SetA(t *testing.T) {
	type fields struct {
		a int
		b int
	}
	type args struct {
		a int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatorBase{
				a: tt.fields.a,
				b: tt.fields.b,
			}
		})
	}
}

func TestOperatorBase_SetB(t *testing.T) {
	type fields struct {
		a int
		b int
	}
	type args struct {
		b int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OperatorBase{
				a: tt.fields.a,
				b: tt.fields.b,
			}
		})
	}
}

func TestPlusOperatorFactory_Create(t *testing.T) {
	tests := []struct {
		name string
		want Operator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := PlusOperatorFactory{}
			if got := o.Create(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlusOperator_Result(t *testing.T) {
	type fields struct {
		OperatorBase *OperatorBase
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := PlusOperator{
				OperatorBase: tt.fields.OperatorBase,
			}
			if got := o.Result(); got != tt.want {
				t.Errorf("Result() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/