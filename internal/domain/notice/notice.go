package notice

import "time"

type Notice struct {
	NoticeID    string
	NoticeTitle string
	Context     string
	NoticeTime  time.Time
	IsRead      bool
	UserId      string
}
