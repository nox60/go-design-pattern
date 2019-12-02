package factoryMethod

type Bmw struct {
	CarType CarType
}

func (b Bmw) Drive() string {
	return b.CarType.CarName + "," + b.CarType.CarBrand
}
