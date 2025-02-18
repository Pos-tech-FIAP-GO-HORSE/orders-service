package contract

type PaymentEvent struct {
	ID       string `json:"id"`
	PublicID string `json:"public_id"`
	Status   string `json:"status"`
	QRCode   string `json:"qr_code"`
}
