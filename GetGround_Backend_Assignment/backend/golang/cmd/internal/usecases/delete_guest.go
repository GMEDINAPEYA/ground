package usecases

import "context"

func (g GetGroundServiceImpl) DeleteGuest(ctx context.Context, name string) error {
	err := g.guestRepository.DeleteGuest(name)
	if err != nil {
		return err
	}
	return nil
}
