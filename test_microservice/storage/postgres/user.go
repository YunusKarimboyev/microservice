package postgres

import (
	"log"

	u "github.com/double/test_microservice/genproto/user"
)

func (r *UserRepo) CreateUser(user *u.UserRequest) (*u.UserResponse, error) {
	var res u.UserResponse
	err := r.db.QueryRow(`
	INSERT INTO 
	  users(first_name, last_name, email) 
	VALUES
	  ($1, $2, $3) 
	RETURNING 
	  id, first_name, last_name, email`,
		user.FirstName, user.LastName, user.Email).
		Scan(&res.Id, &res.FirstName, &res.LastName, &res.Email)
	if err != nil {
		log.Println("Error inserting user info")
		return &u.UserResponse{}, err
	}
	return &res, nil
}

func (r *UserRepo) GetUserById(user *u.UserId) (*u.UserResponse, error) {
	var res u.UserResponse
	err := r.db.QueryRow(`
	SELECT
	  id, first_name, last_name, email
	FROM 
	  users
	WHERE
	  id=$1
	`, user.Id).Scan(
		&res.Id, &res.FirstName, &res.LastName, &res.Email,
	)
	if err != nil {
		return &u.UserResponse{}, err
	}
	return &res, nil
}

func (r *UserRepo) GetUsersAll(user *u.UserListReq) (*u.Users, error){
	var res u.Users
	query := `
	SELECT id, first_name, last_name, email
	FROM 
		users limit $1`
	rows, err := r.db.Query(query, user.Limit)
	if err != nil {
		return &u.Users{}, err
	}
	for rows.Next(){
		temp := u.UserResponse{}
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Email)

		if err != nil {
			return &u.Users{}, err
		}
		res.Users = append(res.Users, &temp)
	}
	return &res, nil
}

func (r *UserRepo) UpdateUser(user *u.UserUpdateReq) (*u.UserResponse, error){
	var res u.UserResponse
	err := r.db.QueryRow(`
	UPDATE users SET
		first_name = $1, last_name = $2, email = $3, updated_at = NOW()
	WHERE id = $4 AND deleted_at IS NULL
	RETURNING id, first_name, last_name, email
	`, user.FirstName, user.LaastName, user.Email, user.Id).Scan(
		&res.Id, &res.FirstName, &res.LastName, &res.Email)
	
		if err != nil {
			return &u.UserResponse{}, err
		}
	return &res, nil
}


func (r *UserRepo) DeleteUser(user *u.UserDeleteReq) (*u.Users, error){
	var res u.Users
	query := `DELETE FROM users WHERE id = $1`
	rows, err := r.db.Query(query, user.Id)
	if err != nil {
		return &u.Users{}, err
	}
	for rows.Next(){
		temp := u.UserResponse{}
		err = rows.Scan(&temp.Id, &temp.FirstName, &temp.LastName, &temp.Email)

		if err != nil {
			return &u.Users{}, err
		}
		res.Users = append(res.Users, &temp)
	}
	return &res, nil
}

func (r *UserRepo) GetUserByPostId(id int64) (*u.UserResponseForPost, error) {
	res := u.UserResponseForPost{}
	err := r.db.QueryRow(`
	SELECT 
	  id, post_id, first_name 
	FROM 
	  users 
	WHERE 
	  post_id=$1`, id).Scan(&res.Id, &res.PostId, &res.FirstName)
	if err != nil {
	  return &u.UserResponseForPost{}, err
	}
	return &res, nil
  }