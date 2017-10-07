package main

// User  is the database structure for the users table
type User struct {
	ID    int
	Email string
}

// InsertUser is used for the createUser mutation call
func InsertUser(user *User) error {
	var id int
	err := db.QueryRow(`
		INSERT INTO users(email)
		VALUES ($1)
 RETURNING id
	`, user.Email).Scan(&id)
	if err != nil {
		return err
	}
	user.ID = id
	return nil
}

// GetUserByID is used for the user query call
func GetUserByID(id int) (*User, error) {
	var email string
	err := db.QueryRow("SELECT email FROM users WHERE id=$1", id).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

// RemoveUserByID is used for the removeUser mutation call
func RemoveUserByID(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}

// Follow is used for the follow mutation call
func Follow(followerID, followeeID int) error {
	_, err := db.Exec(`
		INSERT INTO followers(follower_id, followee_id)
		VALUES ($1, $2)
	`, followerID, followeeID)
	return err
}

// Unfollow is used for the unfollow mutation call
func Unfollow(followerID, followeeID int) error {
	_, err := db.Exec(`
		DELETE
		  FROM followers
		 WHERE follower_id=$1
		   AND followee_id=$2
	`, followerID, followeeID)
	return err
}

// GetFollowerByIDAndUser is used for the follower query call
func GetFollowerByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := db.QueryRow(`
		SELECT u.email
		  FROM users AS u, followers AS f
		 WHERE u.id = f.follower_id
		   AND f.follower_id=$1
		   AND f.followee_id=$2
		 LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

// GetFollowersForUser is used for the follower query call
func GetFollowersForUser(id int) ([]*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.email
		  FROM users AS u, followers AS f
		 WHERE u.id=f.follower_id
		   AND f.followee_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}

// GetFolloweeByIDAndUser is used for the followee query call
func GetFolloweeByIDAndUser(id int, userID int) (*User, error) {
	var email string
	err := db.QueryRow(`
		SELECT u.email
		  FROM users AS u, followers AS f
		 WHERE u.id = f.followee_id
		   AND f.followee_id=$1
		   AND f.follower_id=$2
		 LIMIT 1
	`, id, userID).Scan(&email)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:    id,
		Email: email,
	}, nil
}

// GetFolloweesForUser is used for the followee query call
func GetFolloweesForUser(id int) ([]*User, error) {
	rows, err := db.Query(`
		SELECT u.id, u.email
		  FROM users AS u, followers AS f
		 WHERE u.id=f.follower_id
		   AND f.follower_id=$1
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var (
		users = []*User{}
		uid   int
		email string
	)
	for rows.Next() {
		if err = rows.Scan(&uid, &email); err != nil {
			return nil, err
		}
		users = append(users, &User{ID: id, Email: email})
	}
	return users, nil
}
