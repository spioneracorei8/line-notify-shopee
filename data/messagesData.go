package data

type SSendMessage struct {
	Message  []string
	ImageURL []string
}

var Data SSendMessage

func MessagesData() {
	Data = SSendMessage{
		Message: []string{
			"ซื้อของ 1 บาทครับพรี่!!",
			"ใกล้จะถึงเวลาซื้อของแล้วววลูกพรี่!!!!",
			"เห็นข้อความไหมมลูกพรี่!!",
			"เห็นข้อความเถอะลูกพรี่!",
			"ได้ซื้อของแล้วใช่ไหมม 🥺",
		},
		ImageURL: []string{
			"https://i.pinimg.com/564x/12/67/97/126797f1e5f5cfb82c2e5d51bda95639.jpg",
			"https://i.pinimg.com/564x/f1/cb/6d/f1cb6d236f7106246cac52514835f169.jpg",
			"https://i.pinimg.com/564x/5c/d6/87/5cd6876281c90330ebdee4269d6135aa.jpg",
			"https://i.pinimg.com/564x/db/4a/33/db4a333e8791e3e8e5db7bdd5a0f284f.jpg",
			"https://i.pinimg.com/564x/2b/c6/72/2bc672dbb357088610428f98c7e57a8d.jpg",
		},
	}
}
