package models

type Update struct {
	ChatMember struct {
		Chat struct {
			ID       int    `json:"id"`
			Title    string `json:"title"`
			Type     string `json:"type"`
			Username string `json:"username"`
		} `json:"chat"`
		Date int `json:"date"`
		From struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			IsBot     bool   `json:"is_bot"`
			IsPremium bool   `json:"is_premium"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"from"`
		InviteLink struct {
			CreatesJoinRequest bool `json:"creates_join_request"`
			Creator            struct {
				FirstName    string `json:"first_name"`
				ID           int    `json:"id"`
				IsBot        bool   `json:"is_bot"`
				LanguageCode string `json:"language_code"`
				LastName     string `json:"last_name"`
				Username     string `json:"username"`
			} `json:"creator"`
			InviteLink string `json:"invite_link"`
			IsPrimary  bool   `json:"is_primary"`
			IsRevoked  bool   `json:"is_revoked"`
			Name       string `json:"name"`
		} `json:"invite_link"`
		NewChatMember struct {
			Status string `json:"status"`
			User   struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				IsPremium bool   `json:"is_premium"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"user"`
		} `json:"new_chat_member"`
		OldChatMember struct {
			Status string `json:"status"`
			User   struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				IsBot     bool   `json:"is_bot"`
				IsPremium bool   `json:"is_premium"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"user"`
		} `json:"old_chat_member"`
	} `json:"chat_member"`
	UpdateID int `json:"update_id"`
}

type Updates struct {
	Updates []Update `json:"result"`
	Ok      bool     `json:"ok"`
}
