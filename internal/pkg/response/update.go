package response

// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateId int64   `json:"update_id"`
	Message  Message `json:"message"`
}

// https://core.telegram.org/bots/api#getupdates
type GetUpdates Response[GetUpdatesResult]
type GetUpdatesResult []Update
