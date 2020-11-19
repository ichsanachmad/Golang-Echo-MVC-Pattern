package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
)

type ExampleModel struct {
	db settings.Database
}

// Get Example Post
func (ExampleModel ExampleModel) GetPosts(secret string) []entity.ExampleEntity {
	posts := []entity.ExampleEntity{
		{
			Title:       "NewsOne",
			Description: "NewsOneDescription",
		},
		{
			Title:       "NewsTwo",
			Description: "NewsTwoDescription",
		},
		{
			Title:       secret,
			Description: "NewsTwoDescription",
		},
	}

	return posts
}

// Example For Insert Update and Delete
// Get Example Post
//func (UserModel UserModel) IsEmailOrNoTelponExist(identity string) (bool, error) {
//	sqlStatement := "SELECT is_email_no_telp_exist($1)"
//	isExist := false
//
//	err := UserModel.db.GetDatabase().QueryRow(context.Background(),
//		sqlStatement,
//		identity,
//	).Scan(&isExist)
//
//	if err != nil {
//		utils.LogError(err, utils.DetailFunction())
//	}
//
//	return isExist, err
//}

// Example For Get From DB
//func (UserModel UserModel) GetUserBussinessTypesByUserID(userID string) (entity.BussinessTypeList, error) {
//	sqlStatement := "SELECT * FROM get_jenis_usaha($1)"
//
//	bisnisTypeList := entity.BussinessTypeList{}
//	res, err := UserModel.db.GetDatabase().Query(context.Background(),
//		sqlStatement,
//		userID,
//	)
//
//	defer utils.ResClose(res)
//
//	if err != nil {
//		utils.LogError(err, utils.DetailFunction())
//		return bisnisTypeList, err
//	}
//
//	for res.Next() {
//		tipeBisnis := entity.BussinessType{}
//		err := res.Scan(
//			&tipeBisnis.UserBussinessTypeID,
//			&tipeBisnis.UserBussinessType,
//		)
//		// Exit if we get an error
//		if err != nil {
//			utils.LogError(err, utils.DetailFunction())
//			return bisnisTypeList, err
//		}
//		bisnisTypeList.BussinessType = append(bisnisTypeList.BussinessType, tipeBisnis)
//	}
//
//	return bisnisTypeList, nil
//}
