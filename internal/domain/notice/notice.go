package notice

import "time"

// Notice 通知
type Notice struct {
	NoticeID    string
	NoticeTitle string
	Context     string
	NoticeTime  time.Time
	IsRead      bool
	UserID      string
}
