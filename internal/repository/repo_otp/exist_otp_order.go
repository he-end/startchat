package repootp

import "github.com/hend41234/startchat/internal/repository"

func ExistOtpOrder(email string) (ok bool, err error) {
	tx, err := repository.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return false, err
	}
	{
		var result bool
		err = tx.QueryRow(queExistOtpOrder, email).Scan(&result)
		if err != nil {
			tx.Rollback()
			// fmt.Println("result : ", result)
			// log.Println("error get data : ", err)
			return false, err
		}
		{
			return result, nil
		}
	}
}
