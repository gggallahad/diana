package handler

import "github.com/gggallahad/gui"

func (h *handler) MoveCamera(ctx *gui.Context, eventType gui.Event) {
	switch event := eventType.(type) {
	case *gui.EventKey:
		if event.Key == gui.KeyArrowUp {
			err := h.moveCamera(ctx, 0, -1)
			if err != nil {
				return
			}
		}
		if event.Key == gui.KeyArrowDown {
			err := h.moveCamera(ctx, 0, 1)
			if err != nil {
				return
			}
		}
		// if event.Key == gui.KeyArrowLeft {
		// 	err := h.moveCamera(ctx, -1, 0)
		// 	if err != nil {
		// 		return
		// 	}
		// }
		// if event.Key == gui.KeyArrowRight {
		// 	err := h.moveCamera(ctx, 1, 0)
		// 	if err != nil {
		// 		return
		// 	}
		// }
	}
}

func (h *handler) moveCamera(ctx *gui.Context, viewPositionOffsetX, viewPositionOffsetY int) error {
	h.updateViewPosition(viewPositionOffsetX, viewPositionOffsetY)
	err := h.updateViewContent(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (h *handler) updateViewPosition(viewPositionOffsetX, viewPositionOffsetY int) {
	h.viewPositionX += viewPositionOffsetX
	h.viewPositionY += viewPositionOffsetY

	if h.viewPositionX < 0 {
		h.viewPositionX = 0
	}
	if h.viewPositionY < 0 {
		h.viewPositionY = 0
	}
}

func (h *handler) updateViewContent(ctx *gui.Context) error {
	ctx.SetViewPosition(h.viewPositionX, h.viewPositionY)

	h.drawMutex.Lock()
	defer h.drawMutex.Unlock()

	err := ctx.UpdateViewContent()
	if err != nil {
		return err
	}

	return nil
}
