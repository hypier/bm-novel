package notice

import "time"

type Notice struct {
	NoticeId    string
	NoticeTitle string
	Context     string
	NoticeTime  time.Time
	IsRead      bool
	UserId      string
}
