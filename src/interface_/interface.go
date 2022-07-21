package interface_

type openDB interface {
	GetData() string
	WriteData(data string)
	Close() error
}
