package queue

type Client interface {
	Publish(queue string, body []byte) error
}
