package http

type RequestOptions struct {
	Method      string
	URL         string
	QueryParams map[string]string
	Headers     map[string]string
	Result      any // ponteiro para struct de resposta
}
