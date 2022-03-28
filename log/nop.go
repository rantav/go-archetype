package log

type NopLogger struct{}

func (n NopLogger) Debugf(string, ...interface{}) {}
func (n NopLogger) Infof(string, ...interface{})  {}
func (n NopLogger) Warnf(string, ...interface{})  {}
func (n NopLogger) Errorf(string, ...interface{}) {}
func (n NopLogger) Fatalf(string, ...interface{}) {}
