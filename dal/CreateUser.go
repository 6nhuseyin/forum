package dal

func CreateUser(Username, Email, Password string) error{
	var UserQuery string = "INSERT INTO Users(Username,Email,Password) VALUES(?,?,?)"


	//Prepare the SQL statement
	PrepUserStatement, err := db.Prepare(UserQuery)
	if err != nil {
		return err
	}

	//add hashing of password

	
	defer PrepUserStatement.Close()

	_,err = PrepUserStatement.Exec(Username,Email,Password)
	if err != nil{
		return err
	}

	return nil
}