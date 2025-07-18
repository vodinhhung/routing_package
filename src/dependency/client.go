package dependency

type Client struct {
	ID    uint64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func GetClientByID(clientID uint64) (*Client, error) {
	var client Client
	if err := db.First(&client, clientID).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
