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
