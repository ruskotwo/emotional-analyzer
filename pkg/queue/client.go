package queue

type CallbackFunc func(payload interface{})

type Client interface {
	Publish(queue string, body []byte) error
	Consume(queue string, callback CallbackFunc)
}
