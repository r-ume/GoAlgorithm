package singleton

type singleton struct {
}

var instance *singleton

// GetInstance get the same instance
func GetInstance() *singleton.singleton {
	if instance == nil {
		instance = &singleton{}
	}

	return singleton
}
