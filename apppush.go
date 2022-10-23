package apppush

type Client interface {
	Send(
		id string,
		title string,
		body string,
		data map[string]string,
	) error
	SendMulticast(
		ids []string,
		title string,
		body string,
		data map[string]string,
	) error
}
