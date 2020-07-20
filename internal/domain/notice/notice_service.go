package notice

type INoticeService interface {
	Create(notice *Notice)
	Read()
}
