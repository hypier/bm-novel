package notice

// INoticeService 小说服务接口
type INoticeService interface {
	Create(notice *Notice)
	Read()
}
