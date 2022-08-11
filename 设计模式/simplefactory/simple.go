package simplefactory


import "fmt"

//简单工厂，我们将要创建的对象叫做“产品”，把创建产品的对象叫工厂，如果创建的产品不多，只需要一个工厂类就可以完场，这种模式就叫简单工厂
//简单工厂模式有一个具体的工厂类
//简单工厂里，多一种产品就增加一个具体的工厂类，所以不够好


type API interface {
	Say(name string) string
}

type hiAPI struct {}

func (hi *hiAPI) Say(name string) string  {
	fmt.Println()
	return fmt.Sprintf("Hi,%s",name)
}

type HellpAPI struct {}

func (hello *HellpAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s",name)
}

func NewAPI(i int) API {
	if i==1{
		return &hiAPI{}
	}
	if i==2{
		return &HellpAPI{}
	}
	return nil
}