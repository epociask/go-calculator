package calculator

type Calculator struct {
	lastCalculation float64
}

func New(initCalcs ...float64) (*Calculator, error) {

	if len(initCalcs) == 0 {
		return &Calculator{lastCalculation: 0.0}, nil
	}
	if len(initCalcs) == 1 {
		return &Calculator{lastCalculation: initCalcs[0]}, nil
	}

	return nil, &InitError{}

}

func (c *Calculator) GetLastCalculation() float64 {
	return c.lastCalculation
}

func getFloat64(number interface{}) (float64, error) {

	switch value := number.(type) {

	case int:
		return float64(value), nil
	case float64:
		return value, nil

	default:
		return 0.0, &InputError{}
	}
}

func (c *Calculator) Add(num1 interface{}, num2 interface{}) (float64, error) {

	var numFloat1, numFloat2 float64
	var err error

	if numFloat1, err = getFloat64(num1); err != nil {
		return 0.0, err
	}

	if numFloat2, err = getFloat64(num2); err != nil {
		return 0.0, err
	}

	result := numFloat1 + numFloat2
	c.lastCalculation = result
	return result, nil
}

func (c *Calculator) AddToLastCalculation(num interface{}) (float64, error) {

	numFloat1, err := getFloat64(num)

	if err != nil {
		return 0.0, err
	}

	result := numFloat1 + c.lastCalculation
	c.lastCalculation = result
	return result, nil
}

func (c *Calculator) Subtract(num1 interface{}, num2 interface{}) (float64, error) {

	var numFloat1, numFloat2 float64
	var err error

	if numFloat1, err = getFloat64(num1); err != nil {
		return 0.0, err
	}

	if numFloat2, err = getFloat64(num2); err != nil {
		return 0.0, err
	}

	result := numFloat1 - numFloat2
	c.lastCalculation = result
	return result, nil
}

func (c *Calculator) SubtractFromLastCalculation(num interface{}) (float64, error) {

	numFloat1, err := getFloat64(num)

	if err != nil {
		return 0.0, err
	}

	result := c.lastCalculation - numFloat1
	c.lastCalculation = result
	return result, nil
}
