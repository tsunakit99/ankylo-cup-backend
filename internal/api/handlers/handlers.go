package handlers

import (
	"context"
	"errors"
)

func getRoomIDFromContext(ctx context.Context) (int32, error) {
    roomIDValue := ctx.Value("room_id")
    if roomIDValue == nil {
        return 0, errors.New("room_id not found in context")
    }
    roomID, ok := roomIDValue.(int32)
    if !ok {
        return 0, errors.New("room_id has invalid type")
    }
    return roomID, nil
}
