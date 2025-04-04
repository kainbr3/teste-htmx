package brazaauth

type GenerateTokenRequest struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type GenerateTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
type DecodeTokenRequest struct {
	AccessToken string `json:"accessToken"`
}

type DecodeTokenResponse struct {
	Sub      string `json:"sub"`
	TokenUse string `json:"token_use"`
	Scope    string `json:"scope"`
	AuthTime int    `json:"auth_time"`
	Iss      string `json:"iss"`
	Exp      int    `json:"exp"`
	Iat      int    `json:"iat"`
	Version  int    `json:"version"`
	Jti      string `json:"jti"`
	ClientID string `json:"client_id"`
}
