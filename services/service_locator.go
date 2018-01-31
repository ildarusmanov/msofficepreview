package services

type ServiceLocator struct {
	services map[string]interface{}
}

func CreateServiceLocator() *ServiceLocator {
	return &ServiceLocator{services: make(map[string]interface{})}
}

func (s *ServiceLocator) Set(key string, value interface{}) {
	s.services[key] = value
}

func (s *ServiceLocator) Get(key string) interface{} {
	if value, ok := s.services[key]; ok {
		return value
	}

	return nil
}
