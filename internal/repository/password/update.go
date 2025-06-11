package password

func (r *repository) UpdateByUser(
	userID string, newPassword string,
) error {
	if _, err := r.DB.Exec(
		`UPDATE passwords SET password = $1 WHERE "user" = $2`,
		newPassword, userID,
	); err != nil {
		return err
	}

	return nil
}
