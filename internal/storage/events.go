package storage

import "time"

type Event struct {
	ID        int
	Title     string
	Service   string
	Token     string
	CreatedAt time.Time
}

func GetEventByService(service string) (*Event, error) {
	event := &Event{}
	row := DB.QueryRow("SELECT id, title, service, token, created_at FROM events WHERE service = ?", service)

	err := row.Scan(&event.ID, &event.Title, &event.Service, &event.Token, &event.CreatedAt)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetEventByToken(token string) (*Event, error) {
	event := &Event{}
	row := DB.QueryRow("SELECT id, title, service, token, created_at FROM events WHERE token = ?", token)

	err := row.Scan(&event.ID, &event.Title, &event.Service, &event.Token, &event.CreatedAt)

	if err != nil {
		return nil, err
	}
	return event, nil
}

func GetListEvents() ([]Event, error) {
	rows, err := DB.Query("SELECT id, title, service, token, created_at FROM events ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var e Event

		err := rows.Scan(&e.ID, &e.Title, &e.Service, &e.Token, &e.CreatedAt)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func CreateEvent(title string, service string) (string, error) {
	token, err := GenerateToken()
	if err != nil {
		return "", err
	}
	_, err = DB.Exec("INSERT INTO events (title, service, token) VALUES (?, ?, ?)", title, service, token)
	return token, err
}
