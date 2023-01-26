package usecases

import "context"

func (g TableUseCaseImpl) GetEmptySeats(ctx context.Context) (int, error) {
	emptySeats, err := g.tableRepository.GetEmptySeats()
	if err != nil {
		return 0, err
	}
	return emptySeats, nil
}
