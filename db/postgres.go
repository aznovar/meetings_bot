package db

import (
	"database/sql"
	"errors"
	"schedulertgbot/models"
	"time"
)

type MeetingRepository struct {
	DB *sql.DB
}

func NewMeetingRepository(db *sql.DB) *MeetingRepository {
	return &MeetingRepository{DB: db}
}

func (r *MeetingRepository) AddMeeting(title string, date time.Time, participants string) error {
	_, err := r.DB.Exec("INSERT INTO meetings (title, date, participants) VALUES ($1, $2, $3)", title, date, participants)
	return err
}

func (r *MeetingRepository) GetMeetings() ([]models.Meeting, error) {
	rows, err := r.DB.Query("SELECT id, title, date, participants, summary FROM meetings")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if errors.Is(err, rows.Scan(&meeting.ID, &meeting.Title, &meeting.Date, &meeting.Participants, &meeting.Summary)) {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}
	return meetings, nil
}

func (r *MeetingRepository) AddSummary(meetingID int, summary string) error {
	_, err := r.DB.Exec("UPDATE meetings SET summary = $1 WHERE id = $2", summary, meetingID)
	return err
}

func (r *MeetingRepository) GetUpcomingMeetings() ([]models.Meeting, error) {
	rows, err := r.DB.Query("SELECT id, title, date, participants FROM meetings WHERE date > now() AND date < now() + interval '1 day'")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var meetings []models.Meeting
	for rows.Next() {
		var meeting models.Meeting
		if err := rows.Scan(&meeting.ID, &meeting.Title, &meeting.Date, &meeting.Participants); err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)

	}
	return meetings, nil
}
