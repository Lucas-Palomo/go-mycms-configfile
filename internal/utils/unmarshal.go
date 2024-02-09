package utils

func Unmarshal(
	data []byte,
	model interface{},
	method func(data []byte, model interface{}) error) error {

	err := method(data, model)
	if err != nil {
		return err
	}

	return nil
}

func Marshal(
	model interface{},
	method func(model interface{}) ([]byte, error)) ([]byte, error) {

	data, err := method(model)
	if err != nil {
		return data, err
	}

	return data, nil
}
