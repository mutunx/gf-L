package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GBgrV4cxIAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjLNEJCW8iY/0/m4nsq693nfI2afKt7Y1GelW7+TPlb6mU7Xx7SCXPIOvRmbK2VSsvWfHeCFjoUeNhwFZy4NiyziaGqetX1S/zu1RvqKq7NPeSgNo3LpeqqfM+//jx8ePJ4pjvv8ObFDiU5jKu4uqcIXjk4vVJMzKiJ847muSWkWzT3VPKlvJMirvTupjvZMDyC8bV1/qCTRhmL/qrGDQ/UXT2mvM3jiW92s41N3LlSkfzC+F269sZ2iY6GyVwvF6Z2VV6kf232sP+bIE1Mz49Pr502TLGSeWbM/NjPv/UNmRwUTDd4KfIG7OZXXLqkoV+/AsctIRTdny5c3LfD5OERXozTgUHMC9ktL1zmyXtbkb+UpEZrlZFiy+H8PkKOXxab78wYn7Dlfkdt/0VLwRzqKn1eSbzRz16PoNTPjE4fN5dr/vVD9Lq1hh7lEwQ3tfyxNqjhZ/z4JOX626lv91+p/b11d4rLCGCqfM+7GEx3eDH/zNGM6w5gjfJ9IXN11WMQffYXTxXTlDdnKY3/Uoj2+TJPz8u6RH9IdAp4RH3OW36ebkll4OipEqlVv+5nvu65rDVzY3bu64djd+z7ax+XU6s2Y9zQvqs316l61iulTebe+xPxeZ9GffX3N7tmr7y2OnXZ1/bLi86sb2v8mXXeblb9XLsurcZWDImOBrNSfBg3bznYd4sD0nHXVJhq5lvX1WTMjugyyP+8qorB1/lx00zMuQZ2A5NYKnLVmISNVQ2SDrUESo75XfEyWTV98x7XFTPJ16t/v5vddL2f3qWVxcVf2PXv15RqhfBN2+zjvXW4OX/+n+/3rDv2qo52+t+29qJHosSyvZ8fnK5qibr8SIjg6wHs1bu+9ZyruBtjLvlYcmGIGbhb2znCw8+/t9vwqO77V3DPZMzCUKrIvvntHm4im1fsIP1K2OcF/eneXn/j79YeOF4jMyTx68Dt19qVF++0EajKC/bQCEx48RrVba4mpK9zbF5pYdXhr29u+hD35SNdXHPuK0PebS9n2+Rzn7o+jKVpQ89lnFXlZ9QK5VUnusoKlTfxezlHhLj2OW1pZh9c1mOrpzqLPHZNQuTdGU/fJKzv5UT9+9QevbZX163c+PXR9y2v+u+dfIdDx9TD82AxelXtfRv/w4/63ZUK+z944sv+SdWCKxQYHr6gLlHQDOr4uU+rT/SDAz//wd4s3PM13zh4cfMwHBXmoEBljEZMDImOyJjgvPibNEJCSDdyGoCvBmZRJgRGRvZZFDGhoFtjSASbzZHGIXdKRAgwPDfcSYzAxaHsbKB5JkYmBg6GRgYrjODeIAAAAD//8LmtpB2BAAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
