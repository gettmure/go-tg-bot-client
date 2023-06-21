package response

// https://core.telegram.org/bots/api#making-requests
type Response[T Result] struct {
	Ok          bool    `json:"ok"`
	Description *string `json:"description"`
	Result      *T      `json:"result"`
}

type Result interface {
	GetMeResult | GetUpdatesResult | SendMessageResult | SendAudioResult
}
