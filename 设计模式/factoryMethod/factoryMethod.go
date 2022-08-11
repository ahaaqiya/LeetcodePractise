package factoryMethod

//用户只需要知道具体工厂的名称就可以得到索要的产品，无需知道产品的具体创建过程
//新产品创建，只需要多写一个相应的工厂类
//工厂方法模式包含：抽象工厂，抽象产品，具体工厂，具体产品
//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

type OperatorBase struct {
	a,b int
}

func (o *OperatorBase) SetA(a int)  {
	o.a = a
}

func (o *OperatorBase) SetB(b int)  {
	o.b = b
}

type PlusOperator struct {
	*OperatorBase
}

//Plus工厂类
type PlusOperatorFactory struct {}

type MinusOperator struct {
	*OperatorBase
}

type MinusOperatorFactory struct {}

func (o MinusOperator) Result() int {
	return o.a-o.b
}

func (o PlusOperator) Result() int  {
	return o.a+o.b
}

func (o MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}

}

func (o PlusOperatorFactory) Create() Operator  {
	return &PlusOperator{
		&OperatorBase{},
	}
}

