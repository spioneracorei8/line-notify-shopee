package data

type SSendMessage struct {
	Message  []string
	ImageURL []string
}

var Data SSendMessage

func MessagesData() {
	Data = SSendMessage{
		Message: []string{
			"แอปชอบปี้ครับพรี่!!!",
			"ซื้อของ 1 บาทครับพรี่!!",
			"ใกล้จะถึงเวลาซื้อของแล้วววลูกพรี่!!!!",
			"เห็นข้อความเถอะลูกพรี่!",
			"ได้ซื้อของแล้วใช่ไหมม 🥺",
		},
		ImageURL: []string{
			"https://i.pinimg.com/564x/4c/62/25/4c6225863c7cd8f363aaa0702c4e87e5.jpg",
			"https://i.pinimg.com/564x/ee/6a/9d/ee6a9d2d83dbf83cb81479eb7fba6ad8.jpg",
			"https://i.pinimg.com/564x/01/7c/94/017c944ee7e0ba6d6b2309ee2b48f2c8.jpg",
			"https://i.pinimg.com/564x/24/e9/39/24e9399036719c39a419540f27d694ad.jpg",
			"https://i.pinimg.com/736x/e0/7d/e3/e07de3286a1cc4b6ae7c1bdbb19c9898.jpg",
		},
	}
}
