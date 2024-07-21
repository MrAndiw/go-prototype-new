package response

type HomeResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Favicon     string `json:"favicon"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
}
