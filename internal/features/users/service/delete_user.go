package core_http_service

import "context"

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := s.userRepository.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}
