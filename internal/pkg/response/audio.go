package response

// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
}

type SendAudioResult Audio
