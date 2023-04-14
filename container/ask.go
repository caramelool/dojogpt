package container

import "dojogpt/controller"

var askController controller.Ask

func Ask() controller.Ask {
	if askController == nil {
		askController = controller.NewAsk(
			OpenAI(),
		)
	}
	return askController
}
